import unittest
import source.cart as cart
class TestCart(unittest.TestCase):
    def setUp(self):
        # Khởi tạo một đối tượng User cho mỗi test case
        # self.user.login()
        self.cart = cart.Cart('b77a02d5-9c62-4e15-9d20-3e3c05d287ac','7c39f4c4-06e0-4b91-972a-7eee52176169')
        print("====== set up ======")
    def test_get_cart(self):
        cart_temp = self.cart.get_cart() 
        print(cart_temp)
        self.assertTrue(len(cart_temp) > 1)
        
    # def a_cart_is_empty(self):
    #     cart_temp = self.cart.get_cart() 
    #     self.assertTrue(len(cart_temp) == 1)
        
    def test_add_item(self):
        print("====================test_add_cart==================")
        self.cart.add_item()
        all_cart = self.cart.get_cart()
        check_add_cart = any(item["productID"] == "df64bd0d-809a-4321-93b9-8551552d17b4" for item in all_cart)
        self.assertTrue(check_add_cart)
    # def test_delete_item(self): 
    #     print("====================test_delete_cart==================")
    #     self.cart.add_item()
    #     print(self.cart.id)
    #     r = self.cart.delete_item()
    #     print(r)
    #     all_cart = self.cart.get_cart()
    #     check_delete_item = any(item["ID"] == self.cart.id for item in all_cart)
    #     self.assertFalse(check_delete_item)
        