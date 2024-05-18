
let CurrentProductName = "";
let ProductsOfExport = [];
jQuery('#myModal').modal('show')
// add  New product to export list

jQuery("#AddProductToExport").on("click", function () {
    // jQuery.ajax({
    //     method: "POST",
    //     url: "/Dashboard/export",
    //     data: JSON.stringify({ name: "John", location: "Boston" }),
    //     contentType: "application/json; charset=utf-8",
    // })
    //     .done(function (msg) {
    //         console.log(msg);
    //     });
    var ID = jQuery("#ProductIs").val()
    // var Name=jQuery("#ProductIs").html()
    var Count = jQuery("#ProductBox input[name='Count']").val()
    var MeterPrice = jQuery("#ProductBox input[name='Meter']").val()
    var RolePrice = jQuery("#ProductBox input[name='RolePrice']").val()
    var MeterPrice = jQuery("#ProductBox input[name='MeterPrice']").val()
    var TotalPrice = jQuery("#ProductBox input[name='TotalPrice']").val()
    var value = '<tr><th scope="row">' + ID + '</th><td>' + CurrentProductName + '</td><td>23423</td><td>' + Count + '</td><td>' + MeterPrice + '</td><td>' + RolePrice + '</td><td>' + TotalPrice + '</td></tr>';
    var newRow = {
        // ExportID: ID,
        Name: CurrentProductName,
        count: Count,
        meterPrice: MeterPrice,
        rolePrice: RolePrice,
        totalPrice: TotalPrice
    };
    ProductsOfExport.push(newRow)
    console.log(ProductsOfExport)
    jQuery("#ExportProductsList tbody").append(value);
    var oldprice = jQuery(".TotalPriceOut td").html();
    oldprice = parseFloat(oldprice);
    TotalPrice = parseFloat(TotalPrice)
    TotalPrice = oldprice + TotalPrice;
    jQuery(".TotalPriceOut td").html(TotalPrice)
    jQuery(".Notfound").slideUp();
    jQuery(".close").click()
})


// Select  Inventory name for fech the produts of inventory

jQuery(".ExportPeoducts select#InventoryIS").on("change", function () {
    jQuery("#ProductIs").prop("disabled", true);
    var ID = this.value;
    // CurrentProductName=jQuery(this).find("option:selected").text();
    if (ID != 0) {
        jQuery.ajax({
            method: "POST",
            url: "/Dashboard/getproductbyinventory",
            data: JSON.stringify({ name: "InventoryIS", "id": ID }),
            contentType: "application/json; charset=utf-8",
        })
            .done(function (msg) {

                console.log(msg.result.length)
                setTimeout(function () {
                    jQuery("#ProductIs").empty();
                    jQuery("#ProductIs").append('<option value="0">لطفا یک گزینه را انتخاب کنید</option>')
                    if (msg.result.length > 0) {
                        msg.result.forEach(item => {
                            jQuery("#ProductIs").append('<option value="' + item.ID + '">' + item.Name + '</option>')
                        });
                    }
                }, 200)
                setTimeout(function () {
                    jQuery("#ProductIs").prop("disabled", false);
                    jQuery(".ProductIs").slideDown();
                }, 210)
            });
    } else {
        jQuery(".modal-footer").slideUp();
        jQuery(".ProductIs").slideUp();
    }
})

// Select  Product name for fech the detail of product

jQuery(".ExportPeoducts select#ProductIs").on("change", function () {
    var ID = this.value;
    CurrentProductName = jQuery(this).find("option:selected").text();
    if (ID != 0) {
        jQuery.ajax({
            method: "POST",
            url: "/Dashboard/getproductbyid",
            data: JSON.stringify({ name: "ProductIs", "id": ID }),
            contentType: "application/json; charset=utf-8",
        })
            .done(function (msg) {
                console.log(msg.result.length)
                if (msg.result.length > 0) {
                    var product = msg.result[0];
                    jQuery(".ExportPeoducts .ProductsCount").html(product.Count + "تعداد موجود")
                    jQuery(".ExportPeoducts .ProductNumber").html(product.Number)
                    jQuery(".ExportPeoducts input[name='RolePrice']").val(product.RolePrice)
                    jQuery(".ExportPeoducts input[name='MeterPrice']").val(product.MeterPrice)
                    jQuery(".ExportPeoducts input[name='TotalPrice']").val(product.MeterPrice)
                    jQuery(".ExportPeoducts .Content").slideDown();
                }
                console.log(msg.result[0]);
            });
    } else {
        jQuery(".modal-footer").slideUp();
        jQuery(".Content").slideUp();
    }
})



