package routers

import (
	"customer_managenment/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/api/v2/admin/customer/create", &controllers.CustomerController{}, "post:CreateCustomer")
	beego.Router("/api/v2/admin/customer/delete", &controllers.CustomerController{}, "post:DeleteCustomer")
	beego.Router("/api/v2/admin/customer/create-follow", &controllers.CustomerController{}, "post:CreateCustomerFollow")
	beego.Router("/api/v2/admin/customer/delete-follow", &controllers.CustomerController{}, "post:DeleteCustomerFollow")
	beego.Router("/api/v2/admin/customer/show", &controllers.CustomerController{}, "get:ShowCustomerDetail")

	//beego.Router("/api/v2/admin/update-customer", &controllers.CustomerController{}, "post:UpdateCustomer")



	//beego.Router("/api/v2/admin/show-customer", &controllers.CustomerController{}, "get:ShowCustomer")
	//beego.Router("/api/v2/admin/get-customers", &controllers.CustomerController{}, "get:PageCustomers")
	//beego.Router("/api/v2/admin/update-customer", &controllers.CustomerController{}, "post:UpdateCustomer")
	//beego.Router("/api/v2/admin/search-customer", &controllers.CustomerController{}, "post:SearchCustomer")
}
