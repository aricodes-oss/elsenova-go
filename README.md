# elsenova

[![publish](https://github.com/aricodes-oss/elsenova-go/actions/workflows/publish.yml/badge.svg)](https://github.com/aricodes-oss/elsenova-go/actions/workflows/publish.yml)

Your friendly neighborhood Axiom Verge discord server bot!

---

## Building

```
go build
```

## Running

`elsenova` has a few different subsystems. The main one, providing the chatbot interface, can be started with

```
./elsenova start
```

## Configuring

Place a file named `elsenova.yml` in the same directory as the executable. See [elsenova.example.yml](./elsenova.example.yml) for all of the configuration options supported.

At a minimum, you are **required** to give a `token` and `guild_id` for the bot to connect.

## Contributing

Install [gowatch](https://github.com/silenceper/gowatch) -

```sh
go install github.com/silenceper/gowatch@latest
```

and run `gowatch` in the root of the project to start the bot.

Code must be formatted with `go fmt` before pull requests will be reviewed.
