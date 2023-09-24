from infra.db.postgre.config import db


class BaseRepo:
    def __init__(self, model):
        self.model = model

    @classmethod
    def create_bulk(self, payloads):
        session = db.session()
        with session.begin() as trx:
            try:
                trx.execute(self.model.__table__.insert().values(payloads))
                trx.flush()
                trx.commit()
            except Exception as e:
                trx.rollback()
                return e


    @classmethod
    def create(self, **kwargs):
        session = db.session()
        with session.begin() as trx:
            try:
                data = self.model(**kwargs)
                trx.add(data)
                trx.commit()
                trx.flush()
            except Exception as e:
                trx.rollback()
                return e


    @classmethod
    def find_by(self, condition):
        try:
            conn = db.conn()
            tb_name = self.model.__tablename__
            query = f"SELECT * FROM {tb_name} where {condition}"
            res = conn.exec_driver_sql(query).one_or_none()
            return res
        except Exception as e:
            raise e
    # @classmethod
    # def find_by(self, query_condition:str=None, operator:bool=None, list_key_value:[]=None):
    #     try:
    #         conn = db.conn()
    #         if isinstance(value, str):
    #             value=f"'{value}'"
    #         tb_name = self.model.__tablename__
    #         query = f"SELECT * FROM {tb_name} where {key} = {value}"
    #         res = conn.exec_driver_sql(query).all()
    #         return res[0]
    #     except Exception as e:
    #         raise e