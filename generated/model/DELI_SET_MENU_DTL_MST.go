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


CREATE TABLE `DELI_SET_MENU_DTL_MST` (
  `DELI_SET_MENU_DTL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '配信設定メニュー詳細ID',
  `DELI_SET_MENU_DTL_NM` varchar(20) NOT NULL DEFAULT '' COMMENT '配信設定メニュー詳細名',
  `DELI_SET_MENU_ID` int(11) NOT NULL DEFAULT '0' COMMENT '配信設定メニューID',
  `PROC_PRG_NM` varchar(256) NOT NULL DEFAULT '' COMMENT '処理プログラム名',
  `DISP_DSQ` int(11) NOT NULL DEFAULT '1' COMMENT '表示順',
  `DEF_MENU_AREA_KBN` varchar(4) NOT NULL,
  `DEF_SET_AREA_KBN` varchar(4) DEFAULT NULL,
  `DELE_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`DELI_SET_MENU_DTL_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8 COMMENT='配信設定メニュー詳細マスタ'

JSON Sample
-------------------------------------
{    "deli_set_menu_dtl_id": 91,    "deli_set_menu_dtl_nm": "emlSKwlhrSqhZCbQSwqvIkFPJ",    "deli_set_menu_id": 88,    "proc_prg_nm": "lwPcfShbFmhhICFwxdyUoNjSt",    "disp_dsq": 46,    "def_menu_area_kbn": "PqEhhjZTJEPIumBeTGacjqeOL",    "def_set_area_kbn": "ntdrSCYYFynKneotaokelcqEM",    "dele_flg": "kbVickrwhqmmqsEfGxbEUoWoN",    "upda_dte": "2238-02-26T02:57:13.544746043+09:00",    "upda_user_id": 64,    "crea_dte": "2112-07-01T09:43:22.45930379+09:00",    "crea_user_id": 67}



*/

// DELI_SET_MENU_DTL_MST_ struct is a row record of the DELI_SET_MENU_DTL_MST table in the anpidb database
type DELI_SET_MENU_DTL_MST_ struct {
	//[ 0] DELI_SET_MENU_DTL_ID                           int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	DELISETMENUDTLID int32 `gorm:"primary_key;AUTO_INCREMENT;column:DELI_SET_MENU_DTL_ID;type:int;"` // 配信設定メニュー詳細ID
	//[ 1] DELI_SET_MENU_DTL_NM                           varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	DELISETMENUDTLNM string `gorm:"column:DELI_SET_MENU_DTL_NM;type:varchar;size:20;"` // 配信設定メニュー詳細名
	//[ 2] DELI_SET_MENU_ID                               int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DELISETMENUID int32 `gorm:"column:DELI_SET_MENU_ID;type:int;default:0;"` // 配信設定メニューID
	//[ 3] PROC_PRG_NM                                    varchar(256)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	PROCPRGNM string `gorm:"column:PROC_PRG_NM;type:varchar;size:256;"` // 処理プログラム名
	//[ 4] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;default:1;"` // 表示順
	//[ 5] DEF_MENU_AREA_KBN                              varchar(4)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 4       default: []
	DEFMENUAREAKBN string `gorm:"column:DEF_MENU_AREA_KBN;type:varchar;size:4;"`
	//[ 6] DEF_SET_AREA_KBN                               varchar(4)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 4       default: []
	DEFSETAREAKBN sql.NullString `gorm:"column:DEF_SET_AREA_KBN;type:varchar;size:4;"`
	//[ 7] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var DELI_SET_MENU_DTL_MSTTableInfo = &TableInfo{
	Name: "DELI_SET_MENU_DTL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "DELI_SET_MENU_DTL_ID",
			Comment:            `配信設定メニュー詳細ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DELISETMENUDTLID",
			GoFieldType:        "int32",
			JSONFieldName:      "deli_set_menu_dtl_id",
			ProtobufFieldName:  "deli_set_menu_dtl_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "DELI_SET_MENU_DTL_NM",
			Comment:            `配信設定メニュー詳細名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "DELISETMENUDTLNM",
			GoFieldType:        "string",
			JSONFieldName:      "deli_set_menu_dtl_nm",
			ProtobufFieldName:  "deli_set_menu_dtl_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DELI_SET_MENU_ID",
			Comment:            `配信設定メニューID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DELISETMENUID",
			GoFieldType:        "int32",
			JSONFieldName:      "deli_set_menu_id",
			ProtobufFieldName:  "deli_set_menu_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "PROC_PRG_NM",
			Comment:            `処理プログラム名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "PROCPRGNM",
			GoFieldType:        "string",
			JSONFieldName:      "proc_prg_nm",
			ProtobufFieldName:  "proc_prg_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DEF_MENU_AREA_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       4,
			GoFieldName:        "DEFMENUAREAKBN",
			GoFieldType:        "string",
			JSONFieldName:      "def_menu_area_kbn",
			ProtobufFieldName:  "def_menu_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DEF_SET_AREA_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       4,
			GoFieldName:        "DEFSETAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "def_set_area_kbn",
			ProtobufFieldName:  "def_set_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "UPDA_USER_ID",
			Comment:            `更新者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDAUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "upda_user_id",
			ProtobufFieldName:  "upda_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "CREA_USER_ID",
			Comment:            `作成者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CREAUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "crea_user_id",
			ProtobufFieldName:  "crea_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DELI_SET_MENU_DTL_MST_) TableName() string {
	return "DELI_SET_MENU_DTL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DELI_SET_MENU_DTL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DELI_SET_MENU_DTL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DELI_SET_MENU_DTL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DELI_SET_MENU_DTL_MST_) TableInfo() *TableInfo {
	return DELI_SET_MENU_DTL_MSTTableInfo
}
