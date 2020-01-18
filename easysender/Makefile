NAME=easysender
PORT=5555

all:
	docker build  -t $(NAME) -f Dockerfile.dev .

run:
	docker run -ti --rm -p $(PORT):5555 --name $(NAME) $(NAME)

stop:
	docker stop `docker ps -aq -f name=$(NAME)`
