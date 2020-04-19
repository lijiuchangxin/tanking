package routers

import (
	"github.com/astaxie/beego"
	"lzsc/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v1/admin/get", &controllers.CustomerController{}, "get:Get")
	beego.Router("/api/v1/admin/post", &controllers.CustomerController{}, "post:Post")

	beego.Router("/api/v1/admin/new-customer", &controllers.CustomerController{}, "post:NewCustomer")
	//beego.Router("/api/v1/admin/show-customer", &controllers.CustomerController{}, "get:ShowCustomer")
	//beego.Router("/api/v1/admin/get-customers", &controllers.CustomerController{}, "get:PageCustomers")
	//beego.Router("/api/v1/admin/update-customer", &controllers.CustomerController{}, "post:UpdateCustomer")
	//beego.Router("/api/v1/admin/del-customer", &controllers.CustomerController{}, "post:DelCustomer")
	//beego.Router("/api/v1/admin/search-customer", &controllers.CustomerController{}, "post:SearchCustomer")
}
