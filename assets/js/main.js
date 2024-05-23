
let CurrentProductName = "";
let ProductsOfExport = [];
let ExportTotalPrice = [];
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
    var ID = jQuery("#ProductIs").val();
    var ExportID = jQuery("input[name='ExportNumber']").val();
    var InventoryNumber = jQuery("#InventoryIS").val();
    // var Name=jQuery("#ProductIs").html()
    var Count = jQuery("#ProductBox input[name='Count']").val()
    var MeterPrice = jQuery("#ProductBox input[name='Meter']").val()
    var RolePrice = jQuery("#ProductBox input[name='RolePrice']").val()
    var MeterPrice = jQuery("#ProductBox input[name='MeterPrice']").val()
    var TotalPrice = jQuery("#ProductBox input[name='TotalPrice']").val()
 
    // var oldprice = jQuery("input[name='TotalPrice']")
    // oldprice=parseFloat(1)
    // TotalPrice = parseFloat(TotalPrice)
    // ExportTotalPrice = jQuery("input[name='ExportTotalPrice']").val()
   
        var val = jQuery(this).html();
        var val = PersianTools.addCommas(TotalPrice);
        var TotalPricefa = PersianTools.digitsEnToFa(val);

    
 
    var value = '<tr><th scope="row">' + ID + '</th><td>' + CurrentProductName + '</td><td>23423</td><td>' + Count + '</td><td class="price">' + MeterPrice + '</td><td>' + RolePrice + '</td><td class="itemtotalprice price">' + TotalPricefa + '</td></tr>';
    var newRow = {
        InventoryNumber: InventoryNumber,
        ProductId: ID,
        ExportID: ExportID,
        Name: CurrentProductName,
        count: Count,
        meterPrice: MeterPrice,
        rolePrice: RolePrice,
        totalPrice: TotalPrice

    };
    var NewPrice={
        price:TotalPrice
    }

    ProductsOfExport.push(newRow)
    ExportTotalPrice.push(NewPrice)
    // console.log(ExportTotalPrice)
    exporttotal_Price="0"
    ExportTotalPrice.forEach(function(e,i){
        exporttotal_Price = parseFloat(exporttotal_Price) +  parseFloat(e.price)
    })

    jQuery("#ExportTotalPrice").attr("value",exporttotal_Price);
    jQuery("#ExportTotalPrice").val(exporttotal_Price);
    jQuery("#ExportProductsList tbody").append(value);
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
                if (msg.result.length > 0) {
                    var product = msg.result[0];
                    jQuery(".ExportPeoducts .ProductsCount").html(product.Count + "تعداد موجود")
                    jQuery(".ExportPeoducts .ProductNumber").html(product.Number)
                    jQuery(".ExportPeoducts input[name='RolePrice']").val(product.RolePrice)
                    jQuery(".ExportPeoducts input[name='MeterPrice']").val(product.MeterPrice)
                    jQuery(".ExportPeoducts input[name='TotalPrice']").val(product.MeterPrice)
                    jQuery(".ExportPeoducts .Content").slideDown();
                    jQuery(".ExportPeoducts  input[type='number']").each(function () {
                        var val = jQuery(this).val();
                        var val = PersianTools.addCommas(val);
                        var convertToFa = PersianTools.digitsEnToFa(val);
                        var numberToWords = PersianTools.numberToWords(val);
                        jQuery(this).parent().closest(".form-group").find(".out").html(convertToFa + "   " + numberToWords);
                    });
                }
            });
    } else {
        jQuery(".modal-footer").slideUp();
        jQuery(".Content").slideUp();
    }
})



// Calculate TotalPrice by Count of Role 

