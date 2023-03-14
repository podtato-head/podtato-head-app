FROM golang:1.20.2-alpine3.16 AS builder

ENV CGO_ENABLED=0

WORKDIR /workspace

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN go build -o /workspace/podtatoserver ./cmd/

FROM gcr.io/distroless/base-debian11:nonroot AS production

LABEL org.opencontainers.image.source="https://github.com/podtato-head/podtato-head-app" \
    org.opencontainers.image.url="https://podtato-head.github.io" \
    org.opencontainers.image.title="PodTatoHead" \
    org.opencontainers.image.vendor="The PodTatoHead Maintainers" \
    org.opencontainers.image.licenses="Apache-2.0"

WORKDIR /
COPY --from=builder /workspace/podtatoserver .
USER 65532:65532

ENTRYPOINT ["/podtatoserver"]