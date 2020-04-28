# -*- coding: utf-8 -*-
"""
数据库注册连接
"""
import pymysql
import os
import configparser as parser

base_dir = str(os.path.dirname(os.path.dirname(__file__)))
base_dir = base_dir.replace('\\', '/')
file_path = base_dir + "/db_config.ini"

cf = parser.ConfigParser()
cf.read(file_path)

host = cf.get("config", "host")
port = cf.get("config", "port")
db = cf.get("config", "db_name")
user = cf.get("config", "user")
password = cf.get("config", "password")


class DB:
    def __init__(self):
        try:
            self.connection = pymysql.connect(
            host=host,
               port=int(port),
               user=user,
               password=password,
               db=db,
               charset='utf8mb4',
               cursorclass=pymysql.cursors.DictCursor)

        except pymysql.err.OperationalError as e:
            print("mysql err %d: %s" % (e.args[0], e.args[1]))


    def clear(self, table_name):
        # real_sql = "truncate table " + table_name + ";"
        real_sql = "delete from " + table_name + ";"
        with self.connection.cursor() as cursor:
            cursor.execute("SET FOREIGN_KEY_CHECKS=0;")
            cursor.execute(real_sql)
        self.connection.commit()

    def insert(self, table_name, table_data):
        for key in table_data:
            table_data[key] = "'" + str(table_data[key]) + "'"
        key = ','.join(table_data.keys())
        value = ','.join(table_data.values())
        real_sql = "INSERT INTO " + table_name + " (" + key + ") VALUES (" + value + ")"
        # print(real_sql)

        with self.connection.cursor() as cursor:
            cursor.execute(real_sql)

        self.connection.commit()

    def init_data(self, datas):
        for table, data in datas.items():
            self.clear(table)
            for d in data:
                self.insert(table, d)
        self.close()

    # close database
    def close(self):
        self.connection.close()


if __name__ == '__main__':
    pass