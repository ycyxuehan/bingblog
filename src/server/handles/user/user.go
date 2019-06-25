package user

//User user
type User struct{
	Username string 		`json:"username" mysql:"username,primary"`
	Password string 		`json:"password" mysql:"password"`
	Enabled int				`json:"enabled" mysql:"enabled"`
	Deleted int				`json:"deleted" mysql:"deleted"`
	AccessLevel int			`json:"accesslevel" mysql:"access_level"`
}

var admin User

func init(){
	admin = User{
		Username: "admin",
		Password: "xuehan123",
		Enabled: 1,
		Deleted: 0,
		AccessLevel: 4,
	}
}