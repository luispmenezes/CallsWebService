FROM golang:1.14.1-alpine3.11

WORKDIR /usr/bin/callws

COPY  . .

RUN go build

CMD [ "./CallClient" ]