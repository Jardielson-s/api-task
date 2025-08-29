package ses

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Jardielson-s/api-task/infra"
	RolesRepo "github.com/Jardielson-s/api-task/modules/roles/repositories"
	RolesService "github.com/Jardielson-s/api-task/modules/roles/services"
	UserRolesRepo "github.com/Jardielson-s/api-task/modules/user_roles/repository"

	"github.com/Jardielson-s/api-task/modules/shared"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ses"
	"github.com/aws/aws-sdk-go-v2/service/ses/types"
)

func getEmails(roleService RolesService.RolesService) []string {
	result, err := roleService.FindByRoleByName(shared.GetManagerRole())
	if err != nil {
		return nil
	}
	return result

}
func SendEmailService(message string) {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
		func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			if service == ses.ServiceID {
				return aws.Endpoint{
					URL: os.Getenv("AWS_ENDPOINT"),
				}, nil
			}
			return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
		})), config.WithRegion("us-east-1"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ses.NewFromConfig(cfg)
	db, _ := infra.InitInfraDb()

	roleRepo := RolesRepo.NewRolesRepository(db)
	userRolesRepo := UserRolesRepo.NewUserRolesRepository(db)
	toAddresses := getEmails(RolesService.NewRolesRepository(
		roleRepo,
		userRolesRepo,
	))
	fromAddress := os.Getenv("EMAIL_APPLICATION") // email from application
	subject := "Notification Task Created"
	bodyText := message

	for _, to := range toAddresses {
		sendEmail(svc, fromAddress, to, subject, bodyText)
	}
}

func sendEmail(svc *ses.Client, fromAddress, toAddress, subject, bodyText string) {
	input := &ses.SendEmailInput{
		Source: aws.String(fromAddress),
		Destination: &types.Destination{
			ToAddresses: []string{
				toAddress,
			},
		},
		Message: &types.Message{
			Subject: &types.Content{
				Charset: aws.String("UTF-8"),
				Data:    aws.String(subject),
			},
			Body: &types.Body{
				Text: &types.Content{
					Charset: aws.String("UTF-8"),
					Data:    aws.String(bodyText),
				},
			},
		},
	}

	_, err := svc.SendEmail(context.TODO(), input)
	if err != nil {
		log.Printf("Failed to send email to %s: %v", toAddress, err)
		return
	}

	fmt.Printf("Email sent to %s with Message: %s\n", toAddress, bodyText)
}
