# syntax=docker/dockerfile:1

FROM golang:latest

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . .

# Build
RUN go build -o testapp ./cmd/main.go

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
EXPOSE 8080
ENV POSTGRES_USER=postgres
ENV POSTGRES_PASSWORD=Wild54323
ENV POSTGRES_DB=postgres
ENV PG_URL=postgres://postgres:Wild54323@postgres:5432/postgres
ENV HTTP_PORT=8080
# Run
CMD ["./testapp"]