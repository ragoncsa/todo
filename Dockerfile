FROM golang:1.17.7-alpine3.15 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o /todo

FROM alpine:3.15
COPY --from=build /todo /todo
COPY --from=build /app/config /config
EXPOSE 8080
RUN adduser -D nonroot
USER nonroot
ENTRYPOINT ["/todo"]