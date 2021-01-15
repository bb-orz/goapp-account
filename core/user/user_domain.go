package user

import (
	"github.com/bb-orz/goinfras/XJwt"
	"github.com/bb-orz/goinfras/XOAuth"
	"github.com/segmentio/ksuid"
	"goapp/common"
	"goapp/dtos"
)

/*
User 领域层：实现用户相关具体业务逻辑
封装领域层的错误信息并返回给调用者
*/
type UserDomain struct {
	userDao      *UsersDAO
	OauthsDAO    *OauthsDAO
	userOAuthDao *UserOAuthDAO
	cache        *UsersCache
}

func NewUserDomain() *UserDomain {
	domain := new(UserDomain)
	domain.userDao = NewUsersDAO()
	domain.userOAuthDao = NewUserOAuthDAO()
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
	hashStr, salt = common.HashPassword(password)
	return
}

// 用户密码验证
func (domain *UserDomain) VerifyPassword(id uint, passwordStr string) (*dtos.UsersDTO, bool, error) {
	// 查找账号是否存在
	userDTO, err := domain.GetUser(id)
	if err != nil {
		return nil, false, common.DomainInnerErrorOnSqlQuery(err, domain.DomainName())
	}

	if userDTO == nil {
		return nil, false, nil
	}

	return userDTO, common.ValidatePassword(passwordStr, userDTO.Salt, userDTO.Password), nil
}

// 用户密码验证ForEmail
func (domain *UserDomain) VerifyPasswordForEmail(email, passwordStr string) (*dtos.UsersDTO, bool, error) {
	// 查找账号是否存在
	userDTO, err := domain.GetUserByEmail(email)
	if err != nil {
		return nil, false, common.DomainInnerErrorOnSqlQuery(err, domain.DomainName())
	}

	if userDTO == nil {
		return nil, false, nil
	}

	return userDTO, common.ValidatePassword(passwordStr, userDTO.Salt, userDTO.Password), nil
}

// 用户密码验证ForPhone
func (domain *UserDomain) VerifyPasswordForPhone(phone, passwordStr string) (*dtos.UsersDTO, bool, error) {
	// 查找账号是否存在
	userDTO, err := domain.GetUserByPhone(phone)
	if err != nil {
		return nil, false, common.DomainInnerErrorOnSqlQuery(err, domain.DomainName())
	}

	if userDTO == nil {
		return nil, false, nil
	}

	return userDTO, common.ValidatePassword(passwordStr, userDTO.Salt, userDTO.Password), nil
}

// 鉴权后生成token
func (domain *UserDomain) GenToken(id uint, no, name, avatar string) (string, error) {
	var err error
	var token string
	// 生成
	var claim = XJwt.UserClaim{
		Id:     id,
		No:     no,
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

	if isExist, err = domain.userDao.IsIdExist(uid); err != nil {
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
	if isExist, err = domain.userDao.IsEmailExist(email); err != nil {
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
	if isExist, err = domain.userDao.IsPhoneExist(phone); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsPhoneExist")
	} else if isExist {
		return true, nil
	}
	return false, nil
}

// 邮箱账号创建用户
func (domain *UserDomain) CreateUserForEmail(dto dtos.CreateUserWithEmailDTO) (*dtos.UsersDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO

	createUserData := dtos.UsersDTO{}
	createUserData.Name = dto.Name
	createUserData.Email = dto.Email
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerify // 初始创建时未验证状态

	if userDTO, err = domain.userDao.Create(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}
	return userDTO, nil
}

// 手机号码创建用户
func (domain *UserDomain) CreateUserForPhone(dto dtos.CreateUserWithPhoneDTO) (*dtos.UsersDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO

	createUserData := dtos.UsersDTO{}
	createUserData.Name = dto.Name
	createUserData.Phone = dto.Phone
	createUserData.No = domain.generateUserNo()
	createUserData.Password, createUserData.Salt = domain.encryptPassword(dto.Password)
	createUserData.Status = UserStatusNotVerify // 初始创建时未验证状态

	if userDTO, err = domain.userDao.Create(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUser(uid uint) (*dtos.UsersDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO
	if userDTO, err = domain.userDao.GetById(uid); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetById")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserByEmail(email string) (*dtos.UsersDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO
	if userDTO, err = domain.userDao.GetByEmail(email); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetByEmail")
	}
	return userDTO, nil
}

func (domain *UserDomain) GetUserByPhone(phone string) (*dtos.UsersDTO, error) {
	var err error
	var userDTO *dtos.UsersDTO
	if userDTO, err = domain.userDao.GetByPhone(phone); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetByPhone")
	}
	return userDTO, nil
}

// 设置用户状态
func (domain *UserDomain) SetStatus(uid, status uint) error {
	if err := domain.userDao.SetUsers(uid, "status", status); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUsers")
	}
	return nil
}

// 设置用户状态
func (domain *UserDomain) SetEmailVerify(uid, status uint) error {
	if err := domain.userDao.SetUsers(uid, "email_verify", status); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUsers")
	}
	return nil
}

