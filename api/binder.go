package api

import (
	"bytes"
	"errors"
	"github.com/ProtocolONE/p1pay.api/database/model"
	"github.com/labstack/echo"
	"io/ioutil"
)

const (
	binderErrorQueryParamsIsEmpty = "required params not found"
	binderErrorPeriodIsRequire = "period is required"
)

type OrderFormBinder struct{}

type OrderJsonBinder struct {}

type OrderRevenueDynamicRequestBinder struct {}

func (cb *OrderFormBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	db := new(echo.DefaultBinder)

	if err = db.Bind(i, ctx); err != nil {
		return err
	}

	params, err := ctx.FormParams()
	addParams := make(map[string]interface{})
	rawParams := make(map[string]string)

	if err != nil {
		return err
	}

	o := i.(*model.OrderScalar)

	for key, value := range params {
		if _, ok := model.OrderReservedWords[key]; !ok {
			addParams[key] = value[0]
		}

		rawParams[key] = value[0]
	}

	o.Other = addParams
	o.RawRequestParams = rawParams

	return
}

func (cb *OrderJsonBinder) Bind(i interface{}, ctx echo.Context) (err error) {
	var buf []byte

	if ctx.Request().Body != nil {
		buf, err = ioutil.ReadAll(ctx.Request().Body)
		rdr := ioutil.NopCloser(bytes.NewBuffer(buf))

		if err != nil {
			return err
		}

		ctx.Request().Body = rdr
	}

	db := new(echo.DefaultBinder)

	if err = db.Bind(i, ctx); err != nil {
		return err
	}

	i.(*model.OrderScalar).RawRequestBody = string(buf)

	return
}

func (cb *OrderRevenueDynamicRequestBinder) Bind(i interface{}, ctx echo.Context) error {
	period := ctx.Param("period")

	if period == "" {
		return errors.New(binderErrorPeriodIsRequire)
	}

	params := ctx.QueryParams()

	if len(params) <= 0 {
		return errors.New(binderErrorQueryParamsIsEmpty)
	}

	return nil
}