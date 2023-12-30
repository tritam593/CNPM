const product = [
    {image: '../data/images/balo/balo1/minibalo2.png', title: 'Mini Album', price: 120, Qty: 2},
    {image: '../data/images/chair/chair2/i1.jpg', title: 'Chair', price: 80, Qty: 1},
    {image: '../data/images/keyboard/keyboard1/keyboard_1_1.jpg', title: 'Keyboard', price: 100, Qty: 2},
    {image: '../data/images/handbag/handbag2/tui_xach_xanh.png', title: 'Handbag', price: 2000, Qty: 1},
];

const category = [...new Set(product.map((item) => ({...item})))];

function delElement(a) {
    category.splice(a, 1);
    displaycart();
}

function quantityChanged(event) {
    var input = event.target;
    let index = input.getAttribute('data-index');
    let newQty = parseInt(input.value);

    if (isNaN(newQty) || newQty <= 0) {
        input.value = 1;
        newQty = 1;
    }

    category[index].Qty = newQty;
    displaycart();
}

function displaycart() {
    let j = 0,
        total = 0;

    document.getElementById("CountA").innerHTML = category.length + " Items";
    document.getElementById("CountB").innerHTML = category.length + " Items";

    if (category.length === 0) {
        document.getElementById("root").innerHTML = "Your cart is empty";
        document.getElementById("Subtotal").innerHTML = "$ 00.00";
        document.getElementById("TotalPrice").innerHTML = "$ 00.00";
    } else {
        let cartHTML = category.map((items, index) => {
            let {image, title, price, Qty} = items;
            let subitem = price * Qty;
            total = total + subitem;
            document.getElementById("Subtotal").innerHTML = "$" + total.toFixed(2);

            if (total > 3000) {
                document.getElementById("Ship").innerHTML = "Freeship";
                document.getElementById("TotalPrice").innerHTML = "$" + total.toFixed(2);
            } else {
                document.getElementById("Ship").innerHTML = '$' + 25 + ".00";
                document.getElementById("TotalPrice").innerHTML = "$" + (total + 25).toFixed(2);
            }

            return `<tr>
                        <td>
                            <i class="fa-regular fa-circle-xmark" onclick='delElement(${j++})'></i>
                        </td>
                        <td>
                            <div class='img-box'><img class="img" src=${image}></div>
                        </td>
                        <td><p style='font-size:15px;'>${title}</p></td>
                        <td><h2 style='font-size:15px; color:red;'>$ ${price}.00 </h2></td>
                        <td>
                            <div class='Quanty'>
                                <input type="number" class="Qty-box" value="${Qty}" data-index="${index}">
                            </div>
                        </td>
                        <td><h2 style='font-size:15px; color:red;'>$ ${subitem.toFixed(2)} </h2></td>
                    </tr>`;
        }).join('');

        document.getElementById('root').innerHTML = cartHTML;

      
        document.querySelectorAll('.Qty-box').forEach((input) => {
            input.addEventListener('change', quantityChanged);
        });
    }
}

displaycart();


