function importCSVFromGoogleDrive() {
  const folderId = "*************";
  const targetFolder = DriveApp.getFolderById(folderId);
  const folders = targetFolder.getFolders();
  while(folders.hasNext()) {
    var folder = folders.next();
    var files = folder.getFiles();
    while (files.hasNext()) {
      var file = files.next();
      const targetFiles = ["time.csv", "nsop.csv", "bop.csv", "allocs.csv"];
      if (targetFiles.includes(file.getName())) {
        importCsv(file, folder.getName() + "/" + file.getName());
      }
    }
  }
}

function importCsv(file, sheetName) {
  var csvData = Utilities.parseCsv(file.getBlob().getDataAsString());
  var sheet = SpreadsheetApp.getActiveSpreadsheet().getSheetByName(sheetName);
  sheet.clear();
  sheet.getRange(1, 1, csvData.length, csvData[0].length).setValues(csvData);
  console.log(sheetName + "has imported.");
}