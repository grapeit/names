FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY names /
COPY names-dictionary.json /
COPY *.html /
EXPOSE 80
ENTRYPOINT ["/names"]