// Calculate TotalPrice by Count of Role 

jQuery("input[name='Count']").on("change", function () {
    var number = parseInt(this.value)
    if (number != 0) {
        jQuery(".modal-footer").slideDown();

    } else {
        jQuery(".modal-footer").slideUp();
    }
    // setTimeout(function(){
    jQuery("input[name='Meter']").val(0);
    RolePrice = parseFloat(jQuery(".ExportPeoducts input[name='RolePrice']").val());
    Count = parseFloat(jQuery(this).val());

    if (isNaN(RolePrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = RolePrice * Count;
    }
    jQuery("input[name='TotalPrice']").val(TotalPrice)
    // },1000)

})

// Calculate TotalPrice by price of meter  

jQuery("input[name='Meter']").on("change", function () {
    var number = parseInt(this.value)
    if (number != 0) {
        jQuery(".modal-footer").slideDown();

    } else {
        jQuery(".modal-footer").slideUp();
    }
    // setTimeout(function(){
    jQuery("input[name='Count']").val(0);
    RolePrice = parseFloat(jQuery(".ExportPeoducts input[name='MeterPrice']").val());
    Count = parseFloat(jQuery(this).val());

    if (isNaN(RolePrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = RolePrice * Count;
    }
    jQuery("input[name='TotalPrice']").val(TotalPrice)
    // },1000)

})
// function getFormData($form){
//     var unindexed_array = $form;
//     var indexed_array = {};

//     jQuery.map(unindexed_array, function(n, i){
//         indexed_array[n['name']] = n['value'];
//     });

//     return indexed_array;
// }

jQuery("form[name='expotform']").submit(function (e) {
   e.preventDefault();
    var formValues = jQuery("form[name='expotform']").find("input, select, textarea").map(function() {
        return $(this).attr("name") + "=" + $(this).val();
      }).get().join("&");
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/export",
        data: JSON.stringify({ Name: "expotform", Content:formValues, Products: ProductsOfExport }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            console.log(msg)
        });

})
function printpage() {
  
 const divToPrint = document.getElementById('card');
const style = document.createElement('style');
style.media = 'print';
style.innerHTML = `
  @media print {
    /* hide all elements except the div */
    * {
      display: none;
    }
    #myDiv {
      display: block;
    }
  }
`;
document.head.appendChild(style);
window.print();
}

// share


    const shareButtons = document.querySelectorAll('.sharebtn');

    // Add click event listener to each button
    shareButtons.forEach(button => {
       button.addEventListener('click', () => {
          // Get the URL of the current page
          const url = window.location.href;
    
          // Get the social media platform from the button's class name
          const platform = button.classList[1];
    
          // Set the URL to share based on the social media platform
          let shareUrl;
          switch (platform) {
             case 'facebook':
             shareUrl = `https://www.facebook.com/sharer/sharer.php?u=${encodeURIComponent(url)}`;
             break;
             case 'twitter':
             shareUrl = `https://twitter.com/share?url=${encodeURIComponent(url)}`;
             break;
             case 'linkedin':
             shareUrl = `https://www.linkedin.com/shareArticle?url=${encodeURIComponent(url)}`;
             break;
             case 'pinterest':
             shareUrl = `https://pinterest.com/pin/create/button/?url=${encodeURIComponent(url)}`;
             break;
             case 'reddit':
             shareUrl = `https://reddit.com/submit?url=${encodeURIComponent(url)}`;
             break;
             case 'whatsapp':
             shareUrl = `https://api.whatsapp.com/send?text=${encodeURIComponent(url)}`;
             break;
          }
    
          // Open a new window to share the URL
          window.open(shareUrl, '_blank');
       });
    });
    