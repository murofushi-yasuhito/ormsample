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


CREATE TABLE `ANKEN_HIS_USER_HUB_MST` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `USER_HUB_ID` int(11) NOT NULL COMMENT 'ユーザ拠点ID',
  `USER_HUB_SEQ` int(11) DEFAULT NULL COMMENT 'ユーザ拠点連番',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `HUB_ID` int(11) NOT NULL COMMENT '拠点ID',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`USER_HUB_ID`),
  KEY `ANKEN_HIS_USER_HUB_MST_IDX1` (`USER_HUB_SEQ`,`USER_ID`),
  KEY `ANKEN_HIS_USER_HUB_MST_IDX2` (`USER_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件履歴ユーザ拠点マスタ'

JSON Sample
-------------------------------------
{    "anken_id": 36,    "user_hub_id": 13,    "user_hub_seq": 46,    "user_id": 10,    "hub_id": 12,    "upda_dte": "2140-07-13T06:44:13.739619712+09:00",    "upda_user_id": 87,    "crea_dte": "2064-05-17T07:53:37.644718712+09:00",    "crea_user_id": 0}



*/

// ANKEN_HIS_USER_HUB_MST_ struct is a row record of the ANKEN_HIS_USER_HUB_MST table in the anpidb database
type ANKEN_HIS_USER_HUB_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] USER_HUB_ID                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERHUBID int32 `gorm:"primary_key;column:USER_HUB_ID;type:int;"` // ユーザ拠点ID
	//[ 2] USER_HUB_SEQ                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERHUBSEQ sql.NullInt64 `gorm:"column:USER_HUB_SEQ;type:int;"` // ユーザ拠点連番
	//[ 3] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 4] HUB_ID                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	HUBID int32 `gorm:"column:HUB_ID;type:int;"` // 拠点ID
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_USER_HUB_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_USER_HUB_MST",
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
			Name:               "USER_HUB_ID",
			Comment:            `ユーザ拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERHUBID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_hub_id",
			ProtobufFieldName:  "user_hub_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "USER_HUB_SEQ",
			Comment:            `ユーザ拠点連番`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERHUBSEQ",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "user_hub_seq",
			ProtobufFieldName:  "user_hub_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
func (a *ANKEN_HIS_USER_HUB_MST_) TableName() string {
	return "ANKEN_HIS_USER_HUB_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_USER_HUB_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_USER_HUB_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_USER_HUB_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_USER_HUB_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_USER_HUB_MSTTableInfo
}
