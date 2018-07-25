#!/bin/bash
# build-and-upload.sh
readonly app=jps_go_app
readonly app_image=jps_go_app:latest
readonly app_docker=Dockerfile

readonly nginx_image=nginx:latest
readonly nginx_docker=nginx.Dockerfile

readonly bzip_file=${app}-latest.tar.bz2

readonly app_key="~/.ssh/toodo-ssh"
readonly app_domain_user="core@toodo.deitagy.com"

clear

# Remove image and ignore 'image does not exist' error
echo "Building app and images ..."
docker rmi -f ${app_image} 2>/dev/null 
env GOOS=linux GOARCH=386 buffalo build -o ${app}
docker build --no-cache -t ${app_image} -f ${app_docker} .

rm ${app}

docker rmi -f ${nginx_image} 2>/dev/null
docker build --no-cache -t ${nginx_image} -f ${nginx_docker} .

echo "Creating deployment .env file..."
./scripts/create-env.sh

echo "saving docker images ..."
docker save ${app_image} ${nginx_image} | bzip2 > ${bzip_file}
ls -lah ${bzip_file}

echo "uploading images to the server ..."

#scp -i ${app_key} ${app}-latest.tar.bz2  docker-compose.yml dep.env ${app_domain_user}:/home/core

#ssh -i ${app_key} ${app_domain_user} << ENDSSH
#cd /home/core
#bunzip2 --stdout ${bzip_file} | docker load
#rm ${bzip_file}
#mv dep.env .env -f
#docker-compose down && docker-compose up -d --force-recreate;
#ENDSSH