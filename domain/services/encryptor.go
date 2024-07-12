package services

type Encryptor interface {
	Encrypt(in string) (string, error)
	Decrypt(in string) (string, error)
}
