package client

import (
	"fmt"
	"sync"
	"user-service/internal/model/entity"
)

type repository struct {
	clients map[string]entity.Client
	mutex   sync.RWMutex
}

func NewRepository() *repository {
	return &repository{
		clients: make(map[string]entity.Client),
	}
}

func (r *repository) Get(id string) (entity.Client, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	client, found := r.clients[id]
	if !found {
		return entity.Client{}, fmt.Errorf("client with id %s does not exist", client.Id)
	}

	return client, nil
}

func (r *repository) Save(client entity.Client) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	_, found := r.clients[client.Id]
	if found {
		return fmt.Errorf("client with id %s already exist", client.Id)
	}

	r.clients[client.Id] = client

	return nil
}
