# go-books

## build & start

```
$ docker-compose build
$ docker-compose up
```

## request example

```
$ curl http://localhost:8080/book | jq
```

```
$ curl -X POST http://localhost:8080/book \
       -H 'content-type: application/json' \
       -d '{"title":"python","category":1,"author":"kosuke"}'
```

```
$ curl -X PUT http://localhost:8080/book/1 \
       -H 'content-type: application/json' \
       -d '{"title":"python","category":1,"author":"shimizu"}'
```
