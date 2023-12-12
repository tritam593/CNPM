FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=1

# Sao chép các tệp SQL khởi tạo vào container (nếu có)
COPY ./init.sql /docker-entrypoint-initdb.d/

