package loginandregister

type User struct {
	Username string `form:"username"`
	Password string `form:"password"`
}
