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


CREATE TABLE `HUB_MST` (
  `HUB_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '拠点ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `HUB_NM` varchar(60) NOT NULL COMMENT '拠点名',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県コード',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村コード',
  `DISP_DSQ` int(11) DEFAULT NULL COMMENT '表示順',
  `BATCH_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '一括処理中フラグ',
  `AUTO_START_FLG` char(1) NOT NULL DEFAULT '1' COMMENT '0:自動起動非対象 1:自動起動対象',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`HUB_ID`),
  KEY `HUB_MST_IDX_1` (`CLIENT_ID`) USING BTREE,
  KEY `BATCH_FLG` (`BATCH_FLG`)
) ENGINE=InnoDB AUTO_INCREMENT=42035 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "hub_id": 65,    "client_id": 16,    "hub_nm": "JMUtkIsdkLNRNdOOLuOLYMeZa",    "pref_cd": "uZWMJAgyAqYBdWXrTOqUwVVHg",    "city_cd": "uQGrCOkAjDqGFCJqvRQcRcTGK",    "disp_dsq": 37,    "batch_flg": "ubBLxoNVPvvmRNALMByCjOlDh",    "auto_start_flg": "iADNdpQlyxQqEfZWrqpXRCgOD",    "upda_dte": "2087-12-24T10:56:53.485070778+09:00",    "upda_user_id": 36,    "crea_dte": "2158-06-09T19:17:57.241819861+09:00",    "crea_user_id": 90}



*/

// HUB_MST_ struct is a row record of the HUB_MST table in the anpidb database
type HUB_MST_ struct {
	//[ 0] HUB_ID                                         int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	HUBID int32 `gorm:"primary_key;AUTO_INCREMENT;column:HUB_ID;type:int;"` // 拠点ID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] HUB_NM                                         varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	HUBNM string `gorm:"column:HUB_NM;type:varchar;size:60;"` // 拠点名
	//[ 3] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[ 4] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[ 5] DISP_DSQ                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ sql.NullInt64 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 6] BATCH_FLG                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	BATCHFLG string `gorm:"column:BATCH_FLG;type:char;size:1;default:0;"` // 一括処理中フラグ
	//[ 7] AUTO_START_FLG                                 char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	AUTOSTARTFLG string `gorm:"column:AUTO_START_FLG;type:char;size:1;default:1;"` // 0:自動起動非対象 1:自動起動対象
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var HUB_MSTTableInfo = &TableInfo{
	Name: "HUB_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "HUB_ID",
			Comment:            `拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "HUBID",
			GoFieldType:        "int32",
			JSONFieldName:      "hub_id",
			ProtobufFieldName:  "hub_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CLIENT_ID",
			Comment:            `クライアントID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
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
			Name:               "BATCH_FLG",
			Comment:            `一括処理中フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "BATCHFLG",
			GoFieldType:        "string",
			JSONFieldName:      "batch_flg",
			ProtobufFieldName:  "batch_flg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "AUTO_START_FLG",
			Comment:            `0:自動起動非対象 1:自動起動対象`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AUTOSTARTFLG",
			GoFieldType:        "string",
			JSONFieldName:      "auto_start_flg",
			ProtobufFieldName:  "auto_start_flg",
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
func (h *HUB_MST_) TableName() string {
	return "HUB_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (h *HUB_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (h *HUB_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (h *HUB_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (h *HUB_MST_) TableInfo() *TableInfo {
	return HUB_MSTTableInfo
}
