FROM golang:alpine AS build

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

WORKDIR /

# Get dependencies first
COPY ./go.mod .
COPY ./go.sum .
RUN	CGO_ENABLED=0 go mod download

# Build the app
COPY . .
RUN CGO_ENABLED=0 go build -gcflags="all=-N -l" -o /app

FROM scratch
COPY --from=build /app /app
CMD ["/app"]
