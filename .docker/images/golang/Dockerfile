FROM golang:1.8-alpine
COPY . /go/src/github.com/devigner/authorisation
RUN go build src/github.com/devigner/authorisation/main.go
CMD ./main