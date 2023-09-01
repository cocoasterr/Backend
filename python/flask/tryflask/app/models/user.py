from bson.objectid import ObjectId
from datetime import datetime


class User:
    def __init__(
        self,
        username,
        email,
        password,
        _id=ObjectId(),
        created_at=None,
        updated_at=datetime.utcnow(),
    ):
        self._id = _id
        self.username = username
        self.email = email
        self.password = password
        self.created_at = created_at if created_at else datetime.now()
        self.updated_at = updated_at
