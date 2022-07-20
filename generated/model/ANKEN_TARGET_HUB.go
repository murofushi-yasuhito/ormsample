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


CREATE TABLE `ANKEN_TARGET_HUB` (
  `ANKEN_TARGET_ID` int(11) NOT NULL COMMENT '送信先対象選択ID',
  `HUB_SEQ` int(11) NOT NULL COMMENT '拠点連番',
  `HUB_ID` int(11) NOT NULL COMMENT '拠点ID',
  `HUB_NM` varchar(60) NOT NULL DEFAULT '' COMMENT '拠点名',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`,`HUB_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 76,    "hub_seq": 92,    "hub_id": 75,    "hub_nm": "dbmEZSvdkqpAsYbXMbFTgXaVd",    "upda_dte": "2238-11-11T16:23:48.350462566+09:00",    "upda_user_id": 16,    "crea_dte": "2249-12-28T20:09:12.781198599+09:00",    "crea_user_id": 29}



*/

// ANKEN_TARGET_HUB_ struct is a row record of the ANKEN_TARGET_HUB table in the anpidb database
type ANKEN_TARGET_HUB_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] HUB_SEQ                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	HUBSEQ int32 `gorm:"primary_key;column:HUB_SEQ;type:int;"` // 拠点連番
	//[ 2] HUB_ID                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	HUBID int32 `gorm:"column:HUB_ID;type:int;"` // 拠点ID
	//[ 3] HUB_NM                                         varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	HUBNM string `gorm:"column:HUB_NM;type:varchar;size:60;"` // 拠点名
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_TARGET_HUBTableInfo = &TableInfo{
	Name: "ANKEN_TARGET_HUB",
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
			IsAutoIncrement:    false,
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
			Name:               "HUB_SEQ",
			Comment:            `拠点連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "HUBSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "hub_seq",
			ProtobufFieldName:  "hub_seq",
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
			IsPrimaryKey:       false,
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
func (a *ANKEN_TARGET_HUB_) TableName() string {
	return "ANKEN_TARGET_HUB"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_HUB_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_HUB_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_HUB_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_HUB_) TableInfo() *TableInfo {
	return ANKEN_TARGET_HUBTableInfo
}
