package category

import (
	"server/connections"
	"fmt"
)

var conn = connections.MySQL

//Category article category
type Category struct{
	ID int 				`json:"id" mysql:"id,auto,primary"`
	Name string			`json:"name" mysql:"name"`
	Description string	`json:"description" mysql:"description"`
	Deleted int			`json:"deleted" mysql:"deleted"`
}

//Read read the article content data from database
func (c *Category)Read()error{
	if c.ID < 1 {
		return fmt.Errorf("id error")
	}
	return conn.FindOne(c)
}

//Write write the article content to database
func (c *Category)Write()(int64, error){
	res, err := conn.Insert(c)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()	
}

//Update update the article content to database
func (c *Category)Update()error{
	_, err := conn.Update(c)
	return err
}

//Delete delete the article to database
func (c *Category)Delete()error{
	_, err := conn.Delete(c)
	return err
}

//List list the categories in database
func list(limit int)([]Category, error){
	c := Category{}
	categiries := []Category{}
	limitStr := ""
	if limit > 0 {
		limitStr = fmt.Sprintf("limit %d", limit)
	}
	rows, err := conn.Find(&c, limitStr, "deleted=0")
	if err != nil {
		return categiries, err
	}
	res, err := conn.UnMarshalSlice(rows, &c)
	if err != nil {
		return categiries, err
	}
	if cs, ok := res.([]Category); ok {
		return cs, nil
	}
	return categiries, fmt.Errorf("convert result type error")
}

