package repository

import (
	"context"
	"database/sql"
	"log"
	"microservice-test/common"
	"microservice-test/domain"
	"time"
)

type mysqlOmdbRepository struct {
	Conn *sql.DB
}

func NewMysqlOmdbRepository(Conn *sql.DB) domain.OmdbRepository {
	return &mysqlOmdbRepository{Conn}
}

func (r *mysqlOmdbRepository) fetch(ctx context.Context, query string, args ...interface{}) (res []domain.Omdb, err error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		log.Printf("Error when querying context, err: %v", err)
		return nil, err
	}

	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Printf("Error when closing rows, err: %v", err)
		}
	}()

	res = make([]domain.Omdb, 0)
	for rows.Next() {
		t := domain.Omdb{}
		err = rows.Scan(
			&t.Id,
			&t.Title,
			&t.Year,
			&t.ImdbID,
			&t.ContentType,
			&t.Poster,
		)

		if err != nil {
			log.Printf("Error when try to scanning rows, err: %v", err)
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (r *mysqlOmdbRepository) Store(ctx context.Context, m *domain.Omdb) (err error) {
	m.CreatedAt = time.Now()
	query := `INSERT log SET title=?, year=?, imdb_id=?, content_type=?, poster=?, created_at=?`
	stmt, err := r.Conn.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error when prepare context, err: %v", err)
		return nil
	}

	res, err := stmt.ExecContext(ctx, m.Title, m.Year, m.ImdbID, m.ContentType, m.Poster, m.CreatedAt)
	if err != nil {
		log.Printf("Error when exec context, err: %v", err)
		return nil
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		log.Printf("Error when get last inserted id, err: %v", err)
		return nil
	}
	m.Id = lastID
	return
}

func (r *mysqlOmdbRepository) FindByImdbID(ctx context.Context, imdbId string) (res domain.Omdb, err error) {
	query := `SELECT id, title, year, imdb_id, content_type, poster FROM log WHERE imdb_id = ?`
	list, err := r.fetch(ctx, query, imdbId)
	if err != nil {
		log.Printf("Error when running fetch func, err: %v", err)
		return
	}

	if len(list) > 0 {
		res = list[0]
	} else {
		return res, common.ErrNotFound
	}
	return
}
