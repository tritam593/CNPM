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

        const url = "http://127.0.0.1:9000/"


   

        $.ajax({
            url: url + "login",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            success: function (response) {
                var fieldValue = response.ID; 
                console.log(fieldValue);
                localStorage.setItem('ID', JSON.stringify(fieldValue));

                var fieldValue = localStorage.getItem('ID');
                console.log(fieldValue);
                // Handle JSON data from the response
                $("#login").text("Result: " + JSON.stringify(response));
                $("#local").text("Val: " + fieldValue);
                // window.location.href = "shop.html"
                window.location.href = "../index.html";
            },
            error: function (xhr, status, error) {
                // Display an error message if there's an issue with the request
                $("#login").text("Fail: " + error.message);
            }
        });




    });
});