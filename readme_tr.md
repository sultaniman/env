[![Go](https://github.com/sultaniman/env/actions/workflows/go.yml/badge.svg)](https://github.com/sultaniman/env/actions/workflows/go.yml)

# Env

Ortam değişkenlerinize kolayca erişin.

> [!NOTE]
> Ortam değişkenlerine basit ve tipli erişime ihtiyacım vardı,
> ancak mevcut çözümlerin hepsi çok fazla şey yapıyordu.

## 💾 Kurulum

```sh
go get -u github.com/sultaniman/env
```

## 📕 Örnekler

```go
import "github.com/sultaniman/env"
```

> [!NOTE]
> `env` büyük/küçük harfe duyarlıdır, bu yüzden `Var_Name` ile `var_name`
> aynı değildir.

### 🤏 Doğrudan erişim

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")

host := env.GetString("HOST")
```

### 🔖 Ortam öneki ile

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```

<p align="center">✨ 🚀 ✨</p>
