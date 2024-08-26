package errorcode

type APIErrorCode uint

const (
	APIErrorCodeUndefined APIErrorCode = iota
	APIErrorCodeNoError
	APIErrorCodeUnauthenticated
)

type APIError struct {
	ErrorCode   APIErrorCode
	DetailError string
}

func (code APIErrorCode) Message() string {
	switch code {
	case APIErrorCodeUnauthenticated:
		return "unauthenticated"
	default:
		return "no error"
	}
}

func (code APIErrorCode) HttpCode() int {
	switch code {
	case APIErrorCodeUnauthenticated:
		return 401
	default:
		return 200
	}
}

func (err *APIError) Error() string {
	return err.ErrorCode.Message()
}
