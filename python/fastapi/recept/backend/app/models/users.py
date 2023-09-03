
from sqlalchemy import Column, String
from sqlalchemy.types import String
from app.database import Base


class Users(Base):
    __tablename__= "users"
    id = Column(String, primary_key=True, index=True)
    username = Column(String)
    email = Column(String)
    password = Column(String)
    person_id = Column(String)
