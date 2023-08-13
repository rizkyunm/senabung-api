package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rizkyunm/senabung-api/campaign"
	"github.com/rizkyunm/senabung-api/helper"
	"github.com/rizkyunm/senabung-api/mail"
	"github.com/rizkyunm/senabung-api/transaction"
	"github.com/rizkyunm/senabung-api/user"
	"log"
	"net/http"
)

// catch parameter from URI
// mapping parameter to input struct
// call service, input struct as parameter
// service, get campaign ID
// repo search data which match with the campaign ID

type transactionHandler struct {
	transactionService transaction.Service
	campaignService    campaign.Service
}

func NewTransactionHandler(transactionService transaction.Service, campaignService campaign.Service) *transactionHandler {
	return &transactionHandler{
		transactionService: transactionService,
		campaignService:    campaignService,
	}
}

func (h *transactionHandler) GetCampaignTransactions(c *gin.Context) {
	var input transaction.GetCampaignTransactionsInput

	if err := c.ShouldBindUri(&input); err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Failed to get campaigns's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.transactionService.GetTransactionsByCampaignID(input)
	if err != nil {
		fmt.Println(err)
		response := helper.APIResponse("Failed to get campaigns's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Campaign's Transaction", http.StatusOK, "success", transaction.FormatCampaignTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetUserTransactions(c *gin.Context) {
	currentUser := c.MustGet("current_user").(user.User)
	userID := currentUser.ID

	transactions, err := h.transactionService.GetTransactionsByUserID(userID)
	if err != nil {
		response := helper.APIResponse("Failed to get user's transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("User's Transaction", http.StatusOK, "success", transaction.FormatUserTransactions(transactions))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) CreateTransaction(c *gin.Context) {
	// input from User
	// handler : catch input and mapping to input struct
	// call service create transactions --> call midtrans system
	// call repository create new transaction data

	var input transaction.CreateTransactionInput

	if err := c.ShouldBindJSON(&input); err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("current_user").(user.User)
	input.User = currentUser

	newTransaction, err := h.transactionService.CreateTransaction(input)
	if err != nil {
		response := helper.APIResponse("Failed to create transaction", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("Transaction has been created", http.StatusOK, "success", transaction.FormatTransaction(newTransaction))
	c.JSON(http.StatusOK, response)
}

func (h *transactionHandler) GetNotification(c *gin.Context) {
	var input transaction.TransactionNotificationInput

	if err := c.ShouldBindJSON(&input); err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	trx, err := h.transactionService.GetTransactionByOrderID(input.OrderID)
	if err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err = h.transactionService.ProcessPayment(input, trx); err != nil {
		response := helper.APIResponse("Failed to process notification", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	if err = mail.SendThankYouEmail(trx, input); err != nil {
		log.Println(err)
	}

	c.JSON(http.StatusOK, input)
}
