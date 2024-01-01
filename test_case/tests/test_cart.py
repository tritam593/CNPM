import unittest
import source.cart as cart
class TestCart(unittest.TestCase):
    def setUp(self):
        # Khởi tạo một đối tượng User cho mỗi test case
        # self.user.login()
        self.cart = cart.Cart('1ae0725f-af6f-49f1-abe7-110be47adc4e')
        self.item = cart.Item("","Balo",'2000','2')
        print("====== set up ======")
    
    def test_get_cart(self):
        cart_temp = self.cart.get_cart() 
        print(cart_temp)
        self.assertTrue(len(cart_temp) > 1)
    def test_add_item(self):
        
         cart_temp = self.cart.get_cart() 
         print(cart_temp)
         self.cart.add_item(self.item)
         cart_curr = self.cart.get_cart() 
         self.assertTrue(len(cart_curr) > len(cart_temp))
    def test_delete_item(self): 
        print("====================test_delete_cart==================")
        r = self.cart.delete_item()
        print(r)
        all_product = self.cart.get_cart()
        check_delete_product = any(item["ID"] == product_id for item in all_product)
        self.assertFalse(check_delete_product)
        