package interfaces

type Getter interface {
	DefaultGetter
	StringGetter
	FloatGetter
	BoolGetter
	IntGetter
}

type DefaultGetter interface {
	Get(key string) interface{}
	GetOrDefault(key string, def string) interface{}
}

type StringGetter interface {
	GetString(key string) string
}

type BoolGetter interface {
	GetBool(key string) bool
}

type IntGetter interface {
	GetInt(key string) int64
}

type FloatGetter interface {
	GetFloat(key string) float64
}