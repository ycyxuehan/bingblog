package user

import (
	"server/router"
)


//RouterMap article api router map
var RouterMap = []*router.Router{
	router.New("/api/v1/login", router.POST, Login),
}