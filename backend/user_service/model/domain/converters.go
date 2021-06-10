package domain

import (
	protopb "github.com/david-drvar/xws2021-nistagram/common/proto"
	"github.com/david-drvar/xws2021-nistagram/user_service/model"
	"github.com/david-drvar/xws2021-nistagram/user_service/model/persistence"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (l LoginRequest) ConvertFromGrpc(request *protopb.LoginRequest) LoginRequest{
	return LoginRequest{
		Email:    request.Email,
		Password: request.Password,
	}
}

func (p Password) ConvertFromGrpc(pass *protopb.Password) Password {
	return Password{
		OldPassword:      pass.OldPassword,
		NewPassword:      pass.NewPassword,
		RepeatedPassword: pass.NewPassword,
		Id:               pass.Id,
	}
}

func (u User) ConvertToGrpc() *protopb.UsersDTO {
	return &protopb.UsersDTO{
		Id:           u.Id,
		FirstName:    u.FirstName,
		LastName:     u.LastName,
		Email:        u.Email,
		Username:     u.Username,
		Role:         u.Role.String(),
		Birthdate:    timestamppb.New(u.BirthDate),
		ProfilePhoto: u.ProfilePhoto,
		PhoneNumber:  u.PhoneNumber,
		Sex:          u.Sex,
		IsActive:     u.IsActive,
		Website:      u.Website,
		Biography:    u.Biography,
		Category:     model.GetUserCategoriesString(u.Category),
		ResetCode:    u.ResetCode,
		ApprovedAccount: u.ApprovedAccount,
		TokenEnd: timestamppb.New(u.TokenEnd),
	}
}

func (u User) ConvertFromGrpc(user *protopb.UsersDTO) User {
	return User{
		Id:           user.Id,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Email:        user.Email,
		Username:     user.Username,
		Role:         model.GetUserRole(user.Role),
		BirthDate:    user.Birthdate.AsTime(),
		ProfilePhoto: user.ProfilePhoto,
		PhoneNumber:  user.PhoneNumber,
		Sex:          user.Sex,
		IsActive:     user.IsActive,
		Website:      user.Website,
		Biography:    user.Biography,
		Category:     model.GetUserCategories(user.Category),
		ResetCode:    user.ResetCode,
		ApprovedAccount: user.ApprovedAccount,
		TokenEnd:  user.TokenEnd.AsTime(),
	}
}

func (u *User) GenerateUserDTO(user persistence.User, userAdditionalInfo persistence.UserAdditionalInfo) {
	u.Id = user.Id
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Username = user.Username
	u.Role = user.Role
	u.BirthDate = user.BirthDate
	u.ProfilePhoto = user.ProfilePhoto
	u.PhoneNumber = user.PhoneNumber
	u.IsActive = user.IsActive
	u.Sex = user.Sex
	u.Biography = userAdditionalInfo.Biography
	u.Category = userAdditionalInfo.Category
	u.Website = userAdditionalInfo.Website
	u.ResetCode=user.ResetCode
	u.ApprovedAccount=user.ApprovedAccount
	u.TokenEnd=user.TokenEnd
}
