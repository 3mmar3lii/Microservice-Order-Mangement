package grpcutil

import (
	"net/http"

	"github.com/ammar-alaa/common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GRPCError(w http.ResponseWriter, err error) error {
	if err != nil {
		rStatus := status.Convert(err)
		if rStatus.Code() == codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return err
		}
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return err
	}
	return nil
}