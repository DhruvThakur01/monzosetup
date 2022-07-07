#!/bin/bash
  
echo "Starting the process....."
echo

kubectl apply -f gateway-client.yaml
kubectl apply -f non-gateway-client.yaml

echo

sleep 300

for((i=1; i<=5; i++))
do
    echo "Creating a 10x spike in the load. Incrementing the pod numbers to 50."
    nohup kubectl scale deploy gateway-client -n gateway --replicas=50 &
    nohup kubectl scale deploy non-gateway-client -n non-gateway --replicas=50 &
    echo 
    sleep 300

    echo "Removing the spike. Scaling down to 5"
    nohup kubectl scale deploy gateway-client -n gateway --replicas=5 &
    nohup kubectl scale deploy non-gateway-client -n non-gateway --replicas=5 &
    echo 
    sleep 600
done

echo "Completed"