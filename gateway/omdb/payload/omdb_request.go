package payload

type CreateOMDBLog struct {
	Title       string `json:"title" binding:"required"`
	Year        string `json:"year" binding:"required"`
	ImdbID      string `json:"imdb_id" binding:"required"`
	ContentType string `json:"content_type" binding:"required"`
	Poster      string `json:"poster" binding:"required"`
}
