package mysql

import (
	"fmt"
	"strings"
	"reflect"

)
//Column mysql table column
type Column struct{
	FieldName string
	Name string
	AutoIncrement bool
	Primary bool
	Ref interface{}
}

//NewColumn new a column from reflect structfield
func NewColumn(sf reflect.StructField)(Column, error){
	f := Column{}
	if tag, ok := sf.Tag.Lookup("mysql"); ok {
		tags := strings.Split(tag, ",")
		f.FieldName = sf.Name
		f.Name = tags[0]
		switch len(tags) {
		case 2:
			if tags[1] == "auto" {
				f.AutoIncrement = true
			} else if tags[1] == "primary" {
				f.Primary = true
			}
			break
		case 3:
			f.Primary = true
			f.AutoIncrement = true
		}
		f.Ref = sf.Name
		return f, nil
	}
	return f, fmt.Errorf("field don't have tag mysql")
}
