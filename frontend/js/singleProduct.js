const url = "http://127.0.0.1:9000/";
$(document).ready(function () {
    // var productsContainer = $("#productsContainer");
    
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
            
            for (let i = 0; i < smallImg.length; i++) {
                smallImg[i].addEventListener("click", () => {
                    mainImg.src = smallImg[i].src;
                });
                console.log("img",smallImg[i]);
            }
            
        },
        error: function (xhr, status, error) {
            // Display an error message if there's an issue with the request
            console.log("fail");
        }
    });

         
});



// script.js
$(document).ready(function() {
    
    // Lắng nghe sự kiện khi người dùng nhấn nút "Add to cart"
    $("#add-to-cart").click(function() {
        // Đọc ID sản phẩm từ localStorage (điều chỉnh theo cách bạn lưu trữ ID)
        var productId = localStorage.getItem("pro-id");
        
        // Đọc số lượng từ input type number
        var quantity = $("#qty").val();
        var id_cart = localStorage.getItem("ID-Carts")
        // Kiểm tra xem ID sản phẩm và số lượng có giá trị hợp lệ không
        if (productId && quantity &&id_cart) {
            // Gửi yêu cầu AJAX
            data = {
                "CartID": id_cart,
                "productId": productId,
                "Qty": parseInt(quantity)
            },
            console.log("DATA POST: ",data)
            $.ajax({
                url: url+"carts",  // Điều chỉnh đường dẫn đến backend của bạn
                type: "POST",
                data: JSON.stringify(data),
                success: function(response) {
                    // Xử lý phản hồi từ server (nếu cần)
                    console.log("POST OK")
                    console.log("Success:", response);
                },
                error: function(error) {
                    // Xử lý lỗi (nếu có)
                    console.error("Error:", error);
                }
            });
        } else {
            console.error("Invalid product ID or quantity.");
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

    
    