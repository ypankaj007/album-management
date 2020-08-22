# AlbumManagement Application
Create, Get, List and Delete albums.
Upload, Get, List and Delete images in album.

## Folder structure

├── README.md
├── api                         // contains api services
│   └── apigateway
│       ├── Dockerfile
│       ├── Makefile            // compile proto etc
│       ├── README.md
│       ├── apigateway          // build of apigateway 
│       ├── docs                // swagger docs
│       │   ├── docs.go
│       │   ├── swagger.json
│       │   └── swagger.yaml
│       ├── go.mod
│       ├── go.sum
│       ├── handler
│       │   ├── api.go          // api controller
│       │   └── swagger.model.go
│       ├── main.go
│       ├── plugin.go
│       ├── proto
│       │   └── api
│       │       ├── api.micro.go
│       │       ├── api.pb.go
│       │       └── api.proto
│       └── util
│           └── validator.go
├── build.sh                    // make build of whole application
├── db
│   └── album.cnf               // docker mysql image cnf (changed mysql port to 1306)
├── docker-compose.yaml         // containerize whole application
└── services                    // contains microservices
    ├── albumservice            // album microservice
    │   ├── Dockerfile          // to make docker image
    │   ├── Makefile
    │   ├── README.md
    │   ├── albumservice        
    │   ├── config              // config used in service
    │   │   └── config.go
    │   ├── db                  // create tables on starting of service
    │   │   └── bootstrap.go
    │   ├── go.mod
    │   ├── go.sum
    │   ├── handler
    │   │   └── album.go        // endpoint of album service. Handles all requests 
    │   ├── main.go
    │   ├── plugin.go           // add plugins if required like(Nats | Kubernates)
    │   ├── proto
    │   │   └── album
    │   │       ├── album.micro.go
    │   │       ├── album.pb.go
    │   │       └── album.proto
    │   ├── repository          // to persists in DB (mysql used)
    │   │   ├── album.go        // persists album
    │   │   └── image.go        // persists image
    │   ├── service
    │   │   ├── broker.go       // broker service emit events on create/delete items
    │   │   └── disc.go         // create, delete album (folder) and image
    │   └── util
    │       └── helper.go
    └── consumer
        ├── Dockerfile          containerized consumer service
        ├── consumer
        ├── go.mod
        ├── go.sum
        └── main.go

## AlbumService
Album service responsible to all crud operations of albums and images
### generate proto 

```bash
$ cd services/albumservice
$ make // it will compile proto file
```
## Consumer
Consumer service receives event when album created, deleted and image created and deleted
### generate proto 

```bash
$ cd services/consumer
$ make // it will compile proto file
```

## API gateway
API gateway handles all coming request from different client (WEB |MOBILE CLI).  
### generate proto 

```bash
$ cd api/apigateway
$ make 
```

### Build and run application

```bash
$ ./build.sh            // make build all application (apis and services)
$ docker-compose up     // run application in docker
```

Application will run at http://localhost:8080
Swagger - http://localhost:8080/swagger/index.html
View uploaded images - http://localhost:8080/static/{uploaded-image-url}