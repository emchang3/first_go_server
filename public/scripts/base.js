var menuOpen = false;

function menuButtonClick() {
  var bottom = document.getElementById('menu-bottom');
  var middle = document.getElementById('menu-middle');
  var top = document.getElementById('menu-top');

  if (menuOpen) {
    bottom.className = 'menu-bottom-bar menu-bar-flat';
    middle.className = 'menu-middle-bar';
    top.className = 'menu-top-bar menu-bar-flat';
    menuOpen = false;
  } else {
    bottom.className = 'menu-bottom-bar menu-bottom-tilt';
    middle.className = 'menu-middle-fade';
    top.className = 'menu-top-bar menu-top-tilt';
    menuOpen = true;
  }
}

document.getElementById('menu-button').addEventListener('click', menuButtonClick);
