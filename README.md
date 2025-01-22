# **Env** [![Go Reference](https://pkg.go.dev/badge/github.com/yisroelshulman/env.svg)](https://pkg.go.dev/github.com/yisroelshulman/env)

A trimmed down Go library of the [godotenv](https://github.com/joho/godotenv) project.

The library provides the means to load the .env file from the directory to ENV for the process. Or
the ability to Read the .env file into a map of key value pairs when the user doesn't want to store
them in environment variables.

Changes to parsing are described in the Usage section.

## Installation

```sh
go get github.com/yisroelshulman/env
```

## Usage

After installation the package should be imported to the file and call Load or Read depending on
what is needed.

```go
package yourpackage

import (
    "fmt"
    "os"
    "github.com/yisroelshulman/env"
)

func main() {
    env.Load() // loads form the .env file into ENV

    myVar := os.Getenv("MY_VAR")
    if myVar == "" {
        // failed to read variable
    }
}
```

Alternatively, the .env file can be read into a map if the user doesn't want to store it in ENV.

```go
func main() {
    vars, err := env.Read()
    if err != nil {
        // failed to read
    }

    for key, value := range vars {
        fmt.Printf("key: %v, value: %v\n", key, value)
    }
}
```

 *IMPORTNAT*

This implementation only accepts variable names that are similar to the c standard except that they
cannot start with an underscore.

Therefore keys have the form:
```
variable : letter [letter | digit | "_" ]*
letter : [a-z | A-Z]
digit : [0-9]
```

## Contributing
Currently not adding any features.

Bug reports will be reviewed to see if any changes are necessary.