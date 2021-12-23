package usecase

import (
	"context"
	"microservice-test/common"
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

func (u *omdbUsecase) Save(c context.Context, m *domain.Omdb) (err error) {
	ctx, cancel := context.WithTimeout(c, u.contextTimeout)
	defer cancel()

	// get by IMDB ID
	existedData, _ := u.omdbRepo.FindByImdbID(ctx, m.ImdbID)
	if existedData != (domain.Omdb{}) {
		return common.ErrConflict
	}

	// saving
	err = u.omdbRepo.Store(ctx, m)
	return
}
