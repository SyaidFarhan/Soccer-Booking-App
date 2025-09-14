package error

import "errors"

var (
	ErrInternalServerError = errors.New("internal server error")
	ErrSQLError            = errors.New("database server failed to respond")
	ErrTooManyRequests     = errors.New("too many requests, please try again later")
	ErrUnauthorized        = errors.New("unauthorized, please login again")
	ErrInvalidToken        = errors.New("invalid token, please login again")
	ErrForbidden           = errors.New("forbidden, you do not have permission to access this resource")
)

var GeneralErrors = []error{
	ErrInternalServerError,
	ErrSQLError,
	ErrTooManyRequests,
	ErrUnauthorized,
	ErrInvalidToken,
	ErrForbidden,
}
