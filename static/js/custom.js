'use strict';

(function ($) {
    // filter for brands
    $("input[name=brand]").click(function() {
        var list = "list";
        $.each($("input[name=brand]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProduct(list, /&b=[^(&|\s|#)]+/, "&b=");
    });

    // filter for sizes
    $("input[name=size]").click(function() {
        var list = "list";
        $.each($("input[name=size]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProduct(list, /&s=[^(&|\s|#)]+/, "&s=");
    });

    // filter for colors
    $("input[name=color]").click(function() {
        var list = "list";
        $.each($("input[name=color]"), function(k, input) {
            if ($(input).prop("checked")) {
                list += ":" + $(input).val();
            }
        });

        filterProduct(list, /&c=[^(&|\s|#)]+/, "&c=");
    });

    function filterProduct(list, regexp, prefix) {
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
})(jQuery);
