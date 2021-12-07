
import os
import configparser


cfg = configparser.ConfigParser()
cfg.read('./env.ini')


#=========================================================================
#                          BROKER INFORMATION 
#=========================================================================
RABBITMQ = cfg['rabbitmq']
RABBITMQ_USER = RABBITMQ['user']
RABBITMQ_PASS = RABBITMQ['pass']
RABBITMQ_PORT = RABBITMQ['port']
RABBITMQ_VHOST = RABBITMQ['vhost']
RABBITMQ_HOST = RABBITMQ['host']

#=========================================================================
#                          BROKER INFORMATION 
#=========================================================================
REDIS = cfg['redis']
REDIS_PASS = REDIS['pass']
REDIS_PORT = REDIS['port']
REDIS_DB = REDIS['db']
REDIS_HOST = REDIS['host']
