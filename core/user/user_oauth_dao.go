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
	var userOAuths UserOAuthModel

	if err = XGorm.XDB().Where(OAuthModel{Platform: platform, OpenId: openId, UnionId: unionId}).Preload("OAuths").First(&userOAuths).Error; err != nil {
		return nil, err
	}

	result = userOAuths.ToDTO()

	return result, nil
}
