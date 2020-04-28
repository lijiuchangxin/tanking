# -*- coding: utf-8 -*-

"""
testing for group
"""
import json
import os
import unittest
import random
from testing.HTMLTestRunner import HTMLTestRunner
from testing.mysql_db import DB
import requests
from time import strftime, localtime, time


class TestGroup(unittest.TestCase):
    create = "http://127.0.0.1:8000/admin/v2/group/create"
    delete = "http://127.0.0.1:8000/admin/v2/group/delete"
    update = "http://127.0.0.1:8000/admin/v2/group/update"
    get = "http://127.0.0.1:8000/admin/v2/group/get"
    lst = "http://127.0.0.1:8000/admin/v2/group/list"
    add_agent = "http://127.0.0.1:8000/admin/v2/group/add_agent"
    del_agent = "http://127.0.0.1:8000/admin/v2/group/del_agent"

    db = DB()

    agent_lst = [6, 8, 9]

    def setUp(self):
        self.group_data = {
            "group_name": ""
        }
        self.db.clear("ut_employee_group")
        self.db.clear("ut_rel_employee_group")


    def new_group(self):
        self.group_data = {
            "group_name": "lzs_group"
        }
        r = requests.post(self.create, data=json.dumps(self.group_data))

        return r.json()

    def test_create_group(self):
        name_lst = []
        count = 0
        for i in range(200):
            self.group_data["group_name"] = "group_{}".format(random.randint(1000, 2000))
            r = requests.post(self.create, data=json.dumps(self.group_data))
            res = r.json()

            if self.group_data["group_name"] not in name_lst:
                name_lst.append(self.group_data["group_name"])
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                count += 1
            else:
                self.assertEqual(res["code"], 2)
                self.assertEqual(res["msg"], "create group failed")

        r = requests.post(self.lst, data=json.dumps({"curr_page": 1, "page_size": 10000000}))
        res = r.json()
        col = len(res["groups"])
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(count, col)


    def test_del_group(self):
        name_lst = []
        # 正常生成删除
        for i in range(200):
            self.group_data["group_name"] = "group_{}".format(random.randint(1000, 2000))
            r = requests.post(self.create, data=json.dumps(self.group_data))
            res = r.json()

            name_lst.append(self.group_data["group_name"])
            self.assertEqual(res["code"], 0)
            self.assertEqual(res["msg"], "success")
            r = requests.post(self.delete, data=json.dumps({"group_id": res["user_group"]["group_id"]})).json()
            self.assertEqual(r["code"], 0)
            self.assertEqual(r["msg"], "success")

        # 删除的id不存在
        r = requests.post(self.delete, data=json.dumps({"group_id": 1})).json()
        self.assertEqual(r["code"], 1)
        self.assertEqual(r["msg"], "group dose not exist")

        # 输入的参数有毛病
        r = requests.post(self.delete, data=json.dumps({"group_id": "1"})).json()
        self.assertEqual(r["code"], 1)
        self.assertEqual(r["msg"], "incoming parameter error")


    def test_update_group(self):
        res_new = self.new_group()
        gid = res_new["user_group"]["group_id"]
        for i in range(20):
            name = "{}test".format(random.randint(999, 9999))
            data = {
                "group_id": gid,
                "group_name": name
            }
            # 修改
            res = requests.post(self.update, data=json.dumps(data)).json()
            self.assertEqual(res["code"], 0)
            self.assertEqual(res["msg"], "success")

            # 查出来康康修改成功没
            res_search = requests.get(self.get, params={"group_id": gid}).json()
            self.assertEqual(res_search["code"], 0)
            self.assertEqual(res_search["msg"], "success")
            self.assertEqual(res_search["user_group"]["group_name"], name)

        # 修改不存在的组
        data["group_id"] = 1
        res = requests.post(self.update, data=json.dumps(data)).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "group dose not exist")

    def test_get_group(self):
        res = self.new_group()
        name = res["user_group"]["group_name"]
        id = res["user_group"]["group_id"]

        # 正常get
        res_search = requests.get(self.get, params={"group_id": id}).json()
        self.assertEqual(res_search["code"], 0)
        self.assertEqual(res_search["msg"], "success")
        self.assertEqual(res_search["user_group"]["group_name"], name)

        # get 不存在的组
        res_search = requests.get(self.get, params={"group_id": 1}).json()
        self.assertEqual(res_search["code"], 1)
        self.assertEqual(res_search["msg"], "group dose not exist")
        self.assertEqual(res_search["user_group"]["group_name"], "")


    def test_lst_group(self):
        count = 0
        name_lst = []
        for i in range(200):
            self.group_data["group_name"] = "test{}".format(random.randint(1000, 2000))
            res = requests.post(self.create, data=json.dumps(self.group_data)).json()
            if self.group_data["group_name"] not in name_lst:
                name_lst.append(self.group_data["group_name"])
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                count += 1
            else:
                self.assertEqual(res["code"], 2)
                self.assertEqual(res["msg"], "create group failed")

        for size in range(1, 200):
            all_page = count // size
            for page in range(20, 50):
                # res = requests.get(self.lst, params={"curr_page": page, "page_size": size}).json()
                data = {
                    "curr_page": page,
                    "page_size": size
                }
                res = requests.post(self.lst, data=data).json()
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                if all_page - page <= 0:
                    if all_page - page == -1:
                        self.assertEqual(len(res["groups"]), count-all_page*size)
                    elif all_page - page == 0:
                        self.assertEqual(len(res["groups"]), size)

                    else:
                        self.assertEqual(len(res["groups"]), 0)
                else:
                    len_group = len(res["groups"].keys())
                    print(size, page, all_page, len_group)

                    self.assertEqual(len_group, size)

    def test_add_agent(self):
        name_lst = []
        gid_dict = {}
        gid_lst = []
        # 生成组
        for i in range(100):
            name = "{}group".format(random.randint(10, 20))
            self.group_data["group_name"] = name
            res = requests.post(self.create, data=json.dumps(self.group_data)).json()
            if name not in name_lst:
                name_lst.append(name)
                gid_dict[res["user_group"]["group_id"]] = 0
                gid_lst.append(name)
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
            else:
                self.assertEqual(res["code"], 2)
                self.assertEqual(res["msg"], "create group failed")

        # 组存在
        # 库中员工id为6  8  9
        for i in range(200):
            gid = random.choice(list(gid_dict.keys()))
            aid = random.randint(1, 10)
            data = {
                "group_id": gid,
                "agent_id": aid
            }
            r = requests.post(self.add_agent, data=json.dumps(data)).json()
            if aid in self.agent_lst:
                gid_dict[gid] = gid_dict[gid] + 1
                self.assertEqual(r["code"], 0)
                self.assertEqual(r["msg"], "success")
            else:
                self.assertEqual(r["code"], 2)
                self.assertEqual(r["msg"], "add agent failed")

        # 组不存在
        r = requests.post(self.add_agent, data=json.dumps({"group_id":123, "agent_id": 6})).json()
        self.assertEqual(r["code"], 1)
        self.assertEqual(r["msg"], "add agent failed, because group dose not exist")


    def test_del_agent(self):
        # 调用上个用例
        self.test_add_agent()
        # res = requests.get(self.lst, params={"curr_page": 1,  "page_size": 9999}).json()
        res = requests.post(self.lst, data=json.dumps({"curr_page": 1,  "page_size": 9999})).json()

        gid_lst = [group["group_id"] for group in res["groups"].values()]
        for i in range(200):
            aid = random.randint(1, 10)
            data = {
                "group_id": random.choice(gid_lst),
                "agent_id": aid
            }
            r = requests.post(self.del_agent,  data=json.dumps(data)).json()
            if aid not in self.agent_lst:
                self.assertEqual(r["code"], 2)
                self.assertEqual(r["msg"], "remove agent failed")
            else:
                self.assertEqual(r["code"], 0)
                self.assertEqual(r["msg"], "success")

if __name__ == "__main__":

    suite = unittest.TestSuite()
    # 获取TestSuite的实例对象
    suite.addTest(TestGroup("test_del_agent"))
    suite.addTest(TestGroup("test_add_agent"))
    suite.addTest(TestGroup("test_lst_group"))
    suite.addTest(TestGroup("test_get_group"))
    suite.addTest(TestGroup("test_update_group"))
    suite.addTest(TestGroup("test_del_group"))
    suite.addTest(TestGroup("test_create_group"))
    suite.addTest(TestGroup("new_group"))

    now = strftime("%Y-%m-%M-%H_%M_%S", localtime(time()))
    # 获取当前时间

    filename = now + "test_group.html"
    # 文件名

    filename = os.path.join(os.path.abspath(".."), "res", filename)

    fp = open(filename, 'wb')
    # 以二进制的方式打开文件并写入结果

    runner = HTMLTestRunner(
        stream=fp,
        verbosity=2,
        title="测试报告的标题",
        description="测试报告的详情")

    runner.run(suite)
