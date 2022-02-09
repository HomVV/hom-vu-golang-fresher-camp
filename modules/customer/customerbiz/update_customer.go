package customerbiz

import (
	"context"
	"demo/modules/customer/customermodel"
	"errors"
)

type UpdateCustomerStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
	UpdateData(
		ctx context.Context,
		id int,
		data *customermodel.CustomerUpdate,
	) error
}
type updateCustomerBiz struct {
	store UpdateCustomerStore
}

func NewUpdateCustomerBiz(store UpdateCustomerStore) *updateCustomerBiz {
	return &updateCustomerBiz{store: store}
}
func (biz *updateCustomerBiz) UpdateCustomer(ctx context.Context, id int, data *customermodel.CustomerUpdate) error {
	oldData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})

	if err != nil {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("data deleted")
	}

	if err := biz.store.UpdateData(ctx, id, data); err != nil {
		return err
	}
	return nil
}
