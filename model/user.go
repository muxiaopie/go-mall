package model

type User struct {
	Model
	Username string 	`json:"username"`
	Nickname string 	`json:"nickname"`
	Password string 	`json:"-"`
	Phone    string 	`json:"phone"`
	Admin  uint 	    `json:"is_admin"` // 是否有后台登陆权限
	Email    string 	`json:"email"`
	Status   uint 		`json:"-" gorm:"default:1"`
	Avatar   string 	`json:"avatar"`
	Introduction string `json:"introduction"`
}
