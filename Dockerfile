FROM golang:1.22-alpine as build
COPY . /app

WORKDIR /app
RUN go build -o microblog .

FROM alpine:3
COPY --from=build /app/microblog /app/microblog
COPY ./templates /app/templates
WORKDIR /app
CMD ["/app/microblog"]
EXPOSE 1323