FROM postgres:11.6-alpine
COPY *.sql /docker-entrypoint-initdb.d/
CMD ["-c", "max_connections=200"]
