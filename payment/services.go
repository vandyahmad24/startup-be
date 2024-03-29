package payment

import (
	"github.com/veritrans/go-midtrans"
	"log"
	"startup/campaign"
	"startup/config"
	"startup/users"
)

type service struct {
	campaignRepo campaign.Repository
}

type Service interface {
	GetPaymentUrl(transaction Transaction, user users.User) (string, error)
}

func NewService(campaignRepo campaign.Repository) *service {
	return &service{campaignRepo: campaignRepo}
}
func (s *service) GetPaymentUrl(transaction Transaction, user users.User) (string, error) {
	cfg := config.NewConfig()
	midclient := midtrans.NewClient()
	midclient.ServerKey = cfg.MIDTRANS_SERVER_KEY
	midclient.ClientKey = cfg.MIDTRANS_ACCESS_KEY
	midclient.APIEnvType = midtrans.Sandbox

	snapGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	snapReq := &midtrans.SnapReq{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  transaction.Code,
			GrossAmt: int64(transaction.Amount),
		},
		EnabledPayments: nil,
		Callbacks:       nil,
		Items:           nil,
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
	}

	resp, err := snapGateway.GetToken(snapReq)
	if err != nil {
		log.Println("error midtrans ", err)
		return "", err
	}
	return resp.RedirectURL, nil
}
