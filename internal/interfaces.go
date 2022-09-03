package internal

// Добавьте в пакет iternal интерфейс Discounter
type Discounter interface {
	CalcDiscount() (int, error)
}
