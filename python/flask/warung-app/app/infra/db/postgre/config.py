from sqlalchemy import Engine, create_engine
from sqlalchemy.orm import Session, sessionmaker
from sqlalchemy.ext.declarative import declarative_base
import os


class DatabaseSession:
    def __init__(self):
        self.psql_engine = None
        self.psql_session = None

    def init(self):
        db_url = os.getenv("DB_URL_POSTGRES")
        try:
            self.psql_engine = create_engine(str(db_url))
            self.psql_session = sessionmaker(bind=self.psql_engine, expire_on_commit=False, autoflush=False)
            print("Connected to Postgres!\n\n")
        except Exception as e:
            print("Unable to connect to Postgres!\n\n")
            print("error : ", e)

    def conn(self):
        try:
            conn = self.psql_engine.connect()
            return conn
        except Exception as e:
            print("Unable to connect to Postgres!")
            
    def session(self):
        try:
            return sessionmaker(bind=self.psql_engine, expire_on_commit=False, autoflush=False)
        except Exception as e:
            print("Unable to connect to Postgres!")

Base = declarative_base()
db = DatabaseSession()

# def get_db():
#     try:
#         data = db.session()
#         return data
#     finally:
#         data.close()


def commit_rollback(session):
    try:
        session.commint()
        session.flush()
    except Exception as e:
        session.rollback()
