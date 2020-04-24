package controllers

import (
	"github.com/astaxie/beego"
)

type Resp struct {
	beego.Controller
}

type ResponseStru struct {
	Code  int
	Mess  string
	Value interface{}
}

const CODE_SUCCESS = 200
const SUCCESS = "success "
const CODE_FAIL = 400
const FAIL = "fail "

func (t *ResponseStru) RespSet(code int, mess string, data interface{}) {
	t.Code = 200
	t.Mess = mess
	t.Value = data
}

func (t *ResponseStru) RespSuccessSet(data interface{}) {
	t.Code = CODE_SUCCESS
	t.Mess = SUCCESS
	t.Value = data
}

func (t *ResponseStru) RespFailSet(data interface{}) {
	t.Code = CODE_FAIL
	t.Mess = FAIL
	t.Value = data
}
