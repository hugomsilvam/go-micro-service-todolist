# Todo.List.Micro Service

This is the Todo.List.Micro service (Go micro) sample with CRUD operations.
* You can create a new Todo.
* Get a specific Todo by Id.
* Get a list of all Todos.
* Update a Todo.
* Delete a Todo.

Generated with

```
micro new github.com/hugomsilvam/todo.list.micro --namespace=go.micro --type=srv
```

If you want change the protobuf example you will need to generate using proto-gen-go inside your project directory

```
cd /home/<user>/go/src/github.com/hugomsilvam/todo.list.micro
protoc --proto_path=. --go_out=. --micro_out=. proto/example/example.proto
```


## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: go.micro.srv.todo.list.micro
- Type: srv
- Alias: todo.list.micro

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

To run and to test this project

```
go run main.go
```

Use micro web dashboard for getting more information about yours micro services on the browser (always better than on the console :smiley:)

```
go web
```

To call the Service, the web dashboard should work but in my case isn't working, so I run the command on the console

```
micro call <service-name> <endpoint> <request-parameters>
micro call go.micro.srv.todo.list.micro TodoService.Create '{"todo":{"id":1,"title":"bananas","description":"apples + bananas","completed":true}}'
```

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./todo.list.micro-srv
```

Build a docker image
```
make docker
```