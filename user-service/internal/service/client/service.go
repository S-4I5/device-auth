package client

import (
	"user-service/internal/model/entity"
	"user-service/internal/repository"
)

type service struct {
	clientRepository repository.ClientRepository
}

func NewService(clientRepository repository.ClientRepository) *service {
	return &service{
		clientRepository: clientRepository,
	}
}

func (s *service) Create(client entity.Client) error {
	return s.clientRepository.Save(client)
}

func (s *service) Get(id string) (entity.Client, error) {
	return s.clientRepository.Get(id)
}
