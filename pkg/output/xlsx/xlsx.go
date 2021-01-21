package xlsx

import (
	"fmt"
	"time"

	"github.com/tealeg/xlsx"
	"github.com/varshaprasad96/operator-sdk-data-collector/pkg/fields"
)

func GetOutput(data fields.OperatorData, outputFilePath string) error {
	output := xlsx.NewFile()

	if err := createSheetsAndFillIndexData(output, "community", data.CommunityOperators); err != nil {
		return fmt.Errorf("error writing data for community operators %v", err)
	}

	if err := createSheetsAndFillIndexData(output, "certified", data.CertifiedOperators); err != nil {
		return fmt.Errorf("error writing data for certified operators %v", err)
	}

	if err := createSheetsAndFillIndexData(output, "marketplace", data.MarketplaceOperators); err != nil {
		return fmt.Errorf("error writing data for marketplace operators %v", err)
	}

	if err := createSheetsAndFillIndexData(output, "operatorhub", data.OperatorHub); err != nil {
		return fmt.Errorf("error writing data for operatorhub operators %v", err)
	}

	if err := createSheetsAndFillIndexData(output, "redhat", data.RedHatOperators); err != nil {
		return fmt.Errorf("error writing data for redhat operators %v", err)
	}

	if err := createSheetsAndFillIndexData(output, "prod", data.ProdOperators); err != nil {
		return fmt.Errorf("error writing data for prod operators %v", err)
	}

	defer func() {
		outputName := time.Now().Format("Mon-Jan2-15:04:05PST-2006")
		if err := output.Save(outputFilePath + outputName + ".xlsx"); err != nil {
			fmt.Printf("error whilesaving report")
		}
	}()

	return nil
}

func createSheetsAndFillIndexData(f *xlsx.File, index string, data map[string]fields.ReportColumns) error {
	sheet, err := f.AddSheet(index)
	if err != nil {
		return fmt.Errorf("error creating xlsx sheet")
	}

	initializeReport(sheet)

	for _, value := range data {
		row := sheet.AddRow()

		// Add operator Name
		row.AddCell().Value = value.Operator

		// Add csv timestamp
		row.AddCell().Value = value.CreatedAt

		// Add name of the company
		row.AddCell().Value = value.Company

		// Add operator type
		row.AddCell().Value = value.OperatorType

		// Add sdk version
		row.AddCell().Value = value.SDKVersion
	}

	return nil
}

func initializeReport(sh *xlsx.Sheet) {
	row := sh.AddRow()
	row.AddCell().Value = "Operator name"
	row.AddCell().Value = "CreatedAt - timestamp"
	row.AddCell().Value = "Company"
	row.AddCell().Value = "Operator type"
	row.AddCell().Value = "Sdk Version"
}
