package pkg

import (
	"github.com/egon12/jsonapi"
	"github.com/labstack/echo/v4"
	"strconv"
)

func ErrorResponse(c echo.Context, httpStatus int, err error) error {
	return errorResponse(c, httpStatus, "invalid response", err)
}

func errorResponse(c echo.Context, httpStatusCode int, detail string, err error) error {
	errPayload := jsonapi.ErrorsPayload{
		Errors: []*jsonapi.ErrorObject{
			{
				Code:   "S00",
				Status: strconv.Itoa(httpStatusCode),
				Detail: err.Error(),
			},
		},
	}

	return c.JSON(httpStatusCode, errPayload)
}
