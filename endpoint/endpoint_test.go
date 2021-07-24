package endpoint

import (
	"context"
	"fmt"
	"testing"
)

/*TestChain test chain
=== RUN   TestChain
first pre
second pre
third pre
four pre
my endpoint!
four post
third post
second post
first post
--- PASS: TestChain (0.00s)
PASS
*/
func TestChain(t *testing.T) {
	e := Chain(
		annotate("first"),
		annotate("second"),
		annotate("third"),
		annotate("four"),
	)(myEndpoint)
	if _, err := e(ctx, req); err != nil {
		panic(err)
	}
}

var (
	ctx = context.Background()
	req = struct{}{}
)

func annotate(s string) Middleware {
	return func(next Endpoint) Endpoint {
		return func(ctx context.Context, request interface{}) (interface{}, error) {
			fmt.Println(s, "pre")
			defer fmt.Println(s, "post")

			return next(ctx, request)
		}
	}
}

func myEndpoint(context.Context, interface{}) (interface{}, error) {
	fmt.Println("my endpoint!")
	return struct{}{}, nil
}
