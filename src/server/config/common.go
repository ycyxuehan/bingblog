package config

import (
	"strings"
	"io"
	"bufio"
	"os"
	"fmt"

)
//Config server configuration type
type Config map[string]string
//Get get a value by key
func (c Config)Get(key string)string{
	for k, v := range c {
		if k == key{
			return v
		}
	}
	return ""
}

//Set set value to key
func (c Config)Set(k string, v string){
	c[k] = v
}

//Add add the key and value to configuration, if key is exists, return error
func (c Config)Add(k string, v string)error{
	for key := range c {
		if key == k {
			return fmt.Errorf("key [%s] is exists", k)
		}
	}
	c[k] = v
	return nil
}

//LoadFile load config file
func (c Config)LoadFile(f string)error{
	file, err := os.Open(f)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return err
		}
		if err == io.EOF {
			break
		}
		s := strings.Split(line, "=")
		if len(s) > 0 && s[0] != "" {
			if len(s) > 1 {
				c[s[0]] = strings.Trim(strings.Join(s[1:len(s)], "="), "\n")
			}else {
				c[s[0]] = ""
			}
		}
	}
	return nil
}

//New new a config for a file
func New(f string)(Config, error){
	c := Config{}
	err := c.LoadFile(f)
	return c, err
}

//Conf server configurations
var Conf Config

func init(){
	Conf = make(map[string]string)
}