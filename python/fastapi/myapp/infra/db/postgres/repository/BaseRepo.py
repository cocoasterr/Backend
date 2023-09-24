from infra.db.postgres.config import Config
from exception import *


class BaseRepo:
    def __init__(self, model) -> None:
        self.model = model

    config = Config()
    connection = config.connection()
    cursor = connection.cursor()

    async def create(self, payload):
        key = []
        value = ()
        values = []
        for data in payload:
            key.append(data[0])
            value = value + data[1]
            values.append("%s")
        query = (
            f"INSERT INTO {self.model} ({', '.join(key)}) VALUES ({', '.join(values)})"
        )
        try:
            self.cursor.execute(query, value)
            self.connection.commit()
        except Exception as e:
            self.connection.rollback()
            raise Internal_Server_error(e)
