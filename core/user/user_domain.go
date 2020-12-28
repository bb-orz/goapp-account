package user

import (
	"github.com/bb-orz/goinfras/XGlobal"
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/segmentio/ksuid"
	"goapp-account/common"
	"goapp-account/dtos"
)

/*
User 领域层：实现用户相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type UserDomain struct {
	dao   *userDAO
	cache *userCache
}

func NewUserDomain() *UserDomain {
	domain := new(UserDomain)
	domain.dao = NewUserDAO()
	domain.cache = NewUserCache()
	return domain
}

func (domain *UserDomain) DomainName() string {
	return "UserDomain"
}

// 生成用户编号
func (domain *UserDomain) generateUserNo() string {
	// 采用ksuid的ID生成策略来创建全局唯一的ID
	return ksuid.New().Next().String()
}

// 加密密码，设置密文和盐值
func (domain *UserDomain) encryptPassword(password string) (hashStr, salt string) {
	hashStr, salt = XGlobal.HashPassword(password)
	return
}

// 鉴权后生成token
func (domain *UserDomain) GenToken(no, name, avatar string) (string, error) {
	var err error
	var token string
	// 生成
	var claim = XJwt.UserClaim{
		Id:     no,
		Name:   name,
		Avatar: avatar,
	}
	token, err = XJwt.XTokenUtils().Encode(claim)
	if err != nil {
		return "", common.DomainInnerErrorOnEncodeData(err, claim)
	}

	return token, nil
}

// 移除JWT Token 缓存
func (domain *UserDomain) RemoveTokenCache(token string) error {
	err := XJwt.XTokenUtils().Remove(token)
	if err != nil {
		return common.DomainInnerErrorOnCacheDelete(err, "XJwt.XTokenUtils().Remove(token)")
	}
	return nil
}

// 查找用户id是否已存在
func (domain *UserDomain) IsUserExist(uid uint) (bool, error) {
	var err error
	var isExist bool

	if isExist, err = domain.dao.IsUserIdExist(uid); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsUserIdExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找邮箱是否已存在
func (domain *UserDomain) IsEmailExist(email string) (bool, error) {
	var err error
	var isExist bool
	if isExist, err = domain.dao.IsEmailExist(email); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsEmailExist")
	} else if isExist {
		return true, nil
	}

	return false, nil
}

// 查找手机用户是否已存在
func (domain *UserDomain) IsPhoneExist(phone string) (bool, error) {
	var err error
	var isExist bool
	if isExist, err = domain.dao.IsPhoneExist(phone); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsPhoneExist")
	} else if isExist {
		return true, nil
	}
	return false, nil
}

// 邮箱账号创建用户
func (domain *UserDomain) CreateUserForEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO

	createUserData := dtos.UserDTO{}
	createUserData.Name = dto.Name
	createUserData.Email = dto.Email
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerify // 初始创建时未验证状态

	if userDTO, err = domain.dao.Create(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}
	return userDTO, nil
}

// 手机号码创建用户
func (domain *UserDomain) CreateUserForPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO

	createUserData := dtos.UserDTO{}
	createUserData.Name = dto.Name
	createUserData.Phone = dto.Phone
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerify // 初始创建时未验证状态

	if userDTO, err = domain.dao.Create(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}
	return userDTO, nil
}

// Oauth三方账号绑定创建用户
func (domain *UserDomain) CreateUserOAuthBinding(platform uint, oauthInfo *XOAuth.OAuthAccountInfo) (*dtos.UserOAuthsDTO, error) {
	var err error
	var userOAuthsResult *dtos.UserOAuthsDTO

	// 插入用户信息
	createUserData := dtos.UserOAuthsDTO{}
	createUserData.User.Name = oauthInfo.NickName
	createUserData.User.No = domain.generateUserNo()
	createUserData.User.Status = UserStatusNotVerify // 初始创建时未验证状态
	createUserData.UserOAuths = []dtos.OAuthDTO{
		{
			AccessToken: oauthInfo.AccessToken,
			UnionId:     oauthInfo.UnionId,
			OpenId:      oauthInfo.OpenId,
			NickName:    oauthInfo.NickName,
			Avatar:      oauthInfo.AvatarUrl,
			Gender:      oauthInfo.Gender,
			Platform:    platform,
		},
	}

	if userOAuthsResult, err = domain.dao.CreateUserWithOAuth(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "CreateUserWithOAuth")
	}

	return userOAuthsResult, nil
}

// 获取整个关联的用户信息和三方平台绑定信息
func (domain *UserDomain) GetUserOauths(platform uint, openId, unionId string) (*dtos.UserOAuthsDTO, error) {
	var err error
	var userOAuthsResult *dtos.UserOAuthsDTO

	if userOAuthsResult, err = domain.dao.GetUserOAuths(platform, openId, unionId); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetUserOAuths")
	}

	return userOAuthsResult, nil
}

func (domain *UserDomain) GetUser(uid uint) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO
	if userDTO, err = domain.dao.GetById(uid); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetById")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserByEmail(email string) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO
	if userDTO, err = domain.dao.GetByEmail(email); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetByEmail")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserByPhone(phone string) (*dtos.UserDTO, error) {
	var err error
	var userDTO *dtos.UserDTO
	if userDTO, err = domain.dao.GetByPhone(phone); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetByPhone")
	}
	return userDTO, nil
}

// 设置用户状态
func (domain *UserDomain) SetStatus(uid, status uint) error {
	var err error
	if err = domain.dao.SetUserInfo(uid, "status", status); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUserInfo")
	}
	return nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid uint, field string, value interface{}) error {
	var err error
	if err = domain.dao.SetUserInfo(uid, field, value); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUserInfo")
	}
	return nil
}

// 设置多个用户信息
func (domain *UserDomain) SetUserInfos(uid uint, dto dtos.SetUserInfoDTO) error {
	var err error
	if err = domain.dao.SetUserInfos(uid, dto); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUserInfos")
	}

	return nil
}

// 改变密码
func (domain *UserDomain) ReSetPassword(uid uint, password string) error {
	var err error
	var hashStr, salt string
	hashStr, salt = domain.encryptPassword(password)
	if err = domain.dao.SetPasswordAndSalt(uid, hashStr, salt); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetPasswordAndSalt")
	}
	return nil
}

// 真删除
func (domain *UserDomain) DeleteUser(uid uint) error {
	var err error
	if err = domain.dao.DeleteById(uid); err != nil {
		return common.DomainInnerErrorOnSqlDelete(err, "DeleteById")
	}
	return nil
}

// 伪删除
func (domain *UserDomain) ShamDeleteUser(uid uint) error {
	var err error
	if err = domain.dao.SetDeletedAtById(uid); err != nil {
		return common.DomainInnerErrorOnSqlShamDelete(err, "SetDeletedAtById")
	}
	return nil
}
