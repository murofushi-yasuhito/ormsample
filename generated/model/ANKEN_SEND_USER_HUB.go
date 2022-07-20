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


CREATE TABLE `ANKEN_SEND_USER_HUB` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `HUB_ID` int(11) NOT NULL COMMENT '拠点ID',
  `HUB_NM` varchar(60) NOT NULL COMMENT '拠点名',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`USER_ID`,`HUB_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件送信ユーザ拠点'

JSON Sample
-------------------------------------
{    "anken_id": 99,    "user_id": 46,    "hub_id": 77,    "hub_nm": "JxwJSagwOkDFkEhsubsWtsQtZ",    "pref_cd": "KQIkiQUAPVtBQcfEswsmeqhkx",    "city_cd": "tbsRTPUQMMCOlaoaNkjWyIhuH",    "disp_dsq": 44,    "crea_dte": "2177-06-27T01:43:28.357785774+09:00",    "crea_user_id": 85}



*/

// ANKEN_SEND_USER_HUB_ struct is a row record of the ANKEN_SEND_USER_HUB table in the anpidb database
type ANKEN_SEND_USER_HUB_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] USER_ID                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"primary_key;column:USER_ID;type:int;"` // ユーザID
	//[ 2] HUB_ID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	HUBID int32 `gorm:"primary_key;column:HUB_ID;type:int;"` // 拠点ID
	//[ 3] HUB_NM                                         varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	HUBNM string `gorm:"column:HUB_NM;type:varchar;size:60;"` // 拠点名
	//[ 4] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県
	//[ 5] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村
	//[ 6] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_SEND_USER_HUBTableInfo = &TableInfo{
	Name: "ANKEN_SEND_USER_HUB",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "HUB_ID",
			Comment:            `拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "HUBID",
			GoFieldType:        "int32",
			JSONFieldName:      "hub_id",
			ProtobufFieldName:  "hub_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "HUB_NM",
			Comment:            `拠点名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "HUBNM",
			GoFieldType:        "string",
			JSONFieldName:      "hub_nm",
			ProtobufFieldName:  "hub_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "PREF_CD",
			Comment:            `都道府県`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CITY_CD",
			Comment:            `市区町村`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_SEND_USER_HUB_) TableName() string {
	return "ANKEN_SEND_USER_HUB"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_SEND_USER_HUB_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_SEND_USER_HUB_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_SEND_USER_HUB_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_SEND_USER_HUB_) TableInfo() *TableInfo {
	return ANKEN_SEND_USER_HUBTableInfo
}
