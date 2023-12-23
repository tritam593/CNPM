const pro_img = document.querySelector(".single-pro-img");

const jsonFile = "../single_product.json";
fetch(jsonFile)
    .then((response) => {
        return response.json();
    })
    .then((data) => {
        
        const { ProductImages, Name, Price } = data[0];
        console.log("Product Data:", data);

        // Set the main image source
        pro_img.innerHTML += `
            <img src="../assets/products/f3.jpg" alt="${Name}" width="100%" id="mainImg">
            <div class="small-img-group">
        `;
        const img_group = document.querySelector(".small-img-group");

        // Create small image elements and append them to the small-img-group
        for (let i = 0; i < ProductImages.length; i++) {
            img_group.innerHTML += `
                <div class="small-img-col">
                    <img src="${ProductImages[i].Path}" width="100%" alt="${Name} ${i + 1}" class="small-img">
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
    });
console.log("end");

    
    