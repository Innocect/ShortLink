FROM golang:1.17.2-alpine3.14

RUN mkdir /app
ADD . /app
WORKDIR /app/server

RUN go mod download
RUN go build -o server .

CMD [ "/app/server/server" ]