# test-gitaction
<!-- docker run -d --name "mysql-container" -p 3000:3306 database-mysql -->
![Docker build status](https://github.com/tritam593/se15.1/actions/workflows/docker-image.yml/badge.svg)
![K8s build status](https://github.com/tritam593/se15.1/actions/workflows/k8s.yml/badge.svg)

Clone the repo
```
git clone https://github.com/tritam593/se15.1.git
```

```
cd se15.1
cd frontend
docker build -t frontend-flask .
cd ..

docker build -t backend-golang --target program .
```


Start
```
kubectl apply -f backend/backend-deployment.yaml 
kubectl apply -f backend/backend-service.yaml 
kubectl apply -f frontend/frontend-deployment.yaml 
kubectl apply -f frontend/frontend-service.yaml 
```
Run

```
http://localhost:32410/
```

Delete
```
kubectl delete -f backend/backend-deployment.yaml 
kubectl delete -f backend/backend-service.yaml 
kubectl delete -f frontend/frontend-deployment.yaml 
kubectl delete -f frontend/frontend-service.yaml 
```