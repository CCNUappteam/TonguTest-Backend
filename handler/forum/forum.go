package forum

type PostRequest struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type CommentRequest struct {
	PostId  uint   `json:"post_id"`
	Content string `json:"content"`
}
