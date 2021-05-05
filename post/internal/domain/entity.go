package domain


type Post struct{
	Description string
	UserId string
}




func NewPost(Description string, UserId string) Post {
	return Post{Description: Description, UserId: UserId}
}






