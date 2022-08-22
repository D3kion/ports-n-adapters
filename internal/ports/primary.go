package ports

// import (
// 	"context"
// )

type GRPCPort interface {
	Run()
	GetAdd()
	GetSubtract()
	GetMultiply()
	GetDivide()
}
