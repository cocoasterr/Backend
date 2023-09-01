from typing import Generic, TypeVar
from app.database import db, commit_rollback
from sqlalchemy import update


class BaseRepo:
    model = Generic[TypeVar('T')]
    
    @classmethod
    async def create(cls, session, **kwargs):
        model = cls.model(**kwargs)
        session.add(model)
        commit_rollback(session)
        return model
    
    @classmethod
    async def bulkCreate(cls, session, new_data:list):
        session.bulk_save_objects(new_data, return_defaults=True)
        commit_rollback(session)
        return 'success!'

    @staticmethod
    async def getById(id: str, tb_name:str):
        query = f"SELECT * FROM {tb_name} WHERE id = '{id}'"
        DB = await db.conn()
        res = DB.exec_driver_sql(query).one_or_none()
        return res
    
    @classmethod
    async def update(cls, session, model_id: str, new_data):
        data_query = session.query(cls.model).filter(cls.model.id == model_id)
        data_query.update(new_data, synchronize_session=False)
        commit_rollback(session)
        return 'success!'
    
    @staticmethod
    async def getAll(tb_name:str, page, limit, search):
        offset = (page - 1) * limit
        query = f"SELECT * FROM {tb_name} {search} ORDER BY created_at DESC LIMIT {limit} OFFSET {offset}"
        DB = await db.conn()
        res = DB.exec_driver_sql(query).all()
        query_total = f"SELECT COUNT(id) FROM {tb_name}"
        total = DB.exec_driver_sql(query_total).scalar()
        return res, total
    
    @classmethod
    async def delete(cls, id: str, tb_name:str):
        query_del = f"delete FROM {tb_name} WHERE id = '{id}'"
        DB = await db.conn()
        DB.exec_driver_sql(query_del)
        return 'success!'