# Test Coverage

Look at the coverage of all test cases:
```
go test -cover
```

Get the cover profile and write it out to cover.out:
```
go test -coverprofile cover.out
```

Get the result in a web page:
```
go tool cover -html cover.out
```
