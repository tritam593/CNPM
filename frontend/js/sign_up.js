$(document).ready(function () {
    $("#submit").click(function () {
        var email = document.getElementById("email").value;
        var pass = document.getElementById("password").value;
        var add = document.getElementById("addr").value;
        var firstname = document.getElementById("fname").value;
        var lastname = document.getElementById("lname").value;

        // Tạo một đối tượng JSON từ các giá trị
        var data = {
            email: email,
            password: pass,
            lastname: lastname,
            firstname: firstname,
            address: add
        };

        // document.getElementById("test").innerHTML = data["email"];
        const url = "http://127.0.0.1:9000/"


        $.ajax({
            url: "http://127.0.0.1:9000/register",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            crossDomain: true,
            success: function (response) {
                // Handle JSON data from the response
                
                $("#regis").text("Result: " + JSON.stringify(response));
                
                if (response.Check === "ok") {
                    $("#regis").text("Result: " + JSON.stringify(response));
                    window.location.href = "../sub-pages/sign-in.html";
                } else {
                    $("#regis").text("Fail: Check is not ok.");
                }
            },
            error: function (xhr, status, error) {
                // Display an error message if there's an issue with the request
                
                $("#regis").text("Fail: " + error.message);
            }
        });

    });
});