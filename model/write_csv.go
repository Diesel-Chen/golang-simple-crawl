package model

import (
	"encoding/csv"
	"os"
)

func WriteToCSV(persons []Person) error {
	csvFile, err := os.Create("data.csv")
	if err != nil {
		return err
	}
	defer csvFile.Close()
	write := csv.NewWriter(csvFile)
	csvFile.WriteString("\xEF\xBB\xBF")
	defer write.Flush()
	headers := []string{"id", "昵称", "头像", "性别", "年龄", "城市", "居住地", "月薪", "高度", "婚姻状态", "个性签名"}
	if err = write.Write(headers); err != nil {
		return err
	}
	for _, person := range persons {
		var row []string
		row = append(row, person.Id)
		row = append(row, person.Nick)
		row = append(row, person.Avatar)
		row = append(row, person.Gender)
		row = append(row, person.Age)
		row = append(row, person.City)
		row = append(row, person.Residence)
		row = append(row, person.Salary)
		row = append(row, person.Height)
		row = append(row, person.MarriageStatus)
		row = append(row, person.Signature)
		if err = write.Write(row); err != nil {
			return err
		}
	}
	return nil
}
