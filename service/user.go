package service

import (
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"mall/config"
	"mall/model"
	"mall/query"
	"mall/repository"
	"mall/utils"
)

type UserSrv interface {
	List(req *query.ListQuery)(users []*model.User,err error)
	GetTotal(req *query.ListQuery) (total int,err error)
	Get(user model.User) (*model.User,error)
	Exist(user model.User) *model.User
	ExistByUserID(id string) *model.User
	Add(user model.User) (*model.User,error)
	Edit(user model.User) (bool,error)
	Delete(id string)(bool, error)
}
type UserService struct {
	Repo repository.UserRepoInterface //为什么是接口
}

func (srv *UserService) List(req *query.ListQuery) (users []*model.User, err error) {
	if req.PageSize < 1 {
		req.PageSize = config.PAGE_SIZE
	}
	return srv.Repo.List(req)
}

func (srv UserService) GetTotal(req *query.ListQuery) (total int, err error) {
	return srv.Repo.GetTotal(req)
}

func (srv UserService) Get(user model.User) (*model.User, error) {
	return srv.Repo.Get(user)
}

func (srv UserService) Exist(user model.User) *model.User {
	return srv.Repo.Exist(user)
}

func (srv UserService) ExistByUserID(id string) *model.User {
	return srv.Repo.ExistByUserId(id)
}

func (srv UserService) Add(user model.User) (*model.User, error) {
	//根据手机号码判断用户是否存在
	result := srv.Repo.ExistByMobile(user.Mobile)
	if result != nil{
		return nil,errors.New("用户已存在")
	}
	user.UserID = uuid.NewV4().String()
	if user.Password == ""{
		user.Password = utils.Md5("123456")
	}
	user.IsDelete = false
	user.IsLocked = false
	return srv.Repo.Add(user)
}

func (srv UserService) Edit(user model.User) (bool, error) {
	if user.UserID ==""{
		return false,fmt.Errorf("参数错误")
	}
	exist := srv.Repo.ExistByUserId(user.UserID)
	if exist == nil{
		return false, errors.New("参数错误")
	}
	exist.NickName = user.NickName
	exist.Mobile = user.Mobile
	exist.Address = user.Address
	return srv.Repo.Edit(*exist)
}

func (srv UserService) Delete(id string) (bool, error) {
	if id ==""{
		return false,errors.New("参数错误")
	}
	user := srv.ExistByUserID(id)
	if user == nil{
		return false, errors.New("参数错误")
	}
	user.IsDelete = !user.IsDelete //赋值
	return srv.Repo.Delete(*user)
}
