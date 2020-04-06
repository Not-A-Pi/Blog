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