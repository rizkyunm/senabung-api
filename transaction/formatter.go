package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID        uint      `json:"id"`
	OrderNo   string    `json:"order_no"`
	Name      string    `json:"name"`
	Amount    float64   `json:"amount"`
	Donor     string    `json:"donor"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	campaignTransactionFormatter := CampaignTransactionFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		Donor:     transaction.User.Name,
		OrderNo:   transaction.OrderNo,
		CreatedAt: transaction.CreatedAt,
	}

	return campaignTransactionFormatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionFormatter{}
	}

	var transactionsFormatter []CampaignTransactionFormatter

	for _, transaction := range transactions {
		campaignTransactionFormatter := FormatCampaignTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, campaignTransactionFormatter)
	}

	return transactionsFormatter
}

type UserTransactionFormatter struct {
	ID         uint              `json:"id"`
	OrderNo    string            `json:"order_no"`
	Amount     float64           `json:"amount"`
	Status     string            `json:"status"`
	PaymentURL string            `json:"payment_url"`
	CreatedAt  time.Time         `json:"created_at"`
	Campaign   CampaignFormatter `json:"campaign"`
}

type CampaignFormatter struct {
	Name          string `json:"name"`
	CampaignImage string `json:"campaign_image"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	userTransactionFormatter := UserTransactionFormatter{
		ID:         transaction.ID,
		OrderNo:    transaction.OrderNo,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		PaymentURL: transaction.PaymentURL,
		CreatedAt:  transaction.CreatedAt,
	}

	campaignFormatter := CampaignFormatter{
		Name:          transaction.Campaign.Name,
		CampaignImage: transaction.Campaign.CampaignImage,
	}

	userTransactionFormatter.Campaign = campaignFormatter

	return userTransactionFormatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	if len(transactions) == 0 {
		return []UserTransactionFormatter{}
	}

	var transactionsFormatter []UserTransactionFormatter

	for _, transaction := range transactions {
		userTransactionFormatter := FormatUserTransaction(transaction)
		transactionsFormatter = append(transactionsFormatter, userTransactionFormatter)
	}

	return transactionsFormatter
}

type TransactionFormatter struct {
	ID         uint    `json:"id"`
	OrderNo    string  `json:"order_no"`
	CampaignID uint    `json:"campaign_id"`
	UserID     uint    `json:"user_id"'`
	Amount     float64 `json:"amount"`
	Status     string  `json:"status"`
	Code       string  `json:"code"`
	PaymentURL string  `json:"payment_url"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	campaignTransactionFormatter := TransactionFormatter{
		ID:         transaction.ID,
		OrderNo:    transaction.OrderNo,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentURL: transaction.PaymentURL,
	}

	return campaignTransactionFormatter
}
