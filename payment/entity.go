package payment

type Transaction struct {
	ID     string
	Amount float64 `gorm:"type:decimal(10,2)"`
}
