# elsenova

[![publish](https://github.com/aricodes-oss/elsenova-go/actions/workflows/publish.yml/badge.svg)](https://github.com/aricodes-oss/elsenova-go/actions/workflows/publish.yml)
[![lint](https://github.com/aricodes-oss/elsenova-go/actions/workflows/lint.yml/badge.svg)](https://github.com/aricodes-oss/elsenova-go/actions/workflows/lint.yml)
<img alt="tags-badge" id="tags" src="https://ghcr-badge.egpl.dev/aricodes-oss/elsenova-go/tags?color=%2344cc11&amp;ignore=&amp;n=3&amp;label=image+tags&amp;trim=">
<img alt="size-badge" id="size" src="https://ghcr-badge.egpl.dev/aricodes-oss/elsenova-go/size?color=%2344cc11&amp;tag=latest&amp;label=image+size&amp;trim=">

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

At a minimum, you are **required** to give a `token` and `guildID` for the bot to connect.

## Contributing

Install [gowatch](https://github.com/silenceper/gowatch) -

```sh
go install github.com/silenceper/gowatch@latest
```

and run `gowatch` in the root of the project to start the bot.

Code must be formatted with `go fmt` before pull requests will be reviewed.
