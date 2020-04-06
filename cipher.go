package interfaces

// Cipherable is an interface which allows to encypt and decrypt data
type Cipherable interface {
	Encryptable
	Decryptable
}

// Encryptable is an interface which is used to encrypt data
type Encryptable interface {
	Encrypt(text string, key string) (value string, err error)
}

// Decryptable is an interface which is used to decrypt data
type Decryptable interface {
	Decrypt(text string, key string) (value string, err error)
}
