# Streaming HTML

An example showing [streaming responses](https://developer.mozilla.org/en-US/docs/Web/API/Response/body) to the
[Shadow DOM](https://developer.mozilla.org/en-US/docs/Web/API/ShadowRoot/mode) from a Go server.
Enables asynchronous DOM updates with no Javascript needed!

## Test

```shell
go test ./...
```

## Run

```shell
go run cmd/app/app.go
```

Then visit [localhost:8777](http://localhost:8777).
