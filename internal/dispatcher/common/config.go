package common

type Auth1 struct {
	Issuer       string `envconfig:"AUTH1_ISSUER" default:"https://dev-auth1.tst.protocol.one"`
	ClientId     string `envconfig:"AUTH1_CLIENTID" required:"true"`
	ClientSecret string `envconfig:"AUTH1_CLIENTSECRET" required:"true"`
	RedirectUrl  string `envconfig:"AUTH1_REDIRECTURL" required:"true"`
}

type Config struct {
	Auth1

	AwsAccessKeyIdAgreement     string `envconfig:"AWS_ACCESS_KEY_ID_AGREEMENT" required:"true"`
	AwsSecretAccessKeyAgreement string `envconfig:"AWS_SECRET_ACCESS_KEY_AGREEMENT" required:"true"`
	AwsRegionAgreement          string `envconfig:"AWS_REGION_AGREEMENT" default:"eu-west-1"`
	AwsBucketAgreement          string `envconfig:"AWS_BUCKET_AGREEMENT" required:"true"`

	AwsAccessKeyIdReporter     string `envconfig:"AWS_ACCESS_KEY_ID_REPORTER" required:"true"`
	AwsSecretAccessKeyReporter string `envconfig:"AWS_SECRET_ACCESS_KEY_REPORTER" required:"true"`
	AwsRegionReporter          string `envconfig:"AWS_REGION_REPORTER" default:"eu-west-1"`
	AwsBucketReporter          string `envconfig:"AWS_BUCKET_REPORTER" required:"true"`

	LimitDefault          int32 `default:"100"`
	OffsetDefault         int32 `default:"0"`
	LimitMax              int32 `default:"1000"`
	DisableAuthMiddleware bool

	OrderInlineFormUrlMask string `envconfig:"ORDER_INLINE_FORM_URL_MASK" required:"true"`

	AllowOrigin string `envconfig:"ALLOW_ORIGIN" default:"*"`
}
