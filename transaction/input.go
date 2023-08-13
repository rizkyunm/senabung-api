package transaction

import "github.com/rizkyunm/senabung-api/user"

type GetCampaignTransactionsInput struct {
	ID   uint `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount     float64 `json:"amount" binding:"required"`
	CampaignID uint    `json:"campaign_id" binding:"required"`
	User       user.User
}

type TransactionNotificationInput struct {
	TransactionStatus string `json:"transaction_status"`
	OrderID           string `json:"order_id"`
	PaymentType       string `json:"payment_type"`
	FraudStatus       string `json:"fraud_status"`
	TransactionTime   string `json:"transaction_time"`
	GrossAmount       string `json:"gross_amount"`
}
