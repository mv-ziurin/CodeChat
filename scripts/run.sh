SERVICE_NAME=$1

echo \> docker run --rm --name $SERVICE_NAME webber1580/$SERVICE_NAME:$CI_PIPELINE_ID
docker run --rm --name $SERVICE_NAME webber1580/$SERVICE_NAME:$CI_PIPELINE_ID