package interfaces

// Bag manages key-value pairs
type Bag interface {
	// Set allows to set value for a proposed key
	Set(key string, value interface{})

	// Remove deletes a specific key from Bag
	Remove(key string)

	// Has helps to check if a key exists
	Has(key string) bool

	// All returns all key-value of bag
	All() map[string]interface{}

	Getter
}
