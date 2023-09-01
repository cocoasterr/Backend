from bson.objectid import ObjectId
from datetime import datetime


class Product:
    def __init__(
        self,
        name,
        username,
        stock=0,
        _id=ObjectId(),
        created_at=None,
        updated_at=datetime.utcnow(),
    ):
        self._id = _id
        self.name = name
        self.stock = stock
        self.username = username
        self.created_at = created_at if created_at else datetime.now()
        self.updated_at = updated_at
