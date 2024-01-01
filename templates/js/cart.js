 /* const product = [
    {image: '../data/images/balo/balo1/minibalo2.png', title: 'Mini Album', price: 120, Qty: 2},
    {image: '../data/images/chair/chair2/i1.jpg', title: 'Chair', price: 80, Qty: 1},
    {image: '../data/images/keyboard/keyboard1/keyboard_1_1.jpg', title: 'Keyboard', price: 100, Qty: 2},
    {image: '../data/images/handbag/handbag2/tui_xach_xanh.png', title: 'Handbag', price: 2000, Qty: 1},
];

const category = [...new Set(product.map((item) => ({...item})))]; */
$(document).ready(function () {
const url = "http://127.0.0.1:9000/";
const pro_id = localStorage.getItem("pro-id");
console.log("ID pro: ",pro_id);
function delElement(a) {
    $.ajax({
        url: url + "carts/" + pro_id,
        method: 'POST',
        data: { index: a },
        success: function() {
            category.splice(a, 1);
            displaycart(); // Refresh the display after removal
        },
        error: function(xhr, status, error) {
            console.error('Error:', status, error);
        }
    });
}



function displaycart() {
    var fieldValue = response.ID; 
    console.log(fieldValue);
    localStorage.setItem('ID', JSON.stringify(fieldValue));

    var fieldValue = localStorage.getItem('ID');
    $.ajax({
        url: url + "carts/" + fieldValue,
        method: 'GET',
        dataType: 'json',
        success: function(data) {
        console.log("ok",data);
        let j = 0,
        total = 0;

    document.getElementById("CountA").innerHTML = category.length + " Items";
    document.getElementById("CountB").innerHTML = category.length + " Items";

    if (category.length === 0) {
        document.getElementById("root").innerHTML = "Your cart is empty";
        document.getElementById("Subtotal").innerHTML = "$ 00.00";
        document.getElementById("TotalPrice").innerHTML = "$ 00.00";
    } else {
        let cartHTML = category.map((item, index) => {
            const { image, title, price, Qty } = data[index];
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
                    <div class='img-box'><img class="img" id="image" src=${image}></div>
                </td>
                <td><p style='font-size:15px;' id="Name">${title}</p></td>
                <td><h2 id='price' style='font-size:15px; color:red;'>$ ${price}.00 </h2></td>
                <td>
                    <button type="submit" class="qtyminus" field="quantity" name="minus" id="value-minus2" onclick="minusqty()">-</button> 
                    <input type="text" name="quantity" class="qty" id="value2" value="${Qty}">
                    <button type="submit" class="qtyplus" field="quantity" name="plus" id="value-plus2" onclick="plusqty()">+</button>  
            
                </td>
                <td><h2 style='font-size:15px; color:red;'>$ ${subitem.toFixed(2)} </h2></td>
            </tr>`;
        }).join('');

        document.getElementById('root').innerHTML = cartHTML;

      
        function minusqty() {//update when press button -
            var quantityVal = $("input[name='quantity']").val();
            var qtyVal = quantityVal-1; 
            var idVal = $("input[name='id']").val();
            var itemidVal = $("input[name='item_id']").val();
            var cartidVal = $("input[name='cart_id']").val();
            $.ajax({
                url: 'cart-update.php',
                method: 'GET',
                data: {Qty: qtyVal, item_id: itemidVal, id: idVal, cart_id: cartidVal },
                cache: false,
                dataType: 'html',
                success: function(data) {
            
                },
            });
            }
            function plusqty() {////update when press button +
                var quantityVal2 = $("input[name='quantity']").val();
                var qtyVal2 = quantityVal2-(-1); 
                var idVal2 = $("input[name='id']").val();
                var itemidVal2 = $("input[name='item_id']").val();
                var cartidVal2 = $("input[name='cart_id']").val();
                $.ajax({
                    url: 'cart-update.php',
                    method: 'GET',
                    data: {qty: qtyVal2, item_id: itemidVal2, id: idVal2, cart_id: cartidVal2 },
                    cache: false,
                    dataType: 'html',
                    success: function(data) {
                
                    },
                });
                }
},
error: function(xhr, status, error) {
    console.error('Error fetching cart data:', status, error);
}
    });


displaycart();
}


