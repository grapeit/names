Build
```bash
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a
docker build --platform linux/amd64 -t names .
```
