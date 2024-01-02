import requests

      
class Cart: 
    def __init__(self,user_id,cart_item_id):
        self.user_id = user_id
        self.cart_item_id = cart_item_id
        self.id = ""
        self.data = {
                    "CartID": "7c39f4c4-06e0-4b91-972a-7eee52176169",
                    "productID": "df64bd0d-809a-4321-93b9-8551552d17b4",
                    "Qty": 10,
                }

    def get_cart(self):
        r = requests.get(url=f"http://127.0.0.1:9000/carts/{self.user_id}")
        if r.status_code == 200: 
            self.user_cart = r.json()
            return self.user_cart
        
        raise requests.HTTPError
    def add_item(self):
        r = requests.post(url=f"http://127.0.0.1:9000/carts",json=self.data)
        all_cart = self.get_cart()
        self.id = next((item["ID"] for item in all_cart if item["productID"]== "df64bd0d-809a-4321-93b9-8551552d17b4"), None)
        if r.status_code == 200:
           return r.json()
        raise requests.HTTPError
        
    
    def delete_item(self):
        r = requests.delete(url=f"http://127.0.0.1:9000/carts/{self.cart_item_id}")
        if r.status_code == 200:
           return r.json()
        raise requests.HTTPError
            
        
        