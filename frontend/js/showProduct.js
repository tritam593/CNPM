const container = document.querySelector(".pro-container");

const jsonFile = "../products.json";
fetch(jsonFile)
    .then((response) => {
        return response.json();
    })
    .then((data) => {
        data.forEach((product) => {
        const { ProductImages, Name, Price } = product;

                container.innerHTML += `
                <div class="pro" onclick="window.location.href='sproduct.html'">
                    <img src="${ProductImages[0].Path}" alt="Product Image">
                    <div class="des">
                        <span>adidas</span>
                        <h5>${Name}</h5>
                        <div class="star">
                            <i class="fa-solid fa-star"></i>
                            <i class="fa-solid fa-star"></i>
                            <i class="fa-solid fa-star"></i>
                            <i class="fa-solid fa-star"></i>
                            <i class="fa-solid fa-star"></i>
                        </div>
                        <h4>${Price}</h4>
                    </div>
                    <a href="#" class="cart"><i class="fa-solid fa-cart-shopping"></i></a>
                </div>
                `;

                
        });
});

