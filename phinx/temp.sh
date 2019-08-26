#!/bin/bash
docker exec -it php-nginx /bin/bash -c "cd /data/www/fm&& php vendor/robmorgan/phinx/bin/phinx  create YuboTest2"
