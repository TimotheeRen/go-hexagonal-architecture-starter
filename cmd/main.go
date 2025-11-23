package main

import (
	"16/internal/adaptaters/secondary"
	"16/internal/infrastructure"
	"fmt"
)

func main() {
	db, err := infrastructure.Migrate()
	if err != nil {
		panic(err)
	}
	var UserRepository secondary.UserRepositoryImpl
	//UserRepository.CreateUser(db, "AnonymeUSB")
	user, err := UserRepository.GetUser(db, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
