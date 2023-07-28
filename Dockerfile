FROM scratch
COPY names /
COPY names-dictionary.json /
COPY your-name.html /

ENTRYPOINT ["names"]