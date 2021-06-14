package controllers

import (
	"context"
	"github.com/david-drvar/xws2021-nistagram/common"
	"github.com/david-drvar/xws2021-nistagram/common/logger"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	"github.com/david-drvar/xws2021-nistagram/user_service/services"
	"gorm.io/gorm"
)

type VerificationGrpcController struct {
	service    *services.VerificationService
	jwtManager *common.JWTManager
	logger     *logger.Logger
}

func NewVerificationController(db *gorm.DB, jwtManager *common.JWTManager, logger *logger.Logger) (*VerificationGrpcController, error) {
	service, err := services.NewVerificationService(db)
	if err != nil {
		return nil, err
	}

	return &VerificationGrpcController{
		service,
		jwtManager,
		logger,
	}, nil
}

func (s *VerificationGrpcController) SubmitVerificationRequest(ctx context.Context, in *protopb.VerificationRequest) (*protopb.EmptyResponse, error) {
	span := tracer.StartSpanFromContextMetadata(ctx, "SubmitVerificationRequest")
	defer span.Finish()
	ctx = tracer.ContextWithSpan(context.Background(), span)

	s.logger.ToStdoutAndFile("CreateUser", "User registration attempt: ", logger.Info)

	//todo - napravi konvertor u domenski ili persistence, napisi u user additional info category, dekodiraj sliku, upisi verification request

	return &protopb.EmptyResponse{}, nil
}
