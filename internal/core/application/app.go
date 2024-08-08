package application

import (
	"context"
	l "pub-hf-order-p5/external/logger"
	ps "pub-hf-order-p5/external/strings"
	"pub-hf-order-p5/internal/core/domain/broker"
	"pub-hf-order-p5/internal/core/domain/entity/dto"
)

type Application interface {
	GetOrderByID(msgID string, uuid string) error
	SaveOrder(msgID string, order dto.RequestOrder) error
	UpdateOrderByID(msgID string, id string, order dto.RequestOrder) error
	GetOrders(msgID string, order dto.RequestOrder) error
}

type application struct {
	ctx           context.Context
	orderBroker broker.OrderBroker
}

func NewApplication(ctx context.Context, orderBroker broker.OrderBroker) Application {
	return application{
		ctx:         ctx,
		orderBroker: orderBroker,
	}
}

func (app application) GetOrders(msgID string, order dto.RequestOrder) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetOrdersApp: ", " | ", ps.MarshalString(order))

	inputBroker := dto.OrderBroker{
		UUID:          	  order.UUID,
		MessageID:     	  msgID,
		VoucherUUID:      order.VoucherUUID,
		Items:            order.itemOutList,
		Status:           order.Status,
		VerificationCode: order.VerificationCode,
		CreatedAt:        order.CreatedAt,
	}

	if err := app.orderBroker.GetOrders(inputBroker); err != nil {
		l.Errorf(msgID, "GetOrdersApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetOrdersApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) GetOrderByID(msgID string, uuid string) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "GetOrderByIDApp: ", " | ", uuid)

	inputBroker := dto.OrderBroker{
		UUID:      uuid,
		MessageID: msgID,
	}

	if err := app.orderBroker.GetOrderByID(inputBroker); err != nil {
		l.Errorf(msgID, "GetOrderByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "GetOrderByIDApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) SaveOrder(msgID string, order dto.RequestOrder) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "SaveOrderApp: ", " | ", ps.MarshalString(order))

	inputBroker := dto.OrderBroker{
		UUID:          order.UUID,
		MessageID:     msgID,
		ClientUUID:    order.ClientUUID,
		VoucherUUID:   order.VoucherUUID,
		Items:		   order.Items,
	}

	if err := app.orderBroker.SaveOrder(inputBroker); err != nil {
		l.Errorf(msgID, "SaveOrderApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "SaveOrderApp output: ", " | ", "message sent with success!")
	return nil
}

func (app application) UpdateOrderByID(msgID string, id string, order dto.RequestOrder) error {
	app.setMessageIDCtx(msgID)
	l.Infof(msgID, "UpdateOrderByIDApp: ", " | ", id, " | ", ps.MarshalString(order))

	inputBroker := dto.OrderBroker{
		UUID:          id,
		MessageID:     msgID,
		ClientUUID:    order.ClientUuid,
  		VoucherUUID:   order.VoucherUuid,
  		Item:		   order.Item,
  		Status:		   order.Status,
  		VerificationCode: order.VerificationCode,
  		CreatedAt:	   order.CreatedAt,	
	}

	if err := app.orderBroker.UpdateOrderByID(inputBroker); err != nil {
		l.Errorf(msgID, "UpdateOrderByIDApp error: ", " | ", err)
		return err
	}

	l.Infof(msgID, "UpdateOrderByIDApp output: ", " | ", "message sent with success!")
	return nil
}


func (app application) setMessageIDCtx(msgID string) {
	if app.ctx == nil {
		app.ctx = context.WithValue(context.Background(), l.MessageIDKey, msgID)
		return
	}
	app.ctx = context.WithValue(app.ctx, l.MessageIDKey, msgID)
}
