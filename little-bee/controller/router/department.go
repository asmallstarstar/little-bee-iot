package router

import (
	"controller/config"

	"github.com/gin-gonic/gin"
)

const URL_ROUTER_DEPARTMENT = "/config/department"
const URL_ROUTER_DEPARTMENT_BY_ID = "/config/department/:department-id"
const URL_ROUTER_DEPARTMENT_QUERY = "/config/department/query"
const URL_ROUTER_DEPARTMENT_ALL = "/config/department/all"

func DepartmentRouter(r *gin.RouterGroup) *gin.RouterGroup {
	r.POST(URL_ROUTER_DEPARTMENT, config.CreateDepartment)
	r.GET(URL_ROUTER_DEPARTMENT_BY_ID, config.RetrieveDepartment)
	r.PUT(URL_ROUTER_DEPARTMENT, config.UpdateDepartment)
	r.DELETE(URL_ROUTER_DEPARTMENT_BY_ID, config.DeleteDepartment)
	r.PUT(URL_ROUTER_DEPARTMENT_BY_ID, config.DeleteDepartmentWithFlag)
	r.POST(URL_ROUTER_DEPARTMENT_QUERY, config.QueryDepartment)
	r.GET(URL_ROUTER_DEPARTMENT_ALL, config.GetAllDepartments)
	return r
}