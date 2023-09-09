package main

import (
	"log"
	"net/http"

	"ex03/internal/api/handlers"
	"ex03/internal/api/logreq"
	"ex03/internal/api/rps"
	"ex03/internal/config"
	"ex03/internal/storage"
	"ex03/internal/storage/inmemory"
)

func main() {
	// init config
	cfg := initConfig()

	// init rpsLimiter
	rps := initRPSLimiter(cfg)

	// init storage
	store := initStorage()

	//init handlers
	initHandlers(store, rps)

	//init server
	initServer(cfg)
}

func initConfig() *config.Config {
	defer log.Println("Config init successful!")
	return config.NewConfig()
}

func initRPSLimiter(cfg *config.Config) *rps.Limiter {
	defer log.Println("Limiter init successful!")
	return rps.NewRPSLimiter(cfg.RPSLimit, cfg.RPSDuration)
}

func initStorage() storage.I {
	defer log.Println("Storage init successful!")
	return inmemory.NewMemoryStorage()
}

func initHandlers(store storage.I, rps *rps.Limiter) {
	// POST
	http.HandleFunc("/set", logreq.LogRequest(rps.RPSLimit(handlers.Set(store))))

	// DELETE
	http.HandleFunc("/delete", logreq.LogRequest(rps.RPSLimit(handlers.Delete(store))))

	log.Println("Handlers init successful!")
}

func initServer(cfg *config.Config) {
	log.Println("Server listening")
	http.ListenAndServe(":"+cfg.Port, nil)
}
