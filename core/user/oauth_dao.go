package user

import (
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"goapp/dtos"
	"gorm.io/gorm"
)

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type OAuthDAO struct{}

func NewOAuthDAO() *OAuthDAO {
	dao := new(OAuthDAO)
	return dao
}

func (d *OAuthDAO) isExist(where *OAuthModel) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&OAuthModel{}).Where(where).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return false, nil
		} else {
			// 除无记录外的错误返回
			return false, err
		}
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 查找id是否存在
func (d *OAuthDAO) IsIdExist(id uint) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().First(&OAuthModel{}, id).Count(&count).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return false, nil
		} else {
			// 除无记录外的错误返回
			return false, err
		}
	}

	if count > 0 {
		return true, nil
	}
	return false, nil
}

// 查找是否绑定微信账号
func (d *OAuthDAO) IsWechatAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthModel{OpenId: openId, UnionId: unionId, Platform: WechatOAuthPlatform})
}

// 查找是否绑定微信账号
func (d *OAuthDAO) IsWeiboAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthModel{OpenId: openId, UnionId: unionId, Platform: WeiboOAuthPlatform})
}

// 查找是否绑定微信账号
func (d *OAuthDAO) IsQQAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthModel{OpenId: openId, UnionId: unionId, Platform: QQOAuthPlatform})
}

// 通过Id查找
func (d *OAuthDAO) GetById(id uint) (*dtos.OAuthDTO, error) {
	var err error
	var oauthsResult OAuthModel
	err = XGorm.XDB().Where(id).First(&oauthsResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := oauthsResult.ToDTO()
	return dto, nil
}

// 创建
func (d *OAuthDAO) Create(dto *dtos.OAuthDTO) (int64, error) {
	var err error
	var oauthsDTO *dtos.OAuthDTO
	var oauthsModel OAuthModel

	oauthsModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&oauthsModel).Error; err != nil {
		return -1, err
	}
	oauthsDTO = oauthsModel.ToDTO()
	return int64(oauthsDTO.Id), nil
}

// 设置单个信息字段
func (d *OAuthDAO) SetOAuth(id uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthModel{}).Where("id", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段
func (d *OAuthDAO) UpdateOAuth(id uint, dto dtos.OAuthDTO) error {
	var err error

	if err = XGorm.XDB().Model(&OAuthModel{}).Where("id", id).Updates(&dto).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *OAuthDAO) DeleteById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthModel{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *OAuthDAO) SetDeletedAtById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthModel{}).Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(id).Error; err != nil {
		return err
	}
	return nil
}
