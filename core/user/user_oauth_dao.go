package user

import (
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"goapp/dtos"
)

type UserOAuthDAO struct{}

func NewUserOAuthDAO() *UserOAuthDAO {
	dao := new(UserOAuthDAO)
	return dao
}

// 插入单个用户信息并关联三方平台账户
func (d *UserOAuthDAO) CreateUserWithOAuth(dto *dtos.UserOAuthInfoDTO) (*dtos.UserOAuthInfoDTO, error) {
	var err error
	var result *dtos.UserOAuthInfoDTO
	var model UserOAuthModel

	model.FromDTO(dto)
	if err = XGorm.XDB().Create(&model).Error; err != nil {
		return nil, err
	}
	result = model.ToDTO()
	return result, nil
}

func (d *UserOAuthDAO) GetUserOAuths(platform uint, openId, unionId string) (*dtos.UserOAuthInfoDTO, error) {
	var err error
	var result *dtos.UserOAuthInfoDTO
	var oauthRs OAuthsModel
	var userRs UsersModel

	if err = XGorm.XDB().Model(OAuthsModel{}).Where(OAuthsModel{Platform: platform, OpenId: openId, UnionId: unionId}).First(&oauthRs).Error; err != nil {
		return nil, err
	}

	if err = XGorm.XDB().Model(UsersModel{}).First(&userRs, oauthRs.UserId).Error; err != nil {
		return nil, err
	}

	result = &dtos.UserOAuthInfoDTO{
		Id:            userRs.ID,
		No:            userRs.No,
		Name:          userRs.Name,
		Age:           userRs.Age,
		Gender:        userRs.Gender,
		Avatar:        userRs.Avatar,
		Email:         userRs.Email,
		EmailVerified: userRs.EmailVerified,
		Phone:         userRs.Phone,
		PhoneVerified: userRs.PhoneVerified,
		Status:        userRs.Status,
		CreatedAt:     userRs.CreatedAt,
		UpdatedAt:     userRs.UpdatedAt,
		DeletedAt:     userRs.DeletedAt.Time,
		OAuths:        []dtos.OauthsDTO{*oauthRs.ToDTO()},
	}

	return result, nil
}
