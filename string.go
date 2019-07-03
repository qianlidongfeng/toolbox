package toolbox

import (
	"bytes"
	"github.com/headzoo/surf/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/transform"
	"io/ioutil"
	"regexp"
)

func GbkToUtf8b(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewDecoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func Utf8ToGbk(s string) (string, error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewEncoder())
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func GbkToUtf8(s string) (rs string,err error) {
	reader := transform.NewReader(bytes.NewReader([]byte(s)), simplifiedchinese.GBK.NewDecoder())
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return
	}
	rs=string(b)
	return
}

func Utf8ToGbkb(s []byte) ([]byte, error) {
	reader := transform.NewReader(bytes.NewReader(s), simplifiedchinese.GBK.NewEncoder())
	d, e := ioutil.ReadAll(reader)
	if e != nil {
		return nil, e
	}
	return d, nil
}

func IsUtf8(buf []byte) bool{
	nBytes := 0
	for i:= 0;i<len(buf);i++{
		if nBytes == 0{
			if (buf[i] & 0x80) != 0 {
				for (buf[i] & 0x80) != 0 {
					buf[i] <<= 1
					nBytes++
				}

				if nBytes < 2 || nBytes > 6 {
					return false
				}

				nBytes--
			}
		}else{
			if buf[i] & 0xc0 != 0x80{
				return false
			}
			nBytes--
		}
	}
	return nBytes == 0
}

func GetMapByReg(reg string,content string) (m map[string]string,err error){
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(e.(string))
			return
		}
	}()
	exp:=regexp.MustCompile(reg)
	dict:=exp.FindStringSubmatch(content)
	groupNames := exp.SubexpNames()
	m = make(map[string]string)
	for i, key := range groupNames {
		if i != 0 && key != "" {
			m[key] = dict[i]
		}
	}
	return
}