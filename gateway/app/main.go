package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/qustavo/dotsql"
	"log"
	"microservice-test/common"
	_omdbHandler "microservice-test/omdb/delivery/http"
	_omdbRepo "microservice-test/omdb/repository"
	_omdbUsecase "microservice-test/omdb/usecase"
	"net/http"
	"net/url"
	"time"
)

func init() {
	common.LoadEnv()
}

func main() {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", common.Env.DbUser, common.Env.DbPassword, common.Env.DbHost, common.Env.DbPort, common.Env.DbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		log.Fatalf("Error when create sql connection, err: %v", err)
	}

	err = doMigrate(dbConn)
	if err != nil {
		log.Fatal(err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatalf("Error when trying to ping database, err: %v", err)
	}

	defer func() {
		err := dbConn.Close()
		if err != nil {
			log.Fatalf("Error when trying to close db connection, err: %v", err)
		}
	}()

	r := mux.NewRouter()

	//e := echo.New()
	//mid := middleware.InitMiddleware()
	//e.Use(mid.CORS)

	timeoutContext := time.Duration(common.Env.ContextTimeout) * time.Second
	omdbRepository := _omdbRepo.NewMysqlOmdbRepository(dbConn)
	omdbUsecase := _omdbUsecase.NewOmdbUsecase(omdbRepository, timeoutContext)
	_omdbHandler.NewOmdbHandler(r, omdbUsecase)

	addr := fmt.Sprintf(":%d", common.Env.AppPort)
	log.Println("[+] Gateway running on ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func doMigrate(conn *sql.DB) (err error) {
	// migrate from file
	dot, err := dotsql.LoadFromFile("./migrations/sql/0-create-log-table.sql")
	if err != nil {
		message := fmt.Sprintf("Error when loading sql migration, err: %v", err)
		return errors.New(message)
	}

	// last, create new table
	_, err = dot.Exec(conn, "create-log-table")
	if err != nil {
		message := fmt.Sprintf("Error when create table from sql migration, err: %v", err)
		return errors.New(message)
	}

	return nil
}
