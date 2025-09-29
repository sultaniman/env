[![Go](https://github.com/sultaniman/env/actions/workflows/go.yml/badge.svg)](https://github.com/sultaniman/env/actions/workflows/go.yml)

# Env

Simply access your environment variables.

> [!NOTE]
> I needed a simple and typed access to environment variables
> but all existing solutions did too much.

## 💾 Installation

```sh
go get -u github.com/sultaniman/env
```

## 📕 Examples

```go
import "github.com/sultaniman/env"
```

> [!NOTE]
> `env` is case sensitive so `Var_Name` and `var_name` are not the same.

### 🤏 Direct access

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")

host := env.GetString("HOST")
```

### 🔖 With environment prefix

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```

<p align="center">✨ 🚀 ✨</p>

## 🌐 Translations

- [Deutsch README](readme_de.md)
- [Türkçe README](readme_tr.md)
- [Qyrgyz README](readme_ky.md)
