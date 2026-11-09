package types

type User struct {
	ID       int    `json:"id"`
	Username string `json:"name"`
}

type Post struct {
	ID            int    `json:"id"`
	Author        int    `json:"author"`
	Content       string `json:"content"`
	LikesCount    int    `json:"likes_count"`
	CommentsCount int    `json:"comments_count"`
}

type Comment struct {
	ID     int    `json:"id"`
	UserId int    `json:"user_id"`
	PostId int    `json:"post_id"`
	Body   string `json:"body"`
}

type Like struct {
	UserId int `json:"user_id"`
	PostId int `json:"post_id"`
}

type BadRequest struct {
	Status int
	Msg    string
}

var _ error = *new(BadRequest)


// Error implements [error].
func (br BadRequest) Error() string {
	return br.Msg
}
