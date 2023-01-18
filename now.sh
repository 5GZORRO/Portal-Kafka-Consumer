bash k8s.sh

export KUBECONFIG=k8s/bcn_kubeconfig.yaml      

kubectl delete -f k8s/portal-kafka-consumer.yaml -n portal-kafka-consumer

kubectl apply -f k8s/portal-kafka-consumer.yaml -n portal-kafka-consumer
sleep 4

kubectl get pods -n portal-kafka-consumer

# watch kubectl logs -n portal-kafka-consumer portal-kafka-consumer-deployment-6db59fc7cb-9zlhx