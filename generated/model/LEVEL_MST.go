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


CREATE TABLE `LEVEL_MST` (
  `LEVEL_ID` int(11) NOT NULL COMMENT 'レベルID',
  `LEVEL_VALUE` int(11) NOT NULL COMMENT 'レベル',
  `LEVEL_NM` varchar(20) NOT NULL COMMENT 'レベル名称',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`LEVEL_ID`,`LEVEL_VALUE`),
  UNIQUE KEY `LEVEL_MST_IX1` (`LEVEL_ID`,`LEVEL_VALUE`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='レベルマスタ'

JSON Sample
-------------------------------------
{    "level_id": 23,    "level_value": 71,    "level_nm": "lseuJZwBocKHHVMsWpYrixCJQ",    "upda_dte": "2100-11-29T18:27:23.475812376+09:00",    "upda_user_id": 3,    "crea_dte": "2305-04-04T03:58:58.646001035+09:00",    "crea_user_id": 82}



*/

// LEVEL_MST_ struct is a row record of the LEVEL_MST table in the anpidb database
type LEVEL_MST_ struct {
	//[ 0] LEVEL_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	LEVELID int32 `gorm:"primary_key;column:LEVEL_ID;type:int;"` // レベルID
	//[ 1] LEVEL_VALUE                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	LEVELVALUE int32 `gorm:"primary_key;column:LEVEL_VALUE;type:int;"` // レベル
	//[ 2] LEVEL_NM                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	LEVELNM string `gorm:"column:LEVEL_NM;type:varchar;size:20;"` // レベル名称
	//[ 3] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 4] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 6] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var LEVEL_MSTTableInfo = &TableInfo{
	Name: "LEVEL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "LEVEL_ID",
			Comment:            `レベルID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELID",
			GoFieldType:        "int32",
			JSONFieldName:      "level_id",
			ProtobufFieldName:  "level_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "LEVEL_VALUE",
			Comment:            `レベル`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELVALUE",
			GoFieldType:        "int32",
			JSONFieldName:      "level_value",
			ProtobufFieldName:  "level_value",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "LEVEL_NM",
			Comment:            `レベル名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "LEVELNM",
			GoFieldType:        "string",
			JSONFieldName:      "level_nm",
			ProtobufFieldName:  "level_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},
	},
}

// TableName sets the insert table name for this struct type
func (l *LEVEL_MST_) TableName() string {
	return "LEVEL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (l *LEVEL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (l *LEVEL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (l *LEVEL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (l *LEVEL_MST_) TableInfo() *TableInfo {
	return LEVEL_MSTTableInfo
}
