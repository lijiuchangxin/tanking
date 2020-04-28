# -*- coding: utf-8 -*-

"""
testing for department
"""
import json
import os
import unittest
import random
import requests
from time import strftime, localtime, time

from testing.HTMLTestRunner import HTMLTestRunner
from testing.mysql_db import DB


class TestDepartment(unittest.TestCase):
    create = "http://127.0.0.1:8000/admin/v2/department/create"
    delete = "http://127.0.0.1:8000/admin/v2/department/delete"
    get = "http://127.0.0.1:8000/admin/v2/department/get"
    update = "http://127.0.0.1:8000/admin/v2/department/update"
    lst = "http://127.0.0.1:8000/admin/v2/department/lst"
    db = DB()
    agent_lst = [6, 8, 9]

    department_data = {
        "department_name": " ",
        "parent_id": 0
    }

    def setUp(self):
        self.department_data = {
            "department_name": ""
        }
        self.db.clear("ut_department")


    def new_department(self):
        data = {
            "department_name": "lzstest"
        }
        r = requests.post(self.create, data=json.dumps(data)).json()
        return r

    def test_create_department(self):
        name_lst = []
        count = 0
        for i in range(200):
            self.department_data["department_name"] = "department_{}".format(random.randint(1, 50))
            r = requests.post(self.create, data=json.dumps(self.department_data))
            res = r.json()

            if self.department_data["department_name"] not in name_lst:
                name_lst.append(self.department_data["department_name"])
                self.assertEqual(res["code"], 0)
                self.assertEqual(res["msg"], "success")
                count += 1
            else:
                self.assertEqual(res["code"], 1)
                self.assertEqual(res["msg"], "department already exist")

        r = requests.get(self.lst)
        res = r.json()
        col = len(res["departments"])
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(count, col)



    def test_delete_department(self):
        # 删除新增的 无级联关系
        r = self.new_department()
        res = requests.post(self.delete, data=json.dumps({"department_id": r["department_id"]})).json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")

        # 删除不存在的部门
        res = requests.post(self.delete, data=json.dumps({"department_id": 1})).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "department dose not exist")

        # 删除带有级联关系的部门
        pid = self.new_department()["department_id"]
        data = {
            "department_name": "test2",
            "parent_id": pid
        }
        requests.post(self.create, data=json.dumps(data)).json()
        res = requests.post(self.delete, data=json.dumps({"department_id": pid})).json()
        self.assertEqual(res["code"], 2)
        self.assertEqual(res["msg"], "delete department failed, maybe have child department or employee")


    def test_get_department(self):
        pid = self.new_department()["department_id"]
        data = {
            "department_name": "",
            "parent_id": pid
        }
        count = 0
        for i in range(100):
            data["department_name"] = "{}test".format(random.randint(1, 99999))
            r = requests.post(self.create, data=json.dumps(data)).json()
            if r["code"] == 0:
                count += 1

        res = requests.get(self.get, params={"department_id": pid}).json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(len(res["user_department"]["departments"]), count)


    def test_update_department(self):
        res = self.new_department()
        data = {
            "department_name":  "jian_fei",
            "department_id": res["department_id"]
        }
        res = requests.post(self.update, data=json.dumps(data)).json()
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        res = requests.get(self.get, params={"department_id": res["department_id"]}).json()
        did = res["user_department"]["department_id"]
        self.assertEqual(res["code"], 0)
        self.assertEqual(res["msg"], "success")
        self.assertEqual(res["user_department"]["department_name"], data["department_name"])

        data["department_id"] = 1
        res = requests.get(self.get, params={"department_id": data["department_id"]}).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "department dose not exist")

        # 修改重复的name
        new_data = {
            "department_name": "jian_fei",
            "department_id": did
        }
        res = requests.post(self.update, data=json.dumps(new_data)).json()
        self.assertEqual(res["code"], 1)
        self.assertEqual(res["msg"], "department name repeat")

    # def test_lst_department(self):

if __name__ == "__main__":

    suite = unittest.TestSuite()
    # 获取TestSuite的实例对象
    # 获取TestSuite的实例对象
    suite.addTest(TestDepartment("new_department"))
    suite.addTest(TestDepartment("test_create_department"))
    suite.addTest(TestDepartment("test_delete_department"))
    suite.addTest(TestDepartment("test_get_department"))
    suite.addTest(TestDepartment("test_update_department"))

    now = strftime("%Y-%m-%M-%H_%M_%S", localtime(time()))
    # 获取当前时间

    filename = now + "test_department.html"
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
