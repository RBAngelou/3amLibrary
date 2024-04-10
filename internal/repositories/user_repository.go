package repositories

import (
	"encoding/json"
	"log"

	"strconv"

	"github.com/RBAngelou/3amLibrary/internal/models"
	"github.com/bradfitz/gomemcache/memcache" // Add the import statement for the User type
	// Add the import statement for the cache package
)

// CreateUser implements UserRepository.

func (u *userRepo) CreateUser(user models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	err = u.usercache.Set(&memcache.Item{Key: user.Email, Value: userBytes})
	return err
}

// DeleteUser implements UserRepository.
func (u *userRepo) DeleteUser(userID int) error {
	u.usercache.Delete(strconv.Itoa(userID))

	return nil
}

// GetUserByEmail implements UserRepository.
func (u *userRepo) GetUserByEmail(email string) (*models.User, error) {
	retreivedUser, err := u.usercache.Get(email)
	if err != nil {
		return nil, err
	}
	//convert retrieved user to User struct
	var user models.User
	err = json.Unmarshal(retreivedUser.Value, &user)
	return &user, err
}

// GetUserByID implements UserRepository.
func (u *userRepo) GetUserByID(userID int) (*models.User, error) {
	retreivedUser, err := u.usercache.Get(strconv.Itoa(userID))
	if err != nil {
		return nil, err
	}
	//convert retrieved user to User struct
	var user models.User
	err = json.Unmarshal(retreivedUser.Value, &user)
	return &user, nil
}

// UpdateUser implements UserRepository.
func (u *userRepo) UpdateUser(user models.User) error {
	userBytes, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	u.usercache.Set(&memcache.Item{Key: strconv.Itoa(user.ID), Value: userBytes})
	u.usercache.Set(&memcache.Item{Key: user.Email, Value: userBytes})
	return nil
}

type userRepo struct {
	usercache *memcache.Client
}

type UserRepository interface {
	CreateUser(user models.User) error                 // Use the fully qualified User type
	GetUserByID(userID int) (*models.User, error)      // Use the fully qualified User type
	GetUserByEmail(email string) (*models.User, error) // Use the fully qualified User type
	UpdateUser(user models.User) error                 // Use the fully qualified User type
	DeleteUser(userID int) error
}

// q: please teach me how to add a parameter for the constructor of the UserRepository interface
// a: Add a parameter to the NewUserRepository function that accepts a memcache.Client instance and use it to create a new userRepo instance.
func NewUserRepository(cache *memcache.Client) UserRepository {
	return &userRepo{
		usercache: cache,
	}
}
