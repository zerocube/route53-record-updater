FROM golang:1.9
WORKDIR /go/src/app
RUN go-wrapper download github.com/aws/aws-sdk-go \
  && go-wrapper install github.com/aws/aws-sdk-go
COPY . .
RUN go build -v -x -o gorecord gorecord.go
ENTRYPOINT "./gorecord"