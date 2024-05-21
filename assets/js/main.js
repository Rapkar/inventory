
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
    var ID = jQuery("#ProductIs").val();
    var InventoryNumber= jQuery("#InventoryIS").val();
    // var Name=jQuery("#ProductIs").html()
    var Count = jQuery("#ProductBox input[name='Count']").val()
    var MeterPrice = jQuery("#ProductBox input[name='Meter']").val()
    var RolePrice = jQuery("#ProductBox input[name='RolePrice']").val()
    var MeterPrice = jQuery("#ProductBox input[name='MeterPrice']").val()
    var TotalPrice = jQuery("#ProductBox input[name='TotalPrice']").val()
    var value = '<tr><th scope="row">' + ID + '</th><td>' + CurrentProductName + '</td><td>23423</td><td>' + Count + '</td><td>' + MeterPrice + '</td><td>' + RolePrice + '</td><td>' + TotalPrice + '</td></tr>';
    var newRow = {
        InventoryNumber:InventoryNumber,
        ProductId: ID,
        // ExportID: ID,
        Name: CurrentProductName,
        count: Count,
        meterPrice: MeterPrice,
        rolePrice: RolePrice,
        totalPrice: TotalPrice
    };
    ProductsOfExport.push(newRow)
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
    jQuery("input[name='Count']").val(0);
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
        data: JSON.stringify({ Name: "expotform", Content: formValues, Products: ProductsOfExport }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            console.log(msg)
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
jQuery("#find").on("click",function(e){
    e.preventDefault()
    var value=jQuery("#findval").val()
    // console.log(value)
    // value="حسین سلطانیان"
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/export-find",
        data: JSON.stringify({ term: value}),
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
                    html += '<td dir="ltr" class="Edit" style="text-align:right;"><a href="./edituser?user-id=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-pen" viewBox="0 0 16 16">';
                    html += '<path d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44 .854A1.5 1.5 0 0 1 11.5 .796a1.5 1.5 0 0 1 1.998-.001m-.644 .766a.5 .5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a .5 .5 0 0 0 0-.708z"/></svg></a></td>';
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