

TEMP="$(dirname "$0")"
BASE_DIR=$(echo "$(basename ${TEMP})")
CURRENT_DIR=${PWD}/${BASE_DIR}

kubectl apply -f "${CURRENT_DIR}/"bds-configmap.yaml
kubectl apply -f "${CURRENT_DIR}/"mysql-persistent-storage.yaml
kubectl apply -f "${CURRENT_DIR}/"mysql-persistent-storage-claim.yaml
kubectl apply -f "${CURRENT_DIR}/"mysql.yaml
kubectl apply -f "${CURRENT_DIR}/"bds-go-auth-service.yaml
kubectl apply -f "${CURRENT_DIR}/"bds-service.yaml
