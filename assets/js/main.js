let CurrentProductName = "";
let ProductsOfExport = [];
let ExportTotalPrice = [];
let Payments = [];
let checks = [];
let ExportID
if (document.getElementById("ExportNumber")) {
    ExportID = document.getElementById("ExportNumber").value
}
let tax = parseFloat(jQuery("input[name='Tax']").val())
let draft = false;

const directpaynumber = "PMT-" + Math.floor(Math.random() * 1000000)
function fetchAndLogPayments() {
    if (Payments.length < 1) {
        return jQuery.ajax({
            method: "POST",
            url: "/Dashboard/getpaymentsbyexportid",
            data: JSON.stringify({ "ExportNumber": ExportID }),
            contentType: "application/json; charset=utf-8",
        }).done(function (msg) {
            if (msg.sucess) {

                // Payments = msg.data;
                console.log("مدل دریافتی", msg)

                msg.data.forEach((element, i) => {
                    if (element.Method != "نقدی") {
                        var checkData = {
                            date: element.CreatedAt,
                            bank: element.Name,
                            serial: element.Number,
                            amount: element.TotalPrice,
                            status: element.Status,
                        };
                        checks.push(checkData);
                    } else {
                        jQuery("#directpay").val(element.TotalPrice)
                    }
                });

            }
            renderChecksTable();
            console.log("مقایر کل", Payments, checks)

        })
    }
}


fetchAndLogPayments().then(function () {

    // اینجا می‌توانید از Payments استفاده کنید
});

jQuery('#myModal').modal('show')
// add  New product to export list

jQuery("#AddProductToExport").on("click", function () {

    var ID = jQuery("#ProductIs").val();
    var InventoryNumber = jQuery("#InventoryIS").val();
    var Meter = jQuery("#ProductBox input[name='Meter']").val() || "0";
    var Weight = jQuery("#ProductBox input[name='Weight']").val() || "0";
    var Barrel = jQuery("#ProductBox input[name='Barrel']").val() || "0";
    var Count = jQuery("#ProductBox input[name='Count']").val() || "0";
    var Rolle = jQuery("#ProductBox input[name='Rolle']").val() || "0";
    var RollePrice = jQuery("#ProductBox input[name='RollePrice']").val() || "0";
    var MeterPrice = jQuery("#ProductBox input[name='MeterPrice']").val() || "0";
    var WeightPrice = jQuery("#ProductBox input[name='WeightPrice']").val() || "0";
    var BarrelPrice = jQuery("#ProductBox input[name='BarrelPrice']").val() || "0";
    var CountPrice = jQuery("#ProductBox input[name='CountPrice']").val() || "0";
    var TotalPrice = jQuery("#ProductBox input[name='TotalPrice']").val() || "0";

    hasproducts = jQuery("#ExportProductsList tbody tr");
    hasproductsValue = [];
    hasproducts.each(function () {
        var number = parseInt(jQuery(this).attr("attr-id"))
        hasproductsValue.push(number);
    });
    // var isDuplicate = ProductsOfExport.some(item =>
    //     item.ProductID === ID || item.ProductID === ExportID
    // );
    var isDuplicate = false;
    var isDuplicate = jQuery("#ExportProductsList tbody tr").toArray().some(function(row) {
        return jQuery(row).attr("attr-id") == ID;
    });
    if (isDuplicate) {
        alert("این محصول قبلاً اضافه شده است!");
        return false; // این کار از اجرای بقیه کد جلوگیری می‌کند
    }
    var edit = `<td dir="ltr" class="Edit" style="text-align:right;">
     <a class="me-3 remove"  href="./deleteExport?ExportId={{.ID}}"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">
       <path d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5M11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47M8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5"/>
     </svg></a>
     </td>`;
    var value = '<tr attr-id="' + ID + '"><td   scope="row">' + ID + '</td><td>' + CurrentProductName + '</td><td class="prn">' + Rolle + '</td><td class="price">' + RollePrice + '</td><td class="price">' + Meter + '</td><td class="price">' + MeterPrice + '</td><td class="price">' + Weight + '</td><td class="price">' + WeightPrice + '</td><td class="price">' + Count + '</td><td class="price">' + CountPrice + '</td><td class="price">' + Barrel + '</td><td class="price">' + BarrelPrice + '</td><td class="itemtotalprice price">' + TotalPrice + '</td>' + edit + '</tr>';
    var newRow = {
        InventoryID: InventoryNumber,
        ProductID: ID,
        ExportID: ExportID,
        Name: CurrentProductName,
        Count: Count,
        Meter: Meter,
        Weight: Weight,
        MeterPrice: MeterPrice,
        RollePrice: RollePrice,
        TotalPrice: TotalPrice

    };
    var NewPrice = {
        ProductId: ID,
        price: TotalPrice
    }
    ExportTotalPrice.push(NewPrice)
    ProductsOfExport.push(newRow)


    // console.log(ExportTotalPrice)
    exporttotal_Price = "0"
    // ExportTotalPrice.forEach(function(e,i){
    //     exporttotal_Price = parseFloat(exporttotal_Price) +  parseFloat(e.price)
    // })
    res = GetExportTotalPrice(ExportTotalPrice);


    jQuery("tfoot td").html(res);
    jQuery("#ExportTotalPrice").val(res);
    jQuery("#ExportProductsList tbody").append(value);
    jQuery(".TotalPriceOut td").html(TotalPrice)
    jQuery(".Notfound").slideUp();
    jQuery(".close").click()
    // برای هر نوع قیمت به صورت جداگانه
    var selectors = {
        "td.TotalPrice": true,
        "td.price": true,
        ".price": true,
        "td.prn": true,
        "td.Tax": true,
        "span.price": true
    };

    jQuery.each(selectors, function (selector) {
        jQuery(selector).each(function () {
            var originalValue = jQuery(this).html().trim();
            var cleanedValue = originalValue.replace(/[^\d.,]/g, ''); // حذف کاراکترهای غیرعددی
            var numberValue = parseFloat(cleanedValue.replace(/,/g, ''));

            if (!isNaN(numberValue)) {
                var formattedValue = PersianTools.addCommas(numberValue.toString());
                var convertedValue = PersianTools.digitsEnToFa(formattedValue);
                jQuery(this).html(convertedValue);
            }
        });
    });
    // }

})
// directpay.addEventListener("click",function(){
// console.log(directpay.value && directpay.value.trim().length > 0)
// })

