from infra.db.postgre.config import db


class BaseRepo:
    # DB = get_db()
    # DB = get_db()
    # DB = db.session()
    # conn = db.conn()

    def __init__(self, model):
        self.model = model

    @classmethod
    def create_bulk(self, payloads):
        with self.DB.begin() as trx:
            try:
                # trx = self.DB.begin()
                self.DB.execute(self.model.__table__.insert().values(payloads))
                self.DB.flush()
                trx.commit()
            except Exception as e:
                trx.rollback()

    @classmethod
    def create(self, **kwargs):
        DB = db.session()
        with self.DB.begin() as trx:
            try:
                data = self.model(**kwargs)
                trx.add(data)
                trx.commit()
                trx.flush()
            except Exception as e:
                trx.rollback()
                return e

    @classmethod
    def find_by(self, key, value):
        try:
            conn = db.conn()
            if isinstance(value, str):
                value=f"'{value}'"
            tb_name = self.model.__tablename__
            query = f"SELECT * FROM {tb_name} where {key} = {value}"
            res = conn.exec_driver_sql(query).all()
            return res[0]
        except Exception as e:
            raise e