


TEMP="$(dirname "$0")"
BASE_DIR=$(echo "$(basename ${TEMP})")
CURRENT_DIR=${PWD}/${BASE_DIR}


kubectl delete -f "${CURRENT_DIR}"/bds-go-auth-service.yaml
kubectl delete -f "${CURRENT_DIR}"/bds-service.yaml
kubectl delete -f "${CURRENT_DIR}"/mysql.yaml
kubectl delete -f "${CURRENT_DIR}"/bds-configmap.yaml
kubectl delete -f "${CURRENT_DIR}"/mysql-persistent-storage-claim.yaml
kubectl delete -f "${CURRENT_DIR}"/mysql-persistent-storage.yaml
kubectl get pod
