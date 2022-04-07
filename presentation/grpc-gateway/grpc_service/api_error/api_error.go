package api_error

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type ApiErrorType string

const (
	errorDomain                     = "pocket7878.go-ddd-learning"
	UnknownApiError    ApiErrorType = "ApiError::Unknown"
	FailedToCreateUser ApiErrorType = "ApiError::FailedToCreateUser"
)

func (e ApiErrorType) Error() string {
	return string(e)
}

func (e ApiErrorType) ToGrpcErrorDetail() *errdetails.ErrorInfo {
	return &errdetails.ErrorInfo{
		Domain: errorDomain,
		Metadata: map[string]string{
			"api_error_type": string(e),
		},
	}
}
