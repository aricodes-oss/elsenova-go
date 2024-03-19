FROM golang:latest as build

WORKDIR /elsenova

# Copy source code
COPY . .

# Build statically linked binary
RUN go build

###
FROM golang:latest

WORKDIR /

COPY --from=build /elsenova/elsenova .
ENTRYPOINT ["/elsenova"]
