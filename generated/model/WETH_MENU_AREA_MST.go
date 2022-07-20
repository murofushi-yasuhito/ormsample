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


CREATE TABLE `WETH_MENU_AREA_MST` (
  `WETH_MENU_AREA_CD` char(3) NOT NULL COMMENT '気象メニューエリアコード',
  `WETH_MENU_AREA_LVL` int(1) NOT NULL DEFAULT '0' COMMENT '気象メニューエリアレベル',
  `WETH_MENU_AREA_NM` varchar(255) NOT NULL COMMENT '気象エリアメニュー名',
  `UP_WETH_MENU_AREA_CD` char(3) DEFAULT NULL COMMENT '親エリアコード',
  `PREF_CD` varchar(2) NOT NULL COMMENT '都道府県コード',
  `DISP_DSQ` int(11) DEFAULT NULL COMMENT '表示順',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) NOT NULL COMMENT '作成者コード',
  PRIMARY KEY (`WETH_MENU_AREA_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='気象エリアメニュー名マスタ'

JSON Sample
-------------------------------------
{    "weth_menu_area_cd": "YHgsOHKbsiUIFlVBgtWHDuNyS",    "weth_menu_area_lvl": 33,    "weth_menu_area_nm": "TPIhUyDULQkJsXbmdfqSvFZpF",    "up_weth_menu_area_cd": "AuGehgEsPpBRMqWvWdrWQkROG",    "pref_cd": "WhvFuobsTbLRxeSdSXoubHeoI",    "disp_dsq": 70,    "crea_dte": "2122-12-25T12:35:21.952342854+09:00",    "crea_user_cd": "cRKNxWJsswEGYEYZJwEcqDMvW"}



*/

// WETH_MENU_AREA_MST_ struct is a row record of the WETH_MENU_AREA_MST table in the anpidb database
type WETH_MENU_AREA_MST_ struct {
	//[ 0] WETH_MENU_AREA_CD                              char(3)              null: false  primary: true   isArray: false  auto: false  col: char            len: 3       default: []
	WETHMENUAREACD string `gorm:"primary_key;column:WETH_MENU_AREA_CD;type:char;size:3;"` // 気象メニューエリアコード
	//[ 1] WETH_MENU_AREA_LVL                             int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	WETHMENUAREALVL int32 `gorm:"column:WETH_MENU_AREA_LVL;type:int;default:0;"` // 気象メニューエリアレベル
	//[ 2] WETH_MENU_AREA_NM                              varchar(255)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 255     default: []
	WETHMENUAREANM string `gorm:"column:WETH_MENU_AREA_NM;type:varchar;size:255;"` // 気象エリアメニュー名
	//[ 3] UP_WETH_MENU_AREA_CD                           char(3)              null: true   primary: false  isArray: false  auto: false  col: char            len: 3       default: []
	UPWETHMENUAREACD sql.NullString `gorm:"column:UP_WETH_MENU_AREA_CD;type:char;size:3;"` // 親エリアコード
	//[ 4] PREF_CD                                        varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	PREFCD string `gorm:"column:PREF_CD;type:varchar;size:2;"` // 都道府県コード
	//[ 5] DISP_DSQ                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ sql.NullInt64 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 7] CREA_USER_CD                                   varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD string `gorm:"column:CREA_USER_CD;type:varchar;size:10;"` // 作成者コード

}

var WETH_MENU_AREA_MSTTableInfo = &TableInfo{
	Name: "WETH_MENU_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "WETH_MENU_AREA_CD",
			Comment:            `気象メニューエリアコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "WETHMENUAREACD",
			GoFieldType:        "string",
			JSONFieldName:      "weth_menu_area_cd",
			ProtobufFieldName:  "weth_menu_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WETH_MENU_AREA_LVL",
			Comment:            `気象メニューエリアレベル`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "WETHMENUAREALVL",
			GoFieldType:        "int32",
			JSONFieldName:      "weth_menu_area_lvl",
			ProtobufFieldName:  "weth_menu_area_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "WETH_MENU_AREA_NM",
			Comment:            `気象エリアメニュー名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(255)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       255,
			GoFieldName:        "WETHMENUAREANM",
			GoFieldType:        "string",
			JSONFieldName:      "weth_menu_area_nm",
			ProtobufFieldName:  "weth_menu_area_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "UP_WETH_MENU_AREA_CD",
			Comment:            `親エリアコード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(3)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       3,
			GoFieldName:        "UPWETHMENUAREACD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "up_weth_menu_area_cd",
			ProtobufFieldName:  "up_weth_menu_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "string",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CREA_USER_CD",
			Comment:            `作成者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_MENU_AREA_MST_) TableName() string {
	return "WETH_MENU_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_MENU_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_MENU_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_MENU_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_MENU_AREA_MST_) TableInfo() *TableInfo {
	return WETH_MENU_AREA_MSTTableInfo
}
