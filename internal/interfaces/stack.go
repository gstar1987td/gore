package interfaces

type Stack interface {
	Push(value interface{})
	Pop() (value interface{}, ok bool)
	Peek() (value interface{}, ok bool)

	Empty() bool
	Size() int
	Clear()
	Values() []interface{}
	Close()
}
