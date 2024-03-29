package grpcCodes

// Code for common errors
type Code uint32

// Transport codes used for common errors.
// For more info, check google.golang.org/grpc/codes
const (
	// OK ...
	OK Code = iota
	// Canceled ...
	Canceled
	// Unknown ...
	Unknown
	// InvalidArgument ...
	InvalidArgument
	// DeadlineExceeded ...
	DeadlineExceeded
	// NotFound ...
	NotFound
	// AlreadyExists ...
	AlreadyExists
	// PermissionDenied ...
	PermissionDenied
	// ResourceExhausted ...
	ResourceExhausted
	// FailedPrecondition ...
	FailedPrecondition
	// Aborted ...
	Aborted
	// OutOfRange ...
	OutOfRange
	// Unimplemented ...
	Unimplemented
	// Internal ...
	Internal
	// Unavailable ...
	Unavailable
	// DataLoss ...
	DataLoss
	// Unauthenticated ...
	Unauthenticated
)
