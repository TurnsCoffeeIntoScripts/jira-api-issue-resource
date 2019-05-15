package domain

type Fields struct {
	Versions []Version        `json:"verions"`
	Status   Status           `json:"status"`
	Comment  CommentContainer `json:"comment"`
}
