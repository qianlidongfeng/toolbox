package toolbox

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func LoadConfig(config string,obj interface{}) error{
	data, err := ioutil.ReadFile(config)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(data, obj)
	if err != nil{
		return err
	}
	err=yaml.Unmarshal(data, obj)
	if err != nil{
		return err
	}
	return nil
}