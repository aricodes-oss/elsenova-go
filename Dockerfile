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
FROM golang:latest as backend-build

WORKDIR /elsenova

# Retrieve and install dependencies
COPY go.mod go.sum .
RUN go mod download

# Copy source code
COPY . .

# Build statically linked binary
RUN go build

### Deployment
FROM node:latest

RUN apt-get update && apt-get install -y supervisor

ENV GIN_MODE=release
WORKDIR /code

COPY supervisord.conf .
COPY --from=frontend-build /code .
COPY --from=backend-build /elsenova/elsenova .

CMD ["supervisord", "-c",  "./supervisord.conf"]x
