from exception import Internal_Server_error
from datetime import datetime


class PGService:
    def __init__(self, repo):
        self.repo = repo

    @classmethod
    def create_service(self, payloads, user):
        try:
            res = []
            for payload in payloads:
                now = datetime.now()
                payload["created_at"] = now
                payload["updated_at"] = now
                payload["username"] = user["username"]
                res.append(**payload)
            self.repo.create(**payload)
        except Exception as e:
            return e

    @classmethod
    def create_service_without_token(self, payload):
        try:
            self.repo.create(**payload)
        except Exception as e:
            return e

    @classmethod
    def find_by(self, key, value):
        try:
            return self.repo.find_by(key, value)
        except Exception as e:
            raise 
    
    def row_to_dict(self, row):
        res = {}
        for key,value in zip(row._fields, row._data):
            res[key] = value
        return res