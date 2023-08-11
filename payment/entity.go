package payment

type Transaction struct {
	ID     uint
	Amount float64 `gorm:"type:decimal(10,2)"`
}
