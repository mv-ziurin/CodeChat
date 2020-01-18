echo \> docker stack ls
docker stack  ls  

echo \> docker node ls
docker node   ls 

echo \> docker stack deploy  -c ./docker_stack.yml  codechat  --with-registry-auth 
docker stack deploy  -c ./docker_stack.yml  codechat  --with-registry-auth 
