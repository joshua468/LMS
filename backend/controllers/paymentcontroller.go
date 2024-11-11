package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/joshua468/lms-backend/services"
)

// ProcessPayment handles payment requests using Stripe
func ProcessPayment(c *gin.Context) {
	var paymentData services.PaymentData
	if err := c.ShouldBindJSON(&paymentData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := services.ProcessPayment(paymentData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Payment processing failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Payment processed successfully"})
}
