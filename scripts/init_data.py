import logging

from init_schema import execute_sql_file, load_sql_files


def init_data(conn, files=None):
    if files is None:
        files = load_sql_files("data")

    logging.info("开始导入初始数据")
    for path, content in files:
        execute_sql_file(conn, path, content)
    logging.info("全部初始数据导入完成")
