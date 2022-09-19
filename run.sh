# Closes running services
docker-compose down 

# Runs everything
docker-compose up -d --build 

# Pretty Console
#printf "\033c"
echo "running on 8080"

#Docker logging
docker logs --follow portal-kafka-consumer