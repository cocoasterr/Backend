import os
from pymongo import mongo_client


# MongoDB
db_url = os.getenv("MONGO_DATABASE_URL")
db_name = os.getenv("MONGO_INITDB_DATABASE")
mongodb_client = mongo_client.MongoClient(db_url, serverSelectionTimeoutMS=5000)

try:
    mongodb_connection = mongodb_client.server_info()
    print(f'Connected to MongoDB {mongodb_connection.get("version")}')
except Exception:
    print("Unable to connect to the MongoDB server.")

mongo_db = mongodb_client[str(db_name)]

# create tb
tb_product = mongo_db.tb_product
tb_person = mongo_db.tb_person
tb_user = mongo_db.tb_user
