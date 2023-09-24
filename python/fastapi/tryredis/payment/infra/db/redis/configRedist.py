from redis_om import get_redis_connection


class RedisConfig:
    redis = get_redis_connection(
        host="redis-14218.c292.ap-southeast-1-1.ec2.cloud.redislabs.com",
        port="14218",
        password="VQKu94m2u5e6KfLZVH3ocHkVAxzQor9n",
        decode_responses=True,
    )
