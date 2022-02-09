package customerstorage

import (
	"context"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string,
) (*customermodel.Customer, error) {
	var result customermodel.Customer
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(customermodel.Customer{}.TableName()).Where(condition)

	if err := db.Find(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil

}
