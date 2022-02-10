package customerbiz

import (
	"context"
	"demo/common"
	"demo/modules/customer/customermodel"
)

type FindCustomerStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
}
type findCustomerBiz struct {
	store FindCustomerStore
}

func NewFindCustomerBiz(store FindCustomerStore) *findCustomerBiz {
	return &findCustomerBiz{store: store}
}
func (biz *findCustomerBiz) FindCustomer(ctx context.Context, id int) (*customermodel.Customer, error) {
	data, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		if err != common.RecordNotFound {
			return nil, common.ErrCannotGetEntity(customermodel.EntityName, err)
		}
		return nil, common.ErrCannotGetEntity(customermodel.EntityName, err)
	}
	if data.Status == 0 {
		return nil, common.ErrEntityDeleted(customermodel.EntityName, nil)
	}
	return data, nil
}
