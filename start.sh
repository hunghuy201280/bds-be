chmod +x **/*.sh
#minikube start --mount=true --mount-string="$HOME":/home/docker/ --driver=docker
eval $(minikube -p minikube docker-env)
printf "\n\n"
echo "Starting build docker images"
docker compose build
printf "\n\n"
echo "Deploy to minikube cluster"
./k8s-yaml/start.sh

echo "Wait for pod to ready"
sleep 5 # Waits 5 seconds.

minikube tunnel & kubectl get pod -w
