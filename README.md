<p align="center">
  <a href="https://mdxjs.com">
    <img alt="hexcraft logo" src="assets/logo/logo.svg" width="140" />
  </a>
</p>

# hexcraft

A CLI tool for generating hexagonal architecture for microservices in Go. Quickly scaffold new projects with a clean, maintainable structure following hexagonal architecture principles.


## Features

- Generates complete hexagonal architecture structure
- Creates domain models, ports, and adapters
- Includes gRPC server setup with protocol buffers
- Configures database connections and migrations
- Includes logging and configuration management
- Test directory structure


## Folder Structure
```
├── cmd
│   └── main.go
├── pkg
│   ├── logger
│   │   └── 
│   └── util
│       └── 
├── internal
│   ├── adapter
│   │   ├── config
│   │   │   └── config.go
│   │   ├── grpcserver
│   │   │   ├── handler.go
│   │   │   └── server.go
│   │   ├── proto
│   │   │   ├── [generated].pb.go
│   │   │   ├── [proto-file-name].proto
│   │   │   └── [generated].go
│   │   └── storage
│   │       ├── [db-name]
│   │       │   ├── db.go
│   │       │   ├── migrations
│   │       │   │   ├── [version]_[migration_name]_[down | up].[type]
│   │       │   └── repository
│   │       │       └── 
│   │       └── redis
│   └── core
│       ├── domain
│       ├── port
│       │   └── 
│       └── service
│           └── [service].go
└── test
```





## Installation

Using Homebrew:
I created the formula.rb file but it is not ready to be added to the homebrew-core repository. So for now you need to download and install it like this.


```bash
brew tap bedirhangull/hexcraft
brew install hexcraft

```


## Usage

Hexcraft uses a YAML configuration file to generate your project structure. Create a `hexcraft.yaml` file:

```yaml
packageName: "your-service-name"
dockerfile: true
gitignore: true
envFile: true

```

Also you can check it out in repository (hexcraft.yaml file).
You can see other attributes in hex craft.yaml, but for now you can create a microservice using these attributes:

| value       	| type 	|
|-------------	|------	|
| packageName 	| bool 	|
| dockerfile  	| bool 	|
| gitignore   	| bool 	|
| envFile     	| bool 	|

I plan to make other attributes workable soon.



## Author

- [@bedirhngl](https://x.com/bedirhngl)
