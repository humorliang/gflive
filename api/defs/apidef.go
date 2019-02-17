package defs

//requests
type UserCredential struct {
	UserName string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

//data model
type VideoInfo struct {
	Id string
	AuthorId int
	Name string
	DisplayCtime string
}
