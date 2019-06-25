package category

import (
	"server/router"
)

//RouterMap article api router map
var RouterMap = []*router.Router{
	router.New("/api/v1/categories", router.GET, List),
	router.New("/api/v1/category", router.GET, Get),
	router.New("/api/v1/category", router.POST, Add),
	router.New("/api/v1/category", router.PUT, Update),
	router.New("/api/v1/category", router.DELETE, Delete),
}