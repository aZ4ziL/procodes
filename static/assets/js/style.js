$("#toHome").on("click", function () {
    $(this).addClass("active");
    $("#toAboutMe").removeClass("active")
    $("#toProjects").removeClass("active")
    closeCollapse();
    $("html, body").animate({
        scrollTop: $("#home").offset().top - 100,
    }, 500);
})
$("#toAboutMe").on("click", function () {
    $(this).addClass("active");
    $("#toHome").removeClass("active")
    $("#toProjects").removeClass("active")
    closeCollapse();
    $("html, body").animate({
        scrollTop: $("#about-me").offset().top - 100,
    }, 500);
})
$("#toProjects").on("click", function () {
    $(this).addClass("active");
    $("#toHome").removeClass("active")
    $("#toAboutMe").removeClass("active")
    closeCollapse()
    $("html, body").animate({
        scrollTop: $("#projects").offset().top - 100,
    }, 500);
})

// Tonavbar button
$("#toNavbar").on("click", function () {
    closeCollapse();
    $("html, body").animate({
        scrollTop: $("#navbar").offset().top
    })
})

// Close Collapse Navbar
function closeCollapse() {
    $("#navbarNav").collapse("hide");
}


// Typing
document.addEventListener("DOMContentLoaded", function() {
    new TypeIt("#header-text", {
        waitUntilVisible: true,
        speed: 40,
        strings: [
            "Website ini adalah website pribadi yang berjalan pada server Ubuntu Linux. Dan dibangun oleh bahasa pemrograman Go.",
            "Didesain oleh <a href='https://www.facebook.com'>Fajri Fath</a> dengan bantuan framework Bootstrap dan Javascript Vanilla.",
            // "<button class='btn btn-outline-dark'><a href=''><i class='fas fa-arrow-down fa-4x'></a></button>",
        ]
    }).go();
})