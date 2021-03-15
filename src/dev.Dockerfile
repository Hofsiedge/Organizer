FROM golang:alpine AS build

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /
RUN	CGO_ENABLED=0 go get -ldflags "-s -extldflags '-static'" github.com/go-delve/delve/cmd/dlv

COPY ./go.mod .
COPY ./go.sum .
RUN	CGO_ENABLED=0 go mod download

COPY . .
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /app

FROM scratch
COPY --from=build /go/bin/dlv /dlv
COPY --from=build /app /app
ENTRYPOINT [ "/dlv" ]
