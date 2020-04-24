package routers

import (
	"github.com/astaxie/beego"
	"gowebsite/apps/admin/employee/controllers"
)

func Routers() {
	//beego.Router("empl/v2/demo/index", &controllers.RoleController{}, "get:Index")

	//role业务  table:ut_role 、ut_operation 、ur_rel_role_operation
	beego.Router("empl/role/query/?:id", &controllers.RoleController{}, "get:Query")
	beego.Router("empl/role/add", &controllers.RoleController{}, "post:Add")
	beego.Router("empl/role/update", &controllers.RoleController{}, "post:Update")
	beego.Router("empl/role/delete", &controllers.RoleController{}, "post:Delete")

	//empolyee

	beego.Router("empl/user/query", &controllers.EmployeeController{}, "get:Query")
	//beego.Router("empl/user/add", &controllers.EmployeeController{}, "post.Add")

	// group
	beego.Router("/api/v2/admin/group/create", &controllers.GroupController{}, "post:CreateGroup")
	beego.Router("/api/v2/admin/group/delete", &controllers.GroupController{}, "post:RemoveGroup")
	beego.Router("/api/v2/admin/group/update", &controllers.GroupController{}, "post:UpdateGroup")
	beego.Router("/api/v2/admin/group/get", &controllers.GroupController{}, "get:GetGroupDetail")
	beego.Router("/api/v2/admin/group/list", &controllers.GroupController{}, "post:GroupListByPage")
	beego.Router("/api/v2/admin/group/add_agent", &controllers.GroupController{}, "post:GroupAddAgent") 	// add agent to group
	beego.Router("/api/v2/admin/group/del_agent", &controllers.GroupController{}, "post:GroupRemoveAgent") 	// remove agent from group

	// department


}

//func Routers() {}
