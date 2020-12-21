package user

import (
	"github.com/bb-orz/goinfras/XStore/XGorm"
	"github.com/bb-orz/goinfras/XValidate"
	"goinfras-sample-account/services"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestUserService_CreateUser(t *testing.T) {
	Convey("User Service Create User Testing:", t, func() {
		var err error
		err = XValidate.CreateDefaultValidater(nil)
		So(err, ShouldBeNil)
		err = XGorm.CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		dto := services.CreateUserWithEmailDTO{
			Name:       "fun",
			Email:      "123456@qq.com",
			Password:   "123456",
			RePassword: "123456",
		}
		service := new(services.UserService)
		userDTO, err := service.CreateUserWithEmail(dto)
		So(err, ShouldBeNil)

		Println("New User:", userDTO)
	})
}

func TestUserService_GetUserInfo(t *testing.T) {
	Convey("User Service Get User Info Testing:", t, func() {
		var err error
		err = XValidate.CreateDefaultValidater(nil)
		So(err, ShouldBeNil)
		err = XGorm.CreateDefaultDB(nil)
		So(err, ShouldBeNil)

		service := new(services.UserService)
		userDTO, err := service.GetUserInfo(services.GetUserInfoDTO{1})
		So(err, ShouldBeNil)
		Println("Get User Info:", userDTO)
	})
}
