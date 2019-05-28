package interfaces

type Cipherable interface {
	Encryptable
	Decryptable
}

type Encryptable interface {
	encrypt(text string, key string) (value string, err error)
}

type Decryptable interface {
	decrypt(text string, key string) (value string, err error)
}