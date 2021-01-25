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

type UserDAO struct{}

func NewUserDAO() *UserDAO {
	dao := new(UserDAO)
	return dao
}

func (d *UserDAO) isExist(where *UserModel) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&UserModel{}).Where(where).Count(&count).Error
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
func (d *UserDAO) IsNameExist(name string) (bool, error) {
	return d.isExist(&UserModel{Name: name})
}

// 查找邮箱是否存在
func (d *UserDAO) IsEmailExist(email string) (bool, error) {
	return d.isExist(&UserModel{Email: email})
}

// 查找手机号码是否存在
func (d *UserDAO) IsPhoneExist(phone string) (bool, error) {
	return d.isExist(&UserModel{Phone: phone})
}

// 查找id是否存在
func (d *UserDAO) IsIdExist(id uint) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&UserModel{}).Where("id = ?", id).Count(&count).Error
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

// 查找邮箱是否存在
func (d *UserDAO) IsEmailBinding(id uint, email string) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&UserModel{}).Where("id = ? AND email = ? ", id, email).Count(&count).Error
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

// 查找手机号码是否存在
func (d *UserDAO) IsPhoneBinding(id uint, phone string) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Model(&UserModel{}).Where("id = ? AND phone = ? ", id, phone).Count(&count).Error
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

// 通过Id查找
func (d *UserDAO) GetById(id uint) (*dtos.UserDTO, error) {
	var err error
	var userResult UserModel
	err = XGorm.XDB().Model(&UserModel{}).Where("id = ?", id).First(&userResult).Error
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
func (d *UserDAO) GetByEmail(email string) (*dtos.UserDTO, error) {
	var err error
	var userResult UserModel
	err = XGorm.XDB().Where(&UserModel{Email: email}).First(&userResult).Error
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
func (d *UserDAO) GetByPhone(phone string) (*dtos.UserDTO, error) {
	var err error
	var userResult UserModel
	err = XGorm.XDB().Where(&UserModel{Phone: phone}).First(&userResult).Error
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
func (d *UserDAO) Create(dto *dtos.UserDTO) (int64, error) {
	var err error
	var userDTO *dtos.UserDTO
	var userModel UserModel

	userModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&userModel).Error; err != nil {
		return -1, err
	}
	userDTO = userModel.ToDTO()
	return int64(userDTO.Id), nil
}

// 设置单个信息字段
func (d *UserDAO) SetUser(id uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Where("id", id).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个信息字段
func (d *UserDAO) UpdateUser(id uint, dto dtos.UserInfoDTO) error {
	var err error
	var userModel UserModel
	userModel.FromInfoDTO(&dto)
	if err = XGorm.XDB().Model(&UserModel{}).Where("id", id).Updates(&userModel).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *UserDAO) SetPasswordAndSaltById(id uint, passHash, salt string) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Where("id", id).UpdateColumns(&UserModel{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *UserDAO) SetPasswordAndSaltByEmail(email string, passHash, salt string) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Where("email", email).UpdateColumns(&UserModel{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *UserDAO) DeleteById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Delete(id).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *UserDAO) SetDeletedAtById(id uint) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(id).Error; err != nil {
		return err
	}
	return nil
}
