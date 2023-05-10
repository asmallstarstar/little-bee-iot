package user

import (
	"dao/database"
	"message/littlebee"
	"message/querycommand"
	"util/queryclause"

	"github.com/fatih/structs"
)

func VerifyUser(name string, md5 string) ([]*littlebee.User, int64, error) {
	var users []*littlebee.User
	result := database.DB.Where("user_name = ? AND password = ? AND status = 0 AND is_deleted = 0", name, md5).Find(&users)

	return users, result.RowsAffected, result.Error
}

func ModifyPassword(userId int32, oldMd5Password string, newMd5Password string) (int32, error) {
	result := database.DB.Model(&littlebee.User{}).Where("user_id = ? AND password = ? AND status = 0 AND is_deleted = 0",
		userId, oldMd5Password).Update("password", newMd5Password)
	return int32(result.RowsAffected), result.Error
}

func CreateUser(u *littlebee.User) (*littlebee.User, error) {
	u.UserId = 0
	result := database.DB.Omit("UserId", "UpdatedAt", "UpdatedBy", "IsDeleted", "DeletedAt", "DeletedBy").Create(u)
	return u, result.Error
}

func RetrieveUser(userId int32) (*littlebee.User, error) {
	u := &littlebee.User{}
	result := database.DB.First(&u, userId)
	return u, result.Error
}

func UpdateUser(u *littlebee.User) error {
	uMap := structs.Map(u)
	uMap["created_at"] = u.CreatedAt.AsTime().UTC()
	uMap["updated_at"] = u.UpdatedAt.AsTime().UTC()
	uMap["deleted_at"] = u.DeletedAt.AsTime().UTC()
	result := database.DB.Model(u).Omit("UserId", "Password", "CreatedAt", "CreatedBy").Updates(uMap)
	return result.Error
}

func DeleteUserWithFlag(u *littlebee.User) error {
	result := database.DB.Model(u).Updates(littlebee.User{Status: u.Status, IsDeleted: u.IsDeleted, DeletedAt: u.DeletedAt, DeletedBy: u.DeletedBy})
	return result.Error
}

func DeleteUser(u *littlebee.User) error {
	result := database.DB.Delete(u)
	return result.Error
}

func QueryUser(q *querycommand.QueryCommand) (*littlebee.UserList, error) {
	var users []*littlebee.User
	s := "SELECT user_id, department_id, user_name, " +
		"user_nick, phone, email, sex, " +
		"avatar, status,remark," +
		"created_at,created_by,updated_at, updated_by,is_deleted, deleted_at, deleted_by FROM users "
	w, v := queryclause.QueryCommandString(q)
	s = s + w
	result := database.DB.Raw(s, v...).Scan(&users)
	return &littlebee.UserList{Items: users}, result.Error
}

func GetAll() (*littlebee.UserList, error) {
	var users []*littlebee.User
	result := database.DB.Select("user_id", "department_id", "user_name",
		"user_nick", "phone", "email", "sex",
		"avatar", "status", "remark",
		"created_at", "created_by", "updated_at", "updated_by", "is_deleted", "deleted_at", "deleted_by").Find(&users)
	return &littlebee.UserList{Items: users}, result.Error
}

func GetAllAlias() (*littlebee.UserAliasList, error) {
	var users []*littlebee.UserAlias
	result := database.DB.Raw("SELECT user_id, user_nick FROM users").Scan(&users)
	return &littlebee.UserAliasList{Items: users}, result.Error
}
