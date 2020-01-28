package models

type DataLoader interface {
}

type DataChecker interface {
}

type MapDataLoader struct {
}

func NewMapDataLoader() (DataLoader, error) {
	return MapDataLoader{}, nil
}

type SimpleDataChecker struct {
}

func NewSimpleDataChecker() (DataChecker, error) {
	return SimpleDataChecker{}, nil
}
