docker image build -f Dockerfile -t ascii_image .
docker container run -p 3000:8080 --detach --name ascii_container ascii_image