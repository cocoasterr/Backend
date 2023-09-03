from sqlalchemy import Column, String
from infra.db.postgre.config import Base
from sqlalchemy.types import String

class Users(Base):
    __tablename__= "users"
    id = Column(String, primary_key=True, index=True)
    username = Column(String)
    email = Column(String)
    password = Column(String)
    status = Column(String)
