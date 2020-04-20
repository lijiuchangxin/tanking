package routers

import (
	"customer_managenment/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v2/admin/new-customer", &controllers.CustomerController{}, "post:NewCustomer")
	beego.Router("/api/v2/admin/del-customer", &controllers.CustomerController{}, "post:DeleteCustomer")

	//beego.Router("/api/v2/admin/show-customer", &controllers.CustomerController{}, "get:ShowCustomer")
	//beego.Router("/api/v2/admin/get-customers", &controllers.CustomerController{}, "get:PageCustomers")
	//beego.Router("/api/v2/admin/update-customer", &controllers.CustomerController{}, "post:UpdateCustomer")
	//beego.Router("/api/v2/admin/search-customer", &controllers.CustomerController{}, "post:SearchCustomer")
}
