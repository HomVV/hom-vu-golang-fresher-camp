package customerbiz

import (
	"context"
	"demo/modules/customer/customermodel"
	"errors"
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
		return nil, err
	}
	if data == nil {
		return nil, errors.New("not available customer")
	}
	return data, nil
}
