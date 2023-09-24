from infra.db.redis.configRedist import RedisConfig
from redis_om import HashModel

db = RedisConfig()


class Product(HashModel):
    name: str
    qty: int
    price: int

    class Meta:
        database = db.redis
