package interfaces

type Monitor interface {
	Init()

	AddFeature(prefix, key string, value interface{}, err error, option ...interface{})

	Close()
}
