package customerstorage

import (
	"context"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) Create(ctx context.Context, data *customermodel.CustomerCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return err
	}
	return nil
}
