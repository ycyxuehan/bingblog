package category

import (
	"server/net/responese"
	"strconv"
	"github.com/gin-gonic/gin"
)

//Get [GET] get the category
func Get(c *gin.Context){
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	ca := Category{ID:id}
	err = ca.Read()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "successed", &ca)
}

//List [GET] get the categories
func List(c *gin.Context){
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 0
	}
	res, err := list(limit)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "successed", res)
}

//Add [POST] add the category
func Add(c *gin.Context){
	ca := Category{}
	err := c.BindJSON(&ca)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	n, err := ca.Write()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, n, nil)
}

//Update [PUT] update the category
func Update(c *gin.Context){
	ca := Category{}
	err := c.BindJSON(&ca)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	err = ca.Update()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "updated", nil)
}

//Delete [DELETE] delete the category
func Delete(c *gin.Context){
	idStr := c.Query("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	ca := Category{ID:id}
	err = ca.Delete()
	if err != nil {
		responese.Error(c, err, nil)
		return
	}
	responese.Success(c, "deleted", nil)
}
