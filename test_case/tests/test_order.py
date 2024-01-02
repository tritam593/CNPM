import unittest
import source.cart as cart
import source.user as user
import source.product as product
import source.order as order


class TestOrder(unittest.TestCase):
    def setUpClass(cls):
        print("=================Start set up class==================")
        cls.user = user.User("testuser@example.com", "a", "Test", "User", "as,ndsalkdsalk")
        

        cls.product1 = product.Product()
        cls.product1.data["Name"] = "mouse_test1"
        cls.product2 = product.Product()
        cls.product2.data["Name"]= "mouse_test2"
        cls.product3 = product.Product()
        cls.product3.data["Name"]= "mouse_test3"

        response = cls.user.register()
        print("User response",response)
        cls.product1.post_new_product()
        cls.product2.post_new_product()
        cls.product3.post_new_product()

        id_user = cls.user.login()["ID"]
        
        cls.cart = cart.Cart(id_user)
        cls.cart.data["productID"]
        cls.cart.add_item()


        

        assert 1 == len(response)
        assert "Check" in (response.keys())
        assert "ok" == response["Check"]
        print("=================END set up class==================")


    def setUp(self):
        # Khởi tạo một đối tượng User cho mỗi test case
        print("======STAR set up===============")
    

        
        self.cart.add_item()
        self.order = order.Order(id_user)

        all_cart = self.cart.get_cart()
        self.cart.cart_item_id = all_cart["ID"]
        self.cart.data["CartID"] = self.cart.cart_item_id

        self.cart.product_id = self.product.id
        self.cart.data["productID"] = self.cart.product_id


        print("ID USER",self.cart.user_id)
        print("ID CARTS",self.cart.cart_item_id)
        print("ID PRODUCT",self.cart.product_id)
        print("======END set up ======")


    @classmethod
    def tearDownClass(cls):
        print("===========Start tear down class===========")
        cls.user.login()
        print(cls.user.id)
        response = cls.user.delete_user(cls.user.id)
        print("DELETE USER ",response)
        response2 = cls.product.delete_product()
        print("DELETE PRODUCT ",response2)
        assert 1 == len(response)
        assert "Check" in (response.keys())
        assert "ok" == response["Check"]
        print("===========END tear down class===========")