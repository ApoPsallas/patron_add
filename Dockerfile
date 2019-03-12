FROM golang:latest as builder
RUN cd ..
RUN mkdir micro
WORKDIR micro
COPY . ./
ARG version=dev
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -ldflags "-X main.version=$version" -o micro ./cmd/micro/main.go 

FROM scratch
COPY --from=builder /go/micro/micro .
CMD ["./micro"]
