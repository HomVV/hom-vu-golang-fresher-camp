package customerstorage

import (
	"context"
	"demo/common"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) Create(ctx context.Context, data *customermodel.CustomerCreate) error {
	db := s.db
	if err := db.Create(data).Error; err != nil {
		return common.ErrDB(err)
	}
	return nil
}
