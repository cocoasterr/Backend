from pymongo import MongoClient
from pymongo.collection import Collection as mongo_col
from bson.objectid import ObjectId
from pymongo.errors import OperationFailure
from exceptions import InvalidInputException, InternalServerErrorException


class Repository:
    def __init__(self, mongo_db: mongo_col, collection_db):
        self.mongo_db = mongo_db
        self.collection_db = collection_db

    def bulk_create(self, payloads):
        with MongoClient().start_session() as session:
            with session.start_transaction():
                try:
                    self.mongo_db.insert_many(
                        [self.collection_db(**payload).__dict__ for payload in payloads]
                    )
                    session.commit_transaction()
                except Exception as e:
                    session.abort_transaction()
                    raise InternalServerErrorException(f"Transaction aborted: {e}")

    def create(self, payload: dict):
        with MongoClient().start_session() as session:
            with session.start_transaction():
                try:
                    self.mongo_db.insert_one(payload)
                    session.commit_transaction()
                except Exception as e:
                    session.abort_transaction()
                    raise InternalServerErrorException(f"Transaction aborted: {e}")

    def update(self, payload, id):
        try:
            rec_id = ObjectId(id)
        except Exception as e:
            raise InvalidInputException(str(e))
        res = self.mongo_db.find_one({"_id": rec_id})
        if res:
            with MongoClient().start_session() as session:
                try:
                    with session.start_transaction():
                        self.mongo_db.update_one({"_id": rec_id}, {"$set": payload})
                        session.commit_transaction()
                        return res
                except Exception as e:
                    session.abort_transaction()
                    raise InternalServerErrorException(f"Transaction aborted: {e}")

    def delete(self, id):
        try:
            rec_id = ObjectId(id)
        except Exception as e:
            raise InvalidInputException(str(e))
        res = self.mongo_db.find_one_and_delete({"_id": rec_id})
        return res

    def find_by_id(self, id):
        try:
            rec_id = ObjectId(id)
        except Exception as e:
            raise InvalidInputException(str(e))
        res = self.mongo_db.find_one({"_id": rec_id})
        if res:
            res["_id"] = str(res["_id"])
        return res

    def index(self, page: int, limit: int):
        cursor = (
            self.mongo_db.find()
            .skip((page - 1) * limit)
            .limit(limit)
            .sort("created_at", -1)
        )
        result = []
        for res in cursor:
            if res["_id"]:
                res["_id"] = str(res["_id"])
            result.append(res)
        total = self.mongo_db.count_documents({})

        return result, total
