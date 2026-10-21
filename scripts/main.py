import logging
import sys

from init_data import init_data
from init_schema import connect_database, init_schema, load_sql_files


def main():
    logging.basicConfig(level=logging.INFO, format="%(levelname)s: %(message)s")
    conn = None
    try:
        # 写数据库之前，确认两组 SQL 都存在且能读取。
        schema_files = load_sql_files("schema")
        data_files = load_sql_files("data")
        conn = connect_database()

        # 顺序执行：所有表创建成功后，才开始导入数据。
        init_schema(conn, schema_files)
        init_data(conn, data_files)
        logging.info("数据库初始化完成")
        return 0
    except Exception as exc:
        logging.error("初始化失败：%s", exc)
        logging.error("已完成的建表和已提交的数据不会自动撤销；修复后请使用空库重试。")
        return 1
    finally:
        if conn is not None:
            conn.close()


if __name__ == "__main__":
    sys.exit(main())
