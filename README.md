# Clean Architecture in Golang
This is a sample repository for a Golang project using the Clean Architecture pattern.

# Project Structure
The project structure follows the Clean Architecture pattern, where the domain layer contains the business logic and the application layer contains the use cases. The infrastructure layer is responsible for the implementation details of the application, such as database access and external APIs.

```
├── [ 351]  Dockerfile
├── [ 788]  README.md
├── [4.0K]  assets
├── [4.0K]  cmd
│   ├── [2.7K]  config.go
│   ├── [ 637]  main.go
│   ├── [  95]  nsq.go
│   └── [3.7K]  server.go
├── [4.0K]  delivery
│   └── [4.0K]  rest
│       ├── [4.0K]  entity
│       │   ├── [ 844]  get_count_view_Entityfolio.go
│       │   └── [ 334]  handler.go
│       └── [1.4K]  route.go
├── [ 400]  docker-compose.yaml
├── [4.0K]  domain
│   ├── [4.0K]  entity
│   │   ├── [  67]  model.go
│   │   └── [ 161]  repository.go
│   ├── [4.0K]  orm
│   │   ├── [ 944]  model.go
│   │   └── [ 891]  repository.go
│   ├── [4.0K]  redis
│   │   └── [ 195]  repository.go
│   ├── [4.0K]  referral
│   │   ├── [2.9K]  model.go
│   │   └── [ 911]  repository.go
│   ├── [4.0K]  samuel
│   │   └── [1.9K]  model.go
│   └── [4.0K]  websocket
│       └── [1.4K]  modelIO.go
├── [2.6K]  go.mod
├── [194K]  go.sum
├── [4.0K]  key
│   └── [3.2K]  id_rsa
├── [4.0K]  libs
│   ├── [4.0K]  aes
│   │   └── [ 870]  aes.go
│   ├── [4.0K]  arr
│   │   └── [ 115]  arr.go
│   ├── [4.0K]  authentication
│   │   ├── [1.8K]  get_authentication.go
│   │   └── [ 633]  model.go
│   ├── [4.0K]  custom_time
│   │   └── [ 176]  custom_time.go
│   ├── [4.0K]  date
│   │   └── [ 289]  date.go
│   ├── [4.0K]  errors
│   │   └── [1.3K]  error_mapping.go
│   ├── [4.0K]  guuid
│   │   └── [ 139]  guuid.go
│   ├── [4.0K]  httpresponse
│   │   └── [2.3K]  httpresponse.go
│   ├── [4.0K]  jwe
│   │   ├── [1.2K]  jwe.go
│   │   └── [1.5K]  rsa_key.go
│   ├── [4.0K]  logging
│   │   └── [4.0K]  logger.go
│   ├── [4.0K]  middleware
│   │   └── [4.0K]  app
│   │       ├── [ 911]  basic_auth_middleware.go
│   │       ├── [1.6K]  cors_middleware.go
│   │       ├── [1.9K]  jwt_middleware.go
│   │       ├── [3.6K]  log_middleware.go
│   │       ├── [ 758]  main_middleware.go
│   │       └── [ 990]  middleware.go
│   ├── [4.0K]  minio
│   │   └── [ 564]  minio.go
│   ├── [4.0K]  notation
│   │   ├── [ 521]  get.go
│   │   └── [1.9K]  mapping_name.go
│   ├── [4.0K]  notification
│   │   └── [1.7K]  referral_notif.go
│   ├── [4.0K]  notifications
│   │   └── [1.8K]  event_reward.go
│   ├── [4.0K]  number
│   │   ├── [ 234]  model.go
│   │   └── [ 970]  number.go
│   ├── [4.0K]  pagination
│   │   └── [ 472]  pagination.go
│   ├── [4.0K]  parser
│   │   ├── [3.4K]  order_status_ord.go
│   │   ├── [4.6K]  order_status_trd.go
│   │   └── [5.1K]  stock.go
│   ├── [4.0K]  random
│   │   └── [ 180]  random_number.go
│   ├── [4.0K]  samuel
│   │   ├── [1.8K]  helper.go
│   │   ├── [1.1K]  model.go
│   │   ├── [1.6K]  samuel_realize_return.go
│   │   └── [4.3K]  samuel_stock_position.go
│   ├── [4.0K]  stock
│   │   ├── [ 576]  get_stock_by_name.go
│   │   └── [ 357]  model.go
│   ├── [4.0K]  str
│   │   └── [2.0K]  str.go
│   ├── [4.0K]  table
│   │   └── [ 177]  table_name.go
│   ├── [4.0K]  technologies
│   │   ├── [ 615]  redis.go
│   │   ├── [ 129]  resty_client.go
│   │   └── [3.3K]  sql.go
│   ├── [4.0K]  try
│   │   ├── [1.1K]  new.go
│   │   └── [ 869]  try.go
│   ├── [4.0K]  uvcr
│   │   └── [1.7K]  event_reward.go
│   └── [4.0K]  validator
│       └── [ 771]  validator.go
├── [4.0K]  repository
│   ├── [4.0K]  psql
│   │   ├── [4.0K]  entity
│   │   │   ├── [ 168]  get_count_view_Entityfolio.go
│   │   │   ├── [ 640]  model.go
│   │   │   ├── [ 298]  psql.go
│   │   │   └── [ 156]  update.go
│   │   ├── [4.0K]  orm
│   │   │   └── [6.3K]  orm.go
│   │   └── [ 560]  psql.go
│   └── [4.0K]  redis
│       ├── [ 160]  get.go
│       ├── [ 181]  redis.go
│       └── [ 780]  set.go
└── [4.0K]  service
    ├── [4.0K]  s_entity
    │   ├── [ 217]  get_count_view.go
    │   ├── [1.2K]  service.go
    │   └── [  18]  upsert_entity.go
    └── [ 694]  service.go
```

## Getting Started
Prerequisites
Golang 1.15 or later
PostgreSQL
# InstallationW
## Install Go
    cd /tmp
    wget https://dl.google.com/go/go1.18.2.linux-amd64.tar.gz

    sudo tar -xvf go1.18.2.linux-amd64.tar.gz
    sudo mv go /usr/local

    sudo nano ~/.bashrc
    // add to end of file
    export GOROOT=/usr/local/go
    export GOPATH=$HOME/$GO_WORKSPACE/go
    export PATH=$GOPATH/bin:$GOROOT/bin:$PATH

    source ~/.bashrc
    source ~/.profile

    go version
    // should return "go version go1.18.2 linux/amd64"

## Clone App from Git
    cd $GOPATH
    mkdir src && cd src

    git clone git@gitlab.com:saham-rakyat/sectask.git
    cd sectask
    mkdir logs

## Run HTTP App
    go mod init

    // build app
    go build cmd/*.go run-server

    // setup database

    // run app
    ./main


License
This project is licensed under the MIT License - see the LICENSE file for details.


