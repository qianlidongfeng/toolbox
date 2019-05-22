package toolbox

import (
	"flag"
	"os"
	"path/filepath"
)

func AppPath() (string,error){
	path,err:=os.Readlink("/proc/self/exe")
	return path,err
}

func AppDir()(string,error){
	path,err:=os.Readlink("/proc/self/exe")
	if err != nil{
		return path,err
	}
	dir, err := filepath.Abs(filepath.Dir(path))
	return dir,err
}

func GetConfigFile() (string,error){
	appPath,err :=AppPath()
	if err != nil{
		return "",err
	}
	configFile:=flag.String("c", appPath+".yaml", "the path of config file")
	flag.Parse()
	return *configFile,nil
}

func GetLogFile() (string,error){
	appPath,err :=AppPath()
	if err != nil{
		return "",err
	}
	configFile:=flag.String("l", appPath+".log", "the path of log file")
	flag.Parse()
	return *configFile,nil
}
