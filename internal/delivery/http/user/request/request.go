package userrequest

type User struct {
	Name     string `json:"name"`
	NIK      string `json:"nik"`
	RoleID   int    `json:"role_id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Login struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
