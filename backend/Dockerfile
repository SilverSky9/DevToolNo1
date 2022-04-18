#Build stage 
FROM golang:1.17-alpine AS Build
# RUN apk add build-base git breezy mercurial gcc
RUN ["go", "get", "github.com/rs/cors"]
COPY . /src
RUN cd /src && go build -o goapp

#Final stage
FROM alpine:3.15.4
WORKDIR /app
COPY --from=Build /src/goapp /app/
ENTRYPOINT [ "./goapp" ]