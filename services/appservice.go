package services

import (
	"github.com/apcichewicz/gotodolist/models"
	"github.com/apcichewicz/gotodolist/repos"
	"github.com/golang-jwt/jwt/v5"
)

type UserRepo interface {
	GetUserByID(id string) (models.User, error)
	GetUserByEmail(email string) (models.User, error)
	CreateUser(user models.User) (models.User, error)
	UpdateUser(user models.User) (models.User, error)
	DeleteUser(id string) error
	Login(email string, password string) (models.User, *jwt.Token, error)
	Register(user models.User) (models.User, *jwt.Token, error)
}

type ToDoRepo interface {
	GetToDoByID(id string) (models.ToDo, error)
	GetToDoByUserID(userID string) ([]models.ToDo, error)
	CreateToDo(toDo models.ToDo) (models.ToDo, error)
	UpdateToDo(toDo models.ToDo) (models.ToDo, error)
	DeleteToDo(id string) error
}

type AppService struct {
	UserRepo repos.UserRepository
	ToDoRepo repos.ToDoRepository
}

func NewAppService(authRepo repos.AuthRepository, userRepo repos.UserRepository, toDoRepo repos.ToDoRepository) *AppService {
	return &AppService{
		UserRepo: *repos.UserRepository,
		ToDoRepo: *repos.ToDoRepository,
	}
}

func (s *AppService) Login(email string, password string) (models.User, *jwt.Token, error) {

}

func (s *AppService) Register(user models.User) (models.User, *jwt.Token, error) {

}

func (s *AppService) GetUserByID(id string) (models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

func (s *AppService) GetUserByEmail(email string) (models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

func (s *AppService) CreateUser(user models.User) (models.User, error) {
	return s.UserRepo.CreateUser(user)
}

func (s *AppService) UpdateUser(user models.User) (models.User, error) {
	return s.UserRepo.UpdateUser(user)
}

func (s *AppService) DeleteUser(id string) error {
	return s.UserRepo.DeleteUser(id)
}

func (s *AppService) GetToDoByID(id string) (models.ToDo, error) {
	return s.ToDoRepo.GetToDoByID(id)
}

func (s *AppService) GetToDoByUserID(userID string) ([]models.ToDo, error) {
	return s.ToDoRepo.GetToDoByUserID(userID)
}

func (s *AppService) CreateToDo(toDo models.ToDo) (models.ToDo, error) {
	return s.ToDoRepo.CreateToDo(toDo)
}

func (s *AppService) UpdateToDo(toDo models.ToDo) (models.ToDo, error) {
	return s.ToDoRepo.UpdateToDo(toDo)
}

func (s *AppService) DeleteToDo(id string) error {
	return s.ToDoRepo.DeleteToDo(id)
}
