package domain

type Fields struct {
	Versions []Version        `json:"versions"`
	Status   Status           `json:"status"`
	Comment  CommentContainer `json:"comment"`
}
