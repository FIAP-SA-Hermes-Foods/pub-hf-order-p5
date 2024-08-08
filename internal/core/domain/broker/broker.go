package broker

import (
	"pub-hf-order-p5/internal/core/domain/entity/dto"
)

type ProductBroker interface {
	GetOrderByID(input dto.ProductBroker) error
	SaveOrder(input dto.ProductBroker) error
	UpdateOrderByID(input dto.ProductBroker) error
	GetOrders(input dto.ProductBroker) error
}
