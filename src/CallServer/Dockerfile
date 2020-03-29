FROM golang:1.14.1-alpine3.11

WORKDIR /usr/src/scoing

COPY  . .

RUN go build

CMD [ "go", "run CallServer" ]git 