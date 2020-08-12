package service

import (
	"encoding/csv"
	"masterplan-backend/models"
	"os"
)

func GeneratExcelFile(masterplanList []models.Masterplan) (err error) {
	columnHeader := []string{"Sl No", "Activity", "Start Date", "End Date"}
	file, err := os.Create(models.Config.MasterplanFilename + ".csv")
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	err = writer.Write(columnHeader) //Add the header
	if err != nil {
		return err
	}

	err = writer.WriteAll(FuseFields(masterplanList)) //Write the masterplan data to the file
	if err != nil {
		return err
	}
	return nil
}

// FuseFields combines the data to be written to [][]string so that we could write data in excel file all at once
func FuseFields(masterplanList []models.Masterplan) (fileData [][]string) {
	column := make([]string, 4, 4) //fixed size decalration
	for _, value := range masterplanList {
		column[0] = value.SlNo
		column[1] = value.Activity
		column[2] = value.StartDate.Format("2006-01-02")
		column[3] = value.EndDate.Format("2006-01-02")
		fileData = append(fileData, column)
		column = make([]string, 4, 4) //assign new slice segment in order to avoid overqritng on previous one
	}
	return fileData
}
