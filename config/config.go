package config

import (
	"github.com/kelseyhightower/envconfig"
)

const (
	DefaultJwtSignAlgorithm = "RS256"
)

type Database struct {
	Host     string `envconfig:"MONGO_HOST"`
	Database string `envconfig:"MONGO_DB"`
	User     string `envconfig:"MONGO_USER"`
	Password string `envconfig:"MONGO_PASSWORD"`
}

type Auth1 struct {
	Issuer       string `envconfig:"AUTH1_ISSUER" required:"true" default:"https://dev-auth1.tst.protocol.one"`
	ClientId     string `envconfig:"AUTH1_CLIENTID" required:"true"`
	ClientSecret string `envconfig:"AUTH1_CLIENTSECRET" required:"true"`
	RedirectUrl  string `envconfig:"AUTH1_REDIRECTURL" required:"true"`
}

type S3 struct {
	AccessKeyId string `envconfig:"S3_ACCESS_KEY" required:"true"`
	SecretKey   string `envconfig:"S3_SECRET_KEY" required:"true"`
	Endpoint    string `envconfig:"S3_ENDPOINT" required:"true"`
	BucketName  string `envconfig:"S3_BUCKET_NAME" required:"true"`
	Region      string `envconfig:"S3_REGION" default:"us-west-2"`
	Secure      bool   `envconfig:"S3_SECURE" default:"false"`
}

type Config struct {
	Database
	Auth1
	S3

	HttpScheme     string `envconfig:"HTTP_SCHEME" default:"https"`
	KubernetesHost string `envconfig:"KUBERNETES_SERVICE_HOST" required:"false"`
	AmqpAddress    string `envconfig:"AMQP_ADDRESS" required:"true" default:"amqp://127.0.0.1:5672"`
	Environment    string `envconfig:"ENVIRONMENT" default:"test"`
}

func NewConfig() (error, *Config) {
	var err error

	config := Config{}

	if err = envconfig.Process("", &config); err != nil {
		return err, nil
	}

	return nil, &config
}
