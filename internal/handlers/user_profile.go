package handlers

import (
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	"github.com/paysuper/paysuper-proto/go/billingpb"
	"net/http"
)

const (
	userProfilePath             = "/user/profile"
	userCommonProfilePath       = "/user/profile/common"
	userProfilePathId           = "/user/profile/:id"
	userProfilePathFeedback     = "/user/feedback"
	userProfileConfirmEmailPath = "/user/confirm_email"
)

type UserProfileRoute struct {
	dispatch common.HandlerSet
	cfg      common.Config
	provider.LMT
}

func NewUserProfileRoute(set common.HandlerSet, cfg *common.Config) *UserProfileRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "UserProfileRoute"})
	return &UserProfileRoute{
		dispatch: set,
		LMT:      &set.AwareSet,
		cfg:      *cfg,
	}
}

func (h *UserProfileRoute) Route(groups *common.Groups) {
	groups.AuthProject.GET(userProfilePath, h.getUserProfile)
	groups.AuthProject.GET(userCommonProfilePath, h.getUserCommonProfile)
	groups.SystemUser.GET(userProfilePathId, h.getUserProfile)
	groups.AuthProject.PATCH(userProfilePath, h.setUserProfile)
	groups.AuthProject.POST(userProfilePathFeedback, h.createFeedback)
	groups.Common.PUT(userProfileConfirmEmailPath, h.confirmEmail)
}

// @summary Get the user profile
// @desc Get the user profile
// @id userProfilePathGetUserProfile
// @tag User Profile
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.UserProfile Returns the user's personal and company data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 403 {object} grpc.ResponseErrorMessage Access denied
// @failure 404 {object} grpc.ResponseErrorMessage The user not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/profile [get]

// @summary Get the user profile
// @desc Get the user profile
// @id userProfilePathIdGetUserProfile
// @tag User Profile
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.UserProfile Returns the user's personal and company data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 403 {object} grpc.ResponseErrorMessage Access denied
// @failure 404 {object} grpc.ResponseErrorMessage The user not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /system/api/v1/user/profile [get]
func (h *UserProfileRoute) getUserProfile(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &billingpb.GetUserProfileRequest{
		UserId:    authUser.Id,
		ProfileId: ctx.Param(common.RequestParameterId),
	}
	err := h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.GetUserProfile(ctx.Request().Context(), req)

	if err != nil {
		common.LogSrvCallFailedGRPC(h.L(), err, billingpb.ServiceName, "GetUserProfile", req)
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Get the common user profile
// @desc Get the user's main profile data, role, permissions
// @id userCommonProfilePathGetUserCommonProfile
// @tag User Profile
// @accept application/json
// @produce application/json
// @success 200 {object} grpc.CommonUserProfile Returns the user's main profile data, role and permissions
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 403 {object} grpc.ResponseErrorMessage Access denied
// @failure 404 {object} grpc.ResponseErrorMessage The user not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/profile/common [get]
func (h *UserProfileRoute) getUserCommonProfile(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &billingpb.CommonUserProfileRequest{UserId: authUser.Id}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	res, err := h.dispatch.Services.Billing.GetCommonUserProfile(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "GetCommonUserProfile")
	}

	if res.Status != billingpb.ResponseStatusOk {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Profile)
}

// @summary Create or update the user profile
// @desc Create or update the user profile
// @id userProfilePathSetUserProfile
// @tag User Profile
// @accept application/json
// @produce application/json
// @body grpc.UserProfile
// @success 200 {object} grpc.UserProfile Returns the user's personal and company data
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 403 {object} grpc.ResponseErrorMessage Access denied
// @failure 404 {object} grpc.ResponseErrorMessage The user not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/profile [patch]
func (h *UserProfileRoute) setUserProfile(ctx echo.Context) error {
	authUser := common.ExtractUserContext(ctx)
	req := &billingpb.UserProfile{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = authUser.Id
	req.Email = &billingpb.UserProfileEmail{
		Email: authUser.Email,
	}

	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreateOrUpdateUserProfile(ctx.Request().Context(), req)

	if err != nil {
		h.L().Error(common.InternalErrorTemplate, logger.WithFields(logger.Fields{"err": err.Error()}))
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.JSON(http.StatusOK, res.Item)
}

// @summary Confirm the user's email
// @desc Confirm the user's email using the user token
// @id userProfileConfirmEmailPathConfirmEmail
// @tag User Profile
// @accept application/json
// @produce application/json
// @body grpc.ConfirmUserEmailRequest
// @success 200 {string} Returns an empty response body if the user's email address has been successfully confirmed
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 404 {object} grpc.ResponseErrorMessage The user not found
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /api/v1/user/confirm_email [put]
func (h *UserProfileRoute) confirmEmail(ctx echo.Context) error {
	req := &billingpb.ConfirmUserEmailRequest{}

	if err := h.dispatch.BindAndValidate(req, ctx); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	res, err := h.dispatch.Services.Billing.ConfirmUserEmail(ctx.Request().Context(), req)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ConfirmUserEmail")
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	req2 := &billingpb.OnboardingRequest{
		User: &billingpb.MerchantUser{
			ProfileId:        res.Profile.Id,
			Id:               res.Profile.UserId,
			Email:            res.Profile.Email.Email,
			RegistrationDate: res.Profile.CreatedAt,
		},
	}

	if err = h.dispatch.Validate.Struct(req2); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res2, err := h.dispatch.Services.Billing.ChangeMerchant(ctx.Request().Context(), req2)

	if err != nil {
		return h.dispatch.SrvCallHandler(req, err, billingpb.ServiceName, "ChangeMerchant")
	}

	if res2.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}

// @summary Send the feedback
// @desc Create and send the feedback using the page URL
// @id userProfilePathFeedbackCreateFeedback
// @tag User Profile
// @accept application/json
// @produce application/json
// @body grpc.CreatePageReviewRequest
// @success 200 {string} Returns an empty response body if the feedback was successfully sent
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Internal Server Error
// @router /auth/api/v1/user/feedback [post]
func (h *UserProfileRoute) createFeedback(ctx echo.Context) error {

	authUser := common.ExtractUserContext(ctx)
	if authUser.Id == "" {
		return echo.NewHTTPError(http.StatusUnauthorized, common.ErrorMessageAccessDenied)
	}

	req := &billingpb.CreatePageReviewRequest{}
	err := ctx.Bind(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	req.UserId = authUser.Id
	err = h.dispatch.Validate.Struct(req)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, common.GetValidationError(err))
	}

	res, err := h.dispatch.Services.Billing.CreatePageReview(ctx.Request().Context(), req)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorUnknown)
	}

	if res.Status != http.StatusOK {
		return echo.NewHTTPError(int(res.Status), res.Message)
	}

	return ctx.NoContent(http.StatusOK)
}
