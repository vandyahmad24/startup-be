package transaction

type service struct {
	repository Repository
}
type Service interface {
	GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionCampaignInput) ([]Transaction, error) {
	transaction, err := s.repository.GetByCampaignId(input.Id)
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
