# Go AWS Lambda Testing

Uses dependency injection to start a lambda handler:

```
func main() {
	deps := dependencies{}
	lambda.Start(deps.handler)
}
```

...where dependencies is a struct:

```
type dependencies struct {
    ...
}
```

...and the handler is a pointer receiver of the dependencies:

```
func (d *dependencies) handler() {
    ...
}
```

## Running Tests with Coverage

```
go test -v ./... -coverprofile=c.out
```

```
go tool cover -html=c.out
```