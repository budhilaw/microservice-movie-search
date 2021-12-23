package usecase

import (
	"context"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"microservice-test/domain"
	mock_domain "microservice-test/mocks"
	"testing"
	"time"
)

func Test_Get(t *testing.T) {
	mockOmdb := &domain.Omdb{
		Id:          1,
		Title:       "test-1",
		Year:        "2021",
		ImdbID:      "test-imdb-id",
		ContentType: "movie",
		Poster:      "n/a",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		c := gomock.NewController(t)
		m := mock_domain.NewMockOmdbRepository(c)
		service := omdbUsecase{omdbRepo: m, contextTimeout: 2}

		m.EXPECT().FindByImdbID(mock.MatchedBy(func(ctx context.Context) bool { return true }), mock.Anything).Return(mockOmdb, nil)
		result, err := service.Get(context.TODO(), mock.Anything)

		assert.NoError(t, err)
		assert.NotNil(t, result)
	})

	t.Run("error", func(t *testing.T) {
		c := gomock.NewController(t)
		m := mock_domain.NewMockOmdbRepository(c)
		service := omdbUsecase{omdbRepo: m, contextTimeout: 2}

		m.EXPECT().FindByImdbID(mock.MatchedBy(func(ctx context.Context) bool { return true }), mock.Anything).Return(nil, errors.New("Unexpected"))
		result, err := service.Get(context.TODO(), mock.Anything)

		assert.Error(t, err)
		assert.Nil(t, result)
	})
}

func Test_Save(t *testing.T) {
	mockOmdb := &domain.Omdb{
		Id:          1,
		Title:       "test-1",
		Year:        "2021",
		ImdbID:      "test-imdb-id",
		ContentType: "movie",
		Poster:      "n/a",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	t.Run("success", func(t *testing.T) {
		c := gomock.NewController(t)
		m := mock_domain.NewMockOmdbRepository(c)
		service := omdbUsecase{omdbRepo: m, contextTimeout: 2}

		m.EXPECT().Store(mock.MatchedBy(func(ctx context.Context) bool { return true }), mockOmdb).Return(nil)
		err := service.Save(context.TODO(), mockOmdb)

		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		c := gomock.NewController(t)
		m := mock_domain.NewMockOmdbRepository(c)
		service := omdbUsecase{omdbRepo: m, contextTimeout: 2}

		m.EXPECT().Store(mock.MatchedBy(func(ctx context.Context) bool { return true }), mockOmdb).Return(errors.New("Unexpected"))
		err := service.Save(context.TODO(), mockOmdb)

		assert.Error(t, err)
	})
}
