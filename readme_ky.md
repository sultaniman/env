[![Go](https://github.com/sultaniman/env/actions/workflows/go.yml/badge.svg)](https://github.com/sultaniman/env/actions/workflows/go.yml)

# Env

Ajlana çöjrödögü özgörmölörgö cönököj kirüü.

> [!NOTE]
> Mağa ajlana çöjrö özgörmölörünö cönököj, tipteştirilgen kirüü kerek boldu,
> biraq bar čeçimder aşıqça köp iş atqarat.

## 💾 Ornotuu

```sh
go get -u github.com/sultaniman/env
```

## 📕 Mysaldar

```go
import "github.com/sultaniman/env"
```

> [!NOTE]
> `env` çoñ jana kiçi tamğalardy acyratıp qarajt, demek `Var_Name` menen `var_name`
> birdej emes.

### 🤏 Tüzdön-tüz kirüü

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")

host := env.GetString("HOST")
```

### 🔖 Ajlana prefiksi menen

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```

<p align="center">✨ 🚀 ✨</p>
