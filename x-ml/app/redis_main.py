from redis import Redis
from settings import config


redis_conn = Redis(host=config.REDIS_HOST, port=config.REDIS_PORT, password=config.REDIS_PASS, db= config.REDIS_DB)
