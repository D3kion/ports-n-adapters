package ports

type APIPort interface {
	GetAdd(a, b int32) (int32, error)
	GetSubtract(a, b int32) (int32, error)
	GetMultiply(a, b int32) (int32, error)
	GetDivide(a, b int32) (int32, error)
}
