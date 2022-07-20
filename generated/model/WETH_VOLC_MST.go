package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
	"github.com/satori/go.uuid"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
	_ = uuid.UUID{}
)

/*
DB Table Details
-------------------------------------


CREATE TABLE `WETH_VOLC_MST` (
  `WETH_VOLC_CD` char(3) NOT NULL COMMENT '火山コード',
  `WETH_VOLC_NM` varchar(60) NOT NULL COMMENT '火山名',
  `WETH_VOLC_AREA_CD` char(2) NOT NULL COMMENT '火山エリアコード',
  `WETH_OBSE_CD` char(3) DEFAULT NULL COMMENT '気象台コード（管区気象台）',
  `VOLC_KIND` char(1) NOT NULL COMMENT '火山種別	 1:噴火警戒レベル導入火山 2:噴火警戒レベル未導入火山 3:海底火山',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`WETH_VOLC_CD`),
  KEY `WETH_VOLC_MST_IDX_1` (`WETH_VOLC_AREA_CD`,`WETH_VOLC_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "weth_volc_cd": "uQUATGiuSIMSNBdopONrQjmfv",    "weth_volc_nm": "XLWsfdPmqBuyqbsWsEfLCBmqP",    "weth_volc_area_cd": "HWOTjMEeUSpgrWvKMkgcnaXis",    "weth_obse_cd": "nWSDjkibykenDqWRSTguVrGeA",    "volc_kind": "JQVyCQCrVTVLeWZhDEkOYcrSJ",    "disp_dsq": 38,    "dele_flg": "LtSHjgyFCThjKXLtMfrfZoNUa",    "upda_dte": "2300-08-24T15:23:35.160994506+09:00",    "upda_user_cd": "kZYGqMEkHZyNwwVJtPdnfdeRq",    "crea_dte": "2164-02-15T10:47:02.663011507+09:00",    "crea_user_cd": "aQXDNFsjbkcMDdgskqAumbXkA"}



*/

// WETH_VOLC_MST_ struct is a row record of the WETH_VOLC_MST table in the anpidb database
type WETH_VOLC_MST_ struct {
	//[ 0] WETH_VOLC_CD                                   char(3)              null: false  primary: true   isArray: false  auto: false  col: char            len: 3       default: []
	WETHVOLCCD string `gorm:"primary_key;column:WETH_VOLC_CD;type:char;size:3;"` // 火山コード
	//[ 1] WETH_VOLC_NM                                   varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	WETHVOLCNM string `gorm:"column:WETH_VOLC_NM;type:varchar;size:60;"` // 火山名
	//[ 2] WETH_VOLC_AREA_CD                              char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	WETHVOLCAREACD string `gorm:"column:WETH_VOLC_AREA_CD;type:char;size:2;"` // 火山エリアコード
	//[ 3] WETH_OBSE_CD                                   char(3)              null: true   primary: false  isArray: false  auto: false  col: char            len: 3       default: []
	WETHOBSECD sql.NullString `gorm:"column:WETH_OBSE_CD;type:char;size:3;"` // 気象台コード（管区気象台）
	//[ 4] VOLC_KIND                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	VOLCKIND string `gorm:"column:VOLC_KIND;type:char;size:1;"` // 火山種別	 1:噴火警戒レベル導入火山 2:噴火警戒レベル未導入火山 3:海底火山
	//[ 5] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 6] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"`
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[10] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"`
}

var WETH_VOLC_MSTTableInfo = &TableInfo{
	Name: "WETH_VOLC_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "WETH_VOLC_CD",
			Comment:            `火山コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHVOLCCD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_volc_cd",
			ProtobufFieldName:  "weth_volc_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WETH_VOLC_NM",
			Comment:            `火山名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "WETHVOLCNM",
			GoFieldType:        "string",
			JSONFieldName:      "weth_volc_nm",
			ProtobufFieldName:  "weth_volc_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "WETH_VOLC_AREA_CD",
			Comment:            `火山エリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "WETHVOLCAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_volc_area_cd",
			ProtobufFieldName:  "weth_volc_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "WETH_OBSE_CD",
			Comment:            `気象台コード（管区気象台）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHOBSECD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_obse_cd",
			ProtobufFieldName:  "weth_obse_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "VOLC_KIND",
			Comment: `火山種別	 1:噴火警戒レベル導入火山 2:噴火警戒レベル未導入火山 3:海底火山`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "VOLCKIND",
			GoFieldType:        "string",
			JSONFieldName:      "volc_kind",
			ProtobufFieldName:  "volc_kind",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "int32",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "UPDA_DTE",
			Comment:            `更新日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "UPDADTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "upda_dte",
			ProtobufFieldName:  "upda_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "UPDA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CREA_DTE",
			Comment:            `作成日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "timestamp",
			DatabaseTypePretty: "timestamp",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "timestamp",
			ColumnLength:       -1,
			GoFieldName:        "CREADTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "crea_dte",
			ProtobufFieldName:  "crea_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "CREA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_VOLC_MST_) TableName() string {
	return "WETH_VOLC_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_VOLC_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_VOLC_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_VOLC_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_VOLC_MST_) TableInfo() *TableInfo {
	return WETH_VOLC_MSTTableInfo
}
