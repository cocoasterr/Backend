from bson.objectid import ObjectId
from datetime import datetime


class Agent:
    def __init__(
        self,
        product_name,
        username,
        category,
        stock=0,
        _id=ObjectId(),
        created_at=None,
        updated_at=datetime.utcnow(),
    ):
        self._id = _id
        self.product_name = product_name
        self.stock = stock
        self.category = category
        self.username = username
        self.created_at = created_at if created_at else datetime.now()
        self.updated_at = updated_at
