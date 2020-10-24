#!/bin/sh

set -x
cat /etc/apache2/sites-available/default-ssl.conf

#
exec apachectl -DFOREGROUND "$@"
