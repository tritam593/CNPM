# server.Router.HandleFunc("/orders", server.CreateOrder).Methods("POST")
# server.Router.HandleFunc("/orders/{id}", server.GetOrderByUserID).Methods("GET") // user id
# server.Router.HandleFunc("/orders/order/{id}", server.GetOrderByOrderID).Methods("GET") // order id
# server.Router.HandleFunc("/orders/{id}", server.UpdateOrder).Methods("PUT")
# http://127.0.0.1:9000/order/e05f5f7e-3c80-4ee7-be3c-5d19a34abc4f

import requests

class Order:
    def __init__(self):
        

        

        
        
        
    def get_order(self):
        r = requests.get(url=f"http://127.0.0.1:9000/orders/{self.user_id}")
        if r.status_code == 200:
            return r.json()
        raise requests.HTTPError
    
    # def post_order(self):
    #     r = requests.post(url=f"http://127.0.0.1:9000/orders/{self.user_id}",json=self.data)
    #     if r.status_code == 200:
    #         return r.json()
    #     raise requests.HTTPError
