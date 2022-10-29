chmod +x **/*.sh
#minikube start --mount=true --mount-string="$HOME":/home/docker/ --driver=docker
eval $(minikube -p minikube docker-env)
./k8s-yaml/stop.sh && ./k8s-yaml/start.sh


