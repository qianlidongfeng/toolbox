package toolbox

import (
	"os"
	"runtime"
	"syscall"
)

func RediRectOutPutToLog() (*os.File,error){
	logFile,err:=GetLogFile()
	if err != nil{
		return nil,err
	}
	stdout, err := os.OpenFile(logFile, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	if err != nil{
		return nil,err
	}
	switch runtime.GOOS {
	case "windows":
		os.Stdout = stdout
		os.Stderr = stdout
	default:
		syscall.Dup2(int(stdout.Fd()), 1)
		syscall.Dup2(int(stdout.Fd()), 2)
	}
	return stdout,nil
}
