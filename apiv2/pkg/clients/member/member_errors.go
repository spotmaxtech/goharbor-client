package member

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/spotmaxtech/goharbor-client/v5/apiv2/internal/api/client/member"
	"github.com/spotmaxtech/goharbor-client/v5/apiv2/pkg/errors"
)

// handleSwaggerProjectErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with a proper message.
func handleSwaggerMemberErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusCreated:
			// Harbor sometimes return 201 instead of 200 despite the swagger spec
			// not declaring it.
			return nil
		case http.StatusBadRequest:
			return &errors.ErrBadRequest{}
		case http.StatusUnauthorized:
			return &errors.ErrUnauthorized{}
		case http.StatusForbidden:
			return &errors.ErrForbidden{}
		case http.StatusNotFound:
			return &errors.ErrNotFound{}
		}
	}

	switch in.(type) {
	default:
		return in
	case *member.CreateProjectMemberConflict:
		return &errors.ErrMemberAlreadyExists{}
	}
}
