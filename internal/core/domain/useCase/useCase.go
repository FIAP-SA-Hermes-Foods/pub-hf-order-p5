package useCase

import "pub-hf-order-p5/internal/core/domain/entity/dto"

type ProductUseCase interface {
	SaveOrder(reqOrder dto.RequestOrder) error
	GetOrderByID(uuid string) error
	UpdateOrderByID(uuid string, order dto.RequestOrder) error
	GetOrders(order dto.ResquestOrder) error
}
