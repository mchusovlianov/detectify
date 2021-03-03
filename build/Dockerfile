FROM golang:1.16-alpine AS build

WORKDIR /src/
ADD . /src/
RUN apk update && apk add ca-certificates

RUN go mod download
RUN CGO_ENABLED=0 go build -o /src/dist/detector /src/cmd/main.go


FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /src/dist/detector /bin/detector
ENTRYPOINT ["/bin/detector"]