FROM golang:1.17-alpine
ENV CGO_ENABLED 0
# RUN go test
COPY . /go/bin/app
WORKDIR /go/bin/app
RUN ["go", "get", "github.com/rs/cors"]
RUN ["go", "get", "-v", "github.com/cosmtrek/air"]
CMD ["go", "mod", "tidy"]
ENTRYPOINT ["air"]
