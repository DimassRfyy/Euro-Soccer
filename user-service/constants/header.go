package constants

import "net/textproto"

var (
	XServiceName = textproto.CanonicalMIMEHeaderKey("x-service-name")
	xApiKey 	= textproto.CanonicalMIMEHeaderKey("x-api-key")
	XRequestAt 	= textproto.CanonicalMIMEHeaderKey("x-request-at")
	Authorization = textproto.CanonicalMIMEHeaderKey("authorization")
)