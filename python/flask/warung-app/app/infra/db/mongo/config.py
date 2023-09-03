import os
from pymongo import mongo_client


db_url = os.getenv("MONGO_DATABASE_URL")
db_name = os.getenv("MONGO_INITDB_DATABASE")

mongodb_client = mongo_client.MongoClient(db_url, serverSelectionTimeoutMS=5000)


def test_conn():
    try:
        test_conn = mongodb_client.server_info()
        print(f"\nConnected to Mongo DB {test_conn.get('version')}\n\n")
    except Exception:
        print("Unable to connect Mongo DB")


mongo_db = mongodb_client[str(db_name)]

# create tb
tb_wh_supplier = mongo_db.tb_wh_supplier
tb_wh_agent = mongo_db.tb_wh_agent
