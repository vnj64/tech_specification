package services

type Config interface {
	PostgresqlHost() string
	PostgresqlPort() string
	PostgresqlUser() string
	PostgresqlPassword() string
	PostgresqlDatabase() string
	ServerUrl() string
	ServerProtocol() string
	ServerPort() string
}
