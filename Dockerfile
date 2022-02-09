FROM golang:1.17-alpine
RUN mkdir /app
COPY ./backend/ /app/
WORKDIR /app
RUN go get -v github.com/cosmtrek/air
CMD go mod tidy
ENTRYPOINT ["air"]