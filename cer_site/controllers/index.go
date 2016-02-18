package controllers

import (
	//"encoding/json"
	//"fmt"
	"github.com/astaxie/beego"
	"github.com/donnie4w/go-logger/logger"
	//"strings"
	docx "github.com/dcmsy/cer_site/util"
)

//cer info
type IndexInfo struct {
	beego.Controller
	Singletimer  string
	Repeatetimer string
	Transferes   string
	Nosql        string

	F_ip       string
	F_port     string
	F_username string
	F_password string
	F_dbtype   string
	F_driver   string
	F_dbname   string

	T_ip       string
	T_port     string
	T_username string
	T_password string
	T_dbtype   string
	T_driver   string
	T_dbname   string

	Systype  string
	Dbtype   string
	Sysid    string
	Sysname  string
	Company  string
	Platform string
	Dirname  string
}

// Ret entity
type Ret struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//index
func (this *IndexInfo) Index() {
	list := make([]*IndexInfo, 1, 1)
	this.Data["list"] = list
	this.TplNames = "cer/index.html"
}

//down cer file
func (this *IndexInfo) CerFile() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("读取模板文件出错：", err)
		}
	}()
	r, err := docx.ReadDocxFile("./static/docx/template.docx")
	if err != nil {
		panic(err)
	}

	docx1 := r.Editable()
	docx1.Replace("old_1_1", "new_1_1", -1)
	docx1.Replace("old_1_2", "new_1_2", -1)
	docx1.WriteToFile("./static/docx/new_result_1.docx")

	docx2 := r.Editable()
	docx2.Replace("old_2_1", "new_2_1", -1)
	docx2.Replace("old_2_2", "new_2_2", -1)
	docx2.WriteToFile("./static/docx/new_result_2.docx")

	r.Close()
}
