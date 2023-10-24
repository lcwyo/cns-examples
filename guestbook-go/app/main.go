package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"github.com/urfave/negroni"
)

var rdb *redis.Client

const (
	redisAddr     = "localhost:6379"
	redisPassword = ""
	redisDB       = 0
	serverAddr    = ":8080"
)

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       redisDB,
	})
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
		// You can choose to exit the application gracefully or take other recovery actions here.
	}

	router := mux.NewRouter()
	router.HandleFunc("/", ShowIndex)
	router.HandleFunc("/lrange/{key}", ListRangeHandler).Methods("GET")
	router.HandleFunc("/lrange", ListAllHandler).Methods("GET")
	router.HandleFunc("/rpush/{key}/{value}", ListPushHandler).Methods("POST")
	router.HandleFunc("/info", InfoHandler).Methods("GET")
	router.HandleFunc("/env", EnvHandler)

	n := negroni.Classic()
	n.UseHandler(router)

	srv := &http.Server{
		Addr:         serverAddr,
		Handler:      n,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	gracefulShutdown(srv)
}
func HandleError(w http.ResponseWriter, err error, context string, statusCode int) bool {
	if err != nil {
		errorMsg := fmt.Sprintf("Error in context '%s': %v", context, err)
		log.Printf("[%s] %s\n", time.Now().Format(time.RFC3339), errorMsg)
		http.Error(w, errorMsg, statusCode)
		return true
	}
	return false
}

func gracefulShutdown(srv *http.Server) {
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)

	<-sigCh

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %v\n", err)
	}

	log.Println("Server stopped gracefully.")
}

// Other handlers and utility functions...
func ShowIndex(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "public/index.html")
}

func ListRangeHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	list, err := rdb.LRange(context.Background(), key, 0, -1).Result()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(list); err != nil {
		handleError(w, err)
	}
}

func handleError(w http.ResponseWriter, err error) {
	panic("unimplemented")
}
func ListAllHandler(w http.ResponseWriter, r *http.Request) {
	iter := rdb.Scan(context.Background(), 0, "prefix:*", 0).Iterator()
	var keys []string
	for {
		if !iter.Next(context.Background()) {
			break
		}
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(strings.Join(keys, "\n")))
}

func ListPushHandler(w http.ResponseWriter, r *http.Request) {
	key := mux.Vars(r)["key"]
	value := mux.Vars(r)["value"]

	if err := rdb.RPush(context.Background(), key, value).Err(); err != nil {
		handleError(w, err)
		return
	}

	w.WriteHeader(http.StatusOK)
	ListRangeHandler(w, r)
}
func InfoHandler(w http.ResponseWriter, r *http.Request) {
	info, err := rdb.Info(context.Background()).Result()
	if err != nil {
		handleError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(environment); err != nil {
		handleError(w, err)
	}
}
