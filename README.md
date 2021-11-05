#### 1. docker nginx
```
docker run --rm -d --name my-nginx -p 8080:80 nginx
```

```
$ docker pull nginx
$ docker run --rm -d --name my-nginx -p 8080:80 nginx
$ docker exec -it my-nginx bash
/etc/nginx #config file
/usr/share/nginx/html #html file
```
#### bind mount volume on docker 

```docker run --rm -d --name my-nginx -p 8080:80 -v /Users/sing3demons/_Dev_Code/Nginx/nginx-bootcamp/demo:/usr/share/nginx/html nginx
```

```
docker run --rm -d --name my-nginx -p 8080:80 -v $(pwd)/demo:/usr/share/nginx/html nginx
```
docker run --rm -d --name my-nginx -p 8080:80 -v $(pwd)/demo:/usr/share/nginx/html -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf nginx
```