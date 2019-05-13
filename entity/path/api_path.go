package path

// item path
var Item = "/item"
var ItemById = Item + "/:id"
var EntryItem = "/entry"
var OutcomeItem = "/outcome"

// order path
var Order = "/order"

// stock path
var Stock = "/stock"

// report path
var ReportValue = "/report-value"
var ReportSales = "/report-sales"

// migration path
var Migration = "/migration"
var MigrationExport = Migration + "/export"
var ExportValueReport = MigrationExport + "/report-value/:id"
var ExportSaleReport = MigrationExport + "/report-sales/:id"
var MigrationImport = Migration + "/import"