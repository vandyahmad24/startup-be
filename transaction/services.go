package transaction

import (
	"errors"
	"fmt"
	"startup/campaign"
	"startup/payment"
	"time"
)

type service struct {
	repository     Repository
	paymentService payment.Service
	campaignRepo   campaign.Repository
}
type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionData) error
}

func NewService(repository Repository, paymentService payment.Service, campaignRepo campaign.Repository) *service {
	return &service{repository: repository, paymentService: paymentService, campaignRepo: campaignRepo}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error) {
	transaction, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) GetTransactionByUserId(userId int) ([]Transaction, error) {
	transaction, err := s.repository.GetByUserId(userId)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (s *service) CreateTransaction(input CreateTransactionInput) (Transaction, error) {
	transaction := Transaction{
		CampaignId: input.CampaignId,
		UserId:     input.User.ID,
		Amount:     input.Amount,
		Status:     "pending",
		Code:       fmt.Sprintf("ORDER-%v", time.Now().Unix()),
	}

	transaction, err := s.repository.Save(transaction)
	if err != nil {
		return transaction, err
	}

	paymentURL, err := s.paymentService.GetPaymentUrl(payment.Transaction{
		Code:   transaction.Code,
		Amount: transaction.Amount,
	}, input.User)

	if err != nil {
		return transaction, err
	}

	transaction.PaymentUrl = paymentURL
	transaction, err = s.repository.Update(transaction)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (s *service) ProcessPayment(input TransactionData) error {
	transaction, err := s.repository.GetByCode(input.OrderID)
	if err != nil {
		return err
	}

	if transaction.Status == "paid" {
		return errors.New("transaction already paid")
	}
	if transaction.Status == "deny" {
		return errors.New("transaction already deny")
	}

	if input.PaymentType == "credit_card" && input.FraudStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == "cancel" {
		transaction.Status = "deny"
	}

	updateTransaction, err := s.repository.Update(transaction)
	if err != nil {
		return err
	}

	campaignRes, err := s.campaignRepo.FindByID(updateTransaction.CampaignId)
	if err != nil {
		return err
	}

	if updateTransaction.Status == "paid" {
		campaignRes.BackerCount = campaignRes.BackerCount + 1
		campaignRes.CurrentAmount = campaignRes.CurrentAmount + updateTransaction.Amount

		_, err = s.campaignRepo.Update(campaignRes)
		if err != nil {
			return err
		}
	}

	return nil
}
