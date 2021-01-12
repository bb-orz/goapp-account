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

type UsersDAO struct{}

func NewUsersDAO() *UsersDAO {
	dao := new(UsersDAO)
	return dao
}

func (d *UsersDAO) isExist(where *UsersModel) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&OAuthsModel{}).Where(where).First(&UsersModel{}).Count(&count).Error
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
func (d *UsersDAO) IsIdExist(id uint) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().First(&UsersModel{}, id).Count(&count).Error
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

// 查找用户名是否存在
func (d *UsersDAO) IsNameExist(name string) (bool, error) {
	return d.isExist(&UsersModel{Name: name})
}

// 查找邮箱是否存在
func (d *UsersDAO) IsEmailExist(email string) (bool, error) {
	return d.isExist(&UsersModel{Email: email})
}

// 查找手机号码是否存在
func (d *UsersDAO) IsPhoneExist(phone string) (bool, error) {

	return d.isExist(&UsersModel{Phone: phone})
}

// 通过Id查找
func (d *UsersDAO) GetById(id uint) (*dtos.UsersDTO, error) {
	var err error
	var usersResult UsersModel
	err = XGorm.XDB().Where(id).First(&usersResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := usersResult.ToDTO()
	return dto, nil
}

// 通过邮箱账号查找
func (d *UsersDAO) GetByEmail(email string) (*dtos.UsersDTO, error) {
	var err error
	var userResult UsersModel
	err = XGorm.XDB().Where(&UsersModel{Email: email}).First(&userResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}

	dto := userResult.ToDTO()
	return dto, nil
}

// 通过邮箱账号查找
func (d *UsersDAO) GetByPhone(phone string) (*dtos.UsersDTO, error) {
	var err error
	var userResult UsersModel
	err = XGorm.XDB().Where(&UsersModel{Phone: phone}).First(&userResult).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			// 无记录
			return nil, nil
		} else {
			// 除无记录外的错误返回
			return nil, err
		}
	}
	dto := userResult.ToDTO()
	return dto, nil
}

// 创建
func (d *UsersDAO) Create(dto *dtos.UsersDTO) (*dtos.UsersDTO, error) {
	var err error
	var usersDTO *dtos.UsersDTO
	var usersModel UsersModel

	usersModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&usersModel).Error; err != nil {
		return nil, err
	}
	usersDTO = usersModel.ToDTO()
	return usersDTO, nil
}

// 设置单个信息字段
func (d *UsersDAO) SetUsers(id uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&UsersModel{}).Where("id", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段
func (d *UsersDAO) UpdateUsers(id uint, dto dtos.UserInfoDTO) error {
	var err error

	if err = XGorm.XDB().Model(&UsersModel{}).Where("id", id).Updates(&dto).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *UsersDAO) SetPasswordAndSalt(uid uint, passHash, salt string) error {
	var err error
	if err = XGorm.XDB().Model(&UsersModel{}).Where("id", uid).UpdateColumns(&UsersModel{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *UsersDAO) DeleteById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&UsersModel{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *UsersDAO) SetDeletedAtById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&UsersModel{}).Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(id).Error; err != nil {
		return err
	}
	return nil
}