import unittest
import source.cart as cart
import source.user as user


class TestCart(unittest.TestCase):
    @classmethod
    def setUpClass(cls):
        cls.user = user.User("testuser@example.com", "a", "Test", "User", "as,ndsalkdsalk")
        print(cls.user.data)
        response = cls.user.register()
        print(response)
        
        assert 1 == len(response)
        assert "Check" in (response.keys())
        assert "ok" == response["Check"]
        print("set up class")

    # test delete account
    @classmethod
    def tearDownClass(cls):
        print("===========tear down class===========")
        cls.user.login()
        print(cls.user.id)
        response = cls.user.delete_user(cls.user.id)
        print(response)
        assert 1 == len(response)
        assert "Check" in (response.keys())
        assert "ok" == response["Check"]
        print("===========tear down class===========")





    def setUp(self):
        # Khởi tạo một đối tượng User cho mỗi test case
        print("======STAR set up===============")
        response2 = self.user.login()
        id_user = response2["ID"]

        self.cart = cart.Cart(id_user)
        all_cart = self.cart.get_cart()
        self.cart.cart_item_id = all_cart["ID"]
        self.cart.data["CartID"] = self.cart.cart_item_id


        print("======END set up ======")
    

    def test_get_cart(self):
        cart_temp = self.cart.get_cart() 
        print(cart_temp)
        self.assertTrue(len(cart_temp) > 1)
        
        
    def test_add_item(self):
        print("====================test_add_cart==================")
        #add
        self.cart.add_item()
        all_cart = self.cart.get_cart()
        #check
        check_add_cart = any(item["ProductID"] == "d50daef9-7447-472b-8afe-8fe48a7b4793" for item in all_cart["CartItems"])

        #delete
        id_cart_item = [item["ID"] for item in all_cart["CartItems"]][0]
        self.cart.delete_item(id_cart_item)
        
        self.assertTrue(check_add_cart)


    def test_delete_item(self): 


        print("====================test_delete_cart==================")
        #add
        self.cart.add_item()
        all_cart = self.cart.get_cart()
        id_cart_item = [item["ID"] for item in all_cart["CartItems"]][0]


        #check delete
        self.cart.delete_item(id_cart_item)
        all_cart = self.cart.get_cart()
        check_delete_item = any(item["ProductID"] == id_cart_item for item in all_cart["CartItems"])

        self.assertFalse(check_delete_item)
        