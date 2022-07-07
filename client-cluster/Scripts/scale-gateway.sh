#!/bin/bash
  
echo "Starting the process....."
echo

kubectl apply -f gateway-client.yaml 
kubectl apply -f non-gateway-client.yaml

echo

sleep 30

for((i=1; i<=99; i++))
do
    curr_pods=$(kubectl get pods -n gateway | grep "gateway" | wc -l)
    let new_pods=$curr_pods+1

    echo "Incrementing the pod numbers to $new_pods"
    kubectl scale deploy gateway-client -n gateway --replicas=$new_pods
    kubectl scale deploy non-gateway-client -n non-gateway --replicas=$new_pods
    echo
    sleep 30
done

echo "Completed"