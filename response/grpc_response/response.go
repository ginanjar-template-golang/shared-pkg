package grpc_response

import (
	"context"
	"encoding/json"
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errHandler "github.com/ginanjar-template-golang/shared-pkg/errors"
	"github.com/ginanjar-template-golang/shared-pkg/logger"
	responsepb "github.com/ginanjar-template-golang/shared-pkg/proto/pb/responsepb"
	"github.com/ginanjar-template-golang/shared-pkg/translator"
	"github.com/ginanjar-template-golang/shared-pkg/utils"
)

// =========================
// SUCCESS RESPONSES
// =========================

// Success → 200 OK
func Success(ctx context.Context, messageKey string, data any) (*responsepb.StandardResponse, error) {
	return baseResponse(ctx, 200, codes.OK, messageKey, data)
}

// Created → 201 Created
func Created(ctx context.Context, messageKey string, data any) (*responsepb.StandardResponse, error) {
	return baseResponse(ctx, 201, codes.OK, messageKey, data)
}

// Updated → 200 OK
func Updated(ctx context.Context, messageKey string, data any) (*responsepb.StandardResponse, error) {
	return baseResponse(ctx, 200, codes.OK, messageKey, data)
}

// Deleted → 200 OK (tanpa data)
func Deleted(ctx context.Context, messageKey string) (*responsepb.StandardResponse, error) {
	return baseResponse(ctx, 200, codes.OK, messageKey, nil)
}

// baseResponse helper untuk response sukses standar
func baseResponse(_ context.Context, httpCode int32, grpcCode codes.Code, messageKey string, data any) (*responsepb.StandardResponse, error) {
	reqID := utils.GetRequestID()
	message := translator.GetMessageGlobal(messageKey)

	var jsonBytes []byte
	if data != nil {
		jsonBytes, _ = json.Marshal(data)
	}

	decode, _ := utils.DecodeBytesToJSON(jsonBytes)

	logger.Info("Success", map[string]any{
		"request_id": reqID,
		"http_code":  httpCode,
		"grpc_code":  grpcCode,
		"message":    message,
		"results":    decode,
	})

	return &responsepb.StandardResponse{
		Meta: &responsepb.Meta{
			RequestId: reqID,
			HttpCode:  httpCode,
			Message:   message,
		},
		Results: jsonBytes,
	}, nil
}

// =========================
// PAGINATION SUCCESS
// =========================

type PaginationData struct {
	Page     int32
	Size     int32
	Limit    int32
	TotalRow int32
	Results  any `json:"results,omitempty"`
}

func PaginationSuccess(ctx context.Context, messageKey string, pagination PaginationData) (*responsepb.PaginationResponse, error) {
	reqID := utils.GetRequestID()
	message := translator.GetMessageGlobal(messageKey)

	resultsJSON, _ := json.Marshal(pagination.Results)
	decode, _ := utils.DecodeBytesToJSON(resultsJSON)

	logger.Info("GRPC Pagination Success", map[string]any{
		"request_id": reqID,
		"page":       pagination.Page,
		"limit":      pagination.Limit,
		"total_row":  pagination.TotalRow,
		"message":    message,
		"results":    decode,
	})

	return &responsepb.PaginationResponse{
		Meta: &responsepb.Meta{
			RequestId: reqID,
			HttpCode:  200,
			Message:   message,
		},
		Pagination: &responsepb.Pagination{
			Page:     pagination.Page,
			Size:     pagination.Size,
			Limit:    pagination.Limit,
			TotalRow: pagination.TotalRow,
		},
		Results: resultsJSON,
	}, nil
}

// =========================
// ERROR RESPONSE
// =========================

func FromAppError(ctx context.Context, err error) error {
	reqID := utils.GetRequestID()

	var appErr errHandler.AppError
	switch e := err.(type) {
	case errHandler.AppError:
		appErr = e
	case *errHandler.AppError:
		appErr = *e
	default:
		msg := fmt.Sprintf("[%s] Unexpected error: %v", reqID, err)
		return status.Error(codes.Internal, msg)
	}

	msg := translator.GetMessageGlobal(appErr.MessageKey)
	fullMsg := fmt.Sprintf("[%s] %s", reqID, msg)

	return status.Error(appErr.GrpcCode, fullMsg)
}
