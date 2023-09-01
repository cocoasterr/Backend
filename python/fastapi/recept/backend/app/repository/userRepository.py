from app.models.users import Users
from app.repository.BaseRepo import BaseRepo
from app.database import db

class userRepo(BaseRepo):
    model = Users

    @staticmethod
    async def findUser(email:str = None, username:str = None):
        query = f"SELECT * FROM users WHERE email='{email}'"
        if username:
            query += "OR username = '{username}'"
        DB = await db.conn()
        res = DB.exec_driver_sql(query).one_or_none()
        return res
    
    @staticmethod
    async def findUserLogin(input:str = None):
        try:
            query = f"SELECT * FROM users WHERE email='{input}'"
            DB = await db.conn()
            res = DB.exec_driver_sql(query).one_or_none()
            return res
        except Exception:
            query = f"SELECT * FROM users WHERE username='{input}'"
            DB = await db.conn()
            res = DB.exec_driver_sql(query).one_or_none()
            return res

    