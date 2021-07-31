# go-tasker

## Development

1. Install Local Kubernetes (Docker, Minikube, k3d)
2. Install skaffold
3. Run following command
   `skaffold dev`

## Building

`go build -o ./bin/server server/main.go`

## Running

`./bin/server`

## Operation

```
/task/create - Create Scheduled Task
/task/:id - Get Scheduled Task
/task - Get All Tasks
/task/:id/
/task/delete - Delete Scheduled Task
```

## Maintainer

- Alex Friedrichsen
