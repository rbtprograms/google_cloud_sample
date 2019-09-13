FROM golang:1.13.0-alpine
LABEL robert thompson

ENV SOURCES /go/src/google_cloud_sample

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT [ "google_cloud_sample" ]