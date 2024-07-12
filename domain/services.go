package domain

import "tech/domain/services"

type Services interface {
	Config() services.Config
	Encryptor() services.Encryptor
}
