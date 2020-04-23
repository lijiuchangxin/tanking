import json
import random
import unittest
import requests
from comm.mysql_db import DB



province_lst = ["陕西", "山西", "甘肃", "河北", "河南", "新疆", "浙江"]

city_dict = {
    "陕西": "xian",
    "山西": "taiyuan",
    "甘肃": "lanzhou",
    "河北": "shijiazhuang",
    "河南": "luoyang",
    "新疆": "kashi",
    "浙江": "hangzhou"
}

class TestCustomer(unittest.TestCase):
    create_customer = "http://127.0.0.1:8000/api/v2/admin/customer/create"
    delete_customer = "http://127.0.0.1:8000/api/v2/admin/customer/delete"
    create_follow = "http://127.0.0.1:8000/api/v2/admin/customer/create-follow"
    delete_follow = "http://127.0.0.1:8000/api/v2/admin/customer/delete-follow"
    show_customer = "http://127.0.0.1:8000/api/v2/admin/customer/show"
    update_customer = "http://127.0.0.1:8000/api/v2/admin/customer/update"
    get_customer = "http://127.0.0.1:8000/api/v2/admin/customer/list"
    search_customer = "http://127.0.0.1:8000/api/v2/admin/customer/search"
    last_data = {}
    db = DB()

    def setUp(self):
        self.customer_data = {
            "customer_nike_name": "",
            "desc": "",
            "tag": "",
            "tel_phone": "",
            "cell_phone": "",
            "email": "",
            "is_vip": 1,
            "province": "",
            "city": "",
            "source_channel": "",
            "open_api_token": "{}{}{}{}-1995-0612-XzFG"
        }

        self.follow_data = {
            "customer_id": 0,
            "content": "testing",
            "user_id": 999,

        }

        self.db.clear("ut_customer")


    def new_customer(self):
        """
        调用接口生成一个正经客户
        """
        self.customer_data = {
            "customer_nike_name": "lizhishuang",
            "desc": "正经客户",
            "tag": "正经",
            "tel_phone": "16885885",
            "cell_phone": "186345750906",
            "email": "124379685@.qqcom",
            "is_vip": 1,
            "province": "山西",
            "city": "侯马",
            "source_channel": "api",
            "open_api_token": "1426-1995-0612-XzFG"
        }
        r = requests.post(self.create_customer, json.dumps(self.customer_data))
        return r.json()

    def test_create_customer(self):
        """
        测试创建1000条数据，json数据格式正确
        共请求2500次，期间可能出现重复api
            成功 code=0, msg=success
            失败 code=1, msg=the customer is registered
        """
        api_lst = []
        count = 0
        for i in range(2500):
            self.customer_data["customer_nike_name"] = "test{}".format(i)
            self.customer_data["desc"] = "测试客户0{}".format(i)
            self.customer_data["tag"] = "测试标签0{}".format(i)
            self.customer_data["tel_phone"] = "18688588{}".format(i)
            self.customer_data["cell_phone"] = "1360001{}".format(i)
            self.customer_data["email"] = "{}@qq.com".format(i)
            self.customer_data["is_vip"] = i % 2
            self.customer_data["province"] = random.choice(province_lst)
            self.customer_data["city"] = city_dict[self.customer_data["province"]]
            self.customer_data["source_channel"] = "api"
            self.customer_data["open_api_token"] = "{}-1995-1996-TEST".format(random.randint(1000, 2000))

            r = requests.post(self.create_customer, data=json.dumps(self.customer_data))
            res = r.json()

            if self.customer_data["open_api_token"] not in api_lst:
                api_lst.append(self.customer_data["open_api_token"])
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                count += 1
            else:
                self.assertEqual(res["code"], 1)
                self.assertEqual(res["msg"], "the customer is registered")

        r = requests.get(self.get_customer, params={"curr_page": 1, "page_size": 1000})
        res = r.json()
        col = len(res["customers"])
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(count, col)


    def test_create_para_err(self):
        """
        测试新增客户 无name 无api
        预计结果 code=1 msg=incoming parameter error
        """
        self.customer_data["open_api_token"] = ""
        r = requests.post(self.create_customer, data=json.dumps(self.customer_data))
        res = r.json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "incoming parameter error")
        self.assertEqual(res["customer"]["id"], 0)


    def test_new_follow(self):
        """
        新增客户跟进
        预计结果 code=0 msg="success"
        """
        res = self.new_customer()
        customer_id = res["customer"]["id"]
        self.follow_data["customer_id"] = customer_id
        r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
        res = r.json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(res["customer_follow_up"]["content"], "testing")

    def test_new_follow_cid_err(self):
        """
        新增客户跟进 cid 没写
        预计结果 code=1 msg="ncoming parameter error"
        """
        self.follow_data["customer_id"] = 0
        r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
        res = r.json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "incoming parameter error")

    def test_new_follow_uid_err(self):
        """
        新增客户跟进 uid 没写
        预计结果 code=1 msg="ncoming parameter error"
        """
        self.follow_data["uid"] = 0
        r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
        res = r.json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "incoming parameter error")


    def test_new_follow_content_err(self):
        """
        新增客户跟进 content 没写
        预计结果 code=1 msg="ncoming parameter error"
        """
        self.follow_data["uid"] = 1
        self.follow_data["content"] = ""
        r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
        res = r.json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "incoming parameter error")


    def test_new_follow_customer_none(self):
        """
        新增客户跟进 customer 不存在
        预计结果 code=1 msg="the customer to be followed up dose not exist"
        """
        self.follow_data["customer_id"] = 1
        r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
        res = r.json()

        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "the customer to be followed up dose not exist")


    def test_new_lot_follow(self):
        """
        新建很多follow code=0  msg=success
        """
        res_coustomer = self.new_customer()
        customer_id = res_coustomer["customer"]["id"]
        self.follow_data["customer_id"] = customer_id
        count = 0
        for i in range(999):
            r = requests.post(self.create_follow, data=json.dumps(self.follow_data))
            res = r.json()
            self.assertEqual(res["code"], 0)
            self.assertEqual(res["msg"], "success")
            count += 1
        r = requests.get(self.show_customer, params={"customer_id":customer_id})
        res = r.json()
        all_follow = len(res["customer"]["follow_up"])
        self.assertEqual(all_follow, count)


    def test_del_follow(self):
        """
        删除跟进 code=0  msg=success
        """
        res_customer = self.new_customer()
        customer_id = res_customer["customer"]["id"]
        self.follow_data["customer_id"] = customer_id
        r_follow = requests.post(self.create_follow, data=json.dumps(self.follow_data)).json()

        follow_id = r_follow["customer_follow_up"]["id"]
        post = {"id": follow_id}
        r = requests.post(self.delete_follow, data=json.dumps(post))
        res = r.json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")


    def test_del_follow_none(self):
        """
        删除跟进 code=1  msg=deleted customer follow up does not exist
        """
        post = {"id": 88888888888888}
        r = requests.post(self.delete_follow, data=json.dumps(post))
        res = r.json()
        print(res["msg"])
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "deleted customer follow up does not exist")


    def test_search_customer(self):
        """
        查询 客户 code=0 msg=success
        """
        customer = self.new_customer()
        customer_id = customer["customer"]["id"]
        customer_tk = customer["customer"]["open_api_token"]
        res = requests.get(self.show_customer, params={"customer_id": customer_id}).json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(res["customer"]["id"], customer_id)
        self.assertEqual(res["customer"]["open_api_token"], customer_tk)


    def test_search_customer_none(self):
        """
        查询 客户不存在 code=1 msg=customer dose not exist
        """
        res = requests.get(self.show_customer, params={"customer_id": 999999999999}).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "customer dose not exist")


    def test_search_customer_para_err(self):
        """
        查询 客户id错误 code=1 msg=incoming parameter error
        """
        res = requests.get(self.show_customer, params={"customer_id": 0}).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "incoming parameter error")


    def test_update_customer(self):
        """
        更新客户 code=0 msg=success
        """
        customer = self.new_customer()
        customer_id = customer["customer"]["id"]
        data = {
            "customer_id": customer_id,
            "tel_phone": "888888",
            "desc": "lzs"
        }
        res = requests.post(self.update_customer, data=json.dumps(data)).json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        res_search = requests.get(self.show_customer, params={"customer_id": customer_id}).json()
        self.assertEqual(res_search["customer"]["desc"], "lzs")
        self.assertEqual(res_search["customer"]["tel_phone"], "888888")


    def test_update_customer_none(self):
        """
        更新客户 code=1 msg=update customer failed, because customer_id dose not exist
        """
        data = {
            "customer_id": 8888888888888,
            "tel_phone": "888888",
            "desc": "lzs"
        }
        res = requests.post(self.update_customer, data=json.dumps(data)).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "update customer failed, because customer_id dose not exist")


    def test_page_customer(self):
        """
        客户分页
        """
        count = 0
        api_lst = []
        for i in range(2500):
            self.customer_data["customer_nike_name"] = "test{}".format(i)
            self.customer_data["open_api_token"] = "{}-1995-1996-TEST".format(random.randint(1000, 2000))
            res = requests.post(self.create_customer, data=json.dumps(self.customer_data)).json()
            if self.customer_data["open_api_token"] not in api_lst:
                api_lst.append(self.customer_data["open_api_token"])
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                count += 1
            else:
                self.assertEqual(res["code"], 1)
                self.assertEqual(res["msg"], "the customer is registered")

        for size in range(1, 500):
            all_page = count // size
            for page in range(20, 50):
                res = requests.get(self.get_customer, params={"curr_page": page, "page_size": size}).json()
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                if all_page - page <= 0:
                    if all_page - page == -1:
                        self.assertEqual(len(res["customers"]), count-all_page*size)
                    elif all_page - page == 0:
                        self.assertEqual(len(res["customers"]), size)

                    else:
                        self.assertEqual(len(res["customers"]), 0)
                else:

                    self.assertEqual(len(res["customers"]), size)

