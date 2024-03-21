### Frontend
FROM node:latest as frontend-build

WORKDIR /code

# Retrieve and install dependencies
COPY frontend/package.json frontend/package-lock.json .
RUN npm install

# Bring in project code
COPY frontend .

# Build!
RUN npm run build

### Backend
FROM golang:latest as build

WORKDIR /elsenova

# Retrieve and install dependencies
COPY go.mod go.sum .
RUN go mod download

# Copy source code
COPY . .

# Grab frontend slug
COPY --from=frontend-build /code/dist web/dist

# Build statically linked binary
RUN go build -tags prod

### Deployment
FROM golang:latest

ENV GIN_MODE=release
WORKDIR /

COPY --from=build /elsenova/elsenova .

ENTRYPOINT ["/elsenova"]