function GetExportTotalPrice(ExportTotalPrice) {
    exporttotal_Price = 0
    tax = parseFloat(jQuery("input[name='Tax']").val())

    ExportTotalPrice.forEach(function (e, i) {

        exporttotal_Price = parseFloat(exporttotal_Price) + parseFloat(e.price) + tax
    })
    return exporttotal_Price

}
jQuery("input[name='Tax']").on("keyup", function () {
    res = GetExportTotalPrice(ExportTotalPrice);

    res = res - parseFloat(jQuery(this).val);
    if (res > 0) {
        jQuery("tfoot td").html(res);
    }
})
function CalculateItems() {
    jQuery(".Content input[type='number']").each(function (item) {
        if (jQuery(this).val() == "" || jQuery(this).val() == null) {
            jQuery(this).val(0)
        }
    });
    var target = jQuery(this).val()
    var id = jQuery(this).attr("name")
    var price = jQuery("input[name='" + id + "Price']").val()

    var result = (parseFloat(target) * parseFloat(price))
    jQuery("input[name='TotalPrice']").val(result)
    return result
}

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
                    if (msg.result.length > 0 && msg.result != null) {
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
jQuery("#draft").on("change", function () {
    draft = jQuery(this).prop('checked');
})

jQuery(".draft").on("change", function () {
    var Exportid = jQuery(this).attr("Export-id");
    var draftvalue = jQuery(this).prop('checked');

    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/draft",  // Make sure this matches your backend route
        data: JSON.stringify({
            Exportid: Exportid,
            draftvalue: draftvalue
        }),
        contentType: "application/json",
        dataType: "json"
    })
        .done(function (msg) {
            if (msg.error) {
                alert("Error: " + msg.error);
            } else {
                alert(msg.message);
            }
        })
        .fail(function (jqXHR, textStatus, errorThrown) {
            console.error("AJAX Error:", textStatus, errorThrown);
            alert("Request failed: " + textStatus);
        });
});
// Select  Product name for fech the detail of product in Export Page

function ToggleSIBoxInExport(value) {
    jQuery(".Content .inputs .form-group").each(function (index, element) {
        jQuery(element).slideUp();

        if (jQuery(element).attr("attr-si") == value) {
            jQuery(element).slideDown();
            jQuery(element).removeClass("d-none");

        }

    });
    jQuery(".Content .calculator .form-group").each(function (index, element) {
        jQuery(element).slideUp();

        if (jQuery(element).attr("attr-si") == value) {
            jQuery(element).slideDown();
            jQuery(element).removeClass("d-none");

        }

    });

}


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
                    ToggleSIBoxInExport(product.MeasurementSystem)
                    jQuery(".ExportPeoducts .ProductsRolle").html(product.Roll)
                    jQuery(".ExportPeoducts input[name='Count']").attr("max", product.Count)
                    jQuery(".ExportPeoducts .ProductsMeter").html(product.Meter)
                    jQuery(".ExportPeoducts input[name='Meter']").attr("max", product.Meter)
                    jQuery(".ExportPeoducts input[name='Weight']").attr("max", product.Weight)
                    jQuery(".ExportPeoducts .ProductNumber").html(product.Number)
                    jQuery(".ExportPeoducts .ProductsWeight").html(product.Weight)
                    jQuery(".ExportPeoducts .ProductsBarrel").html(product.Barrel)
                    jQuery(".ExportPeoducts input[name='RollePrice']").attr("value", product.RollePrice)
                    jQuery(".ExportPeoducts input[name='MeterPrice']").attr("value", product.MeterPrice)
                    jQuery(".ExportPeoducts input[name='BarrelPrice']").attr("value", product.BarrelPrice)
                    // jQuery(".ExportPeoducts input[name='TotalPrice']").attr("value", product.MeterPrice)
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
jQuery("span.price").each(function () {
    var priceText = jQuery(this).text().trim(); // "۱,۱۹۷,۹۶۰"

    // تبدیل اعداد فارسی به انگلیسی
    var englishDigits = priceText.replace(/[۰-۹]/g, function (d) {
        return "۰۱۲۳۴۵۶۷۸۹".indexOf(d);
    });

    // حذف کاما و تبدیل به عدد
    var priceNumber = parseInt(englishDigits.replace(/,/g, ''));

    // اگر عدد معتبر بود، تبدیلش کن به حروف
    if (!isNaN(priceNumber)) {
        var numberToWords = PersianTools.numberToWords(priceNumber);

        // console.log(numberToWords);
        jQuery(".wordprice ").html(numberToWords)
        // jQuery(this).after(" (" + numberToWords + ")");
    } else {
        console.log("❌ عدد نامعتبر:", priceText);
    }
});


