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


CREATE TABLE `WETH_SUB_AREA_MST` (
  `WETH_SUB_AREA_CD` char(4) NOT NULL COMMENT '気象細分区域コード',
  `WETH_SUB_AREA_KIND` char(1) NOT NULL COMMENT '気象細分区域種別',
  `WETH_SUB_AREA_NMJ` varchar(40) NOT NULL COMMENT '気象細分区域名',
  `WETH_SUB_AREA_NMC` varchar(40) DEFAULT NULL COMMENT '気象細分区域名カナ',
  `WETH_OFFI_CD` char(4) NOT NULL COMMENT '気象官署コード',
  `WETH_SUB_AREA1_CD` char(4) NOT NULL COMMENT '気象一次細分区域コード',
  `XML_PREF_AREA_CD` char(6) NOT NULL COMMENT 'XML府県予報区コード',
  `XML_SUB_AREA_CD` char(6) NOT NULL COMMENT 'XML一次細分区域コード',
  `XML_CITY_AREA_CD` char(6) NOT NULL COMMENT 'XML市区町村等地域コード',
  `WEEK_YOHOU_USE_FLG` char(1) NOT NULL COMMENT '週間天気予報使用フラグ',
  `WETH_MENU_AREA_CD` char(3) DEFAULT NULL COMMENT '気象メニューエリアコード',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL COMMENT 'UPDA_USER_CD',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL COMMENT 'CREA_USER_CD',
  PRIMARY KEY (`WETH_SUB_AREA_CD`),
  KEY `WETH_SUB_AREA_MST_IDX_1` (`WETH_OFFI_CD`,`WETH_SUB_AREA1_CD`,`WETH_SUB_AREA_CD`) USING BTREE,
  KEY `WETH_SUB_AREA_MST_IDX_2` (`WETH_SUB_AREA1_CD`,`WETH_SUB_AREA_CD`) USING BTREE,
  KEY `WETH_SUB_AREA_MST_IDX_3` (`XML_SUB_AREA_CD`,`XML_CITY_AREA_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='気象細分区域マスタ'

JSON Sample
-------------------------------------
{    "weth_sub_area_cd": "tvEgADDVccljSeGYqBQAlOHfg",    "weth_sub_area_kind": "tpaNQHQjQfcWTulmBeIOMnMFK",    "weth_sub_area_nmj": "xpcNQgkSoRCsuGHMjHYFtYshW",    "weth_sub_area_nmc": "ktGBSiroRVXcBnoAPgRMXUNPw",    "weth_offi_cd": "HZlyrlkcVjeCpiLyKnWfUrGaQ",    "weth_sub_area_1_cd": "rKdFusJehBUuwIZDDGlxlWcxj",    "xml_pref_area_cd": "eJrtxjZbLnuZHVQUvEWSkxerq",    "xml_sub_area_cd": "hFPOSTqlrJFjhIHsavGEZJeCA",    "xml_city_area_cd": "bgyaktsyJkqVAXEuCxMhKisDO",    "week_yohou_use_flg": "DtSyxEDjchgUkCUcEBTvgfJnM",    "weth_menu_area_cd": "wpWqJFSMxQXOtAgoKWHdLKChi",    "dele_flg": "oqEnOlSEVQqOvSOlBMyqDyrlS",    "upda_dte": "2205-01-19T13:00:31.383275651+09:00",    "upda_user_cd": "VHjgjcDOtbDMkQxpNnGgLQmYS",    "crea_dte": "2089-12-24T22:27:52.51416323+09:00",    "crea_user_cd": "QxnoIbwGLtoLEJAHlDIoAbAEU"}



*/

// WETH_SUB_AREA_MST_ struct is a row record of the WETH_SUB_AREA_MST table in the anpidb database
type WETH_SUB_AREA_MST_ struct {
	//[ 0] WETH_SUB_AREA_CD                               char(4)              null: false  primary: true   isArray: false  auto: false  col: char            len: 4       default: []
	WETHSUBAREACD string `gorm:"primary_key;column:WETH_SUB_AREA_CD;type:char;size:4;"` // 気象細分区域コード
	//[ 1] WETH_SUB_AREA_KIND                             char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WETHSUBAREAKIND string `gorm:"column:WETH_SUB_AREA_KIND;type:char;size:1;"` // 気象細分区域種別
	//[ 2] WETH_SUB_AREA_NMJ                              varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WETHSUBAREANMJ string `gorm:"column:WETH_SUB_AREA_NMJ;type:varchar;size:40;"` // 気象細分区域名
	//[ 3] WETH_SUB_AREA_NMC                              varchar(40)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WETHSUBAREANMC sql.NullString `gorm:"column:WETH_SUB_AREA_NMC;type:varchar;size:40;"` // 気象細分区域名カナ
	//[ 4] WETH_OFFI_CD                                   char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	WETHOFFICD string `gorm:"column:WETH_OFFI_CD;type:char;size:4;"` // 気象官署コード
	//[ 5] WETH_SUB_AREA1_CD                              char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	WETHSUBAREA1CD string `gorm:"column:WETH_SUB_AREA1_CD;type:char;size:4;"` // 気象一次細分区域コード
	//[ 6] XML_PREF_AREA_CD                               char(6)              null: false  primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	XMLPREFAREACD string `gorm:"column:XML_PREF_AREA_CD;type:char;size:6;"` // XML府県予報区コード
	//[ 7] XML_SUB_AREA_CD                                char(6)              null: false  primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	XMLSUBAREACD string `gorm:"column:XML_SUB_AREA_CD;type:char;size:6;"` // XML一次細分区域コード
	//[ 8] XML_CITY_AREA_CD                               char(6)              null: false  primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	XMLCITYAREACD string `gorm:"column:XML_CITY_AREA_CD;type:char;size:6;"` // XML市区町村等地域コード
	//[ 9] WEEK_YOHOU_USE_FLG                             char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WEEKYOHOUUSEFLG string `gorm:"column:WEEK_YOHOU_USE_FLG;type:char;size:1;"` // 週間天気予報使用フラグ
	//[10] WETH_MENU_AREA_CD                              char(3)              null: true   primary: false  isArray: false  auto: false  col: char            len: 3       default: []
	WETHMENUAREACD sql.NullString `gorm:"column:WETH_MENU_AREA_CD;type:char;size:3;"` // 気象メニューエリアコード
	//[11] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"` // UPDA_USER_CD
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[15] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"` // CREA_USER_CD

}

var WETH_SUB_AREA_MSTTableInfo = &TableInfo{
	Name: "WETH_SUB_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "WETH_SUB_AREA_CD",
			Comment:            `気象細分区域コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "WETHSUBAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_sub_area_cd",
			ProtobufFieldName:  "weth_sub_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WETH_SUB_AREA_KIND",
			Comment:            `気象細分区域種別`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WETHSUBAREAKIND",
			GoFieldType:        "string",
			JSONFieldName:      "weth_sub_area_kind",
			ProtobufFieldName:  "weth_sub_area_kind",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "WETH_SUB_AREA_NMJ",
			Comment:            `気象細分区域名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WETHSUBAREANMJ",
			GoFieldType:        "string",
			JSONFieldName:      "weth_sub_area_nmj",
			ProtobufFieldName:  "weth_sub_area_nmj",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "WETH_SUB_AREA_NMC",
			Comment:            `気象細分区域名カナ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WETHSUBAREANMC",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_sub_area_nmc",
			ProtobufFieldName:  "weth_sub_area_nmc",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "WETH_OFFI_CD",
			Comment:            `気象官署コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "WETHOFFICD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_offi_cd",
			ProtobufFieldName:  "weth_offi_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "WETH_SUB_AREA1_CD",
			Comment:            `気象一次細分区域コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "WETHSUBAREA1CD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_sub_area_1_cd",
			ProtobufFieldName:  "weth_sub_area_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "XML_PREF_AREA_CD",
			Comment:            `XML府県予報区コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "XMLPREFAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "xml_pref_area_cd",
			ProtobufFieldName:  "xml_pref_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "XML_SUB_AREA_CD",
			Comment:            `XML一次細分区域コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "XMLSUBAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "xml_sub_area_cd",
			ProtobufFieldName:  "xml_sub_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "XML_CITY_AREA_CD",
			Comment:            `XML市区町村等地域コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "XMLCITYAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "xml_city_area_cd",
			ProtobufFieldName:  "xml_city_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "WEEK_YOHOU_USE_FLG",
			Comment:            `週間天気予報使用フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WEEKYOHOUUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "week_yohou_use_flg",
			ProtobufFieldName:  "week_yohou_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "WETH_MENU_AREA_CD",
			Comment:            `気象メニューエリアコード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHMENUAREACD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_menu_area_cd",
			ProtobufFieldName:  "weth_menu_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "UPDA_USER_CD",
			Comment:            `UPDA_USER_CD`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "CREA_USER_CD",
			Comment:            `CREA_USER_CD`,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_SUB_AREA_MST_) TableName() string {
	return "WETH_SUB_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_SUB_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_SUB_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_SUB_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_SUB_AREA_MST_) TableInfo() *TableInfo {
	return WETH_SUB_AREA_MSTTableInfo
}
