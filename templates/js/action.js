$(document).ready(function () {

    const url = "http://127.0.0.1:9000/";
    
        $.ajax({
            url: url + "carts",
            type: "POST",
            contentType: "application/json",
            success: function (response) {
                console.log("Cart updated successfully:", response);
            },
            error: function (xhr, status, error) {
                // Handle error
                console.error("Failed to update cart:", error);
            }
        });
    });