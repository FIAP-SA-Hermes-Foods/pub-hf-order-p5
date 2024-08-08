package broker

import (
	l "pub-hf-order-p5/external/logger"
	ps "pub-hf-order-p5/external/strings"
	sqsBroker "pub-hf-order-p5/internal/core/broker"
	pBroker "pub-hf-order-p5/internal/core/domain/broker"
	"pub-hf-order-p5/internal/core/domain/entity/dto"

	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

var _ pBroker.OrderBroker = (*orderBroker)(nil)

type orderBroker struct {
	queueURL string
	broker   sqsBroker.SQSBroker
}

func NewOrderBroker(broker sqsBroker.SQSBroker, queueURL string) *orderBroker {
	return &orderBroker{broker: broker, queueURL: queueURL}
}

func (p *orderBroker) GetOrders(input dto.OrderBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *orderBroker) GetOrderByID(input dto.OrderBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *orderBroker) SaveOrder(input dto.OrderBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}

func (p *orderBroker) UpdateOrderByID(input dto.OrderBroker) error {
	l.Infof(input.MessageID, "msg broker input: ", " | ", ps.MarshalString(input))

	msgBody := ps.MarshalString(input)

	inPub := &sqs.SendMessageInput{
		QueueUrl:    &p.queueURL,
		MessageBody: &msgBody,
	}

	if _, err := p.broker.Pub(inPub); err != nil {
		return err
	}

	l.Infof(input.MessageID, "Message published successfully: ", " | ", ps.MarshalString(input))
	return nil
}
