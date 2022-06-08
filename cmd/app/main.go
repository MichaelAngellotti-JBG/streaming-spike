package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/mca312/jackbox/server/app"
	"gopkg.in/yaml.v3"
)

type Config struct {
	Addr  string   `yaml:"address"`
	Port  int      `yaml:"port"`
	Users []string `yaml:"users"`
}

// Manager for our dependencies
// we'd put in our logging, tracing, token management or other dependencies we'd need across packages
type Manager struct {
	ctx context.Context
	app.UserServiceApi
}

func NewManager(ctx context.Context) *Manager {
	return &Manager{
		ctx,
		app.NewUserService(ctx),
	}
}

type Server struct {
	*http.Server
}

func main() {
	// load config from yaml file
	cfg := loadConfig()
	ctx := context.Background()

	server, err := setupServer(ctx, cfg)
	if err != nil {
		fmt.Printf("Error setting up server. %+v", err)
	}

	go func() {
		fmt.Printf("Starting server: %s\n", server.Addr)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			panic(err)
		}
	}()

	// properly shutdown server
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}

// setupServer manages the routes and init data
func setupServer(ctx context.Context, cfg *Config) (*Server, error) {
	manager := NewManager(ctx)

	router := setupRouter(manager)

	srv := Server{
		&http.Server{
			Addr:    fmt.Sprintf("%s:%d", cfg.Addr, cfg.Port),
			Handler: router,
		},
	}

	return &srv, nil
}

// setupRouter defines our routes
func setupRouter(manager *Manager) http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		if err := json.NewEncoder(w).Encode(map[string]bool{"ok": true}); err != nil {
			fmt.Printf("Handler: /healtcheck: %v", err)
		}
	})

	r.HandleFunc("/verify", manager.VerifyHandler).
		Methods("POST")

	return r
}

func loadConfig() *Config {
	f, err := ioutil.ReadFile("config.yml")
	if err != nil {
		panic(err)
	}

	var config Config
	if err = yaml.Unmarshal(f, &config); err != nil {
		panic(err)
	}

	return &config
}
