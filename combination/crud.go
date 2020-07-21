package combination

import (
	"fmt"
	"github.com/lvxin0315/golayui/weight"
	"io/ioutil"
)

type Item struct {
	Title string `json:"title"`
	Field string `json:"field"`
	Type  string `json:"type"`
}

const tableTplName = "list.html"
const dataTableTplName = "data_list.html"
const addTplName = "add.html"
const editTplName = "edit.html"

//根据模型生成对应的增删改查html
func Build(itemList []*Item, savePath string) error {
	//普通table
	tableHtml, err := table(itemList)
	if err != nil {
		fmt.Println("Build.table error:", err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", savePath, tableTplName), []byte(tableHtml), 0777)
	if err != nil {
		fmt.Println("Build.write.table error:", err)
		return err
	}
	//数据table
	dataTableHtml, err := dataTable(itemList)
	if err != nil {
		fmt.Println("Build.dataTable error:", err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", savePath, dataTableTplName), []byte(dataTableHtml), 0777)
	if err != nil {
		fmt.Println("Build.write.dataTable error:", err)
		return err
	}
	//新增
	addHtml, err := add(itemList)
	if err != nil {
		fmt.Println("Build.add error:", err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", savePath, addTplName), []byte(addHtml), 0777)
	if err != nil {
		fmt.Println("Build.write.add error:", err)
		return err
	}
	//编辑
	editHtml, err := edit(itemList)
	if err != nil {
		fmt.Println("Build.edit error:", err)
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("%s/%s", savePath, editTplName), []byte(editHtml), 0777)
	if err != nil {
		fmt.Println("Build.write.edit error:", err)
		return err
	}
	return nil
}

func table(itemList []*Item) (string, error) {
	tableWeight := new(weight.TableWeight)
	//表头
	tableTrWeight := new(weight.TableTrWeight)
	tableWeight.Thead = tableTrWeight
	for _, item := range itemList {
		tableTrWeight.TdList = append(tableTrWeight.TdList, &weight.TableTdWeight{
			Content: item.Title,
		})
	}
	//测试数据
	tableTrTestDataWeight := new(weight.TableTrWeight)
	for _, item := range itemList {
		tableTrTestDataWeight.TdList = append(tableTrTestDataWeight.TdList, &weight.TableTdWeight{
			Content: item.Field,
		})
	}
	tableWeight.Tbody = append(tableWeight.Tbody,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight,
		tableTrTestDataWeight)
	return tableWeight.Output()
}

func dataTable(itemList []*Item) (string, error) {
	dataTableWeight := new(weight.DataTableWeight)
	dataTableWeight.Id = "tableId"
	for _, item := range itemList {
		dataTableWeight.FieldList = append(dataTableWeight.FieldList, &weight.DataTableItem{
			Field: item.Field,
			Title: item.Title,
			Sort:  "true",
		})
	}
	return dataTableWeight.Output()
}

func add(itemList []*Item) (string, error) {
	formLabelWeight := new(weight.FormLabelWeight)
	formLabelWeight.Label = "LabelTitle"
	formLabelWeight.FormWeight.Id = "formId"
	for _, item := range itemList {
		fiw := new(weight.FormItemWeight)
		formLabelWeight.FormWeight.Children = append(formLabelWeight.FormWeight.Children, fiw)
		fiw.Label = item.Title
		//根据类型生成对应html内容
		var formItem weight.FormItemWeightImpl
		//id name
		attr := weight.Attr{
			Name: item.Field,
			Id:   item.Field,
		}
		switch item.Type {
		case weight.CheckBox:
			//TODO option
			formItem = &weight.CheckboxWeight{
				OptionList: nil,
			}

		case weight.Select:
			formItem = &weight.SelectWeight{
				Attr:       attr,
				OptionList: nil,
			}
			//TODO option
		case weight.InputPassword:
			formItem = &weight.InputPasswordWeight{
				Attr:        attr,
				Placeholder: item.Title,
			}
		case weight.Textarea:
			formItem = &weight.TextareaWeight{
				Attr:        attr,
				Placeholder: item.Title,
			}
		case weight.Radio:
			formItem = &weight.RadioWeight{
				OptionList: nil,
			}
			//TODO option
		default:
			formItem = &weight.InputTextWeight{
				Attr:        attr,
				Placeholder: item.Title,
				IsDate:      false,
			}
		}
		fiw.Item = formItem
	}
	//添加提交按钮
	formLabelWeight.FormWeight.Children = append(formLabelWeight.FormWeight.Children, &weight.ButtonWeight{
		Attr: weight.Attr{
			Id: "btnId",
		},
		Title: "提交",
	})
	return formLabelWeight.Output()
}

func edit(itemList []*Item) (string, error) {
	//TODO 之后添加js时处理吧
	return add(itemList)
}
