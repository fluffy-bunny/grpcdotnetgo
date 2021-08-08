package error

import (
	"fmt"
	"reflect"

	"github.com/gogo/googleapis/google/rpc"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewBadRequest() *BadRequest {
	return &BadRequest{badRequest: &rpc.BadRequest{}}
}

type BadRequest struct {
	badRequest *rpc.BadRequest
}

func (br *BadRequest) AddViolation(field string, description string) {
	v := &rpc.BadRequest_FieldViolation{
		Field:       field,
		Description: description,
	}

	br.badRequest.FieldViolations = append(br.badRequest.FieldViolations, v)
}

func (br *BadRequest) GetDetails() *rpc.BadRequest {
	fmt.Println(reflect.TypeOf(br.badRequest))
	return br.badRequest
}

func (br *BadRequest) GetStatusError(c codes.Code, msg string) error {
	st := status.New(c, msg)
	if det, err := st.WithDetails(br.GetDetails()); err != nil {
		return st.Err()
	} else {
		return det.Err()
	}
}
