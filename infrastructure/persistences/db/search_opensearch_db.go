package db

import (
	"crypto/tls"
	"log"
	"net/http"
	config_local "products-crud/infrastructure/config"

	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	opensearch "github.com/opensearch-project/opensearch-go/v2"
)

func NewProductOpenSearchDB() (*opensearch.Client, error) {

	// openS := &opensearch_entity.OpenSearch{
	// 	Client: &http.Client{},
	// 	// DomainEndpoint: os.Getenv("AWS_DOMAIN_ENDPOINT"),
	// 	// Username:       os.Getenv("AWS_USER"),
	// 	// Password:       os.Getenv("AWS_PASSWORD"),
	endpoint := config_local.Configuration.GetString("awsOpensearch.domainEndpoint")
	username := config_local.Configuration.GetString("awsOpensearch.opensearch.dev.user")
	password := config_local.Configuration.GetString("awsOpensearch.opensearch.dev.pass")
	// }

	// return openS, nil

	// ctx := context.Background()

	// awsCfg, err := config.LoadDefaultConfig(ctx,
	// 	config.WithRegion("us-east-1"),
	// )
	// if err != nil {
	// 	log.Fatal(err) // Do not log.fatal in a production ready app.
	// }

	// Create an AWS request Signer and load AWS configuration using default config folder or env vars.
	// signer, err := requestsigner.NewSignerWithService(awsCfg, "es")
	// if err != nil {
	// 	log.Fatal(err) // Do not log.fatal in a production ready app.
	// }

	// Create an opensearch client and use the request-signer
	// client, err := opensearch.NewClient(opensearch.Config{
	// 	Addresses: []string{endpoint},
	// 	Signer:    signer,
	// })
	// if err != nil {
	// 	log.Fatal("client creation err", err)
	// }

	client, err := opensearch.NewClient(opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{endpoint},
		Username:  username,
		Password:  password,
	})

	if err != nil {
		log.Fatal("client creation err", err)
	}

	return client, nil
}

func getCredentialProvider(accessKey, secretAccessKey, token string) aws.CredentialsProviderFunc {
	return func(ctx context.Context) (aws.Credentials, error) {
		c := &aws.Credentials{
			AccessKeyID:     accessKey,
			SecretAccessKey: secretAccessKey,
			SessionToken:    token,
		}
		return *c, nil
	}
}
