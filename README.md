# Gilded Rose Kata

Read the problem description in the [requirements](GildedRoseRequirements.txt) file.

## GO Starter

- Run :

```shell
go run texttest_fixture.go [<number-of-days>; default: 2]
```

- Run tests :

```shell
go test ./...
```

- Run tests and coverage :

```shell
go test ./... -coverprofile=coverage.out

go tool cover -html=coverage.out
```