package main

import (
	"tech/domain"
	"tech/domain/services"
)

type ctx struct {
	services   domain.Services
	session    domain.Session
	connection domain.Connection
}

type session struct {
	userId int
	isAuth bool
}

type svs struct {
	encryptor services.Encryptor
	config    services.Config
}

func (s *svs) Config() services.Config {
	return s.config
}

func (s *svs) Encryptor() services.Encryptor {
	return s.encryptor
}

func (s *session) Authorized() bool {
	return s.userId > 0
}

func (s *session) UserId() int {
	return s.userId
}

func (c *ctx) Services() domain.Services {
	return c.services
}

func (c *ctx) Connection() domain.Connection {
	return c.connection
}

func (c *ctx) Session() domain.Session {
	return c.session
}

func (c *ctx) MakeWithSession(s domain.Session) domain.Context {
	return &ctx{
		services:   c.services,
		session:    s,
		connection: c.connection,
	}
}
