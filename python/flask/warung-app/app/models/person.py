# class Person:
#     def __init__(self, fullname, address, birth, gender, user_id, phone_number):
#         self.fullname = fullname
#         self.user_id = user_id
#         self.address = address
#         self.birth = birth
#         self.gender = gender
#         self.phone_number = phone_number
from sqlalchemy import Column, String, Date
from infra.db.postgre.config import Base
from sqlalchemy.types import String

class Person(Base):
    __tablename__= "person"
    id = Column(String, primary_key=True, index=True)
    user_id = Column(String)
    fullname = Column(String)
    address = Column(String)
    birth = Column(Date)
    gender = Column(String)
    phone_number = Column(String)