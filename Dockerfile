FROM golang:1.9
WORKDIR /go/src/app
COPY . .
RUN go-wrapper download \
  && go-wrapper install \
  && go build -o gorecord gorecord.go
ENTRYPOINT "./gorecord"