# test-gitaction

Clone the repo
```
git clone https://github.com/tritam593/se15.1.git
```

```
cd se15.1
cd frontend
docker build -t frontend-flask .
cd ..

docker build . --file ./backend/Dockerfile --tag backend-flask
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