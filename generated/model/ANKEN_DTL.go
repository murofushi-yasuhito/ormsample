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


CREATE TABLE `ANKEN_DTL` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `ANKEN_DTL_ID` int(11) NOT NULL COMMENT '案件詳細ID',
  `STS_CD` char(1) NOT NULL COMMENT '状態コード',
  `STS_MSG` varchar(200) NOT NULL COMMENT '状態メッセージ',
  `STS_MSG_E` varchar(200) DEFAULT NULL COMMENT '状態メッセージ(英)',
  `RES_MSG` varchar(200) DEFAULT NULL COMMENT '選択後メッセージ',
  `RES_MSG_E` varchar(200) DEFAULT NULL COMMENT '選択後メッセージ(英)',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`ANKEN_DTL_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件詳細'

JSON Sample
-------------------------------------
{    "anken_id": 80,    "anken_dtl_id": 55,    "sts_cd": "ZxANGZcWULEIZJVgJMNyDpmeb",    "sts_msg": "cFxtOTDDKDZVJtPGSnOmSGSKY",    "sts_msg_e": "UgUBheJURnIVlpcskwuyXnFYZ",    "res_msg": "uQKPOvIHDYDXtQOfhGmHwHcAe",    "res_msg_e": "nekkXrvyOfOmtGIGmuRJqnCQb",    "crea_dte": "2072-09-11T19:34:00.126796086+09:00",    "crea_user_id": 80}



*/

// ANKEN_DTL_ struct is a row record of the ANKEN_DTL table in the anpidb database
type ANKEN_DTL_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] ANKEN_DTL_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENDTLID int32 `gorm:"primary_key;column:ANKEN_DTL_ID;type:int;"` // 案件詳細ID
	//[ 2] STS_CD                                         char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STSCD string `gorm:"column:STS_CD;type:char;size:1;"` // 状態コード
	//[ 3] STS_MSG                                        varchar(200)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	STSMSG string `gorm:"column:STS_MSG;type:varchar;size:200;"` // 状態メッセージ
	//[ 4] STS_MSG_E                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	STSMSGE sql.NullString `gorm:"column:STS_MSG_E;type:varchar;size:200;"` // 状態メッセージ(英)
	//[ 5] RES_MSG                                        varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	RESMSG sql.NullString `gorm:"column:RES_MSG;type:varchar;size:200;"` // 選択後メッセージ
	//[ 6] RES_MSG_E                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	RESMSGE sql.NullString `gorm:"column:RES_MSG_E;type:varchar;size:200;"` // 選択後メッセージ(英)
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_DTLTableInfo = &TableInfo{
	Name: "ANKEN_DTL",
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
			Name:               "ANKEN_DTL_ID",
			Comment:            `案件詳細ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENDTLID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_dtl_id",
			ProtobufFieldName:  "anken_dtl_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "STS_CD",
			Comment:            `状態コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STSCD",
			GoFieldType:        "string",
			JSONFieldName:      "sts_cd",
			ProtobufFieldName:  "sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "STS_MSG",
			Comment:            `状態メッセージ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "STSMSG",
			GoFieldType:        "string",
			JSONFieldName:      "sts_msg",
			ProtobufFieldName:  "sts_msg",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "STS_MSG_E",
			Comment:            `状態メッセージ(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "STSMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_msg_e",
			ProtobufFieldName:  "sts_msg_e",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "RES_MSG",
			Comment:            `選択後メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "RESMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg",
			ProtobufFieldName:  "res_msg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "RES_MSG_E",
			Comment:            `選択後メッセージ(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "RESMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg_e",
			ProtobufFieldName:  "res_msg_e",
			ProtobufType:       "string",
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
func (a *ANKEN_DTL_) TableName() string {
	return "ANKEN_DTL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_DTL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_DTL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_DTL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_DTL_) TableInfo() *TableInfo {
	return ANKEN_DTLTableInfo
}
