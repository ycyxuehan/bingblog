package mysql

import (
	"fmt"
	"testing"

)

type Student struct{
	ID int `mysql:"id,auto,primary"`
	Name string `mysql:"name"`
	Age int `mysql:"age"`
	Test bool
}

func (s *Student)Show(){
	fmt.Printf("id: %d\nname: %s\nage: %d\ntest:%t\n", s.ID, s.Name, s.Age, s.Test)
}

var mysql *MySQL

func init(){
	mysql = new(MySQL)
	mysql.URI = "root:Hello2019@tcp(127.0.0.1:3306)/test?charset=utf8"
	err := mysql.Connect()
	if err!=nil {
		panic(err)
	}
}
//
func TestFindOne(t *testing.T){
	s := Student{
		ID: 1,
	}
	err := mysql.FindOne(&s)
	if err != nil {
		t.Error(err)
	}
	s.Show()
}

func TestInsert(t *testing.T){
	s := Student{
		ID: 1,
		Name: "lucy",
		Age: 16,
	}
	res, err := mysql.Insert(&s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.RowsAffected())
}

func TestDelete(t *testing.T){
	s := Student{
		ID: 1,
		Name: "tom",
		Age:15,
	}
	res, err := mysql.Delete(&s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.RowsAffected())
}

func TestUpdate(t *testing.T){
	s := Student{
		ID: 2,
		Name: "jimy",
		Age:17,
	}
	res, err := mysql.Update(&s)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(res.RowsAffected())
}