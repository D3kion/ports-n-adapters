package ports

type DbPort interface {
	Close()
	AddToHistory(answer int32, operation string) error
}
