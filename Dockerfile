FROM --platform=$BUILDPLATFORM golang:1.22-alpine AS build

WORKDIR /src
ARG TARGETOS
ARG TARGETARCH
ARG APP_VERSION=dev

COPY go.mod ./
COPY main.go ./

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w -X main.version=$APP_VERSION" -o /out/pipeline-demo .

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=build /out/pipeline-demo /pipeline-demo

EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/pipeline-demo"]
