module github.com/paysuper/paysuper-management-api

require (
	github.com/ProtocolONE/authone-jwt-verifier-golang v0.0.0-20190327070329-4dd563b01681
	github.com/ProtocolONE/geoip-service v0.0.0-20190903084234-1d5ae6b96679
	github.com/ProtocolONE/go-core/v2 v2.1.0
	github.com/PuerkitoBio/purell v1.1.1
	github.com/alexeyco/simpletable v0.0.0-20190222165044-2eb48bcee7cf
	github.com/aws/aws-sdk-go v1.23.16
	github.com/fatih/color v1.7.0
	github.com/forestgiant/sliceutil v0.0.0-20160425183142-94783f95db6c
	github.com/globalsign/mgo v0.0.0-20181015135952-eeefdecb41b8
	github.com/go-log/log v0.1.0
	github.com/go-pascal/iban v0.0.0-20180529131734-f0d46003347e
	github.com/golang/protobuf v1.3.2
	github.com/google/uuid v1.1.1
	github.com/google/wire v0.3.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0 // indirect
	github.com/gurukami/typ/v2 v2.0.1
	github.com/labstack/echo/v4 v4.1.11
	github.com/micro/go-micro v1.18.0
	github.com/micro/go-plugins v1.2.0
	github.com/paysuper/echo-casbin-middleware v0.0.0-20191201120027-e10b488ca709
	github.com/paysuper/paysuper-aws-manager v0.0.1
	github.com/paysuper/paysuper-proto/go/billingpb v0.0.0-20200122193247-3dba658a3cbc
	github.com/paysuper/paysuper-proto/go/casbinpb v0.0.0-20200122152034-6d6bb19abebf
	github.com/paysuper/paysuper-proto/go/recurringpb v0.0.0-20200122152034-6d6bb19abebf
	github.com/paysuper/paysuper-proto/go/reporterpb v0.0.0-20200211080654-935ed390e2e5
	github.com/paysuper/paysuper-proto/go/taxpb v0.0.0-20200122152034-6d6bb19abebf
	github.com/pkg/errors v0.8.1
	github.com/spf13/cobra v0.0.5
	github.com/stretchr/testify v1.4.0
	github.com/ttacon/builder v0.0.0-20170518171403-c099f663e1c2 // indirect
	github.com/ttacon/libphonenumber v1.0.1
	go.uber.org/automaxprocs v1.2.0
	gopkg.in/go-playground/validator.v9 v9.30.0
	gopkg.in/karlseguin/expect.v1 v1.0.1 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.0
	github.com/gogo/protobuf v0.0.0-20190410021324-65acae22fc9 => github.com/gogo/protobuf v1.2.2-0.20190723190241-65acae22fc9d
	github.com/hashicorp/consul => github.com/hashicorp/consul v1.5.1
	github.com/hashicorp/consul/api => github.com/hashicorp/consul/api v1.1.0
	github.com/lucas-clemente/quic-go => github.com/lucas-clemente/quic-go v0.12.0
	github.com/marten-seemann/qtls => github.com/marten-seemann/qtls v0.3.2
	github.com/micro/go-micro => github.com/micro/go-micro v1.8.0
	golang.org/x/sys => golang.org/x/sys v0.0.0-20190927073244-c990c680b611
	gopkg.in/urfave/cli.v1 => github.com/urfave/cli v1.21.0
	sourcegraph.com/sourcegraph/go-diff => github.com/sourcegraph/go-diff v0.5.1
)

go 1.13
