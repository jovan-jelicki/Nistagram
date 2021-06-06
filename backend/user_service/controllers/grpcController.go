package controllers

import (
	"context"
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/common/tracer"
	otgo "github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"io"
)

type Server struct {
	protopb.UnimplementedUsersServer
	protopb.UnimplementedPrivacyServer
	userController    *UserGrpcController
	privacyController *PrivacyGrpcController
	tracer            otgo.Tracer
	closer            io.Closer
}

func NewServer(db *gorm.DB) (*Server, error) {
	newUserController, _ := NewUserController(db)
	newPrivacyController, _ := NewPrivacyController(db)
	tracer, closer := tracer.Init("userService")
	otgo.SetGlobalTracer(tracer)
	return &Server{
		userController:    newUserController,
		privacyController: newPrivacyController,
		tracer:            tracer,
		closer:            closer,
	}, nil
}

func (s *Server) GetTracer() otgo.Tracer {
	return s.tracer
}

func (s *Server) GetCloser() io.Closer {
	return s.closer
}

func (s *Server) CreateUser(ctx context.Context, in *protopb.CreateUserRequest) (*protopb.UsersDTO, error) {
	return s.userController.CreateUser(ctx, in)
}

func (s *Server) GetAllUsers(ctx context.Context, in *protopb.EmptyRequest) (*protopb.UsersResponse, error) {
	return s.userController.GetAllUsers(ctx, in)
}

func (s *Server) UpdateUserProfile(ctx context.Context, in *protopb.CreateUserDTORequest) (*protopb.EmptyResponse, error) {
	return s.userController.UpdateUserProfile(ctx, in)
}

func (s *Server) UpdateUserPassword(ctx context.Context, in *protopb.CreatePasswordRequest) (*protopb.EmptyResponse, error) {
	return s.userController.UpdateUserPassword(ctx, in)
}

func (s *Server) CreatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.CreatePrivacy(ctx, in)
}

func (s *Server) UpdatePrivacy(ctx context.Context, in *protopb.CreatePrivacyRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.UpdatePrivacy(ctx, in)
}

func (s *Server) BlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.BlockUser(ctx, in)
}

func (s *Server) UnBlockUser(ctx context.Context, in *protopb.CreateBlockRequest) (*protopb.EmptyResponsePrivacy, error) {
	return s.privacyController.UnBlockUser(ctx, in)
}

func (s *Server) SearchUser(ctx context.Context, in *protopb.SearchUserDtoRequest) (*protopb.UsersResponse, error) {
	return s.userController.SearchUser(ctx, in)
}

func (s *Server) CheckUserProfilePublic(ctx context.Context, in *protopb.PrivacyRequest) (*protopb.BooleanResponse, error) {
	return s.privacyController.CheckUserProfilePublic(ctx, in)
}
