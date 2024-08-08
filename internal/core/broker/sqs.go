package broker

import "github.com/aws/aws-sdk-go-v2/service/sqs"

type SQSBroker interface {
	Pub(input *sqs.SendMessageInput) (*sqs.SendMessageOutput, error)
	Sub(input *sqs.ReceiveMessageInput) (*sqs.ReceiveMessageOutput, error)
	Delete(input *sqs.DeleteMessageInput) (*sqs.DeleteMessageOutput, error)
}
