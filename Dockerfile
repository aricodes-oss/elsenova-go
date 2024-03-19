FROM golang:latest as build

WORKDIR /elsenova

# Copy source code
COPY . .

# Build statically linked binary
RUN CGO_ENABLED=0 go build

###
FROM scratch

# SSL authority (needed to connect to discord)
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# Application binary
COPY --from=build /elsenova/elsenova /

ENTRYPOINT ["/elsenova"]
