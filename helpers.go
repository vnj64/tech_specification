package main

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"tech/domain"
)

type RawResponse struct {
	error      error
	status     int
	additional interface{}
	payload    interface{}
}

// byteCountSI converts a file size in bytes to a human-readable string
// representation with SI units (e.g. kB, MB, GB). The function rounds to
// one decimal place and uses base 1000 for unit conversion.
func byteCountSI(b int) string {
	const unit = 1000
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%cB", float64(b)/float64(div), "kMGTPE"[exp])
}
func (r *RawResponse) WithPayload(payload any) *RawResponse {
	r.payload = payload
	return r
}

func (r *RawResponse) WithAdditional(additional any) *RawResponse {
	r.additional = additional
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

// Response - ответ клиенту
type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message,omitempty"`
}

// ResponseBody - общее тело ответа клиенту
type ResponseBody struct {
	Response   `json:"response"`
	Additional interface{} `json:"additional,omitempty"`
	Payload    interface{} `json:"payload,omitempty"`
}

func (r *RawResponse) Error() error {
	return r.error
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
