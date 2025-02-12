package main

import (
	"errors"
	"fmt"
	"strings"
)

type User struct {
	Name  string
	Age   int
	Email string
}

type UserBuilder struct {
	user User
}

func newUserBuilder() *UserBuilder {
	return &UserBuilder{}
}

func (b *UserBuilder) SetName(name string) *UserBuilder {
	b.user.Name = name
	return b
}

func (b *UserBuilder) SetAge(age int) *UserBuilder {
	b.user.Age = age
	return b
}

func (b *UserBuilder) SetEmail(email string) *UserBuilder {
	b.user.Email = email
	return b
}

func (b *UserBuilder) Build() (User, error) {
	if b.user.Name == "" {
		return User{}, errors.New("name is required")
	}

	if b.user.Age < 0 || b.user.Age >= 130 {
		return User{}, errors.New("invalid age")
	}

	if b.user.Email == "" || !strings.Contains(b.user.Email, "@") {
		return User{}, errors.New("invalid Email")
	}

	return b.user, nil
}

func main() {
	user, err := newUserBuilder().
		SetName("Eduardo").
		SetAge(30).
		SetEmail("test@mail.com").Build()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("%+v\n", user)
}
