package payment

import (
	"github.com/veritrans/go-midtrans"
	"log"
	"startup/config"
	"startup/transaction"
	"startup/users"
)

type service struct {
}

type Service interface {
	GetToken(transaction transaction.Transaction, user users.User) (string, error)
}

func NewService() *service {
	return &service{}
}
func (s *service) GetToken(transaction transaction.Transaction, user users.User) (string, error) {
	cfg := config.GetConfig()
	midclient := midtrans.NewClient()
	midclient.ServerKey = cfg.Config.MidtransServerKey
	midclient.ClientKey = cfg.Config.MidtransClientKey
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
	return resp.Token, nil
}
