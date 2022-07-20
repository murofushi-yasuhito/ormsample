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


CREATE TABLE `CONT_ANPI_HIS` (
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `PROC_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '処理フラグ	 0:未処理 1:処理済',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SID`),
  KEY `CONT_ANPI_HIS_IDX1` (`PROC_FLG`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツ安否処理履歴'

JSON Sample
-------------------------------------
{    "sid": 91,    "proc_flg": "PTOjxpvsEpqsODPBUnyKXLtGg",    "upda_dte": "2189-06-28T22:47:48.708104583+09:00",    "upda_user_id": 14,    "crea_dte": "2077-11-26T16:06:02.563518303+09:00",    "crea_user_id": 55}



*/

// CONT_ANPI_HIS_ struct is a row record of the CONT_ANPI_HIS table in the anpidb database
type CONT_ANPI_HIS_ struct {
	//[ 0] SID                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"primary_key;column:SID;type:int;"` // 情報番号
	//[ 1] PROC_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	PROCFLG string `gorm:"column:PROC_FLG;type:char;size:1;default:0;"` // 処理フラグ	 0:未処理 1:処理済
	//[ 2] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 3] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 4] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 5] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_ANPI_HISTableInfo = &TableInfo{
	Name: "CONT_ANPI_HIS",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "SID",
			Comment:            `情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SID",
			GoFieldType:        "int32",
			JSONFieldName:      "sid",
			ProtobufFieldName:  "sid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index: 1,
			Name:  "PROC_FLG",
			Comment: `処理フラグ	 0:未処理 1:処理済`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PROCFLG",
			GoFieldType:        "string",
			JSONFieldName:      "proc_flg",
			ProtobufFieldName:  "proc_flg",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_ANPI_HIS_) TableName() string {
	return "CONT_ANPI_HIS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_ANPI_HIS_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_ANPI_HIS_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_ANPI_HIS_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_ANPI_HIS_) TableInfo() *TableInfo {
	return CONT_ANPI_HISTableInfo
}
