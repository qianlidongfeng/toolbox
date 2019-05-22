package toolbox

import "time"

func GetTimeSecond() string{
	return time.Now().Format("2006-01-02 15:04:05")
}

func GetTimeSecondStamp() int64{
	return time.Now().Unix()
}

func GetTimeMilliStamp() int64{
	return time.Now().UnixNano()/1e6
}

func TimeToSecondStamp(t string) (int64,error){
	tm,err:=time.Parse("2006-01-02 15:04:05",t)
	if err!=nil{
		return 0,err
	}
	return tm.Unix(),nil
}

func TimeToMilliStamp(t string)(int64,error){
	tm,err:=time.Parse("2006-01-02 15:04:05",t)
	if err!=nil{
		return 0,err
	}
	return tm.UnixNano()/1e6,nil
}

func SecondStampToTime(s int64) string{
	return time.Unix(s,0).Format("2006-01-02 15:04:05")
}

func MilliStampToTime(s int64) string{
	return time.Unix(s/1000,0).Format("2006-01-02 15:04:05")
}