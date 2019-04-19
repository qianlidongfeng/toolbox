package toolbox

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/headzoo/surf/errors"
)

type MySqlConfig struct{
	User string
	PassWord string
	Address string
	DataBase string
	Table string
	MaxOpenConns int
	MaxIdleConns int
}

func InitMysql(config MySqlConfig) (*sql.DB,error){
	db,err:=sql.Open("mysql",config.User+":"+config.PassWord+"@tcp("+config.Address+")/"+config.DataBase)
	if err != nil{
		return db,err
	}
	err = db.Ping()
	if err != nil{
		return db,err
	}
	db.SetMaxOpenConns(config.MaxOpenConns)
	db.SetMaxIdleConns(config.MaxIdleConns)
	return db,err
}

func CheckTable(db *sql.DB,table string,fields map[string]string) error{
	rows,err:=db.Query(fmt.Sprintf(`SHOW TABLES LIKE '%s'`,table))
	if err != nil{
		return err
	}
	rows.Next()
	var tb string
	err = rows.Scan(&tb)
	if err != nil{
		return errors.New(fmt.Sprintf("table %s not exist",table))
	}
	rows, err = db.Query("SELECT COLUMN_NAME,COLUMN_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='"+table+"'")
	if err != nil{
		return err
	}
	remoteFields:=make(map[string]string)
	for rows.Next(){
		var name,typ string
		err=rows.Scan(&name,&typ)
		if err != nil{
			return err
		}
		remoteFields[name]=typ
	}
	for k,v := range fields{
		columnType,ok:=remoteFields[k]
		if !ok{
			return errors.New(fmt.Sprintf("%s field %s not exist",table,k))
		}
		if columnType != v{
			return errors.New(fmt.Sprintf("table %s field %s error,need %s",table,k,v))
		}
	}
	return err
}


//default use utf8, innodb
func CheckAndFixTable(db *sql.DB,table string,fields map[string]string) error {
	_,err:= db.Exec(fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
		id BIGINT NOT NULL auto_increment,
		PRIMARY KEY (id)
	)ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci;`,table))
	if err != nil{
		return err
	}
	rows, err := db.Query("SELECT COLUMN_NAME,COLUMN_TYPE FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME='" + table + "'")
	if err != nil {
		return err
	}
	remoteFields := make(map[string]string)
	for rows.Next() {
		var name, typ string
		err = rows.Scan(&name, &typ)
		if err != nil {
			return err
		}
		remoteFields[name] = typ
	}
	for k, v := range fields {
		columnType, ok := remoteFields[k]
		if !ok {
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s",table,k,v))
			if err != nil{
				return err
			}
			continue
		}
		if columnType != v {
			_,err=db.Exec(fmt.Sprintf(`ALTER TABLE %s MODIFY COLUMN %s %s`,table,k,v))
			if err != nil{
				return err
			}
		}
	}
	return err
}