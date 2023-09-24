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
            return self.repo.create(**payload)
        except Exception as e:
            return e

    @classmethod
    def find_by(self, key:str=None, value:str=None, custom_condition:str=None, operator:str=None, key_value:dict=None):
        try:
            if key and value:
                if isinstance(value, str):
                    value = f"'{value}'"
                condition = f"{key}={value}"
            elif custom_condition:
                condition = custom_condition
            elif operator:
                res = []
                for i in key_value.items():
                    key = i[0]
                    value = i[1]
                    if isinstance(value, str):
                        value = f"'{value}'"
                    res.append(f"{key}={value}")
                condition = f" {operator} ".join(res)
            else:
                return None
            return self.repo.find_by(condition)
        except Exception as e:
            raise 
    
    def row_to_dict(self, row):
        res = {}
        for key,value in zip(row._fields, row._data):
            res[key] = value
        return res