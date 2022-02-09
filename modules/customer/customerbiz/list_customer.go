package customerbiz

import (
	"context"
	"demo/common"
	"demo/modules/customer/customermodel"
)

type ListCustomerStore interface {
	ListDataByCondition(ctx context.Context,
		conditions map[string]interface{},
		filter *customermodel.Filter,
		paging *common.Paging,
		moreKeys ...string,
	) ([]customermodel.Customer, error)
}
type listCustomerBiz struct {
	store ListCustomerStore
}

func NewListCustomerBiz(store ListCustomerStore) *listCustomerBiz {
	return &listCustomerBiz{store: store}
}
func (biz *listCustomerBiz) ListCustomer(
	ctx context.Context,
	filter *customermodel.Filter,
	paging *common.Paging) ([]customermodel.Customer, error) {

	result, err := biz.store.ListDataByCondition(ctx, nil, filter, paging)
	return result, err
}
