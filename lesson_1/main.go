package main

import "fmt"

type User struct {
	Name   string
	Age    int
	Gender string
}

func (u User) Print() {
	fmt.Println("Name", u.Name)
	fmt.Println("Age", u.Age)
	fmt.Println("Gender", u.Gender)
}

func main() {
	user1 := User{
		Name:   "Egor",
		Age:    27,
		Gender: "man",
	}
	user1.Print()
	user1.SatAge(3)
	user1.PrintAge()
}
func (u User) PrintAge() {
	fmt.Println("Age", u.Age)
}

func (u *User) SatAge(age int) {
	u.Age = age
}
