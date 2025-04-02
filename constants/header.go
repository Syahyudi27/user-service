package constants

import "net/textproto"

var (
	XServiceName  = textproto.CanonicalMIMEHeaderKey("x-service-name")
	XApiKey       = textproto.CanonicalMIMEHeaderKey("api-key")
	RequestAt     = textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)
