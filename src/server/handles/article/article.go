package article

import (
	"fmt"
	"server/connections"
)

var conn = connections.MySQL

//Article blog article
type Article struct {
	ID int 				`json:"id" mysql:"id,auto,primay"`
	Title string		`json:"title" mysql:"title"`
	Category int		`json:"category" mysql:"category"`
	OutLine	string		`json:"outline" mysql:"outline"`
	Reads int			`json:"views" mysql:"views"`
	Deleted int 		`json:"deleted" mysql:"deleted"`
	Content string		`json:"content"`
}

//Read read the article data from database
func (a *Article)Read()error{
	if a.ID < 1 {
		return fmt.Errorf("id error")
	}
	return conn.FindOne(a)
}

//Write write the article to database
func (a *Article)Write()(int64, error){
	res, err := conn.Insert(a)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()	
}

//Update update the article to database
func (a *Article)Update()error{
	_, err := conn.Update(a)
	return err
}

//Delete delete the article to database
func (a *Article)Delete()error{
	_, err := conn.Delete(a)
	return err
}

//AddReads the article reads add 1
func (a *Article)AddReads()error{
	a.Reads ++
	_, err := conn.UpdateVal(a, "reads")
	return err
}

//list get article list from database
func list(categoryID int, limit int, recommend bool)([]Article, error){
	res := []Article{}
	a := Article{}
	if categoryID == 0 && recommend == false {
		return res, fmt.Errorf("category id error")
	}
	limitStr := ""
	if limit > 0 {
		limitStr = fmt.Sprintf("limit %d", limit)
	}
	if categoryID != 0 {
		rows, err := conn.Find(&a, limitStr, "category=? and deleted = 0", categoryID)
		if err != nil {
			return res, err
		}
		result, err := conn.UnMarshalSlice(rows, &a)
		if err != nil {
			return res, err
		}
		if articles, ok := result.([]Article); ok {
			return articles, nil
		}
	}
	if recommend {
		if limit < 1 {
			limit = 5
		}
		limitStr = fmt.Sprintf("order by views desc limit %d", limit)
		rows, err := conn.Find(&a, limitStr, "deleted = 0")
		if err != nil {
			return res, err
		}
		result, err := conn.UnMarshalSlice(rows, &a)
		if err != nil {
			return res, err
		}
		if articles, ok := result.([]Article); ok {
			return articles, nil
		}
	}
	return res, nil
}

//GetCategoryArticles  get article list from database
func GetCategoryArticles(categoryID int, limit int)([]Article, error){
	return list(categoryID, limit, false)
}

//GetRecommendArticles  get article list from database
func GetRecommendArticles(categoryID int, limit int)([]Article, error){
	return list(0, limit, true)
}

//Content article content
type Content struct {
	// ID int				`json:"id" mysql:"id,auto,primary"`
	Article int			`json:"article" mysql:"article,primary"`
	Text string			`json:"content" mysql:"content"`
}

//Read read the article content data from database
func (c *Content)Read()error{
	if c.Article < 1 {
		return fmt.Errorf("id error")
	}
	return conn.FindOne(c)
}

//Write write the article content to database
func (c *Content)Write()(int64, error){
	res, err := conn.Insert(c)
	if err != nil {
		return 0, err
	}
	return res.LastInsertId()	
}

//Update update the article content to database
func (c *Content)Update()error{
	_, err := conn.Update(c)
	return err
}

