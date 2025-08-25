FROM golang:1.25
WORKDIR /go/src/github.com/zerocube/route53-record-updater
COPY go.mod go.sum ./
RUN go mod download -x && go mod verify
COPY . .
RUN go build -v -x -o route53-record-updater
ENTRYPOINT "./route53-record-updater"
