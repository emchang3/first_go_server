namespace main {
  /**
   * menuOpen: This variable keeps track of whether the menu is open.
   */
  let menuOpen : boolean = false;

  /**
   * menuButtonClick: This method is the event handler for 'click' on the Menu Button.
   * @return {void} No return values.
   */
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
      nav.className = 'nav-menu-closed mdl-shadow--16dp';
      header.className = 'header-original mdl-shadow--4dp';
      menuOpen = false;
    } else {
      bottom.className = 'menu-bottom-bar menu-bottom-tilt';
      middle.className = 'menu-middle-fade';
      top.className = 'menu-top-bar menu-top-tilt';
      nav.className = 'nav-menu-open mdl-shadow--16dp';
      header.className = 'header-shifted mdl-shadow--4dp';
      menuOpen = true;
    }
  }

  const menuButton : HTMLElement = document.getElementById('menu-button');
  menuButton.addEventListener('click', menuButtonClick);
}