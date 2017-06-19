package util

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/tealeg/xlsx"

	"pccqcpa.com.cn/components/zlog"
)

func GetExcelTmpPath() string {
	return GetApiPath() + "/xlsTemp/"
}

func GetValueFormExcel(struc interface{}, fileName string) error {
	refValu := reflect.ValueOf(struc)
	strucType := refValu.Elem().Type()

	strucSlice := reflect.MakeSlice(strucType, 0, 0)
	//strucSliceValu := reflect.New(refValu.Type())
	xlFile, err := xlsx.OpenFile(fileName)
	if err == nil {
		sheet := xlFile.Sheets[0]
		for j, row := range sheet.Rows {
			if j > 0 {
				ok := true
				structNew := reflect.New(strucType.Elem())
				for k, cell := range row.Cells {
					if "" == strings.TrimSpace(cell.Value) {
						continue
					}
					ok = false
					fmt.Println("excel有效长度", k, cell.Value, structNew.Elem().NumField())
					if k >= structNew.Elem().NumField() {
						er := fmt.Errorf("excel有效列过多！")
						zlog.Error(er.Error(), er)
						return er
					}
					field := structNew.Elem().Field(k)
					switch field.Type().Name() {
					case "string":

						field.SetString(cell.Value)
					case "float64":
						if "" != cell.Value {
							valu, err := cell.Float()
							if err != nil {
								return fmt.Errorf("第 %d 行,第 %d 列出错,%s", j+1, k, err.Error())
							}
							field.SetFloat(valu)
						}
					case "Time":
						if "" != cell.Value {
							cellFormtTime := cell.Value
							floatTime, err := strconv.ParseFloat(cellFormtTime, 64)
							if nil != err {
								return fmt.Errorf("第 %d 行,第 %d 列出错,%s", j+1, k, err.Error())
							}
							field.Set(reflect.ValueOf(xlsx.TimeFromExcelTime(floatTime, false)))
						}
					case "int":
						if "" != cell.Value {
							valu, err := cell.Int64()
							if err != nil {
								return fmt.Errorf("第 %d 行,第 %d 列出错,%s", j+1, k, err.Error())
							}
							field.SetInt(valu)
						}
					}
				}
				if !ok {
					strucSlice = reflect.Append(strucSlice, structNew.Elem())
				}

			}
		}
		refValu.Elem().Set(strucSlice)
	}
	return err
}

// const (
// 	CellTypeString int = iota
// 	CellTypeFormula
// 	CellTypeNumeric
// 	CellTypeBool
// 	CellTypeInline
// 	CellTypeError
// 	CellTypeDate
// 	CellTypeGeneral
// )
