FROM golang:1.17.7-alpine3.15 AS build
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN GOOS=linux CGO_ENABLED=0 go build -o /todo

FROM public.ecr.aws/lambda/go:1
COPY --from=build /todo ${LAMBDA_TASK_ROOT}
COPY --from=build /app/config ${LAMBDA_TASK_ROOT}/config
ENV ENABLE_GIN_LAMBDA_PROXY=TRUE
ENV ENV=aws

# Command can be overwritten by providing a different command in the template directly.
CMD ["todo"]