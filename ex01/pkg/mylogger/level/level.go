package level

// Level ...
// Тип для хранения уровня лога
type Level int

// Константы уровней лог
const (
	Debug Level = -4
	Info  Level = 0
	Warn  Level = 4
	Error Level = 8
)
