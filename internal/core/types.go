package core

type OperatorType int

const (
	_ OperatorType = iota - 1
	IntOperator
	FloatOperator
	BoolOperator
	StringOperator
	NullOperator
	Undefined

	OperatorTypeNum
)

var OperatorTypeMapping = [OperatorTypeNum][OperatorTypeNum]OperatorType{
	/* int | float | bool | string | null | ud*/
	/*int*/ {IntOperator, FloatOperator, Undefined, IntOperator, NullOperator, Undefined},
	/*float*/ {FloatOperator, FloatOperator, Undefined, FloatOperator, NullOperator, Undefined},
	/*bool*/ {Undefined, Undefined, BoolOperator, NullOperator, Undefined},
	/*string*/ {StringOperator, StringOperator, StringOperator, StringOperator, NullOperator, Undefined},
	/*null*/ {NullOperator, NullOperator, NullOperator, NullOperator, NullOperator, NullOperator},
	/*ud*/ {Undefined, Undefined, Undefined, Undefined, Undefined, Undefined},
}

func GetValueOperatorType(v interface{}) OperatorType {
	switch v.(type) {
	case int, int8, int16, int32, int64:
		return IntOperator
	case float32, float64:
		return FloatOperator
	case bool:
		return BoolOperator
	case string:
		return StringOperator
	case nil:
		return NullOperator
	default:
		return Undefined
	}
}

func GetOperatorType(left, right OperatorType) OperatorType {
	return OperatorTypeMapping[left][right]
}
