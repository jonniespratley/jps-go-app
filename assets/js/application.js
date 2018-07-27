require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require('./lib/pace.min.js')
Pace.restart()


$(() => {

 //   pace.start()
  // var toolbar = mdc.toolbar.MDCToolbar.attachTo(document.querySelector('.mdc-toolbar'));
  // toolbar.fixedAdjustElement = document.querySelector('.mdc-toolbar-fixed-adjust');
  /*
   toolbar.listen('MDCToolbar:change', function(evt) {
    var flexibleExpansionRatio = evt.detail.flexibleExpansionRatio;
    console.log(flexibleExpansionRatio.toFixed(2));
  });
   */

  let drawerOpen = false;
  const app = document.querySelector('.app');
  //const drawer = new mdc.drawer.MDCPersistentDrawer(document.querySelector('.mdc-drawer--persistent'));
  const drawer = new mdc.drawer.MDCTemporaryDrawer(document.querySelector(".mdc-drawer--temporary"));
  const drawerToggle = document.querySelector('.layout-drawer__toggle')

  drawerToggle.addEventListener("click", function () {
    drawer.open = !drawer.open;
    drawerOpen = drawer.open;
    //app.classList.toggle('drawer-open');
  });

    console.log('application.js loaded');
});

