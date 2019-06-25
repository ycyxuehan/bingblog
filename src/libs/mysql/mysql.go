package mysql

import (
	"time"
	"database/sql"
	"fmt"
	"reflect"
	"strings"

	//for mysql
	_ "github.com/go-sql-driver/mysql"
)

//MySQL mysql struct
type MySQL struct {
	URI string
	db  *sql.DB
}

//New new a mysql object
func New(uri string)*MySQL{
	m := MySQL{URI:uri}
	return &m
}

//Connect connect to mysql.URL
func (m *MySQL)Connect() error {
	if m.URI == "" {
		return fmt.Errorf("mysql connection uri error")
	}
	db, err := sql.Open("mysql", m.URI)
	if err != nil {
		return err
	}
	m.db = db
	return nil
}
//ConnectURI connect to mysql.URL
func (m *MySQL)ConnectURI(uri string) error {
	if uri == "" {
		return fmt.Errorf("mysql connection uri error")
	}
	m.URI = uri
	db, err := sql.Open("mysql", m.URI)
	if err != nil {
		return err
	}
	m.db = db
	return nil
}

//SetMaxOpenConnections SetMaxOpenConnections
func (m *MySQL)SetMaxOpenConnections(max int){
	m.db.SetMaxOpenConns(max)
}
//SetMaxIdleConnections SetMaxIdleConnections
func (m *MySQL)SetMaxIdleConnections(max int){
	m.db.SetMaxIdleConns(max)
}
//SetConnMaxLifetime SetConnMaxLifetime
func (m *MySQL)SetConnMaxLifetime(max time.Duration){
	m.db.SetConnMaxLifetime(max)
}

//Connected is the db is connected
func (m *MySQL)Connected()bool{
	return m.db.Ping() == nil
}

//Close close the db connection
func (m *MySQL)Close()error{
	return m.db.Close()
}

//toSQL
func (m *MySQL) toSQL(i interface{}, op Operation) (string, []interface{}, []interface{}, error) {
	table, columns, err := getColumns(i)
	if err != nil {
		return "", nil, nil, err
	}
	cols := []string{}
	vals := []interface{}{}
	where := []string{"1=1"}
	whereVals := []interface{}{}
	values := []string{}
	for _, column := range columns{
		if op == INSERT && !column.AutoIncrement {
			cols = append(cols, column.Name)
			vals = append(vals, column.Ref)
			values = append(values, "?")
		} else if op != INSERT {
			cols = append(cols, column.Name)
			vals = append(vals, column.Ref)
		}
		if column.Primary {
			where = append(where, fmt.Sprintf("%s=?", column.Name))
			whereVals = append(whereVals, column.Ref)
		}

	}
	s := ""
	switch op {
	case SELECT:
		s = fmt.Sprintf("select %s from %s where %s;", strings.Join(cols, ","), table, strings.Join(where, " and "))
		break
	case UPDATE:
		s = fmt.Sprintf("update %s set %s=? where %s;", table, strings.Join(cols, "=?, "), strings.Join(where, " and "))
		break
	case INSERT:
		s = fmt.Sprintf("insert into %s (%s) values(%s);", table, strings.Join(cols, ","), strings.Join(values, ","))
		return strings.ToLower(s), nil, vals, nil
	case DELETE:
		s = fmt.Sprintf("delete from %s where %s;", table, strings.Join(where, " and "))
		break
	case CREATE:
		break
	}
	return strings.ToLower(s), vals, whereVals, nil
}

