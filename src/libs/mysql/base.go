package mysql
import (
	"fmt"
	"reflect"
)


//Operation sql operation like select, update...
type Operation int
const (
	//SELECT select
	SELECT Operation = iota
	//UPDATE update
	UPDATE
	//DELETE delete
	DELETE
	//INSERT insert
	INSERT
	//CREATE create
	CREATE
)

func getColumns(i interface{})(string, []Column, error){
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	v := reflect.ValueOf(i)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	fc := t.NumField()
	fields := []Column{}
	for i := 0; i< fc; i++{
		f := t.Field(i)
		field, err := NewColumn(f)
		if err == nil {
			if f.Name != "" {
				if val := v.FieldByName(f.Name);val.IsValid() {
					field.Ref = val.Addr().Interface()
					fields = append(fields, field)
				}
			}
		}
	}
	return t.Name(), fields, nil
}

func toString(i interface{})(string, error){
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	switch t.Kind() {
	case reflect.String:
		return i.(string), nil
	case reflect.Int:
	case reflect.Int64:
		return fmt.Sprintf("%d", i), nil
	case reflect.Bool:
		return fmt.Sprintf("%t", i), nil
	}
	return "", fmt.Errorf("unkown type")
}