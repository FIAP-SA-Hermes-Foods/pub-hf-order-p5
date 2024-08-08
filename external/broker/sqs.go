package broker

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
)

type sqsBroker struct {
	ctx    context.Context
	config aws.Config
	client *sqs.Client
}

func NewSQSBroker(config aws.Config) *sqsBroker {
	return &sqsBroker{config: config}
}

func (s *sqsBroker) clientSQS() {
	s.client = sqs.NewFromConfig(s.config)
}

func (s *sqsBroker) Pub(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error) {

	if s.client == nil {
		s.clientSQS()
	}

	if s.ctx == nil {
		s.ctx = context.Background()
	}

	return s.client.SendMessage(s.ctx, input)
}

func (s *sqsBroker) Sub(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error) {

	if s.client == nil {
		s.clientSQS()
	}

	if s.ctx == nil {
		s.ctx = context.Background()
	}

	return s.client.ReceiveMessage(s.ctx, input)
}

func (s *sqsBroker) Delete(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error) {

	if s.client == nil {
		s.clientSQS()
	}

	if s.ctx == nil {
		s.ctx = context.Background()
	}

	return s.client.DeleteMessage(s.ctx, input)
}
