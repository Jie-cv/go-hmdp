
import logging
import os
from pathlib import Path

import pymysql
from dotenv import load_dotenv
from pymysql.constants import CLIENT


PROJECT_ROOT = Path(__file__).resolve().parent.parent
SQL_ROOT = Path(__file__).resolve().parent / "sql"

# 显式指定顺序，不依赖文件系统遍历顺序。
MODULE_ORDER = ("user", "shop", "blog", "voucher")


def load_sql_files(group):
    files = []
    for module in MODULE_ORDER:
        path = SQL_ROOT / group / f"{module}.sql"
        content = path.read_text(encoding="utf-8-sig")
        if not content.strip():
            raise ValueError(f"SQL 文件为空：{path}")
        files.append((path, content))
    return files


def connect_database():
    # 通过脚本位置找 .env，从其他工作目录运行也能找到配置。
    # 已有环境变量优先于 .env。
    load_dotenv(PROJECT_ROOT / ".env")

    user = os.getenv("MYSQL_USER", "").strip()
    db_name = os.getenv("MYSQL_DB_NAME", "").strip()
    if not user or not db_name:
        raise ValueError("请在根目录 .env 中设置 MYSQL_USER 和 MYSQL_DB_NAME")

    conn = pymysql.connect(
        host=os.getenv("MYSQL_HOST") or "127.0.0.1",
        port=int(os.getenv("MYSQL_PORT") or "3306"),
        user=user,
        password=os.getenv("MYSQL_PASSWORD", ""),
        charset=os.getenv("MYSQL_CHARSET") or "utf8mb4",
        connect_timeout=10,
        autocommit=True,
        # 仅初始化连接允许整份 SQL 文件包含多条语句。
        client_flag=CLIENT.MULTI_STATEMENTS,
    )
    try:
        # 数据库可能不存在，先连接服务器，再创建并选择数据库。
        # 标识符不能用 %s 参数绑定，需单独转义反引号。
        quoted_name = "`" + db_name.replace("`", "``") + "`"
        with conn.cursor() as cursor:
            cursor.execute(
                f"CREATE DATABASE IF NOT EXISTS {quoted_name} "
                "CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci"
            )
        conn.select_db(db_name)
        logging.info("已连接数据库：%s", db_name)
        return conn
    except Exception:
        conn.close()
        raise


def execute_sql_file(conn, path, content):
    logging.info("正在执行：%s/%s", path.parent.name, path.name)
    try:
        with conn.cursor() as cursor:
            cursor.execute(content)
            # execute 只处理第一个结果；必须消费后续结果，
            # 才能确认本文件全部语句完成，并捕获后续语句的错误。
            # 由 MySQL 处理分号、引号和转义，不自行 split(';')。
            while cursor.nextset():
                pass
    except Exception as exc:
        # 数据文件自带 START TRANSACTION / COMMIT。
        # 文件失败时回滚当前未提交事务，之前文件的提交仍保留。
        conn.rollback()
        raise RuntimeError(f"执行 {path.parent.name}/{path.name} 失败：{exc}") from exc


def init_schema(conn, files=None):
    if files is None:
        files = load_sql_files("schema")

    # 原始 SQL 面向空库；禁止重复执行造成重复建表或重复插入。
    with conn.cursor() as cursor:
        cursor.execute(
            "SELECT COUNT(*) FROM information_schema.tables "
            "WHERE table_schema = DATABASE()"
        )
        table_count = cursor.fetchone()[0]
    if table_count:
        raise RuntimeError(f"当前数据库已有 {table_count} 张表，请使用空库初始化")

    logging.info("开始创建表结构")
    for path, content in files:
        execute_sql_file(conn, path, content)
    logging.info("全部表结构创建完成")
