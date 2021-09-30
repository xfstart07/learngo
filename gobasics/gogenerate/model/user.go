package gogenerate

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) SetAge(age int) {
	u.Age = age
}
