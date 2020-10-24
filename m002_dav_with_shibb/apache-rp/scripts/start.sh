#!/bin/sh

# service shibd start
# exec apachectl -DFOREGROUND "$@"
exec /usr/bin/supervisord --nodaemon