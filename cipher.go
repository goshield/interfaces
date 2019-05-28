package interfaces

type Cipherable interface {
	Encryptable
	Decryptable
}

type Encryptable interface {
	Encrypt(text string, key string) (value string, err error)
}

type Decryptable interface {
	Decrypt(text string, key string) (value string, err error)
}