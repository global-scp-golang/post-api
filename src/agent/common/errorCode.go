package common

const (
	StatusBadRequest                    = 400 // RFC 9110, 15.5.1
	StatusUnauthorized                  = 401 // RFC 9110, 15.5.2
	StatusPaymentRequired               = 402 // RFC 9110, 15.5.3
	StatusForbidden                     = 403 // RFC 9110, 15.5.4
	StatusNotFound                      = 404 // RFC 9110, 15.5.5
	StatusMethodNotAllowed              = 405 // RFC 9110, 15.5.6
	StatusNotAcceptable                 = 406 // RFC 9110, 15.5.7
	StatusProxyAuthRequired             = 407 // RFC 9110, 15.5.8
	StatusRequestTimeout                = 408 // RFC 9110, 15.5.9
	StatusConflict                      = 409 // RFC 9110, 15.5.10
	StatusGone                          = 410 // RFC 9110, 15.5.11
	StatusLengthRequired                = 411 // RFC 9110, 15.5.12
	StatusPreconditionFailed            = 412 // RFC 9110, 15.5.13
	StatusRequestEntityTooLarge         = 413 // RFC 9110, 15.5.14
	StatusRequestURITooLong             = 414 // RFC 9110, 15.5.15
	StatusUnsupportedMediaType          = 415 // RFC 9110, 15.5.16
	StatusRequestedRangeNotSatisfiable  = 416 // RFC 9110, 15.5.17
	StatusExpectationFailed             = 417 // RFC 9110, 15.5.18
	StatusTeapot                        = 418 // RFC 9110, 15.5.19 (Unused)
	StatusMisdirectedRequest            = 421 // RFC 9110, 15.5.20
	StatusUnprocessableEntity           = 422 // RFC 9110, 15.5.21
	StatusLocked                        = 423 // RFC 4918, 11.3
	StatusFailedDependency              = 424 // RFC 4918, 11.4
	StatusTooEarly                      = 425 // RFC 8470, 5.2.
	StatusUpgradeRequired               = 426 // RFC 9110, 15.5.22
	StatusPreconditionRequired          = 428 // RFC 6585, 3
	StatusTooManyRequests               = 429 // RFC 6585, 4
	StatusRequestHeaderFieldsTooLarge   = 431 // RFC 6585, 5
	StatusUnavailableForLegalReasons    = 451 // RFC 7725, 3
	StatusInternalServerError           = 500 // RFC 9110, 15.6.1
	StatusNotImplemented                = 501 // RFC 9110, 15.6.2
	StatusBadGateway                    = 502 // RFC 9110, 15.6.3
	StatusServiceUnavailable            = 503 // RFC 9110, 15.6.4
	StatusGatewayTimeout                = 504 // RFC 9110, 15.6.5
	StatusHTTPVersionNotSupported       = 505 // RFC 9110, 15.6.6
	StatusVariantAlsoNegotiates         = 506 // RFC 2295, 8.1
	StatusInsufficientStorage           = 507 // RFC 4918, 11.5
	StatusLoopDetected                  = 508 // RFC 5842, 7.2
	StatusNotExtended                   = 510 // RFC 2774, 7
	StatusNetworkAuthenticationRequired = 511 // RFC 6585, 6
)

type CustomError struct {
	Error string `json:"error"`
}
