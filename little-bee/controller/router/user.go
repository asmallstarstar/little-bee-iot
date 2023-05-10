package router

import (
	"controller/user"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_USER = "/user"
const URL_ROUTER_USER_BY_ID = "/user/:user-id"
const URL_ROUTER_USER_QUERY = "/user/query"
const URL_ROUTER_USER_ALL = "/user/all"
const URL_ROUTER_USER_ALL_ALIAS = "/user/alias-all"

const URL_ROUTER_USER_LOGIN = "/user/login"
const URL_ROUTER_USER_VERIFY = "/user/verify"

const URL_ROUTER_USER_GRANT_USER = "/user/grant-user"
const URL_ROUTER_USER_GRANT_GROUP = "/user/grant-group"
const URL_ROUTER_USER_GRANT_GROUP_BATCH = "/user/grant-group/batch"
const URL_ROUTER_USER_GRANT_USER_BATCH = "/user/grant-user/batch"
const URL_ROUTER_USER_GRANT_BY_ROLEID = "/user/grant-group/:role-id"
const URL_ROUTER_USER_JOIN_GROUP_BY_USERID = "/user/join-group/:user-id"
const URL_ROUTER_USER_GRANT_USER_BY_USERID = "/user/grant-user/:user-id"
const URL_ROUTER_USER_GRANT_USER_WITH_ROLES_BY_USERID = "/user/grant-user-with-roles/:user-id"
const URL_ROUTER_USER_JOIN_GROUP = "/user/join-group"
const URL_ROUTER_USER_REMOVE_GROUP = "/user/remove-group"
const URL_ROUTER_USER_ASSEMBLED_MENUS = "/user/assembled-menus"
const URL_ROUTER_USER_ALL_ACTIONS = "/user/all-actions"
const URL_ROUTER_USER_MODIFY_PASSWORD = "/user/password"

func UserPublicRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_USER_LOGIN, user.Login)
	return r
}

func UserPrivateRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_USER, user.CreateUser)
	r.GET(URL_ROUTER_USER_BY_ID, user.RetrieveUser)
	r.PUT(URL_ROUTER_USER, user.UpdateUser)
	r.DELETE(URL_ROUTER_USER_BY_ID, user.DeleteUser)
	r.PUT(URL_ROUTER_USER_BY_ID, user.DeleteUserWithFlag)
	r.GET(URL_ROUTER_USER_ALL, user.GetAll)
	r.GET(URL_ROUTER_USER_ALL_ALIAS, user.GetAllAlias)
	r.POST(URL_ROUTER_USER_QUERY, user.QueryUser)
	r.POST(URL_ROUTER_USER_GRANT_USER, user.GrantUser)
	r.POST(URL_ROUTER_USER_GRANT_GROUP, user.GrantGroup)
	r.POST(URL_ROUTER_USER_GRANT_USER_BATCH, user.BatchGrantUser)
	r.POST(URL_ROUTER_USER_GRANT_GROUP_BATCH, user.BatchGrantGroup)
	r.GET(URL_ROUTER_USER_GRANT_BY_ROLEID, user.GetGroupGranted)
	r.GET(URL_ROUTER_USER_JOIN_GROUP_BY_USERID, user.GetUserGroup)
	r.GET(URL_ROUTER_USER_GRANT_USER_BY_USERID, user.GetUserGrant)
	r.GET(URL_ROUTER_USER_GRANT_USER_WITH_ROLES_BY_USERID, user.GetUserActionsWithRoles) //DEACTIVATED
	r.POST(URL_ROUTER_USER_JOIN_GROUP, user.JoinGroup)
	r.POST(URL_ROUTER_USER_REMOVE_GROUP, user.RemoveGroup)
	r.GET(URL_ROUTER_USER_ASSEMBLED_MENUS, user.AssembledMenus)
	r.GET(URL_ROUTER_USER_ALL_ACTIONS, user.GetUserActions)
	r.PUT(URL_ROUTER_USER_MODIFY_PASSWORD, user.ModifyPassword)
	r.POST(URL_ROUTER_USER_VERIFY, user.VerifyUser)
	return r
}
