package responese

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	//SUCCESS success
	SUCCESS = iota
	//WARNING warning
	WARNING
	//ERROR error
	ERROR
)

//Ajax response ajax format data
func Ajax(c *gin.Context, code int, message interface{}, data interface{}){
	c.JSON(http.StatusOK, gin.H{"Code":code, "Message":fmt.Sprintf("%v", message), "Data":data})
}

//Success success responese
func Success(c *gin.Context, msg interface{}, data interface{}){
	Ajax(c, SUCCESS, msg, data)
}

//Warning warning responese
func Warning(c *gin.Context, msg interface{}, data interface{}){
	Ajax(c, WARNING, msg, data)
}

//Error error responese
func Error(c *gin.Context, msg interface{}, data interface{}){
	Ajax(c, ERROR, msg, data)
}
