package handlers

import (
	"fmt"
	"github.com/ProtocolONE/go-core/v2/pkg/logger"
	"github.com/ProtocolONE/go-core/v2/pkg/provider"
	"github.com/labstack/echo/v4"
	awsWrapper "github.com/paysuper/paysuper-aws-manager"
	"github.com/paysuper/paysuper-management-api/internal/dispatcher/common"
	grpc "github.com/paysuper/paysuper-proto/go/billingpb"
	reporterProto "github.com/paysuper/paysuper-proto/go/reporterpb"
	reporter "github.com/paysuper/paysuper-proto/go/reporterpb"
	"net/http"
	"os"
	"strings"
)

const (
	reportFileDownloadPath = "/report_file/download/:file"
)

type ReportFileRoute struct {
	dispatch   common.HandlerSet
	awsManager awsWrapper.AwsManagerInterface
	cfg        common.Config
	provider.LMT
}

func NewReportFileRoute(set common.HandlerSet, awsManager awsWrapper.AwsManagerInterface, cfg *common.Config) *ReportFileRoute {
	set.AwareSet.Logger = set.AwareSet.Logger.WithFields(logger.Fields{"router": "ReportFileRoute"})
	return &ReportFileRoute{
		dispatch:   set,
		LMT:        &set.AwareSet,
		cfg:        *cfg,
		awsManager: awsManager,
	}
}

func (h *ReportFileRoute) Route(groups *common.Groups) {
	groups.AuthUser.GET(reportFileDownloadPath, h.download)
	groups.AuthProject.GET(reportFileDownloadPath, h.download)
}

// @summary Export the report file
// @desc Export the report file into a PDF, CSV, XLSX
// @id reportFileDownloadPathDownload
// @tag Report file
// @accept application/json
// @produce application/pdf, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @success 200 {string} Returns the report file
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data (unable to find the file, the file string is incorrect)
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Unable to download the file because of the internal server error
// @param file_id path {string} true The unique identifier for the report file.
// @param file_type path {string} true The supported file format (PDF, CSV, XLSX).
// @router /auth/api/v1/report_file/download/{file_id}.{file_type} [get]

// @summary Export the report file
// @desc Export the report file into a PDF, CSV, XLSX
// @id reportFileDownloadPathDownload
// @tag Report file
// @accept application/json
// @produce application/pdf, text/csv, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet
// @success 200 {string} Returns the report file
// @failure 400 {object} grpc.ResponseErrorMessage Invalid request data (unable to find the file, the file string is incorrect)
// @failure 401 {object} grpc.ResponseErrorMessage Unauthorized request
// @failure 500 {object} grpc.ResponseErrorMessage Unable to download the file because of the internal server error
// @param file_id path {string} true The unique identifier for the report file.
// @param file_type path {string} true The supported file format (PDF, CSV, XLSX).
// @router /admin/api/v1/report_file/download/{file_id}.{file_type} [get]
func (h *ReportFileRoute) download(ctx echo.Context) error {
	file := ctx.Param("file")

	if file == "" {
		h.L().Error("unable to find the file")
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	params := strings.Split(file, ".")

	if len(params) != 2 {
		h.L().Error("incorrect of file string")
		return echo.NewHTTPError(http.StatusBadRequest, common.ErrorRequestParamsIncorrect)
	}

	fileName := fmt.Sprintf(reporterProto.FileMask, common.ExtractUserContext(ctx).Id, params[0], params[1])
	filePath := os.TempDir() + string(os.PathSeparator) + fileName
	_, err := h.awsManager.Download(ctx.Request().Context(), filePath, &awsWrapper.DownloadInput{FileName: fileName})

	if err != nil {
		h.L().Error("unable to download the file " + fileName + " with message: " + err.Error())
		return echo.NewHTTPError(http.StatusInternalServerError, common.ErrorMessageDownloadReportFile)
	}

	defer os.Remove(filePath)
	return ctx.File(filePath)
}
