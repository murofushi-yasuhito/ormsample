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


CREATE TABLE `USER_HUB_MST` (
  `USER_HUB_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザ拠点ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `USER_HUB_SEQ` int(11) DEFAULT NULL COMMENT 'ユーザ拠点連番',
  `HUB_ID` int(11) NOT NULL COMMENT '拠点ID',
  `CLIENT_ID` int(11) NOT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_HUB_ID`),
  KEY `USER_HUB_MST_IDX_2` (`CLIENT_ID`) USING HASH,
  KEY `USER_HUB_MST_IDX_1` (`USER_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7106591 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "user_hub_id": 96,    "user_id": 38,    "user_hub_seq": 41,    "hub_id": 66,    "client_id": 0,    "upda_dte": "2071-06-02T07:57:27.589119007+09:00",    "upda_user_id": 4,    "crea_dte": "2129-10-09T11:19:50.73027925+09:00",    "crea_user_id": 54}



*/

// USER_HUB_MST_ struct is a row record of the USER_HUB_MST table in the anpidb database
type USER_HUB_MST_ struct {
	//[ 0] USER_HUB_ID                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	USERHUBID int32 `gorm:"primary_key;AUTO_INCREMENT;column:USER_HUB_ID;type:int;"` // ユーザ拠点ID
	//[ 1] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 2] USER_HUB_SEQ                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERHUBSEQ sql.NullInt64 `gorm:"column:USER_HUB_SEQ;type:int;"` // ユーザ拠点連番
	//[ 3] HUB_ID                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	HUBID int32 `gorm:"column:HUB_ID;type:int;"` // 拠点ID
	//[ 4] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"`
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_HUB_MSTTableInfo = &TableInfo{
	Name: "USER_HUB_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_HUB_ID",
			Comment:            `ユーザ拠点ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERHUBID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_hub_id",
			ProtobufFieldName:  "user_hub_id",
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CLIENT_ID",
			Comment:            ``,
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
func (u *USER_HUB_MST_) TableName() string {
	return "USER_HUB_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_HUB_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_HUB_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_HUB_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_HUB_MST_) TableInfo() *TableInfo {
	return USER_HUB_MSTTableInfo
}
