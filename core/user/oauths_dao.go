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

type OauthsDAO struct{}

func NewOauthsDAO() *OauthsDAO {
	dao := new(OauthsDAO)
	return dao
}

func (d *OauthsDAO) isExist(where *OAuthsModel) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&OAuthsModel{}).Where(where).Count(&count).Error
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
func (d *OauthsDAO) IsIdExist(id uint) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().First(&OAuthsModel{}, id).Count(&count).Error
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
func (d *OauthsDAO) IsWechatAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthsModel{OpenId: openId, UnionId: unionId, Platform: WechatOauthPlatform})
}

// 查找是否绑定微信账号
func (d *OauthsDAO) IsWeiboAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthsModel{OpenId: openId, UnionId: unionId, Platform: WeiboOauthPlatform})
}

// 查找是否绑定微信账号
func (d *OauthsDAO) IsQQAccountBindng(openId, unionId string) (bool, error) {
	return d.isExist(&OAuthsModel{OpenId: openId, UnionId: unionId, Platform: QQOauthPlatform})
}

// 通过Id查找
func (d *OauthsDAO) GetById(id uint) (*dtos.OauthsDTO, error) {
	var err error
	var oauthsResult OAuthsModel
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
func (d *OauthsDAO) Create(dto *dtos.OauthsDTO) (*dtos.OauthsDTO, error) {
	var err error
	var oauthsDTO *dtos.OauthsDTO
	var oauthsModel OAuthsModel

	oauthsModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&oauthsModel).Error; err != nil {
		return nil, err
	}
	oauthsDTO = oauthsModel.ToDTO()
	return oauthsDTO, nil
}

// 设置单个信息字段
func (d *OauthsDAO) SetOauths(id uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthsModel{}).Where("id", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段
func (d *OauthsDAO) UpdateOauths(id uint, dto dtos.OauthsDTO) error {
	var err error

	if err = XGorm.XDB().Model(&OAuthsModel{}).Where("id", id).Updates(&dto).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *OauthsDAO) DeleteById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthsModel{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *OauthsDAO) SetDeletedAtById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&OAuthsModel{}).Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(id).Error; err != nil {
		return err
	}
	return nil
}
