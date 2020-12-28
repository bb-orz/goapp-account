package user

import (
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/jinzhu/gorm"
	"goapp-account/dtos"
)

/*
数据访问层，实现具体数据持久化操作
直接返回error和执行结果
*/

type userDAO struct{}

func NewUserDAO() *userDAO {
	dao := new(userDAO)
	return dao
}

// 查找用户名是否存在
func (d *userDAO) IsUserIdExist(uid uint) (bool, error) {
	var err error
	var count int64

	err = XGorm.XDB().Where(uid).First(UserModel{}).Count(&count).Error
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
func (d *userDAO) IsUserNameExist(name string) (bool, error) {
	return d.isExist(&UserModel{Name: name})
}

// 查找邮箱是否存在
func (d *userDAO) IsEmailExist(email string) (bool, error) {
	return d.isExist(&UserModel{Email: email})
}

// 查找手机号码是否存在
func (d *userDAO) IsPhoneExist(phone string) (bool, error) {

	return d.isExist(&UserModel{Phone: phone})
}

func (d *userDAO) isExist(where *UserModel) (bool, error) {
	var err error
	var count int64
	err = XGorm.XDB().Where(where).First(&UserModel{}).Count(&count).Error
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

// 插入单个用户信息
func (d *userDAO) Create(dto *dtos.UserDTO) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO
	var userModel UserModel

	userModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&userModel).Error; err != nil {
		return nil, err
	}
	userDTO = userModel.ToDTO()
	return userDTO, nil
}

// 插入单个用户信息并关联三方平台账户
func (d *userDAO) CreateUserWithOAuth(dto *dtos.UserOAuthsDTO) (*dtos.UserOAuthsDTO, error) {
	var err error
	var userOAuthsDTO *dtos.UserOAuthsDTO
	var userOAuthsModel UserOAuthsModel

	userOAuthsModel.FromDTO(dto)
	if err = XGorm.XDB().Create(&userOAuthsModel).Error; err != nil {
		return nil, err
	}

	userOAuthsDTO = userOAuthsModel.ToDTO()
	return userOAuthsDTO, nil
}

func (d *userDAO) GetUserOAuths(platform uint, openId, unionId string) (*dtos.UserOAuthsDTO, error) {
	var err error
	var oAuthResult OAuthModel
	var userResult UserModel
	var userOAuthDTO *dtos.UserOAuthsDTO
	var authDTOs []dtos.OAuthDTO

	if err = XGorm.XDB().Where(&OAuthModel{Platform: platform, OpenId: openId, UnionId: unionId}).Find(&oAuthResult).Error; err != nil {
		return nil, err
	}

	if err = XGorm.XDB().First(&userResult, oAuthResult.UserId).Error; err != nil {
		return nil, err
	}

	authDTOs = make([]dtos.OAuthDTO, 0)
	authDTOs = append(authDTOs, oAuthResult.ToDTO())

	userOAuthDTO = &dtos.UserOAuthsDTO{}
	userOAuthDTO.UserOAuths = authDTOs
	userOAuthDTO.User = *userResult.ToDTO()

	return userOAuthDTO, nil
}

// 通过Id查找
func (d *userDAO) GetById(id uint) (*dtos.UserDTO, error) {
	var err error
	var userResult UserModel
	err = XGorm.XDB().Where(id).First(&userResult).Error
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
func (d *userDAO) GetByEmail(email string) (*dtos.UserDTO, error) {
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
func (d *userDAO) GetByPhone(phone string) (*dtos.UserDTO, error) {
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

// 设置单个用户信息字段
func (d *userDAO) SetUserInfo(uid uint, field string, value interface{}) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Where("id", uid).Update(field, value).Error; err != nil {
		return err
	}
	return nil
}

// 设置多个用户信息字段
func (d *userDAO) SetUserInfos(uid uint, dto dtos.SetUserInfoDTO) error {
	var err error
	var updater UserModel
	updater.Name = dto.Name
	updater.Avatar = dto.Avatar
	updater.Age = dto.Age
	updater.Gender = dto.Gender
	updater.Status = dto.Status

	if err = XGorm.XDB().Model(&UserModel{}).Where("id", uid).Updates(&updater).Error; err != nil {
		return err
	}
	return nil
}

// 设置用户密码和盐值
func (d *userDAO) SetPasswordAndSalt(uid uint, passHash, salt string) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Where("id", uid).UpdateColumns(&UserModel{Password: passHash, Salt: salt}).Error; err != nil {
		return err
	}
	return nil
}

// 真删除
func (d *userDAO) DeleteById(uid uint) error {
	var err error
	if err = XGorm.XDB().Model(&UserModel{}).Delete(uid).Error; err != nil {
		return err
	}
	return nil
}

// 伪删除
func (d *userDAO) SetDeletedAtById(uid uint) error {
	var err error
	if err = XGorm.XDB().Set("gorm:delete_option", "OPTION (OPTIMIZE FOR UNKNOWN)").Delete(uid).Error; err != nil {
		return err
	}
	return nil
}
