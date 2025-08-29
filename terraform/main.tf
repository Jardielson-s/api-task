terraform {
  required_version = ">= 1.2.0"
  required_providers {
    aws = {
      source  = "hashicorp/aws",
      version = "~> 4.16"
    }
  }
}

provider "aws" {
  access_key                  = "test"
  secret_key                  = "test"
  region                      = "us-east-1"
  skip_credentials_validation = true
  skip_requesting_account_id  = true
  skip_metadata_api_check     = true
  endpoints {
    sqs = "http://localhost:4566"
    ses = "http://localhost:4566"
  }
}

resource "aws_sqs_queue" "test_queue" {
  name = "notification-queue"
}

resource "aws_ses_domain_identity" "test_ses" {
  domain = "example.com"
}

# resource "aws_ses_domain_dkim" "test_dkim" {
#   domain = aws_ses_domain_identity.test_ses.domain
# }

output "sqs_queue_url" {
  value = aws_sqs_queue.test_queue.id
}

output "ses_identity" {
  value = aws_ses_domain_identity.test_ses.domain
}
