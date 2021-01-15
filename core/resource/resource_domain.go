package resource

/*
资源上传下载领域
*/

type ResourceDomain struct {
}

func NewResourceDomain() *ResourceDomain {
	domain := new(ResourceDomain)

	return domain
}

func (domain *ResourceDomain) DomainName() string {
	return "ResourceDomain"
}

// 获取七牛对象存储客户端上传令牌
func (domain *ResourceDomain) GetQiniuClientUploadToken() string {
	return "token string"
}
