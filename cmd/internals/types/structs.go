package types

type User struct{
    ID        int     `json:"id"`
    Username  string  `json:"name"`
}

type Post struct{
	ID              int      `json:"id"`
    Author 		    string   `json:"author"`
	Content         string   `json:"content"`
	Likes_count     int      `json:"likes_count"`
	Comments_count  int      `json:"comments_count"`
}