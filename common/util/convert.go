package util

import (
	"fmt"
	"strconv"
	"time"
)

func ConvertToBoolean(input interface{}) (res bool, err error) {
	switch value := input.(type) {
	case int:
		if value == 1 {
			res = true
		} else if value == 0 {
			res = false
		} else {
			err = fmt.Errorf("invalid to boolean value:%v", input)
		}
	case int32:
		if value == int32(1) {
			res = true
		} else if value == int32(0) {
			res = false
		} else {
			err = fmt.Errorf("invalid to boolean value:%v", input)
		}
	case int64:
		if value == int64(1) {
			res = true
		} else if value == int64(0) {
			res = false
		} else {
			err = fmt.Errorf("invalid to boolean value:%v", input)
		}
	case bool:
		res = value
	case string:
		res, err = strconv.ParseBool(value)
	default:
		err = fmt.Errorf("can not convert to boolean for %v", input)
	}
	return
}

func ConvertToInt64(input interface{}) (res int64, err error) {
	switch value := input.(type) {
	case float64:
		res = int64(value)
	case float32:
		res = int64(value)
	case int:
		res = int64(value)
	case int32:
		res = int64(value)
	case int64:
		res = value
	case string:
		if v, pErr := strconv.ParseInt(value, 10, 64); pErr != nil {
			err = pErr
		} else {
			res = v
		}
	default:
		err = fmt.Errorf("can not convert to int64 for %v", input)
	}
	return
}

func ConvertTofloat64(input interface{}) (res float64, err error) {
	switch value := input.(type) {
	case float64:
		res = value
	case float32:
		res = float64(value)
	case int:
		res = float64(value)
	case int32:
		res = float64(value)
	case int64:
		res = float64(value)
	case string:
		if v, pErr := strconv.ParseFloat(value, 64); pErr != nil {
			err = pErr
		} else {
			res = v
		}
	default:
		err = fmt.Errorf("can not convert to float64 for %v", input)
	}
	return
}

func ConvertToInt(input interface{}) (res int, err error) {
	switch value := input.(type) {
	case int:
		res = value
	case int32:
		res = int(value)
	case int64:
		res = int(value)
	case float32:
		res = int(value)
	case float64:
		res = int(value)
	case string:
		if v, pErr := strconv.ParseFloat(value, 64); pErr != nil {
			err = pErr
		} else {
			res = int(v)
		}
	}
	return
}

const (
	emptyDate             = "0000-00-00 00:00:00"
	TimeStampFormatLayOut = "2006-01-02 15:04:05"
)

func ConvertToTimestamp(input interface{}) (res int64, err error) {
	date, ok := input.(string)
	if !ok {
		err = fmt.Errorf("invalid date string:%v", input)
		return
	}
	if date == "" || date == emptyDate {
		err = fmt.Errorf("empty date string:%v", input)
		return
	}
	d, pErr := time.Parse(TimeStampFormatLayOut, date)
	if pErr != nil {
		err = pErr
		return
	}
	res = d.Unix()
	return
}

func FormatToString(input interface{}) (res string) {
	switch val := input.(type) {
	case string:
		res = val
	case int:
		res = strconv.Itoa(val)
	case uint:
		res = strconv.Itoa(int(val))
	case int8:
		res = strconv.Itoa(int(val))
	case uint8:
		res = strconv.Itoa(int(val))
	case int16:
		res = strconv.Itoa(int(val))
	case uint16:
		res = strconv.Itoa(int(val))
	case int32:
		res = strconv.Itoa(int(val))
	case uint32:
		res = strconv.Itoa(int(val))
	case int64:
		res = strconv.Itoa(int(val))
	case uint64:
		res = strconv.Itoa(int(val))
	case float32:
		res = strconv.FormatFloat(float64(val), 'f', -1, 32)
	case float64:
		res = strconv.FormatFloat(val, 'f', -1, 64)
	case bool:
		if val {
			res = "true"
		} else {
			res = "false"
		}
	default:
		res = fmt.Sprint(val)
	}
	return
}
