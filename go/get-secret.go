// Use this code snippet in your app.
// If you need more information about configurations or implementing the sample code, visit the AWS docs:
// https://aws.github.io/aws-sdk-go-v2/docs/getting-started/

// 21.05.2024 Soon: The new build tag syntax was introduced in Go 1.17 and later version.
// https://github.com/bitfield/script/issues/131
// Hence, please do not use go version <= 1.16. Otherwise, you will encounter the following errors:
// Error: ../../../go/pkg/mod/github.com/aws/aws-sdk-go-v2/credentials@v1.17.15/ssocreds/sso_credentials_provider.go:9:2: //go:build comment without // +build comment
// Error: ../../../go/pkg/mod/github.com/aws/aws-sdk-go-v2/credentials@v1.17.15/ssocreds/sso_token_provider.go:11:2: //go:build comment without // +build comment
// Error: ../../../go/pkg/mod/github.com/aws/aws-sdk-go-v2/credentials@v1.17.15/stscreds/assume_role_provider.go:108:2: //go:build comment without // +build comment

package main

import (
	"context"
	"log"
	// Need to import the fmt package also.
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

func main() {
	secretName := "soon-secret-name"
	region := "us-east-1"

	config, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		log.Fatal(err)
	}

	// Create Secrets Manager client
	svc := secretsmanager.NewFromConfig(config)

	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String("AWSCURRENT"), // VersionStage defaults to AWSCURRENT if unspecified
	}

	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		// For a list of exceptions thrown, see
		// https://docs.aws.amazon.com/secretsmanager/latest/apireference/API_GetSecretValue.html
		log.Fatal(err.Error())
	}

	// Decrypts secret using the associated KMS key.
	var secretString string = *result.SecretString

	// Your code goes here.

        // Expected output: {"soon-secret-key":"soon-secret-value"}
        fmt.Print("Secret: ", secretString)
}
