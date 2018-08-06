
# run: redis
docker run --name redis.local --hostname redis.local -p 36379:6379 -v C:/docker/share/redis/data:/data -d redis 

# run: mysql:5
docker run --name mysql.local --hostname mysql.local -p 33306:3306 -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -e MYSQL_ROOT_PASSWORD= -v C:/docker/share/mysql/data:/var/lib/mysql -d mysql:5 --character-set-server=utf8 --collation-server=utf8_general_ci 

