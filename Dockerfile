FROM alpine:3.5
LABEL robert thompson

ENV SOURCES /app/google_cloud_sample

COPY ./google_cloud_sample ${SOURCES}

RUN chmod +x /app/google_cloud_sample

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/google_cloud_sample