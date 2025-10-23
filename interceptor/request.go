package interceptor

import (
	"context"

	"time"

	"github.com/ginanjar-template-golang/shared-pkg/utils"
	"google.golang.org/grpc/metadata"

	"github.com/ginanjar-template-golang/shared-pkg/logger"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"google.golang.org/grpc"
)

func ExtractMetadata(ctx context.Context) (requestID, lang string) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return utils.NewRequestID(), "en"
	}

	// request_id
	if vals := md.Get("x-request-id"); len(vals) > 0 {
		requestID = vals[0]
	} else {
		requestID = utils.NewRequestID()
	}

	// language
	if langs := md.Get("accept-language"); len(langs) > 0 {
		lang = langs[0]
	} else {
		lang = "en"
	}

	return requestID, lang
}

func UnaryRequestLogger() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req any,
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp any, err error) {
		start := time.Now()

		reqID, lang := ExtractMetadata(ctx)
		utils.SetRequestID(reqID)
		translator.InitGlobalTranslator(lang)

		logger.Trace("gRPC Request Start", map[string]any{
			"request_id": reqID,
			"method":     info.FullMethod,
			"request":    req,
		})

		resp, err = handler(ctx, req)

		duration := time.Since(start)

		logger.Trace("gRPC Request End", map[string]any{
			"request_id":  reqID,
			"duration_ms": duration.Milliseconds(),
			"error":       err,
		})

		return resp, err
	}
}
