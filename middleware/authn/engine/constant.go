package engine

const (
	// BearerWord the bearer key word for authorization
	BearerWord string = "Bearer"

	// BearerFormat authorization token format
	BearerFormat string = "Bearer %s"

	// AuthorizationKey holds the key used to store the token in the request header
	AuthorizationKey = "Authorization"

	// Reason holds the error reason.
	Reason = "UNAUTHORIZED"
)
