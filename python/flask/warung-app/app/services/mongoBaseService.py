from bson.objectid import ObjectId
from exception import *

class MongoService:
    def __init__(self, repository):
        self.repository = repository

    def bulk_create_service(self, payloads):
        try:
            self.repository.bulk_create(payloads)
        except Exception as e:
            raise e

    def create_service(self, payload):
        try:
            self.repository.create(payload)
        except Exception as e:
            raise e

    def update_service(self, payloads, id: str):
        try:
            rec_id = ObjectId(id)
            res = self.repository.update(payloads, rec_id)
            if not res:
                raise Bad_Request
        except Exception as e:
            raise e

    def delete_service(self, id: str):
        try:
            rec_id = ObjectId(id)
            res = self.repository.delete(rec_id)
            if not res:
                raise Bad_Request
        except Exception as e:
            raise e

    def find_by_id_service(self, id: str):
        try:
            rec_id = ObjectId(id)
            res = self.repository.find_by_id(rec_id)
            if not res:
                raise Bad_Request
            return res
        except Exception as e:
            raise e

    def index_service(self, page: int, limit: int):
        try:
            return self.repository.index(page, limit)
        except Exception as e:
            raise e
