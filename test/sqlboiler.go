package test

import (
	"context"
	"database/sql"
	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql" // エイリアスでprefixを省略できる
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"log"
	"net/http"
	"ormSample/models"
	"strconv"
)

type SqlBoiler struct {
	Handler *sql.DB
}

type Summary struct {
	ITEM_CD     string `boil:"ITEM_CD"`
	ITEM_NM     string `boil:"ITEM_NM"`
	ITEM_POS    string `boil:"ITEM_POS"`
	UP_ITEM_CD  string `boil:"UP_ITEM_CD"`
	TOP_ITEM_CD string `boil:"TOP_ITEM_CD"`
	STS_CNT     string `boil:"STS_CNT"`
}

func (s *SqlBoiler) InitSqlBoiler() {
	var err error

	s.Handler, err = sql.Open("mysql", "rescuenow:rescuenow119@tcp(10.60.34.10:3306)/anpidb?parseTime=true")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

}

// パスパラメータ:idの国の該当データを出力
func (s *SqlBoiler) GetCountry(w rest.ResponseWriter, r *rest.Request) {
	ctx := context.Background()

	id := r.PathParam("id")
	countries, err := models.Countries(qm.Where("id=?", id)).All(ctx, s.Handler)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteJson(&countries)
}

// countriesテーブル内のデータ全出力
func (s *SqlBoiler) GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	ctx := context.Background()

	countries, err := models.Countries().All(ctx, s.Handler)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteJson(&countries)
}

// json形式のデータをPOST {name:国名}
func (s *SqlBoiler) PostCountry(w rest.ResponseWriter, r *rest.Request) {
	ctx := context.Background()

	country := models.Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = country.Insert(ctx, s.Handler, boil.Infer())
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データのNameを変更し出力
func (s *SqlBoiler) PutCountry(w rest.ResponseWriter, r *rest.Request) {
	ctx := context.Background()

	country := models.Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	_, err = country.Update(ctx, s.Handler, boil.Infer())
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データを削除
func (s *SqlBoiler) DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	ctx := context.Background()

	pid, _ := strconv.ParseInt(id, 10, 64)
	country := models.Country{ID: pid}

	if _, err := country.Delete(ctx, s.Handler); err != nil {
		rest.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// 生SQLの実施
func (s *SqlBoiler) GetCounttryCond(w rest.ResponseWriter, r *rest.Request) {

	var su []*Summary

	ctx := context.Background()

	sql := "" +
		"\n#EXPLAIN\nSELECT\n\tDEPT_ID AS ITEM_CD,\n\tDEPT_NM AS ITEM_NM,\n\tDEPT_POS AS ITEM_POS,\n\tIFNULL(UP_DEPT_ID, 0) AS UP_ITEM_CD,\n\tCASE\n\t\tWHEN TOP_DEPT_ID = 0 THEN DEPT_ID\n\t\tELSE TOP_DEPT_ID\n\tEND AS TOP_ITEM_CD,\n\tCUR_STS_CD AS STS_CD,\n\tCNT AS STS_CNT\nFROM\n\t(\n\tSELECT\n\t\tUDM.DEPT_ID1 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID1 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID1,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID2 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID2 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID2,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID3 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID3 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID3,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID4 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID4 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID4,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID5 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID5 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID5,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID6 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID6 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID6,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID7 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID7 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID7,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID8 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID8 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID8,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID9 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID9 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID9,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID10 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID10 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID10,\n\t\tASU.CUR_STS_CD\n\n) AS SH\nORDER BY\n\tDISP_DSQ,\n\tDISP_DSQ_SUB_NO,\n\tDEPT_ID,\n\tCUR_STS_CD\n\n\n\n"

	err := queries.Raw(sql).Bind(ctx, s.Handler, &su)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}

	w.WriteJson(&su)
}
