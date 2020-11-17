package common

type Auth1 struct {
	Issuer       string `envconfig:"AUTH1_ISSUER" default:"https://dev-auth1.tst.protocol.one"`
	ClientId     string `envconfig:"AUTH1_CLIENTID" required:"true"`
	ClientSecret string `envconfig:"AUTH1_CLIENTSECRET" required:"true"`
	RedirectUrl  string `envconfig:"AUTH1_REDIRECTURL" required:"true"`
}

type LogsSettings struct {
	AwsCloudWatchAccessKeyId             string `envconfig:"AWS_CLOUDWATCH_ACCESS_KEY_ID" required:"true"`
	AwsCloudWatchSecretAccessKey         string `envconfig:"AWS_CLOUDWATCH_SECRET_ACCESS_KEY" required:"true"`
	AwsCloudWatchRegion                  string `envconfig:"AWS_CLOUDWATCH_REGION" default:"eu-west-1"`
	AwsCloudWatchLogGroupBillingServer   string `envconfig:"AWS_CLOUDWATCH_LOG_GROUP_BILLING_SERVER" required:"true"`
	AwsCloudWatchLogGroupManagementApi   string `envconfig:"AWS_CLOUDWATCH_LOG_GROUP_MANAGEMENT_API" required:"true"`
	AwsCloudWatchLogGroupWebhookNotifier string `envconfig:"AWS_CLOUDWATCH_LOG_GROUP_WEBHOOK_NOTIFIER" required:"true"`
}

type Config struct {
	Auth1
	*LogsSettings

	AwsAccessKeyIdAgreement     string `envconfig:"AWS_ACCESS_KEY_ID_AGREEMENT" required:"true"`
	AwsSecretAccessKeyAgreement string `envconfig:"AWS_SECRET_ACCESS_KEY_AGREEMENT" required:"true"`
	AwsRegionAgreement          string `envconfig:"AWS_REGION_AGREEMENT" default:"eu-west-1"`
	AwsBucketAgreement          string `envconfig:"AWS_BUCKET_AGREEMENT" required:"true"`

	AwsAccessKeyIdReporter     string `envconfig:"AWS_ACCESS_KEY_ID_REPORTER" required:"true"`
	AwsSecretAccessKeyReporter string `envconfig:"AWS_SECRET_ACCESS_KEY_REPORTER" required:"true"`
	AwsRegionReporter          string `envconfig:"AWS_REGION_REPORTER" default:"eu-west-1"`
	AwsBucketReporter          string `envconfig:"AWS_BUCKET_REPORTER" required:"true"`

	AwsAccessKeyIdMerchantDocs     string `envconfig:"AWS_ACCESS_KEY_ID_MERCHANTDOCS" required:"true"`
	AwsSecretAccessKeyMerchantDocs string `envconfig:"AWS_SECRET_ACCESS_KEY_MERCHANTDOCS" required:"true"`
	AwsRegionMerchantDocs          string `envconfig:"AWS_REGION_MERCHANTDOCS" default:"eu-west-1"`
	AwsBucketMerchantDocs          string `envconfig:"AWS_BUCKET_MERCHANTDOCS" required:"true"`

	LimitDefault          int32 `default:"100"`
	OffsetDefault         int32 `default:"0"`
	LimitMax              int32 `default:"1000"`
	DisableAuthMiddleware bool  `envconfig:"DISABLE_AUTH_MIDDLEWARE"`
	DisableCasbinPolicy   bool  `envconfig:"DISABLE_CASBIN_POLICY"`

	OrderInlineFormUrlMask string `envconfig:"ORDER_INLINE_FORM_URL_MASK" required:"true"`

	AllowOrigin string `envconfig:"ALLOW_ORIGIN" default:"*"`
	HttpScheme  string `envconfig:"HTTP_SCHEME" default:"https"`
}
