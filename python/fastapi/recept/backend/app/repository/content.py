from app.database import db
from app.repository.BaseRepo import BaseRepo
from app.models.content import Content

class contentRepo(BaseRepo):
    model = Content

    @staticmethod
    async def searchByTitle(tb_name:str, key:str, value:str = "", page:int=1, limit:int = 10):
        offset = (page - 1) * limit
        query = f"SELECT * FROM {tb_name} WHERE {key} LIKE '%{value}%' LIMIT {limit} OFFSET {offset}"
        DB = await db.conn()
        res = DB.exec_driver_sql(query).all()
        return res
