# go-bar-microservice

Go bar microservice

## Example docker run

```bash
 docker run -d -e MESSAGE="Hi v1 bar" -e PORT=3002 -e ROUTE_PATH="v1/bar" --name bar -p 3002:3002 m2stacklab/go-bar-microservice:latest
```