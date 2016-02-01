addr
====

Small utility for parsing and stringifying host:port addresses.


Install
-------

```
go get github.com/bookerzzz/addr
```


Usage
-----

```go
a, err := addr.Make("localhost:80")
if err != nil {
  // Something went wrong with parsing the port.
}
fmt.Println(a.Host, a.Port, a.String())
// stdout: localhost 80 localhost:80

_ = a.Parse("comediansincarsgettingcoffee.com") // Port is inherited from previous parse
fmt.Println(a.Host, a.Port, a.String())
// stdout: comediansincarsgettingcoffee.com 80 comediansincarsgettingcoffee.com:80

_ = a.Parse("comediansincarsgettingcoffee.com:0") // Clear the port value
fmt.Println(a.Host, a.Port, a.String())
// stdout: comediansincarsgettingcoffee.com 0 comediansincarsgettingcoffee.com

_ = a.Parse(":8080") // only use port
fmt.Println(a.Host, a.Port, a.String())
// stdout:  8080 :8080
```


Test
----

```bash
make test
# stdout: ok  	github.com/bookerzzz/addr	0.007s	coverage: 100.0% of statements
make coverage
# opens go tool cover
```


License
-------

The MIT License (MIT). Copyright (c) 2016 Bookerzzz
