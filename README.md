# test-gitaction
<!-- docker run -d --name "mysql-container" -p 3000:3306 database-mysql -->
![Docker build status](https://github.com/tritam593/se15.1/actions/workflows/docker-image.yml/badge.svg)
![K8s build status](https://github.com/tritam593/se15.1/actions/workflows/k8s.yml/badge.svg)

Clone the repo
```
git clone https://github.com/tritam593/se15.1.git

python3 -m unittest tests.test_user
```

```
server.Router.HandleFunc("/login", server.DoLogin).Methods("POST")
	server.Router.HandleFunc("/register", server.DoRegister).Methods("POST")
	server.Router.HandleFunc("/logout", server.Logout).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.DeleteUser).Methods("DELETE")
	server.Router.HandleFunc("/users/{id}", server.GetUser).Methods("GET")
	server.Router.HandleFunc("/users/{id}", server.UpdateUser).Methods("PUT")

	server.Router.HandleFunc("/products", server.Products).Methods("GET")
	server.Router.HandleFunc("/products/{id}", server.GetProductByID).Methods("GET")
	server.Router.HandleFunc("/products", server.CreateProduct).Methods("POST")
	server.Router.HandleFunc("/products/{id}", server.UpdateProduct).Methods("PUT")
	server.Router.HandleFunc("/products/{id}", server.DeleteProduct).Methods("DELETE")

	server.Router.HandleFunc("/categories", server.GetCategories).Methods("GET")
	server.Router.HandleFunc("/categories", server.CreateCategory).Methods("POST")
	server.Router.HandleFunc("/categories/{id}", server.GetCategoryByID).Methods("GET")
	server.Router.HandleFunc("/categories/{id}", server.DeleteCategory).Methods("DELETE")
	server.Router.HandleFunc("/categories/{id}", server.UpdateCategory).Methods("PUT")

	server.Router.HandleFunc("/product-img/{id}", server.CreateImageByProductID).Methods("POST")
	server.Router.HandleFunc("/product-img/{id}", server.DeleteImageByProductID).Methods("DELETE")

	server.Router.HandleFunc("/carts/{id}", server.GetCart).Methods("GET")
	server.Router.HandleFunc("/carts", server.AddItem).Methods("POST")
	server.Router.HandleFunc("/carts/{id}", server.DeleteItem).Methods("DELETE")

	server.Router.HandleFunc("/orders", server.CreateOrder).Methods("POST")
	server.Router.HandleFunc("/orders/{id}", server.GetOrderByUserID).Methods("GET")
	server.Router.HandleFunc("/orders/{id}", server.UpdateOrder).Methods("PUT")
```

```
cd se15.1
cd frontend
docker build -t frontend-flask .
cd ..

docker build -t backend-golang --target program .
docker build -t mysql-db --target db .
docker-compose up
```


Start
```
kubectl apply -f backend-deployment.yaml 
kubectl apply -f backend-service.yaml 
kubectl apply -f database-deployment.yaml 
kubectl apply -f database-service.yaml 
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