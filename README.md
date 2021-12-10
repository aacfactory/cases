# Cases
String case formatter for Golang

## Install
```shell
go get github.com/aacfactory/cases
```

## Camel
```go
c := Camel()
atoms, err := c.Parse("FooBar")
result := c.Format([]string{"foo", "bar"})
```
## Snake
```go
c := Snake()
atoms, err := c.Parse("foo_bar")
result := c.Format([]string{"foo", "bar"})
```
## Kebab
```go
c := Kebab()
atoms, err := c.Parse("foo-bar")
result := c.Format([]string{"foo", "bar"})
```
## Qualified
```go
c := Qualified()
atoms, err := c.Parse("foo.bar")
result := c.Format([]string{"foo", "bar"})
```
