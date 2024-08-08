package rpc

import (
	"context"
	l "pub-hf-order-p5/external/logger"
	"pub-hf-order-p5/internal/core/application"
	"pub-hf-order-p5/internal/core/domain/entity/dto"
	cp "pub-hf-order-p5/order_pub_proto"
)

type HandlerGRPC interface {
	Handler() *handlerGRPC
}

type handlerGRPC struct {
	app application.Application
	cp.UnimplementedOrderServer
}

// mustEmbedUnimplementedOrderServer implements order_api_proto.OrderServer.
func (h *handlerGRPC) mustEmbedUnimplementedOrderServer() {
	panic("unimplemented")
}

// UpdateOrder implements order_api_proto.OrderServer.
func (h *handlerGRPC) UpdateOrder(context.Context, *cp.UpdateOrderRequest) (*cp.UpdateOrderResponse, error) {
	panic("unimplemented")
}

// mustEmbedUnimplementedOrderServer implements order_api_proto.OrderServer.
func (h *handlerGRPC) mustEmbedUnimplementedOrderServer() {
	panic("unimplemented")
}

func NewHandler(app application.Application) HandlerGRPC {
	return &handlerGRPC{app: app}
}

func (h *handlerGRPC) Handler() *handlerGRPC {
	return h
}

func (h *handlerGRPC) GetOrderByID(ctx context.Context, req *cp.GetOrderByIDRequest) (*cp.GetOrderByIDResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	o, err := h.app.GetOrderByID(req.Id)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.GetOrderByIDResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) CreateOrder(ctx context.Context, req *cp.CreateOrderRequest) (*cp.CreateOrderResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	items := make([]dto.OutputOrderItem, 0)

	for i := range req.Items {
		items = append(items, dto.OutputOrderItem{
			OrderID:     req.Items[i].OrderId,
			ProductUUID: req.Items[i].ProductUuid,
			Quantity:    req.Items[i].Quantity,
		})
	}

	input := dto.RequestOrder{
		ClientUUID:  req.ClientUuid,
		VoucherUUID: req.ClientUuid,
		Items:       items,
	}

	o, err := h.app.SaveOrder(input)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.CreateOrderResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) UpdateProduct(ctx context.Context, req *cp.UpdateProductRequest) (*cp.UpdateProductResponse, error) {
	msgID := ctx.Value(l.MessageIDKey).(string)
	msgID = l.MessageID(msgID)

	items := make([]dto.OutputOrderItem, 0)

	for i := range req.Items {
		items = append(items, dto.OutputOrderItem{
			OrderID:     req.Items[i].OrderId,
			ProductUUID: req.Items[i].ProductUuid,
			Quantity:    req.Items[i].Quantity,
		})
	}

	input := dto.RequestOrder{
		ClientUUID:  req.ClientUuid,
		VoucherUUID: req.ClientUuid,
		Items:       items,
	}

	o, err := h.app.UpdateOrderByID(input.ID, input)

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outItems := make([]*cp.Item, 0)

	for i := range o.Items {
		item := cp.Item{
			ProductUuid: o.Items[i].ProductUUID,
			Quantity:    o.Items[i].Quantity,
		}
		outItems = append(outItems, &item)
	}

	out := &cp.UpdateOrderResponse{
		Id:               o.ID,
		ClientUuid:       o.ClientUUID,
		VoucherUuid:      o.VoucherUUID,
		Items:            outItems,
		VerificationCode: o.VerificationCode,
		Status:           o.Status,
		CreatedAt:        o.CreatedAt,
	}

	return out, nil
}

func (h *handlerGRPC) GetOrder(context.Context, *cp.GetOrderRequest) (*cp.GetOrderResponse, error) {

	o, err := h.app.GetOrders()

	if err != nil {
		return nil, err
	}

	if o == nil {
		return nil, nil
	}

	outOrders := make([]*cp.OrderItem, 0)
	for orderIdx := range o {

		outItems := make([]*cp.Item, 0)

		for i := range o[orderIdx].Items {
			item := cp.Item{
				ProductUuid: o[orderIdx].Items[i].ProductUUID,
				Quantity:    o[orderIdx].Items[i].Quantity,
			}
			outItems = append(outItems, &item)
		}

		oItem := cp.OrderItem{
			Id:               o[orderIdx].ID,
			ClientUuid:       o[orderIdx].ClientUUID,
			VoucherUuid:      o[orderIdx].VoucherUUID,
			Items:            outItems,
			Status:           o[orderIdx].Status,
			VerificationCode: o[orderIdx].VerificationCode,
			CreatedAt:        o[orderIdx].CreatedAt,
		}

		outOrders = append(outOrders, &oItem)

	}

	out := &cp.GetOrderResponse{
		Orders: outOrders,
	}

	return out, nil
}
