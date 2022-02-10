package customerstorage

import (
	"context"
	"demo/common"
	"demo/modules/customer/customermodel"

	"gorm.io/gorm"
)

func (s *sqlStore) FindDataByCondition(
	ctx context.Context,
	condition map[string]interface{},
	moreKeys ...string) (*customermodel.Customer, error) {
	var result customermodel.Customer
	db := s.db
	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(customermodel.Customer{}.TableName()).Where(condition)

	if err := db.First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, common.RecordNotFound
		}
		return nil, common.ErrDB(err)
	}
	return &result, nil

}
