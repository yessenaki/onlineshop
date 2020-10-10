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

    $(".btn-cart").click(function(e) {
        e.preventDefault();
        var data = {
            user_id: 1,
            product_id: parseInt($(this).data("id")),
            quantity: 1
        };
        $.ajax({
            url: "/cart/",
            type: "POST",
            contentType: "application/json",
            data: JSON.stringify(data),
            dataType: "json",
            success: function(msg) {
                alert(msg);
            },
            error: function(err){
                // console.log(err);
                alert("Sorry, something went wrong");
            }
        });
    });
})(jQuery);
