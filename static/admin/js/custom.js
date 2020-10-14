$(document).ready(function() {
    "use strict";

    selectSizes();

    $("#category_id").change(function() {
        selectSizes();
    });

    function selectSizes() {
        var parentID = $("#category_id option:selected").data("parent");
        if (parentID == 3 || parentID == 4) {
            $("#size_id option[value='0']").prop("selected", true);
            $("#size_id").closest(".form-group").hide();
        } else {
            $("#size_id").closest(".form-group").show();
            $.each($("#size_id option"), function(k, option) {
                if ($(option).data("type") == 0) {
                    if (parentID == 1) $(option).show();
                    else $(option).hide();
                } else if ($(option).data("type") == 1) {
                    if (parentID == 1) $(option).hide();
                    else $(option).show();
                }
            });
        }
    }

    $(".form-image i").click(function(e) {
        if (confirm('Are you sure you want to do this?') == false) {
            e.preventDefault();
        }

        var self = $(this);
        $.ajax({
            url: "/admin/products/delete-image/",
            type: "POST",
            data: {"id": self.data("id")},
            dataType: "json",
            success: function(data) {
                // console.log(data);
                if (data) {
                    self.closest(".col-4").remove();
                }
            },
            error: function(err){
                // console.log(err);
                alert("Sorry, something went wrong");
            }
        });
    });

    if ($(location).attr("pathname").indexOf("admin/post") != -1) {
        $("a[data-target='#blog']").trigger("click");
        $("a[data-target='#blog']").parent().addClass("expand");
    }
});
