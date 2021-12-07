from redis import Redis
from redis.exceptions import ConnectionError
from settings import config

def get_redis_password() -> str:
    return config.REDIS_PASS


def get_redis_port() -> str:
    return config.REDIS_PORT

def get_redis_dbnum() -> str:
    return config.REDIS_DB


def get_redis_host() -> str:
    return config.REDIS_HOST

def get_backend_url() -> str:
    pw = get_redis_password()
    port = get_redis_port()
    db = get_redis_dbnum()
    return "redis://{password}{hostname}{port}{db}".format(
        hostname=get_redis_host(),
        password=':' + pw + '@' if len(pw) != 0 else '',
        port=':' + port if len(port) != 0 else '',
        db='/' + db if len(db) != 0 else ''
    )


def is_backend_running() -> bool:
    print(get_backend_url())
    try:
        conn = Redis(
            host=get_redis_host(),
            port=int(get_redis_port()),
            db=int(get_redis_dbnum()),
            password=get_redis_password()
        )
        conn.client_list()  # Must perform an operation to check connection.
    except ConnectionError as e:
        print("Failed to connect to Redis instance at %s", get_redis_host())
        print(repr(e))
        return False

    conn.close()  # type: ignore

    return True
