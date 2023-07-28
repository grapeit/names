FROM scratch
COPY names /
COPY names-dictionary.json /
COPY *.html /

ENTRYPOINT ["/names"]
