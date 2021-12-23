package http

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	"log"
	"microservice-test/common"
	"microservice-test/domain"
	"microservice-test/omdb/delivery/pb"
	"net/http"
	"strconv"
	"time"
)

type OmdbHandler struct {
	OmdbUsecase domain.OmdbUsecase
}

func NewOmdbHandler(e *mux.Router, ou domain.OmdbUsecase) {
	handler := &OmdbHandler{OmdbUsecase: ou}
	e.HandleFunc("/search", handler.Search).Methods("GET")
}

func isRequestValid(m *domain.Omdb) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

func getStatusCode(err error) int {
	if err == nil {
		return http.StatusOK
	}

	log.Println(err)
	switch err {
	case common.ErrInternalServerError:
		return http.StatusInternalServerError
	case common.ErrNotFound:
		return http.StatusNotFound
	case common.ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

func (h *OmdbHandler) Search(w http.ResponseWriter, request *http.Request) {
	movieName := request.URL.Query().Get("title")
	if movieName == "" {
		log.Println("Error when reading URL Query, err: title not found")
		w.Write([]byte("title cannot be empty"))
		return
	}

	page, err := strconv.Atoi(request.URL.Query().Get("page"))
	if err != nil {
		log.Printf("Error when reading URL Query, err: %v", err)
		page = 1
	}

	addr := fmt.Sprintf("%s:%d", common.Env.GrpcAdress, common.Env.GrpcPort)
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Printf("Error when dial gRPC. err: %v", err)
	}
	defer grpcConn.Close()

	c := pb.NewMovieSearchClient(grpcConn)

	req := &pb.SearchRequest{
		MovieName: movieName,
		Page:      int32(page),
	}

	resp, err := c.SearchMovie(context.Background(), req)
	if err != nil {
		log.Println("error when call grpc. err : ", err.Error())
		w.Write([]byte("error when call grpc. Is GRPC up? See logs for more information or contact admin"))
	} else {
		json.NewEncoder(w).Encode(resp)
		ctx := context.Background()

		for _, movie := range resp.MovieList {
			data := domain.Omdb{
				Title:       movie.Title,
				Year:        movie.Year,
				ImdbID:      movie.ImdbID,
				ContentType: movie.GetType(),
				Poster:      movie.Poster,
				CreatedAt:   time.Now(),
			}

			err = h.OmdbUsecase.Save(ctx, &data)
			if err != nil {
				log.Println("error when creating log. err : ", err.Error())
			}
		}
	}
}
