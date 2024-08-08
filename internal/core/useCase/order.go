package useCase

import (
	"errors"
	"pub-hf-order-p5/internal/core/domain/entity/dto"
	vo "pub-hf-order-p5/internal/core/domain/entity/valueObject"
	"pub-hf-order-p5/internal/core/domain/useCase"
	"strings"
)

var _ useCase.OrderUseCase = (*orderUseCase)(nil)

type orderUseCase struct {
}

func NewOrderUseCase() orderUseCase {
	return orderUseCase{}
}

func (p orderUseCase) SaveOrder(reqOrder dto.RequestOrder) error {
	order := reqOrder.Order()

	if err := order.VerificationCode.Validate(); err != nil {
		return err
	}

	reqOrder.Status = order.Status.Value

	return nil
}

func (p orderUseCase) UpdateOrderByID(uuid string, reqOrder dto.RequestOrder) error {
	if len(uuid) < 1 {
		return errors.New("the id is not valid for consult")
	}

	order := reqOrder.Order()

	if err := order.Status.Validate(); err != nil {
		return err
	}

	reqOrder.Category = order.Category.Value

	return nil
}

func (p orderUseCase) GetOrderByID(uuid string) error {
	if len(uuid) < 1 {
		return errors.New("the id is not valid for consult")
	}
	return nil
}

func (p orderUseCase) GetOrders(order string) error {
	return nil
}

