# Env

Simply access your environment variables.

## Installation

```sh
go get -u github.com/sultaniman/env
```

> [!NOTE] 
> `env` is case sensitive so `Var_Name` and `var_name` are not the same.

## Examples

```go
import "github.com/sultaniman/env"
```

### Direct access

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")
```

### With environment prefix

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```