jQuery("input[name='Count']").on("keyup", function () {
    var number = parseInt(this.value)
    if (number != 0) {
        jQuery(".modal-footer").slideDown();

    } else {
        jQuery(".modal-footer").slideUp();
    }
    // setTimeout(function(){
    // jQuery("input[name='Meter']").val(0);
    RolePrice = parseFloat(jQuery(".ExportPeoducts input[name='RolePrice']").val());
    Count = parseFloat(jQuery(this).val());

    if (isNaN(RolePrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = RolePrice * Count;
        // ExportTotalPrice =parseFloat(ExportTotalPrice) + parseFloat(TotalPrice)
    }
    jQuery("#ProductBox input[type='number']").each(function () {
        var val = jQuery(this).val();
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);
        var numberToWords = PersianTools.numberToWords(val);
        jQuery(this).parent().closest(".form-group").find(".out").html(convertToFa + "   " + numberToWords);
    });
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
    // jQuery("input[name='Count']").val(0);
    RolePrice = parseFloat(jQuery(".ExportPeoducts input[name='MeterPrice']").val());
    Count = parseFloat(jQuery(this).val());

    if (isNaN(RolePrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = RolePrice * Count;
    }
    jQuery("input[name='TotalPrice']").val(TotalPrice)

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
    var formValues = jQuery("form[name='expotform']").find("input, select, textarea").map(function () {
        return $(this).attr("name") + "=" + $(this).val();
    }).get().join("&");
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/export",
        data: JSON.stringify({ Name: "expotform", TotalPrice: ExportTotalPrice, Content: formValues, Products: ProductsOfExport }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            if (msg.message == "sucess") {
                 window.location.replace("./export-list");
            }
        });

})



jQuery("#exportspaginate a").on("click", function (e) {

    e.preventDefault();
    jQuery("#exportspaginate .page-item").removeClass("active")
    jQuery("#exportspaginate .page-item").removeClass("inpending")
    var page = jQuery(this).attr("attr-page")
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/export-list",
        data: JSON.stringify({ page: page, offset: "1" }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            var lengthofres = msg.message.length;
            if (lengthofres > 0) {
                let html = "";
                msg.message.forEach(function (index) {
                    html += '<tr>';
                    html += '<td class="' + index.ID + '" style="text-align:right;">' + index.ID + '</td>';
                    html += '<td class="' + index.Name + '" style="text-align:right;">' + index.Name + '</td>';
                    html += '<td class="' + index.Number + '" style="text-align:right;">' + index.Number + '</td>';
                    html += '<td class="' + index.Phonenumber + '" style="text-align:right;">' + index.Phonenumber + '</td>';
                    html += '<td class="' + index.Address + '" style="text-align:right;">' + index.Address + '</td>';
                    html += '<td class="' + index.TotalPrice + '" style="text-align:right;">' + index.TotalPrice + '</td>';
                    html += '<td class="' + index.Tax + '" style="text-align:right;">' + index.Tax + '</td>';
                    html += '<td class="' + index.CreatedAt + '" style="text-align:right;">' + index.CreatedAt + '</td>';
                    html += '<td class="' + index.InventoryNumber + '" style="text-align:right;">' + index.inventory_number + '</td>';
                    html += '<td dir="ltr" class="Edit" style="text-align:right;"><a href="./edituser?user-id=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pen" viewBox="0 0 16 16">';
                    html += '<path d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44 .854A1.5 1.5 0 0 1 11.5 .796a1.5 1.5 0 0 1 1.998-.001m-.644 .766a.5 .5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a .5 .5 0 0 0 0-.708z"/></svg></a></td>';
                    html += '</tr>';
                    jQuery(e.target).parent().closest("li").addClass("active")
                    jQuery(e.target).parent().closest("li").next("li").addClass("inpending");
                    jQuery(e.target).parent().closest("li").prev("li").addClass("inpending");
                    jQuery(this).addClass("active");
                    // jQuery(this).parent("li").addClass("active");
                    // jQuery(this).closest("li").addClass("active")
                    jQuery(this).addClass("active")
                });
                if (html.length > 0) {
                    jQuery("#exportlist tbody").empty()
                    jQuery("#exportlist tbody").append(html)

                }
            }

        });
})

