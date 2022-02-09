package customerbiz

import (
	"context"
	"demo/modules/customer/customermodel"
)

type CreateCustomerStore interface {
	Create(ctx context.Context, data *customermodel.CustomerCreate) error
}
type createCustomerBiz struct {
	store CreateCustomerStore
}

func NewCreateCustomerBiz(store CreateCustomerStore) *createCustomerBiz {
	return &createCustomerBiz{store: store}
}
func (biz *createCustomerBiz) CreateCustomer(ctx context.Context, data *customermodel.CustomerCreate) error {

	if err := data.Validate(); err != nil {
		return err
	}
	err := biz.store.Create(ctx, data)
	return err
}
