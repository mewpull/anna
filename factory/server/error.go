package factoryserver

import (
	"fmt"

	"github.com/juju/errgo"
)

var (
	maskAny = errgo.MaskFunc(errgo.Any)
)

func maskAnyf(cause error, f string, v ...interface{}) error {
	f = fmt.Sprintf("%s: %s", cause.Error(), f)
	err := maskAny(errgo.WithCausef(nil, cause, f, v...))

	if e, _ := err.(*errgo.Err); e != nil {
		e.SetLocation(1)
	}

	return err
}

var invalidFactoryGatewayRequestError = errgo.New("invalid factory gateway request")

func IsInvalidFactoryGatewayRequest(err error) bool {
	return errgo.Cause(err) == invalidFactoryGatewayRequestError
}
