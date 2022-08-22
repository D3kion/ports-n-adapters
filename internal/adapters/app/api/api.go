package api

import (
	"hexarch/internal/ports"
)

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db, arith}
}

func (api Adapter) GetAdd(a, b int32) (int32, error) {
	result, err := api.arith.Add(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(result, "add")
	if err != nil {
		return result, err
	}

	return result, nil
}

func (api Adapter) GetSubtract(a, b int32) (int32, error) {
	result, err := api.arith.Subtract(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(result, "subtract")
	if err != nil {
		return result, err
	}

	return result, nil
}

func (api Adapter) GetMultiply(a, b int32) (int32, error) {
	result, err := api.arith.Multiply(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(result, "multiply")
	if err != nil {
		return result, err
	}

	return result, nil
}
func (api Adapter) GetDivide(a, b int32) (int32, error) {
	result, err := api.arith.Divide(a, b)
	if err != nil {
		return 0, err
	}

	err = api.db.AddToHistory(result, "divide")
	if err != nil {
		return result, err
	}

	return result, nil
}
