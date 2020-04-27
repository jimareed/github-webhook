# github-webhook
Implement a basic webhook in go that listens for new issues added to this repo

build & run container
```
docker build --tag github-webhook-image .
docker run --name github-webhook -p 8080:8080 -d github-webhook-image
```

clean up
```
docker rm -f github-webhook
docker rmi github-webhook-image
```

add to nginx conf
```
vi /etc/nginx/nginx.conf (add)

        location /postreceive {
            proxy_pass http://localhost:8080;
        }

sudo service nginx stop
sudo service nginx start
```