package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/negroni"
)

var (
	rdb *redis.Client
)

func ShowIndex(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-Type", "application/text")
	//w.Header().Set("X-Content-Type-Options", "nosniff")
	http.ServeFile(w, r, "public/index.html")
}

func ListRangeHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	list := rdb.LRange(context.Background(), key, 0, -1)

	var members []string
	err := list.ScanSlice(&members)
	if err != nil {
		HandleError(nil, err)
	}

	membersJSON, err := json.MarshalIndent(members, "", "  ")
	if err != nil {
		HandleError(nil, err)
	}

	w.Write(membersJSON)
}
func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	iter := rdb.Scan(context.Background(), 0, "prefix:*", 0).Iterator()
	var keys []string
	for iter.Next(context.Background()) {
		keys = append(keys, iter.Val())
	}
	if err := iter.Err(); err != nil {
		HandleError(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(strings.Join(keys, "\n")))
	}
}

func ListPushHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value := mux.Vars(r)["value"]
	list := rdb.RPush(context.Background(), key, value)

	err := list.Err()
	if err != nil {
		HandleError(w, err)
	} else {
		w.WriteHeader(http.StatusOK)
		ListRangeHandler(w, r)
	}
}

func InfoHandler(w http.ResponseWriter, r *http.Request) {
	info, err := rdb.Info(context.Background()).Result()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//w.Header().Set("Content-Type", "application/html")
	w.Write([]byte(info))
}

func EnvHandler(w http.ResponseWriter, r *http.Request) {
	environment := make(map[string]string)

	for _, item := range os.Environ() {
		splits := strings.Split(item, "=")
		key := splits[0]
		val := strings.Join(splits[1:], "=")
		environment[key] = val
	}

	envJSON, err := json.MarshalIndent(environment, "", "  ")
	if err != nil {
		HandleError(nil, err)
	}

	w.Write(envJSON)
}

func HandleError(result interface{}, err error) (r interface{}) {
	if err != nil {
		log.Println(err)
		return nil
	}
	return result
}

func main() {
	ctx := context.Background()
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	iter := rdb.Scan(ctx, 0, "prefix:*", 0).Iterator()
	for iter.Next(ctx) {
		fmt.Println("keys", iter.Val())
	}
	if err := iter.Err(); err != nil {
		log.Println(err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ShowIndex)
	router.HandleFunc("/lrange/{key}", func(w http.ResponseWriter, r *http.Request) {
		ListRangeHandler(w, r)
	}).Methods("GET")
	router.HandleFunc("/lrange", ListAllHandler).Methods("GET")
	router.HandleFunc("/rpush/{key}/{value}", func(w http.ResponseWriter, r *http.Request) {
		ListPushHandler(w, r)
	}).Methods("POST")
	router.HandleFunc("/info", func(w http.ResponseWriter, r *http.Request) {
		InfoHandler(w, r)
	}).Methods("GET")
	router.HandleFunc("/env", EnvHandler)

	n := negroni.Classic()
	n.Use(negroni.NewLogger())
	n.UseHandler(router)

	srv := &http.Server{
		Addr:         ":8080",
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
