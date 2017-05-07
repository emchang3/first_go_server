(function() : void {
  let menuOpen : boolean = false;

  const menuButtonClick = function() : void {
    const bottom : HTMLElement = document.getElementById('menu-bottom');
    const header : HTMLElement = document.getElementById('header');
    const middle : HTMLElement = document.getElementById('menu-middle');
    const top : HTMLElement = document.getElementById('menu-top');
    const nav : HTMLElement = document.getElementById('nav-menu');

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

  const menuButton : HTMLElement = document.getElementById('menu-button');

  menuButton.addEventListener('click', menuButtonClick);
})();
