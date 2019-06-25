package article

import (
	"server/router"
)


//RouterMap article api router map
var RouterMap = []*router.Router{
	router.New("/api/v1/articles", router.GET, List),
	router.New("/api/v1/article", router.GET, Get),
	router.New("/api/v1/article", router.POST, Add),
	router.New("/api/v1/article", router.PUT, Update),
	router.New("/api/v1/article", router.DELETE, Delete),
}