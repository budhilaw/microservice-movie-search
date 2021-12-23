package domain

import (
	"context"
	"time"
)

type Omdb struct {
	Id          int64     `json:"id" form:"id"`
	Title       string    `json:"title" form:"title"`
	Year        string    `json:"year" form:"year"`
	ImdbID      string    `json:"imdb_id" form:"imdb_id"`
	ContentType string    `json:"content_type" form:"content_type"`
	Poster      string    `json:"poster" form:"poster"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type OmdbUsecase interface {
	Get(c context.Context, id string) (res *Omdb, err error)
	Save(c context.Context, m *Omdb) (err error)
}

type OmdbRepository interface {
	Store(ctx context.Context, m *Omdb) (err error)
	FindByImdbID(ctx context.Context, imdbId string) (res *Omdb, err error)
}
