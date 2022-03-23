package imparse

import "context"

type Trace interface {
	StartSpanFromContext(ctx context.Context, funcName string) (func(), context.Context)
}
