package main


type AGE interface {
	Age()
}



type User struct {
	Name string
	age int
}

func (u *User) Age(age int) int {
	u.age = age
	return age + 5
}






