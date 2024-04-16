package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/RBAngelou/3amLibrary/internal/repositories"
	"github.com/RBAngelou/3amLibrary/internal/services"
	"github.com/bradfitz/gomemcache/memcache"
)

type Config struct {
	Server string `json:"server"`
	Port   string `json:"port"`
}

func main() {
	// Read configuration from file (replace "config.json" with your actual file path)
	configFile, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Fatal(fmt.Errorf("failed to read config file: %w", err))
	}

	// Parse configuration JSON
	var config Config
	if err := json.Unmarshal(configFile, &config); err != nil {
		log.Fatal(fmt.Errorf("failed to unmarshal config: %w", err))
	}

	mc := memcache.New(fmt.Sprintf("%s:%s", config.Server, config.Port))
	if mc == nil {
		log.Fatal("Failed to create memcache client")
	}

	// Initialize a UserRepository implementation (e.g., using a database)
	userRepository := repositories.NewUserRepository(mc)

	// Create a new instance of UserService
	userService := services.NewUserService(userRepository)

	// Register HTTP handlers
	http.HandleFunc("/register", userService.RegisterHandler)
	http.HandleFunc("/getuser", userService.GetUserByEmailHandler)

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":8080", nil))
}
