package arithmetic

type Adapter struct{}

func NewAdapter() *Adapter {
	return &Adapter{}
}

func (arith Adapter) Add(a int32, b int32) (int32, error) {
	return a + b, nil
}

func (arith Adapter) Subtract(a int32, b int32) (int32, error) {
	return a - b, nil
}

func (arith Adapter) Multiply(a int32, b int32) (int32, error) {
	return a * b, nil
}

func (arith Adapter) Divide(a int32, b int32) (int32, error) {
	return a / b, nil
}
