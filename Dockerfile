FROM golang:1.25 AS build

WORKDIR /build
COPY . ./
RUN go mod download
ENV CGO_ENABLED=0 GOOS=linux
RUN go build -o go-api-example ./cmd/api/

FROM gcr.io/distroless/base-debian12:nonroot
COPY --from=build /build/go-api-example /go-api-example

USER nonroot:nonroot
EXPOSE 3000
ENV PORT=3000 ENV=PROD
ENTRYPOINT [ "/go-api-example" ]
