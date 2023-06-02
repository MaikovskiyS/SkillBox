package model

type User struct {
	Name string `json:"name"`
	Age  uint64 `json:"age"`
	//Friends []User
}
