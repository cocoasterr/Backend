from infra.db.redis.configRedist import RedisConfig
from redis_om import HashModel

db = RedisConfig()


class Order(HashModel):
    product_id: str
    price: int
    fee: int
    total: int
    qty: int
    status: str

    class Meta:
        database = db.redis
