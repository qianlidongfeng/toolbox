package toolbox

import "testing"

func TestAppDir(t *testing.T) {
	dir,err:=AppDir()
	_,_=dir,err
}

type config struct{
	Name string
	Age int
	Lover lover
}
type lover struct{
	Who string
	Do string
	Phone int
	Child []string
}
func TestLoadConfig(t *testing.T) {
	c:=config{}
	configFile,err:=AppPath()
	if err != nil{
		t.Error(err)
	}
	err=LoadConfig(configFile+".yaml",&c)
	if err != nil{
		t.Error(err)
	}
}

func TestCheckTable(t *testing.T) {
	cfg := MySqlConfig{}
	cfg.User="root"
	cfg.PassWord="333221"
	cfg.DataBase="actions"
	cfg.Address="127.0.0.1:3306"
	db,err:=InitMysql(cfg)
	if err != nil{
		t.Error(err)
	}
	fields:=make(map[string]string)
	fields["time"]="datetime"
	fields["label"]="varchar(64)"
	fields["url"]="varchar(255)"
	fields["meta"]="blob"
	fields["respy"]="int(11)"
	fields["parser"]="varchar(64)"
	fields["method"]="varchar(16)"
	fields["postdata"]="text"
	err = CheckTable(db,"action",fields)
	if err != nil{
		t.Error(err)
	}
}

func TestCheckAndFixTable(t *testing.T) {
	cfg := MySqlConfig{}
	cfg.User="root"
	cfg.PassWord="333221"
	cfg.DataBase="actions"
	cfg.Address="127.0.0.1:3306"
	db,err:=InitMysql(cfg)
	if err != nil{
		t.Error(err)
	}
	fields:=make(map[string]string)
	fields["time"]="datetime"
	fields["label"]="varchar(64)"
	fields["url"]="varchar(255)"
	fields["meta"]="blob"
	fields["respy"]="int"
	fields["parser"]="varchar(64)"
	fields["method"]="varchar(16)"
	fields["postdata"]="text"
	err = CheckAndFixTable(db,"action2",fields)
	if err != nil{
		t.Error(err)
	}
}

func TestStampToTime(t *testing.T) {
	n:=GetTimeMilliStamp()
	_=n
}

