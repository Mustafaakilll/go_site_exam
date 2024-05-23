FROM golang:1.22-alpine

WORKDIR /app

COPY . ./

RUN go mod download

RUN go build -o /go-site-exam

EXPOSE 8081

CMD [ "/go-site-exam" ]
