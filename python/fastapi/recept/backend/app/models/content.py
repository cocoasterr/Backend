from sqlalchemy import Column
from app.database import Base
from sqlalchemy.types import String, Integer


class Content(Base):
    __tablename__ = "receipt"
    id = Column(String, primary_key=True, index=True)
    user_id = Column(String)
    username = Column(String)
    title = Column(String)
    like = Column(Integer)
    comment = Column(String)
    status = Column(String)
    ingridient = Column(String)
    intructions = Column(String)
    notes = Column(String)
    created_at = Column(Integer)
    updated_at = Column(Integer)

