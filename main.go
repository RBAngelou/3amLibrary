package main

import (
	"log"
	"net/http"

	"github.com/RBAngelou/3amLibrary/internal/repositories"
	"github.com/RBAngelou/3amLibrary/internal/services"
	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	// q: How to I check the status of the Memcached server?
	// a: You can check the status of the Memcached server by using the memcache.Stats method. This method returns a map of statistics about the Memcached server, such as the number of items stored, the total memory used, and the number of connections. You can use this information to monitor the health and performance of the Memcached server.
	// q: How can I

	mc := memcache.New("memcache:11211")
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
