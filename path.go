package toolbox

import (
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