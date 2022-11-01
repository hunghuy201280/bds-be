chmod +x **/*.sh
CD ./k8s-yaml
./k8s-yaml/stop.sh

kill -9 $(lsof -t -i :3000,3306,2011)
