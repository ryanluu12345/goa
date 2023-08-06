package goa

import "context"

const (
	// MethodKey is the request context key used to store the name of the
	// method as defined in the design. The generated transport code
	// initializes the corresponding value prior to invoking the endpoint.
	MethodKey contextKey = iota + 1

	// ServiceKey is the request context key used to store the name of the
	// service as defined in the design. The generated transport code
	// initializes the corresponding value prior to invoking the endpoint.
	ServiceKey

	// PathPatternKey is the request context key used to store the name of the
	// path pattern. This value is used in a variety of middlewares for metrics
	// and tracing and is preferred to dynamic paths because path patterns don't
	// contain sensitive details or values that can cause metrics cardinality
	// to explode.
	PathPatternKey
)

type (
	// private type used to define context keys.
	contextKey int
)

// Endpoint exposes service methods to remote clients independently of the
// underlying transport.
type Endpoint func(ctx context.Context, request any) (response any, err error)
