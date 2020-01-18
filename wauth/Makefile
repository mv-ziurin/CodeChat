NAME=wauth
PORT=7070

all:
	docker build  -t $(NAME) -f Dockerfile.dev .

run:
	docker run -ti --rm -p $(PORT):7070 --name $(NAME) $(NAME)

stop:
	docker stop `docker ps -aq -f name=$(NAME)`
