from celery import Celery
from redis_config import (is_backend_running, get_backend_url)
from broker_config import (get_broker_url)
from redis_main import redis_conn
from ai.ai import CompleteModel
from settings import (celery_config, config, ml_config)


if not is_backend_running(): exit()


app = Celery(celery_config.QUERY_NAME, broker=get_broker_url(), backend=get_backend_url())

app.config_from_object('settings.celery_config')


@app.task(bind=True, name="{}.test".format(celery_config.QUERY_NAME))
def test():
    pass
