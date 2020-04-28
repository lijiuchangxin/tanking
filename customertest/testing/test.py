# -*- coding: utf-8 -*-
import os
import unittest
from time import strftime, localtime, time

from testing.HTMLTestRunner import HTMLTestRunner
from testing.test_customer import TestCustomer
from testing.test_department import TestDepartment
from testing.test_group import TestGroup

suite = unittest.TestSuite()
# 获取TestSuite的实例对象
suite.addTest(TestCustomer("new_customer"))
suite.addTest(TestCustomer("test_search_customer"))
suite.addTest(TestCustomer("test_create_para_err"))
suite.addTest(TestCustomer("test_new_follow"))
suite.addTest(TestCustomer("test_new_follow_cid_err"))
suite.addTest(TestCustomer("test_new_follow_uid_err"))
suite.addTest(TestCustomer("test_new_follow_content_err"))
suite.addTest(TestCustomer("test_new_follow_customer_none"))
suite.addTest(TestCustomer("test_new_lot_follow"))
suite.addTest(TestCustomer("test_del_follow"))
suite.addTest(TestCustomer("test_del_follow_none"))
suite.addTest(TestCustomer("test_search_customer_none"))
suite.addTest(TestCustomer("test_search_customer_para_err"))
suite.addTest(TestCustomer("test_update_customer"))
suite.addTest(TestCustomer("test_update_customer_none"))
suite.addTest(TestCustomer("test_page_customer"))
# 把测试用例添加到测试容器中


# 获取TestSuite的实例对象
suite.addTest(TestDepartment("new_department"))
suite.addTest(TestDepartment("test_create_department"))
suite.addTest(TestDepartment("test_delete_department"))
suite.addTest(TestDepartment("test_get_department"))
suite.addTest(TestDepartment("test_update_department"))


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

filename = now + "test.html"
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

