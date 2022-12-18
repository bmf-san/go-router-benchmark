function createSheets() {
  const input_sheetname = "static-routes-root/time.csv,static-routes-root/nsop.csv,static-routes-root/bop.csv,static-routes-root/allocs.csv,static-routes-1/time.csv,static-routes-1/nsop.csv,static-routes-1/bop.csv,static-routes-1/allocs.csv,static-routes-5/time.csv,static-routes-5/nsop.csv,static-routes-5/bop.csv,static-routes-5/allocs.csv,static-routes-10/time.csv,static-routes-10/nsop.csv,static-routes-10/bop.csv,static-routes-10/allocs.csv,pathparam-routes-1/time.csv,pathparam-routes-1/nsop.csv,pathparam-routes-1/bop.csv,pathparam-routes-1/allocs.csv,pathparam-routes-5/time.csv,pathparam-routes-5/nsop.csv,pathparam-routes-5/bop.csv,pathparam-routes-5/allocs.csv,pathparam-routes-10/time.csv,pathparam-routes-10/nsop.csv,pathparam-routes-10/bop.csv,pathparam-routes-10/allocs.csv";
  const sheets_name = input_sheetname.split(',');
  const input_formatsheet = Browser.inputBox('Input the name of the sheet to be copied', Browser.Buttons.OK_CANCEL);
  let spreadsheet = SpreadsheetApp.getActiveSpreadsheet();
  let exist_sheets = spreadsheet.getSheets();
  inputsheetloop:for (i = 0; i < sheets_name.length; i++) {
    for(j = 0; j < exist_sheets.length; j++ ){
      if(sheets_name[i] == exist_sheets[j].getSheetName()){
        Browser.msgBox('Skip' + sheets_name[i] + 'since it already exists');
        continue inputsheetloop;
      }
    }
    let formatsheet = spreadsheet.getSheetByName(input_formatsheet);
    let copiedsheet = formatsheet.copyTo(spreadsheet);
    copiedsheet.setName(sheets_name[i]);
  }
}