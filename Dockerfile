FROM mariadb:10

# Set environment variables
ENV MYSQL_ROOT_PASSWORD=password
ENV MYSQL_DATABASE=gocrud
ENV MYSQL_USER=admin
ENV MYSQL_PASSWORD=password

# Copy custom configuration (if needed)
# COPY ./my.cnf /etc/mysql/my.cnf

# Expose the default MariaDB port
EXPOSE 3306

# Start the MariaDB service
CMD ["mysqld"]
