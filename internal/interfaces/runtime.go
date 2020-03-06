package interfaces

type Runtime interface {
	Init() error
	Close()

	GetStack() Stack
	GetMonitor() Monitor
}
