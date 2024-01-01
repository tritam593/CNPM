const url = "http://127.0.0.1:9000/"
$(document).ready(function () {
    
    $("#submit").click(function () {
        var email = document.getElementById("email").value;
        var pass = document.getElementById("password").value;
        // var add = document.getElementById("addr").value;
        // var firstname = document.getElementById("fname").value;
        // var lastname = document.getElementById("lname").value;

        // Tạo một đối tượng JSON từ các giá trị
        var data = {
            email: email,
            password: pass,
            // lastname: lastname,
            // firstname: firstname,
            // address: add
        };
        console.log(data);
        


   

        $.ajax({
            url: url + "login",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: function (response) {
                var fieldValue = response.ID; 
                console.log(fieldValue);
                localStorage.setItem('ID', fieldValue);

                var fieldValue = localStorage.getItem('ID');
                console.log(fieldValue);
                // Handle JSON data from the response
                $("#login").text("Result: " + JSON.stringify(response));
                $("#local").text("Val: " + fieldValue);
                
                getCarts();
                console("GET CARTS OK");
                
                // window.location.href = "shop.html"
                
            },
            error: function (xhr, status, error) {
                // Display an error message if there's an issue with the request
                $("#login").text("Fail: " + error.message);
            }
        });
    });


    
});

function getCarts(){
    var id_user = localStorage.getItem("ID")
    $.ajax({
        
        url: url+ "carts/"+id_user,
        type: "GET",
        contentType: "application/json",
        crossDomain: true,
        success: function (response) {
            // Handle JSON data from the response
            console.log("Result carts: " + JSON.stringify(response));
            var idValue = response.ID;
            localStorage.setItem("ID-Carts",idValue)
            console.log("ID CART: ",localStorage.getItem("ID-Carts"));
            window.location.href = "../index.html";
            
        },
        error: function (xhr, status, error) {
            // Display an error message if there's an issue with the request
            
            console.log("Fail: " + error.message);
        }
    });
}