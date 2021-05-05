package main

import (
	"github.com/pooriaAcademy/golang-repositories/post"
	"github.com/pooriaAcademy/golang-repositories/user"
	"log"
)

func main() {

	log.Println(post.CreatePost("This is a sample post", user.ExampleUser))
	log.Println(user.ExampleUser)

}








