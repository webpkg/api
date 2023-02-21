FROM scratch

COPY ./docker/ /
COPY ./bin/api-linux-amd64 /usr/bin/api
USER nobody
WORKDIR /app
ENTRYPOINT ["api"]

CMD ["serve"]
