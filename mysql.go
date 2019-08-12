package toolbox

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
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
	if config.MaxOpenConns != 0{
		db.SetMaxOpenConns(config.MaxOpenConns)
	}
	if config.MaxIdleConns != 0{
		db.SetMaxIdleConns(config.MaxIdleConns)
	}
	return db,err
}

func CheckTable(db *sql.DB,table string,fields map[string]string) error{
	rows,err:=db.Query(fmt.Sprintf(`SHOW TABLES LIKE '%s'`,table))
	if err != nil{
		return err
	}
	defer rows.Close()
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


type MysqlIndex struct{
	Name string
	Typ string
	Method string
}

//default use utf8, innodb
//索引修复还不完美，索引修复不要和手动混用
func CheckAndFixTable(db *sql.DB,table string,fields map[string]string,index map[string]MysqlIndex) error {
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
	defer rows.Close()
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
			_,err:=db.Exec(fmt.Sprintf("ALTER TABLE `%s` ADD COLUMN `%s` %s",table,k,v))
			if err != nil{
				return err
			}
			continue
		}
		if columnType != v {
			_,err=db.Exec(fmt.Sprintf("ALTER TABLE `%s` MODIFY COLUMN `%s` %s",table,k,v))
			if err != nil{
				return err
			}
		}
	}
	for k,v:=range index{
		_,err:=db.Exec(fmt.Sprintf("ALTER TABLE %s ADD %s INDEX `%s`(`%s`) USING %s",table,v.Typ,v.Name,k,v.Method))
		if err !=nil{
			continue
		}
	}
	return err
}

func SelectMapFromMysql(db *sql.DB,selector string) ([]map[string]interface{},error){
	rows,err:=db.Query(selector)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	cols,err:=rows.ColumnTypes()
	if err != nil{
		return nil,err
	}
	count:=len(cols)
	result:=make([]map[string]interface{},0)
	values := make([]interface{}, count)
	for rows.Next(){
		scans := make([]interface{}, count)
		for i:=range scans{
			switch(cols[i].ScanType().Name()){
			case "uint32":
				scans[i]=new(uint32)
			case "int32":
				scans[i]=new(int32)
			case "uint64":
				scans[i]=new(uint64)
			case "int64":
				scans[i]=new(int32)
			case "int8":
				scans[i]=new(int8)
			case "uint8":
				scans[i]=new(uint8)
			case "int16":
				scans[i]=new(int16)
			case "uint16":
				scans[i]=new(uint16)
			case "float32":
				scans[i]=new(float32)
			case "float64":
				scans[i]=new(float64)
			default:
				scans[i]=&values[i]
			}
		}
		err=rows.Scan(scans...)
		if err != nil{
			return nil,err
		}
		row := make(map[string]interface{})
		for k, v := range scans { //每行数据是放在values里面，现在把它挪到row里
			key := cols[k].Name()
			if value,ok := values[k].([]byte);ok{
				row[key] = string(value)
			} else {
				row[key]=v
			}
		}
		result=append(result,row)
	}
	return result,nil
}

func SelectArrayFromMysql(db *sql.DB,selector string) ([][]interface{},error){
	rows,err:=db.Query(selector)
	if err != nil{
		return nil,err
	}
	defer rows.Close()
	cols,err:=rows.ColumnTypes()
	if err != nil{
		return nil,err
	}
	count:=len(cols)
	result:=make([][]interface{},0)
	values := make([]interface{}, count)
	for rows.Next(){
		scans := make([]interface{}, count)
		for i:=range scans{
			switch(cols[i].ScanType().Name()){
			case "uint32":
				scans[i]=new(uint32)
			case "int32":
				scans[i]=new(int32)
			case "uint64":
				scans[i]=new(uint64)
			case "int64":
				scans[i]=new(int32)
			case "int8":
				scans[i]=new(int8)
			case "uint8":
				scans[i]=new(uint8)
			case "int16":
				scans[i]=new(int16)
			case "uint16":
				scans[i]=new(uint16)
			case "float32":
				scans[i]=new(float32)
			case "float64":
				scans[i]=new(float64)
			default:
				scans[i]=&values[i]
			}
		}
		err=rows.Scan(scans...)
		if err != nil{
			return nil,err
		}
		row := make([]interface{},0)
		for k, v := range scans { //每行数据是放在values里面，现在把它挪到row里
			if value,ok := values[k].([]byte);ok{
				row=append(row,string(value))
			}else{
				row=append(row,v)
			}
		}
		result=append(result,row)
	}
	return result,nil
}