package customerbiz

import (
	"context"
	"demo/modules/customer/customermodel"
	"errors"
)

type SoftDeleteCustomerStore interface {
	FindDataByCondition(
		ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string,
	) (*customermodel.Customer, error)
	SoftDeleteData(
		ctx context.Context,
		id int,
	) error
}
type softDeleteCustomerBiz struct {
	store SoftDeleteCustomerStore
}

func NewSoftDeleteCustomerBiz(store SoftDeleteCustomerStore) *softDeleteCustomerBiz {
	return &softDeleteCustomerBiz{store: store}
}
func (biz *softDeleteCustomerBiz) SoftDeleteCustomer(
	ctx context.Context, id int) error {
	deleteData, err := biz.store.FindDataByCondition(ctx, map[string]interface{}{"id": id})
	if err != nil {
		return err
	}
	if deleteData.Status == 0 {
		return errors.New("customer deleted")
	}
	if err := biz.store.SoftDeleteData(ctx, id); err != nil {
		return err
	}
	return nil
}
