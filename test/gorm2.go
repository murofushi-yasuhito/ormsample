package test

import (
	"fmt"
	"github.com/ant0ine/go-json-rest/rest"
	"gorm.io/gorm/logger"

	//"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql" // エイリアスでprefixを省略できる
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"ormSample/domain"
)

type Gorm struct {
	Db *gorm.DB
}

type SummaryGorm struct {
	ITEM_CD     string `boil:"ITEM_CD"`
	ITEM_NM     string `boil:"ITEM_NM"`
	ITEM_POS    string `boil:"ITEM_POS"`
	UP_ITEM_CD  string `boil:"UP_ITEM_CD"`
	TOP_ITEM_CD string `boil:"TOP_ITEM_CD"`
	STS_CNT     string `boil:"STS_CNT"`
}

func (g *Gorm) InitGorm() {

	//user := os.Getenv("rescuenow")
	//pw := os.Getenv("rescuenow119")
	//db_name := os.Getenv("anpidb")
	user := "rescuenow"
	pw := "rescuenow119"
	db_name := "anpidb"
	var path string = fmt.Sprintf("%s:%s@tcp(10.60.34.10:3306)/%s?charset=utf8&parseTime=true", user, pw, db_name)
	diact := mysql.Open(path)
	var err error
	if g.Db, err = gorm.Open(diact); err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	g.Db.Logger.LogMode(logger.Info)
	fmt.Println("db connected!!")
}

// DBマイグレーション
func (g *Gorm) InitSchema() {
	g.Db.AutoMigrate(&domain.Country{})
}

// パスパラメータ:idの国の該当データを出力
func (g *Gorm) GetCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	country := domain.Country{}
	if g.Db.Table("country").Find(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&country)
}

// countriesテーブル内のデータ全出力
func (g *Gorm) GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	country := []domain.Country{}
	g.Db.Table("country").Find(&country)
	w.WriteJson(country)
}

// json形式のデータをPOST {name:国名}
func (g *Gorm) PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := domain.Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = g.Db.Table("country").Save(&country).Error
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データのNameを変更し出力
func (g *Gorm) PutCountry(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	country := domain.Country{}
	if g.Db.Table("country").First(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	updated := domain.Country{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	country.Name = updated.Name

	if err := g.Db.Table("country").Save(&country).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データを削除
func (g *Gorm) DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	country := domain.Country{}
	if g.Db.Table("country").First(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := g.Db.Table("country").Delete(&country).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// 生SQLの実施
func (g *Gorm) GetCountrySQL(w rest.ResponseWriter, r *rest.Request) {

	//summary := SummaryGorm{}
	result := []SummaryGorm{}
	sql := "" +
		"\n#EXPLAIN\nSELECT\n\tDEPT_ID AS ITEM_CD,\n\tDEPT_NM AS ITEM_NM,\n\tDEPT_POS AS ITEM_POS,\n\tIFNULL(UP_DEPT_ID, 0) AS UP_ITEM_CD,\n\tCASE\n\t\tWHEN TOP_DEPT_ID = 0 THEN DEPT_ID\n\t\tELSE TOP_DEPT_ID\n\tEND AS TOP_ITEM_CD,\n\tCUR_STS_CD AS STS_CD,\n\tCNT AS STS_CNT\nFROM\n\t(\n\tSELECT\n\t\tUDM.DEPT_ID1 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID1 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID1,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID2 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID2 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID2,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID3 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID3 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID3,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID4 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID4 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID4,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID5 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID5 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID5,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID6 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID6 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID6,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID7 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID7 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID7,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID8 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID8 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID8,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID9 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID9 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID9,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID10 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID10 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID10,\n\t\tASU.CUR_STS_CD\n\n) AS SH\nORDER BY\n\tDISP_DSQ,\n\tDISP_DSQ_SUB_NO,\n\tDEPT_ID,\n\tCUR_STS_CD\n\n\n\n"

	rows, err := g.Db.Raw(sql).Rows()
	defer rows.Close()
	for rows.Next() {
		//result = append(result, i.DB.ScanRows(rows, &SummaryGorm{}))
		s := SummaryGorm{}
		g.Db.ScanRows(rows, &s)
		result = append(result, s)
	}

	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(result)
}
