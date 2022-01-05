package xgqlgen

import (
	"fmt"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// GQLError :nodoc:
type GQLError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// ToGQLError self explained
func (e GQLError) ToGQLError() *gqlerror.Error {
	return &gqlerror.Error{
		Message: fmt.Sprintf("%s: %s", e.Code, e.Message),
		Extensions: map[string]interface{}{
			"code":    e.Code,
			"message": e.Message,
		},
	}
}
