function onOpen() {
  var ui = SpreadsheetApp.getUi()
  var menu = ui.createMenu('utils');
  menu.addItem('Create sheets for importing csv files', 'createSheets');
  menu.addItem('Import csv files', 'importCSVFromGoogleDrive');
  menu.addToUi();
}