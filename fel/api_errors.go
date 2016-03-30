package fel

var (
	ErrBadRequest = Err{"bad_request", 400,
		"Bad request",
		"Request body is not well-formed. It must be JSON."}

	ErrNotAcceptable = Err{"not_acceptable", 406,
		"Not Acceptable",
		"Accept header must be set to 'application/json'."}

	ErrUnsupportedMediaType = Err{"unsupported_media_type", 415,
		"Unsupported Media Type",
		"Content-Type header must be set to: 'application/json'."}

	ErrInternalServer = Err{"internal_server_error",
		500, "Internal Server Error", "Something went wrong."}
)
