[![Go](https://github.com/sultaniman/env/actions/workflows/go.yml/badge.svg)](https://github.com/sultaniman/env/actions/workflows/go.yml)

# Env

Simply access your environment variables.

## Installation

```sh
go get -u github.com/sultaniman/env
```

## Examples

```go
import "github.com/sultaniman/env"
```

> [!NOTE]
> `env` is case sensitive so `Var_Name` and `var_name` are not the same.

### Direct access

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")

host := env.GetString("HOST")
```

### With environment prefix

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```
