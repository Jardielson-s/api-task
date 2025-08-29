package sqs

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Jardielson-s/api-task/infra/aws/ses"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"gorm.io/gorm"
)

func CreateSQSClient() (*sqs.Client, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion("us-east-1"), config.WithEndpointResolver(aws.EndpointResolverFunc(func(service, region string) (aws.Endpoint, error) {
		return aws.Endpoint{URL: os.Getenv("AWS_ENDPOINT")}, nil
	})))
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	client := sqs.NewFromConfig(cfg)
	return client, nil
}

func SendMessage(client *sqs.Client, messageBody string) error {
	var queueURL = os.Getenv("AWS_SQS_ENDPOINT")
	_, err := client.SendMessage(context.TODO(), &sqs.SendMessageInput{
		QueueUrl:    &queueURL,
		MessageBody: &messageBody,
	})
	return err
}

func ProcessMessages(client *sqs.Client, db *gorm.DB) {
	var queueURL = os.Getenv("AWS_SQS_ENDPOINT")

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
			go ses.SendEmailService(*message.Body)
			go ProcessJob(queueURL, message, client)
		}
	}
}

func ProcessJob(queueURL string, message types.Message, sqsClient *sqs.Client) {
	deleteParams := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: message.ReceiptHandle,
	}
	_, err := sqsClient.DeleteMessage(context.TODO(), deleteParams)
	if err != nil {
		log.Fatal(err)
	}
}
