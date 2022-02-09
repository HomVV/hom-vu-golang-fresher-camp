package customerstorage

import (
	"context"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) UpdateData(
	ctx context.Context,
	id int,
	data *customermodel.CustomerUpdate,
) error {
	db := s.db
	if err := db.Where("id =?", id).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
