from kombu import Connection
from kombu.exceptions import OperationalError
from settings import config


def get_rabbitmq_userpass() -> str:
    return config.RABBITMQ_USER + ":" + config.RABBITMQ_PASS + "@"


def get_rabbitmq_port() -> str:
    return ":" + config.RABBITMQ_PORT


def get_rabbitmq_vhost() -> str:
    return "/" + config.RABBITMQ_VHOST


def get_rabbitmq_host() -> str:
    return config.RABBITMQ_HOST


def get_broker_url() -> str:
    return "amqp://{userpass}{hostname}{port}{vhost}".format(
        hostname=get_rabbitmq_host(),
        userpass=get_rabbitmq_userpass(),
        port=get_rabbitmq_port(),
        vhost=get_rabbitmq_vhost()
    )


def is_broker_running(retries: int = 3) -> bool:
    try:
        conn = Connection(get_broker_url())
        conn.ensure_connection(max_retries=retries)
    except OperationalError as e:
        print("Failed to connect to RabbitMQ instance at %s", get_rabbitmq_host())
        print(str(e))
        return False

    conn.close()
    return True
