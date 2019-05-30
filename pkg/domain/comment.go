package domain

type CommentContainer struct {
	Total    int       `json:"total"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Body string `json:"body"`
}
