#!/usr/bin/env ash

set -ex

chown -R www.www /var/www/ott_admin
chown -R www.www /var/run/session
chown -R www.www /data/logs/ott_admin

exec php-fpm
