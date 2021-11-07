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

```
docker run -d --name my-nginx -p 8080:80 -v $(pwd)/demo:/usr/share/nginx/html -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf nginx
```

#### install openssl

```
docker exec -it my-nginx bash
```

```
apt update
```

```
apt install openssl
```

```
cd /etc/nginx
```

```
mkdir ssl
```

```
cd ssl
```

```
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout /etc/nginx/ssl/self.key -out /etc/nginx/ssl/self.crt
```

#### Generating a RSA private key
..+++++
.....................................................................+++++
writing new private key to '/etc/nginx/ssl/self.key'
-----
You are about to be asked to enter information that will be incorporated
into your certificate request.
What you are about to enter is what is called a Distinguished Name or a DN.
There are quite a few fields but you can leave some blank
For some fields there will be a default value,
If you enter '.', the field will be left blank.
-----
Country Name (2 letter code) [AU]:TH
State or Province Name (full name) [Some-State]:Bangkok
Locality Name (eg, city) []:Bangkok
Organization Name (eg, company) [Internet Widgits Pty Ltd]:sing3demon.com
Organizational Unit Name (eg, section) []:dev
Common Name (e.g. server FQDN or YOUR name) []:sing3demons
Email Address []:sing3demons@live.com


```
docker stop my-nginx
```
```
docker commit my-nginx nginx-ssl
```
```
docker run -d --name my-nginx-ssl -p 8080:80 -p 443:443 -v $(pwd)/demo:/usr/share/nginx/html -v $(pwd)/conf/nginx.conf:/etc/nginx/nginx.conf nginx-ssl
```