// 设置用户状态
func (domain *UserDomain) SetPhoneVerify(uid, status uint) error {
	if err := domain.userDao.SetUsers(uid, "phone_verify", status); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUsers")
	}
	return nil
}

// 设置单个用户信息
func (domain *UserDomain) SetUserInfo(uid uint, field string, value interface{}) error {
	if err := domain.userDao.SetUsers(uid, field, value); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetUsers")
	}
	return nil
}

// 设置多个用户信息
func (domain *UserDomain) UpdateUsers(uid uint, dto dtos.SetUserInfoDTO) error {
	updateData := dtos.UserInfoDTO{
		Id:     dto.Id,
		Name:   dto.Name,
		Avatar: dto.Avatar,
		Gender: dto.Gender,
		Age:    dto.Age,
	}
	if err := domain.userDao.UpdateUsers(uid, updateData); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "UpdateUsers")
	}

	return nil
}

// 改变密码
func (domain *UserDomain) ReSetPassword(id uint, password string) error {
	hashStr, salt := domain.encryptPassword(password)
	if err := domain.userDao.SetPasswordAndSalt(id, hashStr, salt); err != nil {
		return common.DomainInnerErrorOnSqlUpdate(err, "SetPasswordAndSalt")
	}
	return nil
}

// 真删除
func (domain *UserDomain) DeleteUser(uid uint) error {
	if err := domain.userDao.DeleteById(uid); err != nil {
		return common.DomainInnerErrorOnSqlDelete(err, "DeleteById")
	}
	return nil
}

// 伪删除
func (domain *UserDomain) ShamDeleteUser(uid uint) error {
	if err := domain.userDao.SetDeletedAtById(uid); err != nil {
		return common.DomainInnerErrorOnSqlShamDelete(err, "SetDeletedAtById")
	}
	return nil
}

// 是否qq账号已绑定
func (domain *UserDomain) IsQQAccountBinding(openId, unionId string) (bool, error) {
	if b, err := domain.OauthsDAO.IsQQAccountBindng(openId, unionId); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsQQAccountBinding")
	} else if b {
		return true, nil
	}
	return false, nil
}

// 是否微信账号已绑定
func (domain *UserDomain) IsWechatAccountBinding(openId, unionId string) (bool, error) {
	if b, err := domain.OauthsDAO.IsWechatAccountBindng(openId, unionId); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsWechatAccountBinding")
	} else if b {
		return true, nil
	}
	return false, nil
}

// 是否微博账户已绑定
func (domain *UserDomain) IsWeiboAccountBinding(openId, unionId string) (bool, error) {
	if b, err := domain.OauthsDAO.IsWeiboAccountBindng(openId, unionId); err != nil {
		return false, common.DomainInnerErrorOnSqlQuery(err, "IsWeiboAccountBinding")
	} else if b {
		return true, nil
	}
	return false, nil
}

// Oauth三方账号绑定创建用户
func (domain *UserDomain) CreateUserWithOAuthBinding(platform uint, oauthInfo *XOAuth.OAuthAccountInfo) (*dtos.UserOAuthInfoDTO, error) {
	var err error

	// 插入用户信息
	createUserData := dtos.UserOAuthInfoDTO{}
	createUserData.Name = oauthInfo.NickName
	createUserData.No = domain.generateUserNo()
	createUserData.Status = UserStatusNotVerify // 初始创建时未验证状态
	createUserData.OAuths = []dtos.OauthsDTO{
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

	var result *dtos.UserOAuthInfoDTO
	if result, err = domain.userOAuthDao.CreateUserWithOAuth(&createUserData); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "CreateUserWithOAuth")
	}

	return result, nil
}

// 获取整个关联的用户信息和三方平台绑定信息
func (domain *UserDomain) GetUserOauths(platform uint, openId, unionId string) (*dtos.UserOAuthInfoDTO, error) {
	var err error
	var userOAuthsResult *dtos.UserOAuthInfoDTO

	if userOAuthsResult, err = domain.userOAuthDao.GetUserOAuths(platform, openId, unionId); err != nil {
		return nil, common.DomainInnerErrorOnSqlQuery(err, "GetUserOAuths")
	}

	return userOAuthsResult, nil
}

// Oauth三方账号绑定创建用户
func (domain *UserDomain) CreateOAuthBinding(platform uint, oauthInfo *XOAuth.OAuthAccountInfo) (*dtos.OauthsDTO, error) {
	var err error
	var oauthsDTO *dtos.OauthsDTO
	var result *dtos.OauthsDTO
	oauthsDTO = &dtos.OauthsDTO{
		AccessToken: oauthInfo.AccessToken,
		UnionId:     oauthInfo.UnionId,
		OpenId:      oauthInfo.OpenId,
		NickName:    oauthInfo.NickName,
		Avatar:      oauthInfo.AvatarUrl,
		Gender:      oauthInfo.Gender,
		Platform:    platform,
	}

	oauthsDAO := NewOauthsDAO()
	if result, err = oauthsDAO.Create(oauthsDTO); err != nil {
		return nil, common.DomainInnerErrorOnSqlInsert(err, "Create")
	}

	return result, nil
}
