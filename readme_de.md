[![Go](https://github.com/sultaniman/env/actions/workflows/go.yml/badge.svg)](https://github.com/sultaniman/env/actions/workflows/go.yml)

# Env

Greife einfach auf deine Umgebungsvariablen zu.

> [!NOTE]
> Ich brauchte einen einfachen und typisierten Zugriff auf Umgebungsvariablen,
> aber alle bestehenden Lösungen machten zu viel.

## 💾 Installation

```sh
go get -u github.com/sultaniman/env
```

## 📕 Beispiele

```go
import "github.com/sultaniman/env"
```

> [!NOTE]
> `env` unterscheidet zwischen Groß- und Kleinschreibung, daher sind `Var_Name`
> und `var_name` nicht identisch.

### 🤏 Direkter Zugriff

```go
port := env.GetInt("PORT")
port, err := env.GetIntE("PORT")

host := env.GetString("HOST")
```

### 🔖 Mit Umgebungspräfix

```go
env.SetEnvPrefix("PREFIX_")
port := env.GetInt("PORT")
```

<p align="center">✨ 🚀 ✨</p>
