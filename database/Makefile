NAME=database
PORT=5431

all:
	docker build  -t $(NAME) -f Dockerfile .

run:
	docker run -ti --rm -p $(PORT):5431 --name $(NAME) $(NAME)

stop:
	docker stop `docker ps -aq -f name=$(NAME)`
