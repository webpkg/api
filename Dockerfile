FROM scratch

COPY ./docker/etc/ /etc/
COPY ./bin/api-linux-amd64 /usr/bin/api
USER nobody
WORKDIR /app
EXPOSE 8443
ENTRYPOINT ["api"]

CMD ["serve"]
