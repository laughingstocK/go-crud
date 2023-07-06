# docker run -d --name mariadb \
#   -e MYSQL_ROOT_PASSWORD=password \
#   -e MYSQL_DATABASE=gocrud \
#   -e MYSQL_USER=admin \
#   -e MYSQL_PASSWORD=password \
#   -p 3306:3306 \
#   mariadb:latest
echo "# go-crud" >> README.md
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/laughingstocK/go-crud.git
git push -u origin main