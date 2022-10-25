package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	postgres "learninggolang.com/learning-golang/adapter/postgre"
	"learninggolang.com/learning-golang/di"
	"log"
	"net/http"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	ctx := context.Background()
	conn := postgres.GetConnection(ctx)
	defer conn.Close()

	postgres.RunMigrations()
	productService := di.ConfigProductDI(conn)

	router := mux.NewRouter()
	router.Handle("/skill", http.HandlerFunc(productService.Create)).Methods("POST")

	router.Handle("/skill", http.HandlerFunc(productService.Fetch)).Queries(
		"page", "{page}",
		"itemsPerPage", "{itemsPerPage}",
		"descending", "{descending}",
		"sort", "{sort}",
		"search", "{search}",
	).Methods("GET")

	port := viper.GetString("server.port")
	log.Printf("LISTEN ON PORT: %v", port)
	http.ListenAndServe(fmt.Sprintf(":%v", port), router)
}
