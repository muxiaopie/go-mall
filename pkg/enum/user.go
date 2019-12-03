package enum

// 登陆的字段
const (
	ID        = iota
	USERNAME
	EMAIL
	PHONE
)

//
var FieldMap = map[int]string{
	ID:		  "id",
	USERNAME: "username",
	EMAIL:	  "email",
	PHONE:    "phone",
}
