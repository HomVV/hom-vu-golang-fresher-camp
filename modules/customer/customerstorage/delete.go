package customerstorage

import (
	"context"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) SoftDeleteData(
	ctx context.Context,
	id int,
) error {
	db := s.db
	if err := db.Table(customermodel.Customer{}.TableName()).
		Where("id= ?", id).Updates(map[string]interface{}{
		"status": 0,
	}).Error; err != nil {
		return err
	}
	return nil
}
