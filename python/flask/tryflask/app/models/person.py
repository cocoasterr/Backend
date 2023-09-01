from bson.objectid import ObjectId
from datetime import datetime


class Person:
    def __init__(
        self,
        fullname,
        address,
        phone_number,
        username,
        birth,
        _id=ObjectId(),
        created_at=None,
        updated_at=datetime.utcnow(),
    ):
        self._id = _id
        self.fullname = fullname
        self.address = address
        self.username = username
        self.phone_number = phone_number
        self.birth = birth
        self.created_at = created_at if created_at else datetime.now()
        self.updated_at = updated_at
