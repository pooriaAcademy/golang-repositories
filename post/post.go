package post


import (
	"github.com/pooriaAcademy/golang-repositories/post/internal/domain"
	"github.com/pooriaAcademy/golang-repositories/user"
)



func CreatePost(Description string, User user.User) domain.Post {
	return domain.NewPost(Description, User.Id)
}














