# Meonzi

## prerequisite
- docker, go, mysql, aws(or locally localstack)

## Getting started
```sh
  docker-compose -f local.compose.yml  # using when in your local computer
  docker-compose up                    # using production remote server
```

## Test
```sh
  go test ./... --count=1              # run all test without cache
```
