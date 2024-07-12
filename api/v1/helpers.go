package v1

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"tech/domain"
)

type RawResponse struct {
	error      error
	status     int
	additional interface{}
	payload    interface{}
}

func (r *RawResponse) Error() error {
	return r.error
}

func (r *RawResponse) WithPayload(payload any) *RawResponse {
	r.payload = payload
	return r
}

func (r *RawResponse) Body() *ResponseBody {
	return &ResponseBody{
		Response: Response{
			Status: r.status,
		},
		Additional: r.additional,
		Payload:    r.payload,
	}
}

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

type ResponseBody struct {
	Response   `json:"response"`
	Additional interface{} `json:"additional,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}

func Wrap(handler func(c domain.Context, g *gin.Context) *RawResponse) func(c *gin.Context) {
	return func(c *gin.Context) {
		//var respBody ResponseBody

		ctx, ok := c.MustGet("context").(domain.Context)
		if !ok {
			return
		}

		response := handler(ctx, c)
		body := response.Body()

		status := body.Status

		if err := response.Error(); err != nil {
			var domainErr *domain.Error
			if errors.As(err, &domainErr) {
				body.Message = domainErr.Message(true)

				if domainErr.HttpCode() > 0 {
					status = domainErr.HttpCode()
					body.Status = domainErr.HttpCode()
				}

				if domainErr.ExtraCode() > 0 {
					body.Status = domainErr.ExtraCode()
				}
			}
		}

		c.AbortWithStatusJSON(status, body)
	}
}

func BadRequest(err error) *RawResponse {
	return &RawResponse{
		status: http.StatusBadRequest,
		error:  err,
	}
}

func InternalServerError(err error) *RawResponse {
	return &RawResponse{
		status: http.StatusInternalServerError,
		error:  err,
	}
}

func OK(payload any) *RawResponse {
	out := &RawResponse{
		status: http.StatusOK,
	}

	if payload != nil {
		out.WithPayload(payload)
	}

	return out
}
