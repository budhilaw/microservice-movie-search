package usecase

import (
	"context"
	"microservice-test/domain"
	"time"
)

type omdbUsecase struct {
	omdbRepo       domain.OmdbRepository
	contextTimeout time.Duration
}

func NewOmdbUsecase(or domain.OmdbRepository, timeout time.Duration) domain.OmdbUsecase {
	return &omdbUsecase{omdbRepo: or, contextTimeout: timeout}
}

func (u *omdbUsecase) Get(c context.Context, id string) (m *domain.Omdb, err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// get by IMDB ID
	existedData, err := u.omdbRepo.FindByImdbID(ctx, id)
	if err != nil {
		return nil, err
	}

	return existedData, nil
}

func (u *omdbUsecase) Save(c context.Context, m *domain.Omdb) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// saving
	err = u.omdbRepo.Store(ctx, m)
	return
}
