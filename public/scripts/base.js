var menuOpen = false;

var nav = document.getElementById('nav-menu');

function menuButtonClick() {
  var bottom = document.getElementById('menu-bottom');
  var header = document.getElementById('header');
  var middle = document.getElementById('menu-middle');
  var top = document.getElementById('menu-top');

  if (menuOpen) {
    bottom.className = 'menu-bottom-bar menu-bar-flat';
    middle.className = 'menu-middle-bar';
    top.className = 'menu-top-bar menu-bar-flat';
    nav.className = 'nav-menu-closed mdl-shadow--2dp';
    header.className = 'header-original mdl-shadow--2dp';
    menuOpen = false;
  } else {
    bottom.className = 'menu-bottom-bar menu-bottom-tilt';
    middle.className = 'menu-middle-fade';
    top.className = 'menu-top-bar menu-top-tilt';
    nav.className = 'nav-menu-open mdl-shadow--2dp';
    header.className = 'header-shifted mdl-shadow--2dp';
    menuOpen = true;
  }
}

document.getElementById('menu-button').addEventListener('click', menuButtonClick);
