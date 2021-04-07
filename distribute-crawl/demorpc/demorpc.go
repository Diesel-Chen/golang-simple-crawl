package demorpc

import "errors"

type DemoRpcService struct {
}

type Args struct {
	A, B int
}

func (s DemoRpcService) Div(args Args, results *float64) error {
	if args.B == 0 {
		return errors.New("division by zero")
	}
	*results = float64(args.A) / float64(args.B)
	return nil
}
