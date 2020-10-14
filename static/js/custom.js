'use strict';

(function ($) {
    showFiveItems("brand");
    showFiveItems("size");
    showFiveItems("color");

    function showFiveItems(name) {
        var len = $("input[name=" + name + "]:checked").length;
        var counter = 5 - len;

        $.each($("input[name=" + name + "]"), function(k, input) {
            if (counter <= 0) return false;

            if ($(input).prop("checked") == false) {
                $(input).closest("label").addClass("show");
                counter -= 1;
            }
        });
    }

    $(".brand-more").click(function(e) {
        e.preventDefault();
        showOrHideItems("brand", $(this));
    });

    $(".size-more").click(function(e) {
        e.preventDefault();
        showOrHideItems("size", $(this));
    });

    $(".color-more").click(function(e) {
        e.preventDefault();
        showOrHideItems("color", $(this));
    });

    function showOrHideItems(name, self) {
        if ($(self).data("state") == "close") {
            $.each($("input[name=" + name + "]"), function(k, input) {
                if ($(input).prop("checked") == false) {
                    $(input).closest("label").addClass("show");
                }
            });

            $(self).data("state", "open");
            $(self).html("Hide")
        } else {
            var len = $("input[name=" + name + "]:checked").length;
            var counter = 5 - len;

            $.each($("input[name=" + name + "]"), function(k, input) {
                if ($(input).prop("checked") == false) {
                    if (counter > 0) {
                        counter -= 1;
                        return true;
                    }

                    $(input).closest("label").removeClass("show");
                }
            });

            $(self).data("state", "close");
            $(self).html("Show more");
        }
    }

    // filter for brands
    $("input[name=brand]").click(function() {
        var list = "list";
        $.each($("input[name=brand]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProducts(list, /&b=[^(&|\s|#)]+/, "&b=");
    });

    // filter for sizes
    $("input[name=size]").click(function() {
        var list = "list";
        $.each($("input[name=size]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProducts(list, /&s=[^(&|\s|#)]+/, "&s=");
    });

    // filter for colors
    $("input[name=color]").click(function() {
        var list = "list";
        $.each($("input[name=color]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProducts(list, /&c=[^(&|\s|#)]+/, "&c=");
    });

    function filterProducts(list, regexp, prefix) {
        var url = $(location).attr("href");

        if (regexp.test(url)) {
            if (list == "list") {
                url = url.replace(regexp, "");
            } else {
                url = url.replace(regexp, prefix + list);
            }
        } else {
            url += prefix + list;
        }

        $(location).attr("href", url);
    }

    $(".cart-btn").click(function(e) {
        e.preventDefault();
        if ($("#authUser").length == 0) {
            $(location).attr("href", "/login/");
            return
        }

        var userID = $("#authUser").data("id");
        var productID = $(this).data("id");
        var data = {
            user_id: parseInt(userID),
            product_id: parseInt(productID),
            quantity: 1
        };

        $.ajax({
            url: "/cart/",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            dataType: "json",
            success: function(data) {
                // console.log(data);
                if (data.Status) {
                    var itemQnt = $("#itemQnt").html();
                    var newItemQnt = parseInt(itemQnt) + 1;
                    $("#itemQnt").html(newItemQnt);
                }

                alert(data.Message);
            },
            error: function(err){
                // console.log(err);
                alert("Sorry, something went wrong");
            }
        });
    });

    $(".qtybtn").click(function() {
        var self = $(this);
		var oldValue = self.parent().find("input").val();
		if (self.hasClass("inc")) {
			var newVal = parseFloat(oldValue) + 1;
		} else {
			// Don't allow decrementing below zero
			if (oldValue > 0) {
				var newVal = parseFloat(oldValue) - 1;
			} else {
				newVal = 0;
			}
        }

        var cartID = self.closest("tr").data("cart");
        var productID = self.closest("tr").data("product");
        var data = {
            cart_id: parseInt(cartID),
            product_id: parseInt(productID),
            quantity: newVal
        };

        $.ajax({
            url: "/cart/",
            type: "PUT",
            contentType: "application/json",
            data: JSON.stringify(data),
            dataType: "json",
            success: function(data) {
                // console.log(data);
                self.parent().find("input").val(newVal);
                self.closest("tr").find(".cart__total").html("$ " + data.Subtotal);
                $("#totalSum").html("$ " + data.Total);
            },
            error: function(err) {
                // console.log(err);
                alert("Sorry, something went wrong");
            }
        });
    });

    $(".cart__close").click(function () {
        var self = $(this);
        var cartID = self.closest("tr").data("cart");
        var productID = self.closest("tr").data("product");
        var data = {
            cart_id: parseInt(cartID),
            product_id: parseInt(productID),
        };
    
        $.ajax({
            url: "/cart/",
            type: "DELETE",
            contentType: "application/json",
            data: JSON.stringify(data),
            dataType: "json",
            success: function(data) {
                // console.log(data);
                $("#totalSum").html("$ " + data.Total);
                self.closest("tr").remove();
            },
            error: function(err) {
                // console.log(err);
                alert("Sorry, something went wrong");
            }
        });
    });
})(jQuery);
