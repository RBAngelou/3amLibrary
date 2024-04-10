package services

import (
	"encoding/json"
	"net/http" // Import the package that contains the UserRepository type

	"github.com/RBAngelou/3amLibrary/internal/models" // Import the package that contains the User type
	"github.com/RBAngelou/3amLibrary/internal/repositories"
)

type UserService struct {
	userRepository repositories.UserRepository // Use the fully qualified type name
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{
		userRepository: repo,
	}
}

func (h *UserService) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse request body to extract registration data
	var newUser models.User // Use the fully qualified type name
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.userRepository.CreateUser(newUser)
	if err != nil {
		http.Error(w, "Error registering user"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *UserService) GetUserByEmailHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	if email == "" {
		http.Error(w, "Email parameter is required", http.StatusBadRequest)
		return
	}

	user, err := h.userRepository.GetUserByEmail(email)
	if err != nil {
		http.Error(w, "User not found"+err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(user)
}
