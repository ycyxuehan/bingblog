package connections

import (
	"libs/mysql"
)

//MySQL my sql client
var MySQL *mysql.MySQL
func init(){
	MySQL = &mysql.MySQL{}
}