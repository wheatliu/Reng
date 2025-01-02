package twitter

import "errors"

var (
	ErrTimeout         = errors.New("timeout")
	ErrServerInternal  = errors.New("server internal error")
	ErrBadRequest      = errors.New("bad request")
	ErrTooManyRequests = errors.New("too many requests")
	ErrNotFound        = errors.New("not found")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")

	ErrUserSuspended          = errors.New("user account suspended")
	ErrUserPostsProtected     = errors.New("user posts protected")
	ErrInvalidTimeLineData    = errors.New("invalid timeline data")
	ErrInvalidInstruction     = errors.New("invalid instruction")
	ErrInvalidTimeLineEntries = errors.New("invalid timeline entries")
	ErrInvalidTweetEntry      = errors.New("invalid tweet entry")
	ErrTweetDeleted           = errors.New("tweet deleted")
)
