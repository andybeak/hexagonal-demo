# Hexagonal architecture demo

See Medium for explanatory article.

## Running

Bring up the stack with Docker:

```
docker-compose build
docker-compose up -d
```

Then you can make http queries:

```
 $ curl -X POST -H "Content-Type: application/json" -d '{"name": "Example"}' http://localhost:8200/v1/users
{"id":"004cdd12-7d65-4f78-8e41-def2f2ef1bef","name":"Example"}%   

$ curl http://localhost:8200/v1/users/004cdd12-7d65-4f78-8e41-def2f2ef1bef
{"id":"004cdd12-7d65-4f78-8e41-def2f2ef1bef","name":"Example"}%         