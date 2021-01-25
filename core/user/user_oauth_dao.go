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
func (d *UserOAuthDAO) CreateUserWithOAuth(dto *dtos.UserOAuthInfoDTO) (int64, error) {
	var err error
	var result *dtos.UserOAuthInfoDTO
	var model UserOAuthModel

	model.FromDTO(dto)
	if err = XGorm.XDB().Create(&model).Error; err != nil {
		return -1, err
	}
	result = model.ToDTO()
	return int64(result.Id), nil
}

func (d *UserOAuthDAO) GetUserOAuth(platform uint, openId, unionId string) (*dtos.UserOAuthInfoDTO, error) {
	var err error
	var result *dtos.UserOAuthInfoDTO
	var oauthRs OAuthModel
	var userRs UserModel

	if err = XGorm.XDB().Model(OAuthModel{}).Where(OAuthModel{Platform: platform, OpenId: openId, UnionId: unionId}).First(&oauthRs).Error; err != nil {
		return nil, err
	}

	if err = XGorm.XDB().Model(UserModel{}).First(&userRs, oauthRs.UserId).Error; err != nil {
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
		OAuth:         []dtos.OAuthDTO{*oauthRs.ToDTO()},
	}

	return result, nil
}
