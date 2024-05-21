# Unit tests 

fichier *_test.go

func Test***(t *testing.T) {
}

> -v for verbose
> --count=1 to not use cache

## Test table / subtests

```go

Used for grouping tests for the same function

## Parallel tests

```go
t.Parallel()
```

# Testing Package 

## Test helpers

```go
func Test***(t *testing.T) {
    t.Helper()
}
```

- Make function test only usage
- Skip function in stack trace

## Test Log

```go
    t.Log("message")
```

- Print message only if test fails
- Usefull for debugging

## Test Fail

```go
    t.Fail()
```

- Fail the test
- Continue to run the test

## Test FailNow

```go
    t.FailNow()
```

- Fail the test
- Stop the test
- No other test will be run

## Test Skip

```go
    t.Skip()
```

- Skip the test
- Continue to run the test

## Test Error

```go
    t.Error("message")
```

- Fail the test
- Continue to run the test