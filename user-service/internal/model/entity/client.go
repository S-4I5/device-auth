package entity

import (
	"user-service/internal/api/scope"
	"user-service/internal/config"
)

type Client struct {
	Id     string
	Secret string
	Scopes []scope.Scope
}

func ConfigClientToEntity(client config.Client) (Client, error) {
	mappedClient := Client{
		Id:     client.ClientId,
		Secret: client.ClientSecret,
	}

	scopes := make([]scope.Scope, 1)
	for _, stringScope := range client.Scopes {
		curScope, err := scope.ParseScope(stringScope)
		if err != nil {
			return Client{}, err
		}

		scopes = append(scopes, curScope)
	}

	mappedClient.Scopes = scopes

	return mappedClient, nil
}
