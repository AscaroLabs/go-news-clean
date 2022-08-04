FROM golang:latest
WORKDIR /go-news
COPY ./app.out .
EXPOSE 8080
CMD ["./app.out"]