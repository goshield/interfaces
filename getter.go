package interfaces

// Getter is a simple interface which allow to retrieve
// value by key.
//
// Getter would support to return most possible value's types
type Getter interface {
	DefaultGetter
	StringGetter
	FloatGetter
	BoolGetter
	UIntGetter
	IntGetter
}

// DefaultGetter is an interface which allows to set default value
// if there is not value for a particular key
type DefaultGetter interface {
	Get(key string) interface{}
	GetOrDefault(key string, def interface{}) interface{}
}

// StringGetter is an interface to get string value
type StringGetter interface {
	GetString(key string) string
}

// BoolGetter is an interface to get bool value
type BoolGetter interface {
	GetBool(key string) bool
}

// IntGetter is an interface to get int64 value
type IntGetter interface {
	GetInt(key string) int64
}

// UIntGetter is an interface to get uint64 value
type UIntGetter interface {
	GetUInt(key string) uint64
}

// FloatGetter is an interface to get float64 value
type FloatGetter interface {
	GetFloat(key string) float64
}
