package handlers

import (
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-management-api/internal/mock"
	"github.com/paysuper/paysuper-management-api/internal/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

type WebhookReportsTestSuite struct {
	suite.Suite
	router *WebHookRoute
	caller *test.EchoReqResCaller
}

func Test_Webhook(t *testing.T) {
	suite.Run(t, new(WebhookReportsTestSuite))
}

func (suite *WebhookReportsTestSuite) SetupTest() {
	var e error
	settings := test.DefaultSettings()
	srv := common.Services{
		Billing: mock.NewBillingServerOkMock(),
	}
	suite.caller, e = test.SetUp(settings, srv, func(set *test.TestSet, mw test.Middleware) common.Handlers {
		suite.router = NewWebHookRoute(set.HandlerSet, set.GlobalConfig)
		return common.Handlers{
			suite.router,
		}
	})
	if e != nil {
		panic(e)
	}
}

func (suite *WebhookReportsTestSuite) TestWebhook_sendTest_Ok() {
	json := `{"type":"simple","testing_case":"existing_user","user":{"external_id":"123"},"order_id":"2997247","amount":1,"currency":"EUR","project":"5dca80177ae51a00016d73a3"}`

	res, err := suite.caller.Builder().
		Method(http.MethodPost).
		Path(common.AuthUserGroupPath + testMerchantWebhook).
		Init(test.ReqInitJSON()).
		BodyString(json).
		Exec(suite.T())

	if assert.NoError(suite.T(), err) {
		assert.Equal(suite.T(), http.StatusOK, res.Code)
		assert.NotEmpty(suite.T(), res.Body.String())
	}
}