// Select  Product name for fech the detail of product in Production Page
jQuery(".production select#ProductIs").on("change", function () {
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
                    jQuery(".production span.ProductsCount").html(product.Count)
                    jQuery(".production span.ProductMeter").html(product.Meter)
                    jQuery(".production input[name='ProductsCount']").attr("value", product.Count)
                    jQuery(".production input[name='ProductMeter']").attr("value", product.Meter)
                    jQuery(".production input[name='RolePrice']").attr("value", product.RolePrice)
                    jQuery(".production input[name='RolePrice']").attr("value", product.RolePrice)
                    jQuery(".production input[name='MeterPrice']").attr("value", product.MeterPrice)
                    jQuery(".production  input[type='number']").each(function () {
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



    RolePrice = parseFloat(jQuery(".ExportPeoducts input[name='RolePrice']").val());
    Count = parseFloat(jQuery(this).val());

    if (isNaN(RolePrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = RolePrice * Count;
        // ExportTotalPrice =parseFloat(ExportTotalPrice) + parseFloat(TotalPrice)
    }

    jQuery("input[name='TotalPrice']").val(CalculateItems())
    jQuery("#ProductBox input[type='number']").each(function () {
        var val = jQuery(this).val();
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);
        var numberToWords = PersianTools.numberToWords(val);
        jQuery(this).parent().closest(".form-group").find(".out").html(convertToFa + "   " + numberToWords);
    });


})

// Calculate TotalPrice by price of meter  

jQuery(".Content .inputs input").on("keyup", function () {
    var number = parseInt(this.value)
    var id = jQuery(this).attr("name");
    if (number != 0) {
        jQuery(".modal-footer").slideDown();

    } else {
        jQuery(".modal-footer").slideUp();
    }
    // jQuery("input[name='Count']").val(0);
    targetPrice = parseFloat(jQuery(".ExportPeoducts .calculator input[name='" + id + "Price']").val());
    Count = parseFloat(jQuery(this).val());
    var TotalPrice = 0
    if (isNaN(targetPrice) || isNaN(Count)) {
        console.log("Invalid number");
    } else {
        var TotalPrice = targetPrice * Count;
    }


    jQuery("input[name='TotalPrice']").val(TotalPrice)
    jQuery("#ProductBox input[type='number']").each(function () {
        var val = jQuery(this).val();
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);
        var numberToWords = PersianTools.numberToWords(val);
        jQuery(this).parent().closest(".form-group").find(".out").html(convertToFa + "   " + numberToWords);
    });

})

function calculateTotalPayments(payments) {
    let total = 0;
    const directpay = document.getElementById("directpay");
    if (directpay) {
        if (directpay.value && directpay.value.trim().length > 0) {

            var Payment = {
                Method: "نقدی",
                Name: "نقدی",
                Status: "collected",
                TotalPrice: parseFloat(directpay.value),
                Number: directpaynumber,
                CreatedAt: document.getElementById("checkDate").value
            };

            var existingPaymentIndex = Payments.findIndex(p => p.Number === directpaynumber);

            if (existingPaymentIndex !== -1) {
                Payments[existingPaymentIndex] = Payment;
            } else {
                Payments.push(Payment);
            }
        }
    }

    for (let payment of payments) {
        total += parseFloat(payment.TotalPrice);

    }
    var val = PersianTools.addCommas(total);
    console.log("ttttt", val)

    jQuery("#TotalPayments").html(PersianTools.digitsEnToFa(val))
    return total;
}
jQuery("#directpay").on("keyup", function () {
    calculateTotalPayments(Payments)
})
jQuery("form[name='expotform']").submit(function (e) {
    e.preventDefault();
    var formValues = jQuery("form[name='expotform']").find("input, select, textarea").map(function () {
        var $this = $(this);

        // اگر عنصر checkbox بود
        if ($this.attr('type') === 'checkbox') {
            return $this.attr('name') + "=" + $this.prop('checked');
        }
        // برای سایر عناصر
        return $this.attr("name") + "=" + $this.val();
    }).get().join("&");
    ExportPrice = GetExportTotalPrice(ExportTotalPrice);

    if (Array.isArray(ProductsOfExport) && ProductsOfExport.length > 0) {
        calculateTotalPayments(Payments)
        if (parseFloat(calculateTotalPayments(Payments)) < parseFloat(ExportPrice)) {
            const isConfirmed = confirm("مبلغ پرداختی کمتر از قیمت فاکتور میباشد آیا از ادامه مطمین هستید؟");
            if (!isConfirmed) {
                return;
            }
        }

        jQuery.ajax({
            method: "POST",
            url: "/Dashboard/export",
            data: JSON.stringify({
                Name: "expotform",
                TotalPrice: ExportPrice,
                Content: formValues,
                Products: ProductsOfExport,
                Payments: Payments
            }),
            contentType: "application/json; charset=utf-8",
        })
            .done(function (msg) {
                if (msg.message == "sucess") {
                    window.location.replace("./exportshow?ExportId=" + msg.id);
                }
            })
            .error(function (msg) {
                alert("خطا در ارسال اطلاعات");
            });
    } else {
        alert("لطفا محصولی را اضافه کنید !")
    }

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
                    html += '<a href="./exportshow?ExportId=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-eye" viewBox="0 0 16 16">';
                    html += '<path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8M1.173 8a13 13 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5s3.879 1.168 5.168 2.457A13 13 0 0 1 14.828 8q-.086.13-.195.288c-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5s-3.879-1.168-5.168-2.457A13 13 0 0 1 1.172 8z"/><path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5M4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0"/></svg></a>';
                    html += '<a class="me-3" href="./deleteExport?ExportId=' + index.ID + '"><svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-trash3" viewBox="0 0 16 16">';
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
jQuery("#findpayment").on("click", function (e) {
    e.preventDefault()
    var value = jQuery("#findval").val()
    // console.log(value)
    // value="حسین سلطانیان"
    jQuery.ajax({
        method: "POST",
        url: "/Dashboard/payment-find",
        data: JSON.stringify({ term: value }),
        // data: { Name: "expotform", Content: jQuery("form[name='expotform']").serialize(), Products: ProductsOfExport },
        // contentType: "application/json; charset=utf-8",
    })
        .done(function (msg) {
            var lengthofres = msg.message.length;
            console.log(msg.message)
            if (lengthofres > 0) {
                let html = "";
                msg.message.forEach(function (index) {
                    html += '<tr>';
                    html += '<td class="' + index.ID + '" style="text-align:right;">' + index.ID + '</td>';
                    html += '<td class="' + index.Method + '" style="text-align:right;">' + index.Method + '</td>';
                    html += '<td class="' + index.Number + '" style="text-align:right;">' + index.Number + '</td>';
                    html += '<td class="' + index.Name + '" style="text-align:right;">' + index.Name + '</td>';
                    html += '<td class="' + index.TotalPrice + '" style="text-align:right;">' + index.TotalPrice + '</td>';
                    html += '<td class="' + index.UserName + '" style="text-align:right;">' + index.UserName + '</td>';
                    html += '<td class="' + index.CreatedAt + '" style="text-align:right;">' + index.CreatedAt + '</td>';
                    html += '<td class="' + index.export_number + '" style="text-align:right;">' + index.export_number + '</td>';
                    html += '<td class="' + index.ExportID + '" style="text-align:right;"><a class="me-3" href="./exportshow?ExportId=' + index.ExportID + '">';
                    html += '<svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor" class="bi bi-eye" viewBox="0 0 16 16">';
                    html += '<path d="M16 8s-3-5.5-8-5.5S0 8 0 8s3 5.5 8 5.5S16 8 16 8M1.173 8a13 13 0 0 1 1.66-2.043C4.12 4.668 5.88 3.5 8 3.5s3.879 1.168 5.168 2.457A13 13 0 0 1 14.828 8q-.086.13-.195.288c-.335.48-.83 1.12-1.465 1.755C11.879 11.332 10.119 12.5 8 12.5s-3.879-1.168-5.168-2.457A13 13 0 0 1 1.172 8z" />';
                    html += '<path d="M8 5.5a2.5 2.5 0 1 0 0 5 2.5 2.5 0 0 0 0-5M4.5 8a3.5 3.5 0 1 1 7 0 3.5 3.5 0 0 1-7 0" />';
                    html += '</svg></a></td>';

                    if (index.Status == "collected") {
                        html += '<td dir="" class="InventoryNumber d-none d-md-table-cell bg-success text-center">';
                        html += '<img src="../../assets/images/collected.svg">';
                        html += '</td>';
                    }
                    if (index.Status == "rejected") {
                        html += '<td dir="" class="InventoryNumber d-none d-md-table-cell bg-danger text-center">';
                        html += '<img src="../assets/images/angry.svg">';
                        html += '</td>';
                    }
                    if (index.Status == "pending") {
                        html += '<td dir="" class="InventoryNumber d-none d-md-table-cell bg-warning text-center">';
                        html += '<img src="../assets/images/waite.svg">';
                        html += '</td>';
                    }
                    html += '<td dir="ltr" class="Edit" style="text-align:right;">';
                    html += `<a href="#" data-bs-toggle="modal"
                    data-bs-target="#editModal${index.ID}">
                    <svg xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                        class="bi bi-pen" viewBox="0 0 16 16">
                        <path
                            d="m13.498.795.149-.149a1.207 1.207 0 1 1 1.707 1.708l-.149.148a1.5 1.5 0 0 1-.059 2.059L4.854 14.854a.5.5 0 0 1-.233.131l-4 1a.5.5 0 0 1-.606-.606l1-4a.5.5 0 0 1 .131-.232l9.642-9.642a.5.5 0 0 0-.642.056L6.854 4.854a.5.5 0 1 1-.708-.708L9.44.854A1.5 1.5 0 0 1 11.5.796a1.5 1.5 0 0 1 1.998-.001m-.644.766a.5.5 0 0 0-.707 0L1.95 11.756l-.764 3.057 3.057-.764L14.44 3.854a.5.5 0 0 0 0-.708z" />
                    </svg>
                    </a>

                    <a class="me-3" href="./deletePayments?PaymentId={{.ID}}"><svg
                        xmlns="http://www.w3.org/2000/svg" width="16" height="16" fill="currentColor"
                        class="bi bi-trash3" viewBox="0 0 16 16">
                        <path
                            d="M6.5 1h3a.5.5 0 0 1 .5.5v1H6v-1a.5.5 0 0 1 .5-.5M11 2.5v-1A1.5 1.5 0 0 0 9.5 0h-3A1.5 1.5 0 0 0 5 1.5v1H1.5a.5.5 0 0 0 0 1h.538l.853 10.66A2 2 0 0 0 4.885 16h6.23a2 2 0 0 0 1.994-1.84l.853-10.66h.538a.5.5 0 0 0 0-1zm1.958 1-.846 10.58a1 1 0 0 1-.997.92h-6.23a1 1 0 0 1-.997-.92L3.042 3.5zm-7.487 1a.5.5 0 0 1 .528.47l.5 8.5a.5.5 0 0 1-.998.06L5 5.03a.5.5 0 0 1 .47-.53Zm5.058 0a.5.5 0 0 1 .47.53l-.5 8.5a.5.5 0 1 1-.998-.06l.5-8.5a.5.5 0 0 1 .528-.47M8 4.5a.5.5 0 0 1 .5.5v8.5a.5.5 0 0 1-1 0V5a.5.5 0 0 1 .5-.5" />
                    </svg></a>`;

                    html += `
                    <div class="modal fade" id="editModal${index.ID}" tabindex="-1"
                    aria-labelledby="editModalLabel${index.ID}" aria-hidden="true">
                    <div class="modal-dialog modal-xl">
                        <div class="modal-content">
                            <div class="modal-header">
                                <h5 class="modal-title" id="editModalLabel${index.ID}">Edit Export #${index.ID}</h5>
                                <button type="button" class="btn-close" data-bs-dismiss="modal"
                                    aria-label="Close"></button>
                            </div>
                            <div class="modal-body">
                                <!-- Your edit form goes here -->
                                <form action="./updatepayment" method="POST">
                                    <input type="hidden" name="PaymentID" value="${index.ID}">

                                    <div class="row">
                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportName${index.ID}" class="form-label text-dark">نام
                                                بانک</label>
                                            <input type="text" class="form-control" id="exportName${index.ID}"
                                                name="PaymentName" value="${index.Name}">
                                        </div>

                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportTotalPrice${index.ID}"
                                                class="form-label text-dark">مبلغ
                                            </label>
                                            <input type="number" class="form-control"
                                                id="exportTotalPrice${index.ID}" name="PaymentTotalPrice"
                                                value="${index.TotalPrice}">
                                        </div>
                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportName${index.ID}" class="form-label text-dark">روش
                                                پرداخت</label>
                                            <input type="text" readonly class="form-control"
                                                id="exportMethod${index.ID}" name="Method" value="${index.Method}">
                                        </div>

                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportTotalPrice${index.ID}"
                                                class="form-label text-dark">شماره سریال
                                            </label>
                                            <input type="number" class="form-control"
                                                id="exportTotalPrice${index.ID}" name="PaymentNumber"
                                                value="${index.Number}">
                                        </div>
                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportName${index.ID}"
                                                class="form-label text-dark">تاریخ
                                            </label>
                                            <input type="text" class="form-control" id="CreatedAt${index.ID}"
                                                name="CreatedAt" value="${index.CreatedAt}">
                                        </div>

                                        <div class="mb-3 col-lg-6 ">
                                            <label for="exportTotalPrice${index.ID}"
                                                class="form-label text-dark">وضعیت
                                            </label>
                                            <select class="form-select status-select" name="PaymentStatus"
                                                data-index="${index}">
                                                <option value="pending">در انتظار</option>
                                                <option value="collected">وصول شده</option>
                                                <option value="rejected">برگشت خورده</option>
                                            </select>
                                        </div>
                                    </div>



                                    <!-- Add more fields as needed -->
                            </div>
                            <div class="modal-footer">
                                <button type="button" class="btn btn-secondary"
                                    data-bs-dismiss="modal">بستن</button>
                                <button type="submit" class="btn btn-primary"
                                    onclick="document.forms[0].submit()">ذخیره</button>
                            </div>
                            </form>

                        </div>
                    </div>
                </div>`;

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


jQuery("input[name='Tax']").on("keyup", function (e) {
    oldval = GetExportTotalPrice(ExportTotalPrice);
    newval = parseFloat(this.value)
    res = oldval + newval


    jQuery("tfoot td").html(res);
    jQuery("td.TotalPrice,td.price,.price,td.prn,td.Tax").each(function () {
        var val = jQuery(this).html();
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);

        jQuery(this).html(convertToFa);
    });
})
jQuery(document).on('click', '.remove', function (e) {
    e.preventDefault()
    var id = jQuery(this).parent().closest("tr").find(".id").html();
    id = parseInt(id)
    RemoveItem(this, id);
})




function RemoveItem(target, id) {

    ExportTotalPrice.forEach((element, index) => {
        console.log(index == id, element.ProductId, id)
        if (parseInt(element.ProductId) == id) {
            ExportTotalPrice.splice(index, 1);
        }
    });
    ProductsOfExport.forEach((element, index) => {
        if (parseInt(element.ProductId) == id) {
            ProductsOfExport.splice(index, 1);
        }
    });
    res = GetExportTotalPrice(ExportTotalPrice);

    jQuery("tfoot td").html(res);
    jQuery(target).parent().closest("tr").remove()
    jQuery("td.TotalPrice,td.price,.price,td.prn,td.Tax").each(function () {
        var val = jQuery(this).html();
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);

        jQuery(this).html(convertToFa);
    });
}
function Print() {
    window.print();

}
(function ($) {
    "use strict";

    // Navbar on scrolling
    $(window).scroll(function () {
        if ($(this).scrollTop() > 200) {
            $('.home .navbar').fadeIn('slow').css('display', 'flex');
        } else {
            $('.home  .navbar').fadeOut('slow').css('display', 'none');
        }
    });


    // Smooth scrolling on the navbar links
    $(".navbar-nav a, .btn-scroll").on('click', function (event) {
        if (this.hash !== "") {
            event.preventDefault();

            $('html, body').animate({
                scrollTop: $(this.hash).offset().top - 45
            }, 1500, 'easeInOutExpo');

            if ($(this).parents('.navbar-nav').length) {
                $('.navbar-nav .active').removeClass('active');
                $(this).closest('a').addClass('active');
            }
        }
    });


    // Scroll to Bottom
    $(window).scroll(function () {
        if ($(this).scrollTop() > 100) {
            $('.scroll-to-bottom').fadeOut('slow');
        } else {
            $('.scroll-to-bottom').fadeIn('slow');
        }
    });


    // Portfolio isotope and filter
    if ($('.portfolio-container').length > 0) {
        var portfolioIsotope = $('.portfolio-container').isotope({
            itemSelector: '.portfolio-item',
            layoutMode: 'fitRows'
        });
        $('#portfolio-flters li').on('click', function () {
            $("#portfolio-flters li").removeClass('active');
            $(this).addClass('active');

            portfolioIsotope.isotope({ filter: $(this).data('filter') });
        });
    }


    // Back to top button
    $(window).scroll(function () {
        if ($(this).scrollTop() > 200) {
            $('.back-to-top').fadeIn('slow');
        } else {
            $('.back-to-top').fadeOut('slow');
        }
    });
    $('.back-to-top').click(function () {
        $('html, body').animate({ scrollTop: 0 }, 1500, 'easeInOutExpo');
        return false;
    });


    // Gallery carousel
    if ($(".gallery-carousel").length > 0) {
        $(".gallery-carousel").owlCarousel({
            autoplay: false,
            smartSpeed: 1500,
            dots: false,
            loop: true,
            nav: true,
            navText: [
                '<i class="fa fa-angle-left" aria-hidden="true"></i>',
                '<i class="fa fa-angle-right" aria-hidden="true"></i>'
            ],
            responsive: {
                0: {
                    items: 1
                },
                576: {
                    items: 2
                },
                768: {
                    items: 3
                },
                992: {
                    items: 4
                },
                1200: {
                    items: 5
                }
            }
        });

    }
    // Testimonials carousel
    if ($(".testimonial-carousel").length > 0) {
        $(".testimonial-carousel").owlCarousel({
            autoplay: true,
            smartSpeed: 1500,
            items: 1,
            dots: false,
            loop: true,
            nav: true,
            navText: [
                '<i class="fa fa-angle-left" aria-hidden="true"></i>',
                '<i class="fa fa-angle-right" aria-hidden="true"></i>'
            ],
        });
    }
})(jQuery);


fetch("/Dashboard/api/allexports")
    .then(response => {
        if (!response.ok) {
            throw new Error('Network response was not ok');
        }
        return response.json(); // or response.text() depending on what the API returns
    })
    .then(data => {
        data.sort((a, b) => new Date(a.CreatedAt) - new Date(b.CreatedAt));
        var labels = [];
        var date = [];
        data.forEach(function (item, index) {
            labels.push(item.CreatedAt);
            date.push(item.TotalPrice);

        })
        const ctx = document.getElementById('myChart');
        if (ctx) {
            new Chart(ctx, {
                type: 'line',  // Changed from 'bar' to 'line'
                data: {
                    labels: labels,
                    datasets: [{
                        label: '# نمودار فروش',
                        data: date,
                        borderWidth: 2,  // Slightly increased for better visibility in line chart
                        backgroundColor: 'rgba(75, 192, 192, 0.2)',  // Added fill color
                        borderColor: 'rgba(75, 192, 192, 1)',  // Added line color
                        tension: 0.1  // Makes the line slightly curved
                    }]
                },
                options: {
                    scales: {
                        y: {
                            beginAtZero: true
                        }
                    }
                }

            });
        }
    })
    .catch(error => {
        console.error('There was a problem with the fetch operation:', error);
    });
const checkForm = document.getElementById('addcheck');
const checksTableBody = document.getElementById('checksTableBody');
let editIndex = null;
// let Payments = []; // Array to store payment data
//     btncheckbox.addEventListener('click', function (e) {
//         var ExportNumber = jQuery("input[name='ExportNumber']").val();
// console.log("payments",Payments,"getPaymentsByExportId",)

//     })
// Load checks when page loads
renderChecksTable();

// Form submission
if (checkForm) {
    checkForm.addEventListener('click', function (e) {
        e.preventDefault();



        // validation

        let errorBox = document.getElementById('form-error');
        if (!errorBox) {
            errorBox = document.createElement('div');
            errorBox.id = 'form-error';
            errorBox.style.color = 'red';
            errorBox.style.margin = '10px 0';
            checkForm.parentNode.insertBefore(errorBox, checkForm.nextSibling);
        }

        const fields = [
            { id: 'checkDate', name: 'تاریخ چک' },
            { id: 'bankName', name: 'نام بانک' },
            { id: 'serialCode', name: 'شماره سریال' },
            { id: 'checkAmount', name: 'مبلغ چک' }
        ];

        let isValid = true;
        errorBox.innerHTML = ''; // پاک کردن خطاهای قبلی

        fields.forEach(field => {
            const input = document.getElementById(field.id);
            const value = input.value.trim();

            if (!value) {
                isValid = false;
                errorBox.innerHTML += `فیلد ${field.name} الزامی است.<br>`;
                input.classList.add('error-field');
            } else {
                input.classList.remove('error-field');
            }
        });

        if (!isValid) {
            // اسکرول به قسمت خطاها
            errorBox.scrollIntoView({ behavior: 'smooth', block: 'center' });
            return;
        }

        // validation
        const checkData = {
            date: document.getElementById('checkDate').value,
            bank: document.getElementById('bankName').value,
            serial: document.getElementById('serialCode').value,
            amount: document.getElementById('checkAmount').value,
            status: 'pending' // Default status
        };

        if (editIndex !== null) {
            // Update existing check
            checks[editIndex] = checkData;
            editIndex = null;
        } else {
            // Add new check
            checks.push(checkData);
        }


        // Refresh table
        renderChecksTable();

        // Reset form
        // checkForm.reset();
        document.getElementById('addcheck').textContent = 'ذخیره چک';
    });
}

function ToFaprice(val) {
    var convertToFa = 0
    if (parseFloat(val)) {
        var val = PersianTools.addCommas(val);
        var convertToFa = PersianTools.digitsEnToFa(val);
    }
    return convertToFa
}
function ToFaDigit(val) {
    var convertToFa = 0
    if (parseFloat(val)) {
        var convertToFa = PersianTools.digitsEnToFa(val);
    }
    return convertToFa
}
function renderChecksTable() {
    if (checksTableBody && Payments.length != 0) {
        checksTableBody.innerHTML = '';
    }
    checks.forEach((check, index) => {
        const row = document.createElement('tr');

        row.innerHTML = `
                    <td>${index + 1}</td>
                    <td>${check.date}</td>
                    <td>${getBankName(check.bank)}</td>
                    <td data-code="${check.serial}" class="serial price serial">${ToFaDigit(check.serial)}</td>
                    <td  class="price" >${ToFaprice(Number(check.amount).toLocaleString())}</td>
                    <td>
                        <select class="form-select status-select" data-index="${index}">
                            <option value="pending" ${check.status === 'pending' ? 'selected' : ''}>در انتظار</option>
                            <option value="collected" ${check.status === 'collected' ? 'selected' : ''}>وصول شده</option>
                            <option value="rejected" ${check.status === 'rejected' ? 'selected' : ''}>برگشت خورده</option>
                        </select>
                    </td>
                    <td>
                        <button class="btn btn-sm btn-outline-danger delete-btn" data-index="${index}">حذف</button>
                    </td>
                `;

        // Create Payment object for each check
        var Payment = {
            Method: "چک",
            Name: getBankName(check.bank),
            ExportID: jQuery("input[name='ExportNumber']").val(),
            Status: check.status,
            TotalPrice: check.amount,
            Number: check.serial,
            CreatedAt: check.date
        };
        var isDuplicate = Payments.some(item =>
            item.Number === check.serial
        );
        if (!isDuplicate) {
            Payments.push(Payment);
        }

        if (checksTableBody) {
            checksTableBody.appendChild(row);
        }
    });


    // add check
    document.querySelectorAll('.edit-btn').forEach(btn => {
        btn.addEventListener('click', function () {
            editIndex = parseInt(this.getAttribute('data-index'));
            const check = checks[editIndex];

            document.getElementById('checkDate').value = check.date;
            document.getElementById('bankName').value = check.bank;
            document.getElementById('serialCode').value = check.serial;
            document.getElementById('checkAmount').value = check.amount;

            document.querySelector('#checkForm button[type="submit"]').textContent = 'ویرایش چک';
        });
    });

    document.querySelectorAll('.delete-btn').forEach(btn => {
        btn.addEventListener('click', function (e) {
            e.preventDefault()
            // دریافت مقدار شماره چک از ستون مربوطه
            var numberElement = jQuery(this).closest("tr").find(".serial").attr("data-code");
            var checkNumber = numberElement.trim(); // یا .val() اگر input باشد

            if (confirm('آیا از حذف این چک مطمئن هستید؟')) {
                const index = parseInt(this.getAttribute('data-index'));
                checks.splice(index, 1);
                console.log("Payments1", Payments)
                // حذف از آرایه Payments با مقایسه صحیح
                Payments.forEach((element, i) => {
                    if (element.Number.toString() === checkNumber) {
                        Payments.splice(i, 1);
                        return; // برای توقف حلقه بعد از حذف
                    }
                });
                console.log("Payments2", Payments)
                renderChecksTable();
            }
        });
    });

    document.querySelectorAll('.status-select').forEach(select => {
        select.addEventListener('change', function () {
            const index = parseInt(this.getAttribute('data-index'));
            checks[index].status = this.value;
            renderChecksTable(); // Refresh to update Payments array
        });
    });
    calculateTotalPayments(Payments)
}

// Bank Names
function getBankName(bankCode) {
    const banks = {
        'melli': 'ملی',
        'mellat': 'ملت',
        'saderat': 'صادرات',
        'tejarat': 'تجارت',
        'saman': 'سامان',
        'shahr': 'شهر',
        'pasargad': 'پاسارگاد',
        'sepah': 'سپه',
        'keshavarzi': 'کشاورزی',
        'parsian': 'پارسیان',
        'eghtesad-novin': 'اقتصاد نوین',
        'ansar': 'انصار',
        'karafarin': 'کارآفرین',
        'sina': 'سینا',
        'sarmayeh': 'سرمایه',
        'tosee': 'توسعه',
        'tosee-saderat': 'توسعه صادرات',
        'tosee-taavon': 'توسعه تعاون',
        'day': 'دی',
        'hekmat': 'حکمت ایرانیان',
        'ayandeh': 'آینده',
        'ghavamin': 'قوامین',
        'khavar': 'خاورمیانه',
        'mehr-iran': 'مهر ایران',
        'mehr-eqtesad': 'مهر اقتصاد',
        'post': 'پست بانک',
        'qarzolqasaneh': 'قرض‌الحسنه مهر ایران',
        'qarzolqasaneh-resalat': 'قرض‌الحسنه رسالت',
        'iran-zamin': 'ایران زمین',
        'kosar': 'کوثر',
        'markazi': 'مرکزی',
        'reffah': 'رفاه',
        'tourism': 'گردشگری',
        'industry': 'صنعت و معدن'
    };
    return banks[bankCode] || bankCode;
}

function debounce(func, timeout = 500) {
    let timer;
    return (...args) => {
        clearTimeout(timer);
        timer = setTimeout(() => { func.apply(this, args); }, timeout);
    };
}

function sendSearchRequest() {
    var nameValue = jQuery("input[name='Name']").val().trim();
    var phoneValue = jQuery("input[name='Phonenumber']").val().trim();
    var nameTerm = "";
    var PhoneTerm = "";

    var isNameValid = nameValue.length > 2;
    var isPhoneValid = phoneValue.length > 10;

    if (isNameValid || isPhoneValid) {
        if (isNameValid) {
            nameTerm = nameValue;
        }
        if (isPhoneValid) {
            PhoneTerm = phoneValue;
        }

        jQuery.ajax({
            method: "POST",
            url: "/Dashboard/users-find",
            contentType: "application/json; charset=utf-8",
            data: JSON.stringify({ name: nameTerm, phone: PhoneTerm })
        })
            .done(function (msg) {
                var messageElement = jQuery("small[for='Phonenumber']");
                messageElement.removeClass('text-danger text-muted text-success');

                if (msg.message.includes("ثبت شده")) {
                    messageElement.html(msg.message).addClass('text-danger');
                    jQuery("button[name='inquiry']").prop("disabled", false);
                    msg.users.forEach(function (item) {
                        html += `<tr>
                        <td>${item.ID}</td>
                        <td>${item.Name}</td>
                        <td>${item.Phonenumber}</td>
                        <td><a target="_blank" href="/Dashboard/user/details?user-id=${item.ID}">مشاهده جزییات</a></td>
                    </tr>`;
                    });
                    jQuery("#inquirybox").html(html);

                } else if (Array.isArray(msg.users) && msg.users.length > 0) {
                    messageElement.html(msg.message).addClass('text-success');
                    jQuery("button[name='inquiry']").prop("disabled", false);

                    var html = "";
                    msg.users.forEach(function (item) {
                        html += `<tr>
                        <td>${item.ID}</td>
                        <td>${item.Name}</td>
                        <td>${item.Phonenumber}</td>
                        <td><a target="_blank" href="/Dashboard/user/details?user-id=${item.ID}">مشاهده جزییات</a></td>
                    </tr>`;
                    });
                    jQuery("#inquirybox").html(html);
                } else {
                    messageElement.html(msg.message).addClass('text-muted');
                    jQuery("button[name='inquiry']").prop("disabled", true);
                    jQuery("#inquirybox").html("");
                }
            })
            .fail(function (jqXHR, textStatus, errorThrown) {
                console.error("AJAX request failed: " + textStatus, errorThrown);
                jQuery("button[name='inquiry']").prop("disabled", true);
                jQuery("small[for='Phonenumber']").html('خطا در ارتباط با سرور').addClass('text-danger');
            });
    } else {
        jQuery("button[name='inquiry']").prop("disabled", true);
        jQuery("small[for='Phonenumber']").html('مثال: 09123456789').addClass('text-muted');
    }
}

jQuery("input[name='Name'], input[name='Phonenumber']").on("keyup", debounce(sendSearchRequest, 500));

jQuery(document).ready(function () {
    jQuery("#checkDate").pDatepicker({
        format: 'YYYY/MM/DD',
        autoClose: true
    });
});

function convertToEnglishNumbers(input) {
    const persianNumbers = [/۰/g, /۱/g, /۲/g, /۳/g, /۴/g, /۵/g, /۶/g, /۷/g, /۸/g, /۹/g];
    const arabicNumbers = [/٠/g, /١/g, /٢/g, /٣/g, /٤/g, /٥/g, /٦/g, /٧/g, /٨/g, /٩/g];
    if (typeof input === 'string' || typeof input === 'number') {
        for (let i = 0; i < 10; i++) {
            input = input.replace(persianNumbers[i], i).replace(arabicNumbers[i], i);
        }
    }
    return input;
}

function forceEnglishNumbers(e) {
    const input = e.target;
    let caretPos = input.selectionStart;
    let convertedValue = convertToEnglishNumbers(input.value);

    // بررسی آیا مقدار شامل حروف غیر عددی است
    const hasNonNumeric = /[^0-9]/.test(convertedValue);

    if (hasNonNumeric) {
        alert("لطفاً فقط عدد وارد کنید!");
        input.value = convertedValue.replace(/[^0-9]/g, ''); // حذف تمام غیر اعداد
        caretPos = Math.max(0, caretPos - 1); // تنظیم موقعیت کارت
        input.setSelectionRange(caretPos, caretPos);
        return;
    }

    if (convertedValue !== input.value) {
        input.value = convertedValue;
        input.setSelectionRange(caretPos, caretPos);
    }
}

document.addEventListener('DOMContentLoaded', function () {
    const numberInputs = document.querySelectorAll('input.price, input[name="Phonenumber"]');

    numberInputs.forEach(input => {
        input.addEventListener('input', forceEnglishNumbers);
        input.addEventListener('keyup', forceEnglishNumbers);
        input.addEventListener('paste', function (e) {
            setTimeout(() => {
                forceEnglishNumbers(e);
            }, 0);
        });
    });
});

document.addEventListener('DOMContentLoaded', function () {
    const measurementSystem = document.querySelector('select[name="MeasurementSystem"]');

    if (measurementSystem) {
        // For edit mode
        measurementSystem.addEventListener('change', function () {
            const selectedValue = this.value;

            // Hide all sections first
            document.querySelectorAll('[id$="-section"]').forEach(section => {
                section.style.display = 'none';
            });

            // Show selected section
            if (selectedValue) {
                document.getElementById(selectedValue + '-section').style.display = 'block';
            }
        });

        // Trigger change event on page load for edit mode
        if (measurementSystem.value) {
            measurementSystem.dispatchEvent(new Event('change'));
        }
    } else {
        // For add mode
        const addModeSelect = document.getElementById('measurement-system');
        if (addModeSelect) {
            addModeSelect.addEventListener('change', function () {
                const selectedValue = this.value;

                // Hide all sections first
                document.querySelectorAll('[id$="-section"]').forEach(section => {
                    section.style.display = 'none';
                });

                // Show selected section
                if (selectedValue) {
                    document.getElementById(selectedValue + '-section').style.display = 'block';
                }
            });
        }
    }
});

// FechPaymentS

/**
 * دریافت لیست پرداخت‌ها بر اساس ExportID
 * @param {string} ExportID - شناسه صادراتی مورد نظر
 * @param {string} token - توکن احراز هویت
 * @returns {Promise<Object>} - نتیجه درخواست
 */
// async function getPaymentsByExportId() {
//     try {
//         const apiUrl = `/Dashboard/getpaymentsbyexportid?ExportNumber=${encodeURIComponent(ExportID)}`;

//         const response = await fetch(apiUrl, {
//             method: 'GET',
//             headers: {
//                 'Content-Type': 'application/json',
//                 'Accept-Language': 'fa-IR' // برای پیام‌های فارسی
//             }
//         });

//         if (!response.ok) {
//             const errorData = await response.json();
//             throw new Error(errorData.error || 'خطا در دریافت اطلاعات');
//         }

//         const result = await response.json();

//         if (result.message && result.data) {
//             return result.data; // بازگرداندن داده‌های پرداخت
//         } else {
//             throw new Error('فرمت پاسخ سرور نامعتبر است');
//         }
//     } catch (error) {
//         console.error('❌ خطا در دریافت پرداخت‌ها:', error.message);
//         throw error; // پرتاب مجدد خطا برای مدیریت توسط caller
//     }
// }
