package user

import (
	"crypto/md5"
	dao "dao/user"
	"errors"
	"fmt"
	"message/littlebee"
	"message/querycommand"
	"strings"
	"util/jwt"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func Login(u *littlebee.LoginRequest) (*littlebee.LoginResponse, int64, error) {
	p := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(u.Password))))
	l, r, e := dao.VerifyUser(u.UserName, p)
	if e != nil {
		return nil, 0, e
	} else {
		if r != 1 {
			return nil, r, nil
		} else {
			token, err := jwt.GenerateToken(l[0].UserId)
			if err != nil {
				return nil, 1, err
			} else {
				l[0].Password = ""
				return &littlebee.LoginResponse{Token: token, User: l[0]}, r, e
			}
		}
	}
}

func VerifyUser(u *littlebee.UserVerify) (*littlebee.UserVerifyResponse, error) {
	p := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(u.Password))))
	l, r, e := dao.VerifyUser(u.UserName, p)
	if e != nil {
		return nil, e
	} else {
		if r != 1 {
			return nil, nil
		} else {
			l[0].Password = ""
			return &littlebee.UserVerifyResponse{User: l[0]}, nil
		}
	}
}

func ModifyPassword(p *littlebee.UserModifyPasword) (int32, error) {
	if p.NewPassword == p.NewRepeatPassword {
		op := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(p.OldPassword))))
		np := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(p.NewPassword))))
		return dao.ModifyPassword(p.UserId, op, np)
	}
	return 0, errors.New("")
}

func CreateUser(u *littlebee.User, userId int32) (*littlebee.User, error) {
	p := strings.ToLower(fmt.Sprintf("%x", md5.Sum([]byte(u.Password))))
	u.Password = p
	u.UserId = 0
	u.CreatedAt = timestamppb.Now()
	u.CreatedBy = userId
	n, e := dao.CreateUser(u)
	n.Password = ""
	return n, e
}

func RetrieveUser(userId int32) (*littlebee.User, error) {
	u, e := dao.RetrieveUser(userId)
	u.Password = ""
	return u, e
}

func UpdateUser(u *littlebee.User, userId int32) error {
	if u.Status == 1 {
		u.IsDeleted = true
		u.DeletedAt = timestamppb.Now()
		u.DeletedBy = userId
	} else {
		u.IsDeleted = false
	}
	u.UpdatedAt = timestamppb.Now()
	u.UpdatedBy = userId
	return dao.UpdateUser(u)
}

func DeleteUserWithFlag(userId int, deletedByUserId int32) error {
	u := &littlebee.User{
		UserId:    int32(userId),
		Status:    1,
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: deletedByUserId,
	}
	return dao.DeleteUserWithFlag(u)
}

func DeleteUser(userId int, deletedByUserId int32) error {
	u := &littlebee.User{
		UserId:    int32(userId),
		IsDeleted: true,
		DeletedAt: timestamppb.Now(),
		DeletedBy: deletedByUserId,
	}
	return dao.DeleteUser(u)
}

func QueryUser(q *querycommand.QueryCommand) (*littlebee.UserList, error) {
	return dao.QueryUser(q)
}

func GetAll() (*littlebee.UserList, error) {
	return dao.GetAll()
}

func GetAllAlias() (*littlebee.UserAliasList, error) {
	return dao.GetAllAlias()
}
