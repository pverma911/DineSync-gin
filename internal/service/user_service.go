package service

import "myapp/internal/repository"

func GetAllUsers() []string {
    return repository.FetchUsers()
}
