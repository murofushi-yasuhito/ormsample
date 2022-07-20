package main

import (
	"github.com/jinzhu/gorm"
	"log"
	"net/http"
	"time"

	"github.com/ant0ine/go-json-rest/rest"
	_ "github.com/go-sql-driver/mysql" // エイリアスでprefixを省略できる
	"ormSample/test"
	"xorm.io/xorm"
)

type Country struct {
	Id        int64     `json:"id"`
	Name      string    `sql:"size:1024" json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}

type Impl struct {
	DB *gorm.DB
}

type Xorm struct {
	engine *xorm.Engine
}

type Summary struct {
	ITEM_CD     string `xorm:"item_cd"`
	ItemName    string `xorm:"ITEM_NM"`
	ITEM_POS    string `xorm:"ITEM_POS"`
	UP_ITEM_CD  string `xorm:"UP_ITEM_CD"`
	TOP_ITEM_CD string `xorm:"TOP_ITEM_CD"`
	STS_CNT     string `xorm:"STS_CNT"`
}

type SummaryGorm struct {
	ITEM_CD     string `json:"item_cd"`
	ItemName    string `json:"ITEM_NM"`
	ITEM_POS    string `json:"DEPT_POS"`
	UP_ITEM_CD  string `json:"UP_ITEM_CD"`
	TOP_ITEM_CD string `json:"TOP_ITEM_CD"`
	StsCnt      string `json:"STS_CNT"`
}

func (i *Impl) InitDB() {
	var err error
	// MySQLとの接続。ユーザ名：gorm パスワード：password DB名：country
	i.DB, err = gorm.Open("mysql", "rescuenow:rescuenow119@tcp(10.60.34.10:3306)/anpidb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}
	i.DB.LogMode(true)
}

func (x *Xorm) InitDBXorm() {
	var err error
	// MySQLとの接続。ユーザ名：gorm パスワード：password DB名：country
	x.engine, _ = xorm.NewEngine("mysql", "rescuenow:rescuenow119@tcp(10.60.34.10:3306)/anpidb?charset=utf8mb4&parseTime=true")
	if err != nil {
		log.Fatalf("Got error when connect database, the error is '%v'", err)
	}

	x.engine.ShowSQL(true)
}

// DBマイグレーション
func (i *Impl) InitSchema() {
	i.DB.AutoMigrate(&Country{})
}

func main() {
	//controllers.StartWebServer()

	i := Impl{}
	i.InitDB()
	i.InitSchema()

	// xorm 初期生成
	//x := Xorm{}
	//x.InitDBXorm()
	//
	//err := x.engine.Sync2(new(Country))
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//// SqlBoiler 初期設定
	//s := test.SqlBoiler{}
	//s.InitSqlBoiler()
	//
	// Gorm2 初期設定
	g := test.Gorm{}
	g.InitGorm()
	g.InitSchema()

	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/countries", i.GetAllCountries),
		rest.Post("/countries", i.PostCountry),
		rest.Get("/countries/:id", i.GetCountry),
		rest.Put("/countries/:id", i.PutCountry),
		rest.Delete("/countries/:id", i.DeleteCountry),
		rest.Get("/countriesSql", i.GetCountrySQL),
		//rest.Get("/xorm/countries", x.GetAllCountries),
		//rest.Post("/xorm/countries", x.PostCountry),
		//rest.Get("/xorm/countries/:id", x.GetCountry),
		//rest.Put("/xorm/countries/:id", x.PutCountry),
		//rest.Delete("/xorm/countries/:id", x.DeleteCountry),
		//rest.Get("/xorm/countriesSql", x.GetCountrySQL),
		//rest.Get("/sqlBoiler/countries", s.GetAllCountries),
		//rest.Post("/sqlBoiler/countries", s.PostCountry),
		//rest.Get("/sqlBoiler/countries/:id", s.GetCountry),
		//rest.Put("/sqlBoiler/countries/:id", s.PutCountry),
		//rest.Delete("/sqlBoiler/countries/:id", s.DeleteCountry),
		//rest.Get("/sqlBoiler/countriesCond", s.GetCounttryCond),
		rest.Get("/gorm/countries", g.GetAllCountries),
		rest.Post("/gorm/countries", g.PostCountry),
		rest.Get("/gorm/countries/:id", g.GetCountry),
		rest.Put("/gorm/countries/:id", g.PutCountry),
		rest.Delete("/gorm/countries/:id", g.DeleteCountry),
		rest.Get("/gorm/sql", g.GetCountrySQL),
		//rest.Get("/gorm/countries", g.GetAllCountries),
		rest.Post("/gorm/users", g.PostUser),
		rest.Get("/gorm/users/:id", g.GetUser),
		rest.Put("/gorm/users/:id", g.PutUser),
		rest.Delete("/gorm/users/:id", g.DeleteUser),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("server started.")
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))

}

// パスパラメータ:idの国の該当データを出力
func (i *Impl) GetCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	country := Country{}
	if i.DB.Find(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&country)
}

// countriesテーブル内のデータ全出力
func (i *Impl) GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	country := []Country{}
	i.DB.Find(&country)
	w.WriteJson(&country)
}

// json形式のデータをPOST {name:国名}
func (i *Impl) PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = i.DB.Save(&country).Error
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データのNameを変更し出力
func (i *Impl) PutCountry(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")
	country := Country{}
	if i.DB.First(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	updated := Country{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	country.Name = updated.Name

	if err := i.DB.Save(&country).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データを削除
func (i *Impl) DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	country := Country{}
	if i.DB.First(&country, id).Error != nil {
		rest.NotFound(w, r)
		return
	}
	if err := i.DB.Delete(&country).Error; err != nil {
		rest.Error(w, err.Error(), http.StatusInsufficientStorage)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// 生SQLの実施
func (i *Impl) GetCountrySQL(w rest.ResponseWriter, r *rest.Request) {

	//summary := SummaryGorm{}
	result := []SummaryGorm{}
	sql := "" +
		"\n#EXPLAIN\nSELECT\n\tDEPT_ID AS ITEM_CD,\n\tDEPT_NM AS ITEM_NM,\n\tDEPT_POS AS ITEM_POS,\n\tIFNULL(UP_DEPT_ID, 0) AS UP_ITEM_CD,\n\tCASE\n\t\tWHEN TOP_DEPT_ID = 0 THEN DEPT_ID\n\t\tELSE TOP_DEPT_ID\n\tEND AS TOP_ITEM_CD,\n\tCUR_STS_CD AS STS_CD,\n\tCNT AS STS_CNT\nFROM\n\t(\n\tSELECT\n\t\tUDM.DEPT_ID1 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID1 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID1,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID2 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID2 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID2,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID3 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID3 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID3,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID4 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID4 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID4,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID5 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID5 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID5,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID6 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID6 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID6,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID7 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID7 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID7,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID8 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID8 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID8,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID9 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID9 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID9,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID10 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID10 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID10,\n\t\tASU.CUR_STS_CD\n\n) AS SH\nORDER BY\n\tDISP_DSQ,\n\tDISP_DSQ_SUB_NO,\n\tDEPT_ID,\n\tCUR_STS_CD\n\n\n\n"

	rows, err := i.DB.Raw(sql).Rows()
	defer rows.Close()
	for rows.Next() {
		//result = append(result, i.DB.ScanRows(rows, &SummaryGorm{}))
		s := SummaryGorm{}
		i.DB.ScanRows(rows, &s)
		result = append(result, s)
	}

	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(result)
}

// countriesテーブル内のデータ全出力
func (x *Xorm) GetAllCountries(w rest.ResponseWriter, r *rest.Request) {
	countries := []Country{}

	err := x.engine.Find(&countries)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(countries)
}

// パスパラメータ:idの国の該当データを出力
func (x *Xorm) GetCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")
	country := Country{}

	result, err := x.engine.Where("id = ?", id).Get(&country)
	if err != nil {
		log.Fatal(err)
	}
	if !result {
		log.Fatal("Not Found")
	}
	w.WriteJson(&country)
}

// json形式のデータをPOST {name:国名}
func (x *Xorm) PostCountry(w rest.ResponseWriter, r *rest.Request) {
	country := Country{}
	err := r.DecodeJsonPayload(&country)

	_, err = x.engine.Insert(country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteJson(&country)
}

// パスパラメータ:idの国の該当データのNameを変更し出力
func (x *Xorm) PutCountry(w rest.ResponseWriter, r *rest.Request) {

	id := r.PathParam("id")

	updated := Country{}
	if err := r.DecodeJsonPayload(&updated); err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := x.engine.Where("id =?", id).Update(&updated)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteJson(result)
}

// パスパラメータ:idの国の該当データを削除
func (x *Xorm) DeleteCountry(w rest.ResponseWriter, r *rest.Request) {
	id := r.PathParam("id")

	_, err := x.engine.ID(id).Delete(&Country{})
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}

// 生SQLの実施
func (x *Xorm) GetCountrySQL(w rest.ResponseWriter, r *rest.Request) {

	summary := []Summary{}

	sql := "" +
		"\n#EXPLAIN\nSELECT\n\tDEPT_ID AS ITEM_CD,\n\tDEPT_NM AS ITEM_NM,\n\tDEPT_POS AS ITEM_POS,\n\tIFNULL(UP_DEPT_ID, 0) AS UP_ITEM_CD,\n\tCASE\n\t\tWHEN TOP_DEPT_ID = 0 THEN DEPT_ID\n\t\tELSE TOP_DEPT_ID\n\tEND AS TOP_ITEM_CD,\n\tCUR_STS_CD AS STS_CD,\n\tCNT AS STS_CNT\nFROM\n\t(\n\tSELECT\n\t\tUDM.DEPT_ID1 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID1 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID1,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID2 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID2 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID2,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID3 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID3 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID3,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID4 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID4 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID4,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID5 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID5 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID5,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID6 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID6 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID6,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID7 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID7 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID7,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID8 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID8 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID8,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID9 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID9 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID9,\n\t\tASU.CUR_STS_CD\nUNION ALL\n\tSELECT\n\t\tUDM.DEPT_ID10 AS DEPT_ID,\n\t\tDM.DEPT_NM,\n\t\tDM.DEPT_POS,\n\t\tDM.UP_DEPT_ID,\n\t\tDM.TOP_DEPT_ID,\n\t\tDM.DISP_DSQ,\n\t\tDM.DISP_DSQ_SUB_NO,\n\t\tIFNULL(ASU.CUR_STS_CD, '90') AS CUR_STS_CD,\n\t\tCOUNT(*) as CNT\n\tFROM\n\t\t(\n\t\tSELECT\n\t\t\tMAX(STS.ANKEN_ID) AS ANKEN_ID,\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n\t\tFROM\n\t\t\t(\n\t\t\tSELECT\n\t\t\t\tANKEN_ID,\n\t\t\t\tUSER_ID,\n\t\t\t\tCUR_STS_CD,\n\t\t\t\tCASE\n\t\t\t\t\tWHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n\t\t\t\t\tELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL)\n\t\t\t\tEND AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n            ) AS STS\n\t\tJOIN (\n\t\t\tSELECT\n\t\t\t\tUSER_ID,\n\t\t\t\tMAX(CASE\n                       WHEN LENGTH(CUR_STS_CD) = 1 THEN UNIX_TIMESTAMP(CUR_ANS_DTE)\n                       ELSE CAST(IFNULL(CUR_STS_CD, '90') AS DECIMAL) END) AS STS_CD_KEY\n\t\t\tFROM\n\t\t\t\tANKEN_STATUS_CUR\n\t\t\tWHERE\n\t\t\t\tANKEN_ID IN ( 144628, 144627 )\n\t\t\tGROUP BY\n\t\t\t\tUSER_ID\n            ) AS STS2\n                ON\n\t\t\tSTS.USER_ID = STS2.USER_ID\n\t\t\tAND STS.STS_CD_KEY = STS2.STS_CD_KEY\n\t\tGROUP BY\n\t\t\tSTS.USER_ID,\n\t\t\tSTS.CUR_STS_CD\n        ) AS ASU\n\tJOIN ANKEN_HIS_USER_DEPT_MST AS UDM\n            ON\n\t\tASU.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND ASU.USER_ID = UDM.USER_ID\n\t\tAND UDM.USER_DEPT_SEQ = 1\n\tJOIN ANKEN_HIS_DEPT_MST AS DM\n            ON\n\t\tDM.ANKEN_ID = UDM.ANKEN_ID\n\t\tAND UDM.DEPT_ID10 = DM.DEPT_ID\n\tGROUP BY\n\t\tUDM.DEPT_ID10,\n\t\tASU.CUR_STS_CD\n\n) AS SH\nORDER BY\n\tDISP_DSQ,\n\tDISP_DSQ_SUB_NO,\n\tDEPT_ID,\n\tCUR_STS_CD\n\n\n\n"

	err := x.engine.SQL(sql).Find(&summary)
	if err != nil {
		log.Fatal(err)
	}

	w.WriteJson(summary)
}