//Find find
func (m *MySQL) Find(i interface{}, function string, where string, whereArgs ...interface{}) (*sql.Rows, error) {
	table, columns, err := getColumns(i)
	if err != nil {
		return nil, err
	}
	cols := []string{}
	dest := []interface{}{}
	for _, column := range columns{
		cols = append(cols, column.Name)
		dest = append(dest, column.Ref)
	}
	s := fmt.Sprintf("select %s from %s ", strings.Join(cols, ","), table)
	if where != "" {
		s += fmt.Sprintf("where %s", where)
	}
	s = fmt.Sprintf("%s %s;", s, function)
	rows, err := m.db.Query(strings.ToLower(s), whereArgs...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

//FindOne find
func (m *MySQL) FindOne(i interface{}) error {
	s, dests, vals, err := m.toSQL(i, SELECT)
	if err != nil {
		return err
	}
	return m.db.QueryRow(s, vals...).Scan(dests...)
}

//Update update object
func (m *MySQL) Update(i interface{}) (sql.Result, error) {
	s, dests, vals, err := m.toSQL(i, UPDATE)
	if err != nil {
		return nil, err
	}
	args := []interface{}{}
	for _, d := range dests{
		args = append(args, d)
	}
	for _, v := range vals{
		args = append(args, v)
	}
	return m.Exec(s, args...)
}

//Insert insert an interface
func (m *MySQL) Insert(i interface{}) (sql.Result, error) {
	s, _, vals, err := m.toSQL(i, INSERT)
	if err != nil {
		return nil, err
	}
	return m.Exec(s, vals...)
}

//InsertMulti update object
func (m *MySQL) InsertMulti(i interface{}) (interface{}, error) {
	return nil, nil
}

//Delete delete
func (m *MySQL) Delete(i interface{}) (sql.Result, error) {
	s, _, vals, err := m.toSQL(i, DELETE)
	if err != nil {
		return nil, err
	}
	return m.Exec(s, vals...)
}

//Marshal Marshal
func (m *MySQL) Marshal(i interface{}) (string, error) {
	return "", nil
}

//UnMarshal UnMarshal
func (m *MySQL) UnMarshal(src interface{}, dest interface{}) error {
	_, columns, err := getColumns(dest)
	if err != nil {
		return err
	}
	cols := []string{}
	destCols := []interface{}{}
	for _, column := range columns{
		cols = append(cols, column.Name)
		destCols = append(destCols, column.Ref)

	}
	if row, ok := src.(sql.Row); ok {
		return 	row.Scan(destCols...)
	}
	if rows, ok := src.(sql.Rows); ok {
		defer rows.Close()
		return rows.Scan(destCols...)
	}
	return nil
}

//UnMarshalSlice UnMarshalSlice
func (m *MySQL) UnMarshalSlice(src interface{}, i interface{}) (interface{}, error) {
	t := reflect.TypeOf(i)
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	_, columns, err := getColumns(i)
	if err != nil {
		return nil, err
	}
	tSlice := reflect.SliceOf(t)
	s := reflect.MakeSlice(tSlice, 0, 1)
	
	if row, ok := src.(*sql.Row); ok {
		dest := []interface{}{}
		newI := reflect.New(t)
		for _, column := range columns {
			dest = append(dest, newI.FieldByName(column.FieldName).Addr().Interface())
		}
		err = row.Scan(dest...)
		if err != nil {
			return nil, err
		}
		s = reflect.Append(s, newI)
		return s.Interface(), nil
	}
	//if rows
	if rows, ok := src.(*sql.Rows); ok {
		defer rows.Close()
		for {
			dest := []interface{}{}
			newI := reflect.New(t)
			if newI.Kind() == reflect.Ptr {
				newI = newI.Elem()
			}
			for _, column := range columns {
				dest = append(dest, newI.FieldByName(column.FieldName).Addr().Interface())
			}
			if !rows.Next() {
				return s.Interface(), nil
			}
			err = rows.Scan(dest...)
			if err != nil {
				return s.Interface(), err
			}
			s = reflect.Append(s, newI)
		}
	}
	return s.Interface(), nil
}

//Query Query
func (m *MySQL) Query(sql string, args ...interface{}) (*sql.Rows, error) {
	return m.db.Query(sql, args...)
}

//CreateTable CreateTable
func (m *MySQL) CreateTable(i interface{}) error {
	return nil
}

//Exec exec sql
func (m *MySQL) Exec(query string, args ...interface{}) (sql.Result, error) {
	if m.db == nil {
		return nil, fmt.Errorf("not connect to a db")
	}
	tx, err := m.db.Begin()
	if err != nil {
		return nil, err
	}
	res, err := tx.Exec(query, args...)
	if err != nil {
		return nil, err
	}
	
	return res, tx.Commit()
}

//Count count 
func (m *MySQL)Count(i interface{}, where string, whereVals ...interface{})(int, error){
	t := reflect.TypeOf(i)
	s := fmt.Sprintf("select (1) from %s ", t.Name())
	if where != "" {
		s += fmt.Sprintf("where %s ", where)
	}
	s += ";"
	count := -1
	err := m.db.QueryRow(s, whereVals...).Scan(&count)
	return count, err
}

//Exists exists
func (m *MySQL)Exists(i interface{})error{
	_, columns, err := getColumns(i)
	if err != nil {
		return err
	}
	where := " 1=1 "
	whereVals := []interface{}{}
	for _, col := range columns{
		if col.Primary {
			where += fmt.Sprintf("and %s=?", col.Name)
			whereVals = append(whereVals, col.Ref)
		}
	}
	count, err := m.Count(i, where, whereVals...)
	if err != nil {
		return err
	}
	if count < 1 {
		return fmt.Errorf("not found")
	}
	return nil
}

//UpSert update if exists or insert
func (m *MySQL)UpSert(i interface{})(sql.Result, error){
	if err := m.Exists(i); err != nil {
		return m.Insert(i)
	}
	return m.Update(i)
}

//UpdateVal update some vals
func (m *MySQL)UpdateVal(i interface{}, columns ...string)(sql.Result, error){
	if i == nil || len(columns) == 0 {
		return nil, fmt.Errorf("interface is nil or columns is empty")
	}
	table, cols, err := getColumns(i)
	if err != nil {
		return  nil, err
	}
	vals := []interface{}{}
	for _, c := range cols{
		for _, column := range columns {
			if c.Name == column {
				vals = append(vals, c.Ref)
			}
		}
	}
	sql := fmt.Sprintf("update %s set %s =?;", table, strings.Join(columns, "=? "))
	return m.Exec(sql, vals...)
}