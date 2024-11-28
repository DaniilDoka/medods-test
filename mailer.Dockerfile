
FROM golang:alpine as build

WORKDIR /app

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o /main ./cmd/mailer

FROM alpine

WORKDIR /app

COPY --from=build /main /main

EXPOSE 6969
CMD ["/main"]
