package config

import (
	"errors"
	"fmt"
	"os"
	"tech/domain/services"
)

type service struct {
	postgresqlHost     string
	postgresqlPort     string
	postgresqlUser     string
	postgresqlPassword string
	postgresqlDatabase string
	serverUrl          string
	serverProtocol     string
	serverPort         string
}

func Make(encryptor services.Encryptor) (services.Config, error) {
	postgresqlHost := os.Getenv("POSTGRESQL_HOST")
	postgresqlPort := os.Getenv("POSTGRESQL_PORT")
	rawPostgresqlPassword := os.Getenv("POSTGRESQL_PASSWORD")
	rawPostgresqlUser := os.Getenv("POSTGRESQL_USER")
	postgresqlDatabase := os.Getenv("POSTGRESQL_DATABASE")
	serverUrl := os.Getenv("SERVER_URL")
	serverProtocol := os.Getenv("SERVER_PROTOCOL")
	serverPort := os.Getenv("SERVER_PORT")

	if postgresqlHost == "" {
		return nil, errors.New("env POSTGRESQL_HOST is empty")
	}

	if postgresqlPort == "" {
		return nil, errors.New("env POSTGRESQL_PORT is empty")
	}

	if rawPostgresqlPassword == "" {
		return nil, errors.New("env POSTGRESQL_PASSWORD is empty")
	}

	if rawPostgresqlUser == "" {
		return nil, errors.New("env POSTGRESQL_USER is empty")
	}

	if postgresqlDatabase == "" {
		return nil, errors.New("env POSTGRESQL_DATABASE is empty")
	}

	if serverUrl == "" {
		return nil, errors.New("env SERVER_URL is empty")
	}

	if serverProtocol == "" {
		return nil, errors.New("env SERVER_PROTOCOL is empty")
	}

	if serverPort == "" {
		return nil, errors.New("env SERVER_PORT is empty")
	}

	postgresqlUser, err := encryptor.Decrypt(rawPostgresqlUser)
	if err != nil {
		return nil, fmt.Errorf("env POSTGRESQL_USER invalid due [%s]", err)
	}

	postgresqlPassword, err := encryptor.Decrypt(rawPostgresqlPassword)
	if err != nil {
		return nil, fmt.Errorf("env POSTGRESQL_PASSWORD invalid due [%s]", err)
	}

	return &service{
		postgresqlHost:     postgresqlHost,
		postgresqlPort:     postgresqlPort,
		postgresqlUser:     postgresqlUser,
		postgresqlPassword: postgresqlPassword,
		postgresqlDatabase: postgresqlDatabase,
		serverUrl:          serverUrl,
		serverProtocol:     serverProtocol,
		serverPort:         serverPort,
	}, nil
}

func (s *service) PostgresqlHost() string {
	return s.postgresqlHost
}

func (s *service) PostgresqlPort() string {
	return s.postgresqlPort
}

func (s *service) PostgresqlUser() string {
	return s.postgresqlUser
}

func (s *service) PostgresqlPassword() string {
	return s.postgresqlPassword
}

func (s *service) PostgresqlDatabase() string {
	return s.postgresqlDatabase
}

func (s *service) ServerUrl() string {
	return s.serverUrl
}

func (s *service) ServerProtocol() string {
	return s.serverProtocol
}

func (s *service) ServerPort() string {
	return s.serverPort
}
