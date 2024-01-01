import requests
class Item: 
    def __init__(self,img,name,price,qty):
         self.img = img
         self.name = name 
         self.price = price
         self.qty = qty
      
class Cart: 
    def __init__(self,user_id,cart_item_id):
        self.user_id = user_id
        self.cart_item_id = cart_item_id
        self.user_cart = []
    def get_cart(self):
        r = requests.get(url=f"http://127.0.0.1:9000/carts/{self.user_id}")
        if r.status_code == 200: 
            self.user_cart = r.json()
            return self.user_cart
        
        raise requests.HTTPError
    def add_item(self,item):
        r = requests.post(url=f"http://127.0.0.1:9000/carts")
        if r.status_code == 200:
            self.user_cart.append(item)
            return self.user_cart
        raise requests.HTTPError
        
    
    def delete_item(self,item):
        r = requests.delete(url=f"http://127.0.0.1:9000/carts/{self.cart_item_id}")
        if r.status_code == 200:
            return self.user_cart.pop(item)
        raise requests.HTTPError
            
        
        