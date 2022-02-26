package models

type User struct {
	Id   int    `json:"id,omitempty" sql:"ID"`
	Name string `json:"name" sql:"NAME"`
	Age  int    `json:"age" sql:"AGE"`
}
