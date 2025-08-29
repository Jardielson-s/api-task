package sqs

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

var queueURL = "http://localhost:4566/000000000000/notification-queue"

func CreateSQSClient() (*sqs.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"), config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{URL: "http://localhost:4566"}, nil
	})))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := sqs.NewFromConfig(cfg)
	return client, nil
}

func SendMessage(client *sqs.Client, messageBody string) error {
	_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &queueURL,
		MessageBody: &messageBody,
	})
	return err
}

func ProcessMessages(client *sqs.Client) {
	for {
		results, err := client.ReceiveMessage(
			context.Background(),
			&sqs.ReceiveMessageInput{
				QueueUrl:            &queueURL,
				MaxNumberOfMessages: 10,
			},
		)

		if err != nil {
			log.Fatal(err)
		}

		for _, message := range results.Messages {
			go ProcessJob(message, client)
		}
	}
}

func ProcessJob(message types.Message, sqsClient *sqs.Client) {

	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: message.ReceiptHandle,
	}
	log.Println(*message.Body)

	_, err := sqsClient.DeleteMessage(context.TODO(), deleteParams)

	if err != nil {
		log.Fatal(err)
	}
}
