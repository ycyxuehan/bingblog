package article

import (
	"server/net/responese"
	"strconv"
	"github.com/gin-gonic/gin"
)

//ResponseData respones data
type ResponseData struct {
	ID int					`json:"id"`
	Name string				`json:"name"`
	Content string			`json:"content"`
	Category int			`json:"category"`
}

//Get [GET] get article
func Get(c *gin.Context){
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	a := Article{ID:id}
	err = a.Read()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	ac := Content{Article:a.ID}
	err = ac.Read()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	a.Content = ac.Text
	a.Reads++
	a.Update()
	responese.Success(c, "successed", &a)
}

//List [GET] get articles
func List(c *gin.Context){
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}
	categoryStr :=  c.Query("category")
	recommendStr := c.Query("recommend")
	categoryID, err := strconv.Atoi(categoryStr)
	if err != nil && categoryStr == "false" {
		responese.Error(c, err, nil)
		return
	}
	res, err := list(categoryID, limit, recommendStr == "true")
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "successed", res)
}

//Add [POST] add article
func Add(c *gin.Context){
	a := Article{}
	err := c.BindJSON(&a)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	n, err := a.Write()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	ac := Content{Article:int(n), Text: a.Content}
	_, err = ac.Write()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, n, nil)
}

//Update [PUT] update article
func Update(c *gin.Context){
	a := Article{}
	err := c.BindJSON(&a)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	err = a.Update()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	ac := Content{Article:a.ID, Text:a.Content}
	err = ac.Update()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "updated", nil)
}

//Delete [DELETE] delete article
func Delete(c *gin.Context){
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	a := Article{ID:id}
	err = a.Delete()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "deleted", nil)
}
