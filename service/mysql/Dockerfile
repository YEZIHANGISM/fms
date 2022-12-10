FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=mysql
ENV MYSQL_DATABASE=fms

COPY fms.sql /docker-entrypoint-initdb.d/

EXPOSE 3306