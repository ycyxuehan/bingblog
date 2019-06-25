package router

import (
	"github.com/gin-gonic/gin"

)
//Method http method
type Method int
const (
	//GET get
	GET Method = iota
	//POST post
	POST 
	//PUT put
	PUT
	//DELETE delete
	DELETE 
)
//Router gin router 
type Router struct {
	Path string //path
	Method Method //method
	Handler gin.HandlerFunc
}

//New new a router
func New(path string, method Method, h gin.HandlerFunc)*Router{
	return &Router{
		Path:path,
		Method: method,
		Handler: h,
	}
}

//Regist regist router to engine
func (r *Router)Regist(e *gin.Engine){
	switch r.Method {
	case GET:
		e.GET(r.Path, r.Handler)
		break
	case POST:
		e.POST(r.Path, r.Handler)
		break
	case PUT:
		e.PUT(r.Path, r.Handler)
		break
	case DELETE:
		e.DELETE(r.Path, r.Handler)
		break
	}
}

//RegistRoute regist route to engine
func RegistRoute(e *gin.Engine, routers ...*Router){
	for _, r := range routers{
		r.Regist(e)
	}
}