// ========== jQuery ==========
//  TODO: Refactor Vanilla
// ============================
$(document).ready(function() {

    // Scroll Control
    $(window).scroll(function() {
        if ($(document).scrollTop() > 400) {
            $('.up').stop().addClass('show');
        } else {
            $('.up').stop().removeClass('show');
        }
    });
    $('a[href*=#]:not([href=#])').click(function() {
        if (location.pathname.replace(/^\//, '') === this.pathname.replace(/^\//, '') || location.hostname === this.hostname) {

            var target = $(this.hash);
            target = target.length ? target : $('[name=' + this.hash.slice(1) + ']');
            if (target.length) {
                $('html,body').animate({
                    scrollTop: target.offset().top
                }, 1000);
                return false;
            }
        }
    });

    // Viewport
    function viewportSpecs() {
        var the_width = $(window).width();
        var the_height = $(window).height();
        $('#width').text(the_width);
        $('#height').text(the_height);
    }
    $(window).resize(function() {
        viewportSpecs();
    });

    // Active/Inactive Tab
    $(".tabs li a").click(function() {
        var id = $(this).attr("data");
        $(".tabs li a").removeClass("active");
        $(this).addClass("active");
        $("#" + id).show().siblings("div").hide();
    });

    // Equal Heights
    $.fn.equalHeight = function() {
        return this.height(Math.max.apply(this,
            $.map(this, function(e) {
                return $(e).height();
            })
        ));
    };
    $(".equal").equalHeight();

    // Include Functionality
    // Ex: <div class="include" data-include="header"></div>
    var includes = $('.include');
    jQuery.each(includes, function() {
        var file = '/' + $(this).data('include') + '.html';
        $(this).load(file);
    });

    // Character Count
    $('input, textarea').on("propertychange keyup input paste",
        function() {
            var limit = $(this).data("limit");
            var remainingChars = limit - $(this).val().length;
            if (remainingChars <= 0) {
                $(this).val($(this).val().substring(0, limit));
                $(this).nextAll('div').css('color', '#cc6666');
            } else {
                $(this).nextAll('div').css('color', '');
            }
            $(this).nextAll('div').text((remainingChars <= 0 ? 0 : remainingChars) + " characters remaining.");
        });

    // Enforce an Inactive Element
    $('.inactive').click(function(event) {
        event.preventDefault();
        return false;
    });

});
