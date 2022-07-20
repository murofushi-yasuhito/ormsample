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


CREATE TABLE `ANKEN_TARGET` (
  `ANKEN_TARGET_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '送信先対象選択ID',
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `TARGET_COND_KBN` char(1) DEFAULT NULL COMMENT '送信先対象区分（1:役職 2:拠点 3:部署 4:地域 5:ユーザ）',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=240 DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 64,    "anken_id": 57,    "seq": 16,    "target_cond_kbn": "WdhgEahaimXybOwlHUNcAcOYb",    "upda_dte": "2080-12-17T18:02:57.384102999+09:00",    "upda_user_id": 93,    "crea_dte": "2260-12-02T21:32:57.957096303+09:00",    "crea_user_id": 62}



*/

// ANKEN_TARGET_ struct is a row record of the ANKEN_TARGET table in the anpidb database
type ANKEN_TARGET_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] ANKEN_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"column:ANKEN_ID;type:int;"` // 案件ID
	//[ 2] SEQ                                            int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"column:SEQ;type:int;"` // 連番
	//[ 3] TARGET_COND_KBN                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	TARGETCONDKBN sql.NullString `gorm:"column:TARGET_COND_KBN;type:char;size:1;"` // 送信先対象区分（1:役職 2:拠点 3:部署 4:地域 5:ユーザ）
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var ANKEN_TARGETTableInfo = &TableInfo{
	Name: "ANKEN_TARGET",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_TARGET_ID",
			Comment:            `送信先対象選択ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENTARGETID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_target_id",
			ProtobufFieldName:  "anken_target_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "SEQ",
			Comment:            `連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "TARGET_COND_KBN",
			Comment:            `送信先対象区分（1:役職 2:拠点 3:部署 4:地域 5:ユーザ）`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "TARGETCONDKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "target_cond_kbn",
			ProtobufFieldName:  "target_cond_kbn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_TARGET_) TableName() string {
	return "ANKEN_TARGET"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_) TableInfo() *TableInfo {
	return ANKEN_TARGETTableInfo
}
