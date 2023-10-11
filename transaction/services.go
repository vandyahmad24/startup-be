package transaction

import (
	"fmt"
	"startup/payment"
	"time"
)

type service struct {
	repository     Repository
	paymentService payment.Service
}
type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error)
	GetTransactionByUserId(userId int) ([]Transaction, error)
	CreateTransaction(input CreateTransactionInput) (Transaction, error)
}

func NewService(repository Repository, paymentService payment.Service) *service {
	return &service{repository: repository, paymentService: paymentService}
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
