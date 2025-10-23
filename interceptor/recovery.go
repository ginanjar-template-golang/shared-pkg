package interceptor

import (
	"context"
	"fmt"

	appError "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	grpcResponse "github.com/ginanjar-template-golang/shared-pkg/response/grpc_response"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
	"google.golang.org/grpc"
)

func UnaryRecovery() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				reqID := utils.GetRequestID()

				logger.Error("🔥 Panic recovered in gRPC", map[string]any{
					"request_id": reqID,
					"panic":      r,
					"method":     info.FullMethod,
				})

				// Bungkus panic jadi AppError agar seragam
				appErr := appError.UnknownError("panicRecovered", fmt.Sprintf("%v", r))

				err = grpcResponse.FromAppError(ctx, appErr)
			}
		}()

		return handler(ctx, req)
	}
}
