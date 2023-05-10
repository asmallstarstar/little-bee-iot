package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_ROLE = "/config/role"
const URL_ROUTER_ROLE_BY_ID = "/config/role/:role-id"
const URL_ROUTER_ROLE_QUERY = "/config/role/query"
const URL_ROUTER_ROLE_ALL = "/config/role/all"

func RoleRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_ROLE, config.CreateRole)
	r.GET(URL_ROUTER_ROLE_BY_ID, config.RetrieveRole)
	r.PUT(URL_ROUTER_ROLE, config.UpdateRole)
	r.DELETE(URL_ROUTER_ROLE_BY_ID, config.DeleteRole)
	r.PUT(URL_ROUTER_ROLE_BY_ID, config.DeleteRoleWithFlag)
	r.POST(URL_ROUTER_ROLE_QUERY, config.QueryRole)
	r.GET(URL_ROUTER_ROLE_ALL, config.GetAllRoles)
	return r
}