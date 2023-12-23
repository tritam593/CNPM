document.addEventListener('DOMContentLoaded', function () {
    var nav = document.getElementById("navbar");



    var isLoggedIn = false; 


    if (isLoggedIn) {
        nav.innerHTML += '<li id="lg-user"><a href=""><i class="far fa-user"></i></a></li>';
        nav.innerHTML += '<li id="logout"><a href="">Logout</a></li>';
    } else {
        nav.innerHTML += '<li id="sign-in"><a href="../sub-pages/sign-in.html">Sign in</a></li>';
        nav.innerHTML += '<li id="sign-up"><a href="../sub-pages/sign-up.html">Sign up</a></li>'
    }
});


