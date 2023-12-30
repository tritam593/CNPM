$(document).ready(function () {
    // var productsContainer = $("#productsContainer");
    const url = "http://127.0.0.1:9000/";
    const pro_img = document.querySelector(".single-pro-img");
    const baseImagePath = "../../";
    const pro_id = localStorage.getItem("pro-id");
    console.log("ID pro: ",pro_id);
    // Gửi yêu cầu GET đến "http://127.0.0.1:9000/products" sử dụng jQuery AJAX
    $.ajax({
        url: url + "products/"+pro_id,
        type: "GET",
        contentType: "application/json",
        success: function (data) {
            console.log("ok",data);
            const { ProductImages, Name, Price } = data;
            console.log("Main src ",ProductImages[0].Path)
            pro_img.innerHTML += `
            <img src="${baseImagePath}${ProductImages[0].Path}" alt="${Name}" width="100%" height="60%" id="mainImg">
            <div class="small-img-group">
            `;
            const img_group = document.querySelector(".small-img-group");

            // Create small image elements and append them to the small-img-group
            for (let i = 0; i < ProductImages.length; i++) {
                img_group.innerHTML += `
                    <div class="small-img-col">
                        <img src="${baseImagePath}${ProductImages[i].Path}" width="100%" alt="${Name} ${i + 1}" class="small-img">
                    </div>
                `;
                console.log(`Image ${i + 1} path: ${ProductImages[i].Path}`);
            }
            var mainImg = document.getElementById("mainImg");
            var smallImg = document.getElementsByClassName("small-img");
            console.log(smallImg);
            console.log('ok2');
            for (let i = 0; i < smallImg.length; i++) {
                smallImg[i].addEventListener("click", () => {
                    mainImg.src = smallImg[i].src;
                });
                console.log("img",smallImg[i]);
            }
            console.log("ok3");
        },
        error: function (xhr, status, error) {
            // Display an error message if there's an issue with the request
            console.log("fail");
        }
    });
});











// const pro_img = document.querySelector(".single-pro-img");

// const jsonFile = "../single_product.json";
// fetch(jsonFile)
//     .then((response) => {
//         return response.json();
//     })
//     .then((data) => {
        
//         const { ProductImages, Name, Price } = data[0];
//         console.log("Product Data:", data);

//         // Set the main image source
//         pro_img.innerHTML += `
//             <img src="../assets/products/f3.jpg" alt="${Name}" width="100%" id="mainImg">
//             <div class="small-img-group">
//         `;
//         const img_group = document.querySelector(".small-img-group");

//         // Create small image elements and append them to the small-img-group
//         for (let i = 0; i < ProductImages.length; i++) {
//             img_group.innerHTML += `
//                 <div class="small-img-col">
//                     <img src="${ProductImages[i].Path}" width="100%" alt="${Name} ${i + 1}" class="small-img">
//                 </div>
//             `;
//             console.log(`Image ${i + 1} path: ${ProductImages[i].Path}`);
//         }

//         var mainImg = document.getElementById("mainImg");
//         var smallImg = document.getElementsByClassName("small-img");
//         console.log(smallImg);
//         console.log('ok2');
//         for (let i = 0; i < smallImg.length; i++) {
//             smallImg[i].addEventListener("click", () => {
//                 mainImg.src = smallImg[i].src;
//             });
//             console.log("img",smallImg[i]);
//         }
//         console.log("ok3");
//     });
// console.log("end");

    
    