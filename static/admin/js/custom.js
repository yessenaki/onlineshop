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
});
