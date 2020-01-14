FROM golang:1.13-alpine

# WORKDIR /app


RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/github.com/mritunjaykumar/users-api
COPY . .


RUN go get -d -v
RUN go build -o /go/bin/usersapi .

EXPOSE 8080
# Run the hello binary.
ENTRYPOINT ["/go/bin/usersapi"]