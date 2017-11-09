[![Build Status](https://travis-ci.org/yowcow/goirccommand.svg?branch=master)](https://travis-ci.org/yowcow/goirccommand)

IRC Command
===========

Yet another IRC command writer.

HOW TO USE
----------

### Import

Do:

```
go get github.com/yowcow/goirccommand
```

and import like:

```go
import (
    command "github.com/yowcow/goirccommand"
)
```

### Write

Pass a writer along with necessary parameters.

```go
err := command.Nick(conn, "foobar")
```

will write a `NICK` command like: `NICK foobar\r\n`

SEE ALSO
--------

[RFC 2812](https://tools.ietf.org/html/rfc2812)
