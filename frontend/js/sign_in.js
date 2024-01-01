const url = "http://127.0.0.1:9000/"
var email = document.getElementById("email");
var password = document.getElementById("password");
$(document).ready(function () {
    
    $("#submit").click(function () {


        var data = {
            email: email.value,
            password: password.value

        };
        const url = "http://127.0.0.1:9000/"
        if (isFormValid()) {
            // Proceed with form submission or any other action
            console.log('Form is valid. Submitting...');

            console.log(data);
            $.ajax({
                url: url + "login",
                type: "POST",
                contentType: "application/json",
                data: JSON.stringify(data),
                success: function (response) {
                    if(Object.keys(response).length == 1){
                        alert("Login fail")
                    }
                    else{
                        var fieldValue = response.ID; 
                    console.log(fieldValue);
                    localStorage.setItem('ID', fieldValue);

                    var fieldValue = localStorage.getItem('ID');
                    console.log(fieldValue);
                    $("#login").text("Result: " + JSON.stringify(response));
                    $("#local").text("Val: " + fieldValue);
                    
                    getCarts();
                    console.log("GET CARTS OK");
                    }
                    
                    
                },
                error: function (xhr, status, error) {
                    // Display an error message if there's an issue with the request
                    $("#login").text("Fail: " + error.message);
                }
        });
        } else {
            console.log('Form is not valid. Please check your inputs.');
        }
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





const setError = (element, message) => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');

    errorDisplay.innerText = message;
    inputControl.classList.add('error');
    inputControl.classList.remove('success')
}

const setSuccess = element => {
    const inputControl = element.parentElement;
    const errorDisplay = inputControl.querySelector('.error');

    errorDisplay.innerText = '';
    inputControl.classList.add('success');
    inputControl.classList.remove('error');
};

const isValidEmail = email => {
    const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
    return re.test(String(email).toLowerCase());
}
const isFormValid = () => {
    const emailValue = email.value.trim();
    const passwordValue = password.value.trim();


    if (emailValue === '' || !isValidEmail(emailValue)) {
        setError(email, 'Provide a valid email address');
        return false;
    } else {
        setSuccess(email);
    }

    if (passwordValue === '' || passwordValue.length < 8) {
        setError(password, 'Password must be at least 8 characters.');
        return false;
    } else {
        setSuccess(password);
    }



    // If all validations pass, return true
    return true;
};