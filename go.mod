module github.com/paysuper/paysuper-management-api

require (
	github.com/ProtocolONE/authone-jwt-verifier-golang v0.0.0-20190327070329-4dd563b01681
	github.com/ProtocolONE/geoip-service v1.0.3-0.20200203172514-41df5c78bf01
	github.com/ProtocolONE/go-core/v2 v2.1.0
	github.com/ProtocolONE/go-micro-plugins/wrapper/select/version v0.0.0-20200213135655-2678e485cc54
	github.com/PuerkitoBio/purell v1.1.1
	github.com/PuerkitoBio/urlesc v0.0.0-20170810143723-de5bf2ad4578 // indirect
	github.com/alexeyco/simpletable v0.0.0-20190222165044-2eb48bcee7cf
	github.com/armon/circbuf v0.0.0-20190214190532-5111143e8da2 // indirect
	github.com/armon/go-metrics v0.0.0-20190430140413-ec5e00d3c878 // indirect
	github.com/armon/go-radix v1.0.0 // indirect
	github.com/aws/aws-sdk-go v1.30.7
	github.com/fatih/color v1.7.0
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-log/log v0.2.0
	github.com/go-pascal/iban v0.0.0-20180529131734-f0d46003347e
	github.com/gogo/protobuf v1.3.1 // indirect
	github.com/golang/protobuf v1.4.0-rc.4.0.20200313231945-b860323f09d0
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.3.0
	github.com/gurukami/typ/v2 v2.0.1
	github.com/hashicorp/consul/api v1.1.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.1.0 // indirect
	github.com/hashicorp/go-msgpack v0.5.5 // indirect
	github.com/hashicorp/go-retryablehttp v0.5.4 // indirect
	github.com/hashicorp/go-rootcerts v1.0.1 // indirect
	github.com/hashicorp/go-sockaddr v1.0.2 // indirect
	github.com/hashicorp/go-version v1.2.0 // indirect
	github.com/hashicorp/mdns v1.0.1 // indirect
	github.com/hashicorp/memberlist v0.1.4 // indirect
	github.com/hashicorp/serf v0.8.3 // indirect
	github.com/karlseguin/expect v1.0.1 // indirect
	github.com/labstack/echo/v4 v4.1.11
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins/broker/rabbitmq v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/client/selector/static v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/registry/etcd v0.0.0-20200119172437-4fe21aa238fd // indirect
	github.com/micro/go-plugins/registry/kubernetes v0.0.0-20200119172437-4fe21aa238fd
	github.com/micro/go-plugins/transport/grpc v0.0.0-20200119172437-4fe21aa238fd
	github.com/mitchellh/gox v1.0.1 // indirect
	github.com/onsi/ginkgo v1.8.0 // indirect
	github.com/onsi/gomega v1.5.0 // indirect
	github.com/paysuper/echo-casbin-middleware v1.0.1-0.20200203133300-6f18edeb3072
	github.com/paysuper/paysuper-aws-manager v0.0.1
	github.com/paysuper/paysuper-proto/go/billingpb v0.0.0-20200415124355-3c76c5947b7d
	github.com/paysuper/paysuper-proto/go/casbinpb v0.0.0-20200203130641-45056764a1d7
	github.com/paysuper/paysuper-proto/go/recurringpb v0.0.0-20200302121221-60194355f1cc
	github.com/paysuper/paysuper-proto/go/reporterpb v0.0.0-20200218090503-e603ee3997a4
	github.com/paysuper/paysuper-proto/go/taxpb v0.0.0-20200203130641-45056764a1d7
	github.com/pkg/errors v0.9.1
	github.com/posener/complete v1.2.1 // indirect
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.5.1
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.0.1
	github.com/wsxiaoys/terminal v0.0.0-20160513160801-0940f3fc43a0 // indirect
	go.uber.org/automaxprocs v1.2.0
	golang.org/x/exp v0.0.0-20190627132806-fd42eb6b336f // indirect
	golang.org/x/image v0.0.0-20190703141733-d6a02ce849c9 // indirect
	golang.org/x/mobile v0.0.0-20190711165009-e47acb2ca7f9 // indirect
	golang.org/x/mod v0.1.0 // indirect
	gopkg.in/go-playground/validator.v9 v9.30.0
	gopkg.in/karlseguin/expect.v1 v1.0.1 // indirect
)

replace (
	github.com/go-playground/locales => github.com/go-playground/locales v0.12.1
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190927073244-c990c680b611
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.21.0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)

go 1.13
