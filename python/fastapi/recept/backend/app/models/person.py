

from sqlalchemy import Enum, Column
from app.database import Base
from sqlalchemy.types import String


class Sex(str, Enum):
    MALE = "MALE"
    FEMALE = "FEMALE"


class Person(Base):
    __tablename__ = "person"
    id = Column(String, primary_key=True, index=True)
    name = Column(String)
    birth = Column(String)
    sex = Column(String)
    profile = Column(String)
    phone_number = Column(String)

