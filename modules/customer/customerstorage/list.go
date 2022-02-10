package customerstorage

import (
	"context"
	"demo/common"
	"demo/modules/customer/customermodel"
)

func (s *sqlStore) ListDataByCondition(ctx context.Context,
	conditions map[string]interface{},
	filter *customermodel.Filter,
	paging *common.Paging,
	moreKeys ...string,
) ([]customermodel.Customer, error) {
	var result []customermodel.Customer

	db := s.db

	for i := range moreKeys {
		db = db.Preload(moreKeys[i])
	}

	db = db.Table(customermodel.Customer{}.TableName()).Where(conditions).Where(map[string]interface{}{"status": 1})

	if v := filter; v != nil {
		if v.Address != "" {
			db = db.Where("address = ?", v.Address)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	if err := db.
		Offset((paging.Page - 1) * paging.Limit).
		Limit(paging.Limit).
		Find(&result).Error; err != nil {
		return nil, common.ErrDB(err)
	}

	return result, nil
}
