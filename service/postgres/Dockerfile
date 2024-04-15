FROM postgres:12

ENV POSTGRES_USER postgres
ENV POSTGRES_PASSWORD postgres
ENV POSTGRES_DB fms

COPY fms.sql /docker-entrypoint-initdb.d/

EXPOSE 5432