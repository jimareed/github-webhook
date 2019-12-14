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
