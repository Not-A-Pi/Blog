$(function() {
  function close() {
    $('body').removeClass('has-active-menu');
    setTimeout(function(){
      $('.nav-slider').removeClass('toggling');
    }, 500);
  }
  
  function open() {
    $('body').addClass('has-active-menu');
    $('.nav-slider').addClass('toggling');
  }
  
  $('.nav-mask').click(close);
  $('.navbar-toggler').click(open);
});

document.addEventListener("DOMContentLoaded", () => {
  // Get all "navbar-burger" elements
  const $navbarBurgers = Array.prototype.slice.call(
    document.querySelectorAll(".navbar-burger"),
    0
  );

  // Check if there are any navbar burgers
  if ($navbarBurgers.length > 0) {
    // Add a click event on each of them
    $navbarBurgers.forEach(el => {
      el.addEventListener("click", () => {
        // Get the target from the "data-target" attribute
        const target = el.dataset.target;
        const $target = document.getElementById(target);

        // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
        el.classList.toggle("is-active");
        $target.classList.toggle("is-active");
      });
    });
  }
});