jQuery("#userspaginate a").on("click", function (e) {

    e.preventDefault();
    jQuery("#userspaginate .page-item").removeClass("active")
    jQuery("#userspaginate .page-item").removeClass("inpending")
    var page = jQuery(this).attr("attr-page")
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/user-list",
        data: JSON.stringify({ page: page, offset: "1" }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            var lengthofres = msg.message.length;
            if (lengthofres > 0) {
                let html = "";
                msg.message.forEach(function (index) {
                    html += '<tr>';
                    html += '<td class="' + index.ID + '" style="text-align:right;">' + index.ID + '</td>';
                    html += '<td class="' + index.Name + '" style="text-align:right;">' + index.Name + '</td>';
                    html += '<td class="' + index.Phonenumber + '" style="text-align:right;">' + index.Phonenumber + '</td>';
                    html += '<td class="' + index.Address + '" style="text-align:right;">' + index.Address + '</td>';
                    html += '<td dir="ltr" class="Edit" style="text-align:right;"><a href="./edituser?user-id=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pen" viewBox="0 0 16 16">';
                    html += '<path d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44 .854A1.5 1.5 0 0 1 11.5 .796a1.5 1.5 0 0 1 1.998-.001m-.644 .766a.5 .5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a .5 .5 0 0 0 0-.708z"/></svg></a></td>';
                    html += '</tr>';
                    jQuery(e.target).parent().closest("li").addClass("active")
                    jQuery(e.target).parent().closest("li").next("li").addClass("inpending");
                    jQuery(e.target).parent().closest("li").prev("li").addClass("inpending");
                    jQuery(this).addClass("active");
                    // jQuery(this).parent("li").addClass("active");
                    // jQuery(this).closest("li").addClass("active")
                    jQuery(this).addClass("active")
                });
                if (html.length > 0) {
                    jQuery("#userlist tbody").empty()
                    jQuery("#userlist tbody").append(html)

                }
            }

        });
})
jQuery("#find").on("click", function (e) {
    e.preventDefault()
    var value = jQuery("#findval").val()
    // console.log(value)
    // value="حسین سلطانیان"
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/export-find",
        data: JSON.stringify({ term: value }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            var lengthofres = msg.message.length;

            if (lengthofres > 0) {
                let html = "";
                msg.message.forEach(function (index) {
                    html += '<tr>';
                    html += '<td class="' + index.ID + '" style="text-align:right;">' + index.ID + '</td>';
                    html += '<td class="' + index.Name + '" style="text-align:right;">' + index.Name + '</td>';
                    html += '<td class="' + index.Number + '" style="text-align:right;">' + index.Number + '</td>';
                    html += '<td class="' + index.Phonenumber + '" style="text-align:right;">' + index.Phonenumber + '</td>';
                    html += '<td class="' + index.Address + '" style="text-align:right;">' + index.Address + '</td>';
                    html += '<td class="' + index.TotalPrice + '" style="text-align:right;">' + index.TotalPrice + '</td>';
                    html += '<td class="' + index.Tax + '" style="text-align:right;">' + index.Tax + '</td>';
                    html += '<td class="' + index.CreatedAt + '" style="text-align:right;">' + index.CreatedAt + '</td>';
                    html += '<td class="' + index.InventoryNumber + '" style="text-align:right;">' + index.InventoryNumber + '</td>';
                    html += '<td dir="ltr" class="Edit" style="text-align:right;">';
                    html +='<a href="./exportshow?ExportId=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-eye" viewBox="0 0 16 16">';
                    html +='<path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8M1.173 8a13 13 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5s3.879 1.168 5.168 2.457A13 13 0 0 1 14.828 8q-.086.13-.195.288c-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5s-3.879-1.168-5.168-2.457A13 13 0 0 1 1.172 8z"/><path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5M4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0"/></svg></a>';
                    html +='<a class="me-3" href="./deleteExport?ExportId='+ index.ID +'"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">';
                    html += '<path d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5M11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47M8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5"/>';
                    html += '</td>';
                    html += '</tr>';
                    // console.log(html)
                });
                // console.log(html)
                if (html.length > 0) {
                    jQuery("#exportlist tbody").empty()
                    jQuery("#exportlist tbody").append(html)

                }
            }

        });
})
jQuery("#Userfind").on("click", function (e) {
    e.preventDefault()
    var value = jQuery("#findval").val()
    // console.log(value)
    // value="حسین سلطانیان"
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/users-find",
        data: JSON.stringify({ term: value }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            var lengthofres = msg.message.length;

            if (lengthofres > 0) {
                let html = "";
                msg.message.forEach(function (index) {
                    
                    html += '<tr>';
                    html += '<td class="' + index.ID + '" style="text-align:right;">' + index.ID + '</td>';
                    html += '<td class="' + index.Name + '" style="text-align:right;">' + index.Name + '</td>';
                    html += '<td class="' + index.Phonenumber + '" style="text-align:right;">' + index.Phonenumber + '</td>';
                    html += '<td class="' + index.Address + '" style="text-align:right;">' + index.Address + '</td>';
                    html += '<td dir="ltr" class="Edit" style="text-align:right;"><a href="./deleteuser?user-id=' + index.ID + '">  <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">';
                    html += '<path d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5M11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47M8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5"/></svg></a></td>';
                    html += '</tr>';

                  
                     
                    // console.log(html)

                });
                // console.log(html)
                if (html.length > 0) {
                    jQuery("#userlist tbody").empty()
                    jQuery("#userlist tbody").append(html)

                }
            }

        });
})

function Print(){
    window.print();

}