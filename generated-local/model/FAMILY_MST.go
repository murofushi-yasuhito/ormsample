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


CREATE TABLE `FAMILY_MST` (
  `FAMILY_ID` int(11) NOT NULL AUTO_INCREMENT,
  `CLIENT_ID` int(11) NOT NULL,
  `USER_ID` int(11) NOT NULL,
  `MAX_USER_NUM` int(11) NOT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `UPDA_USER_ID` int(11) NOT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `CREA_USER_ID` int(11) DEFAULT NULL,
  PRIMARY KEY (`FAMILY_ID`),
  KEY `FAMILY_MST_IDX1` (`USER_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=119832 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "family_id": 49,    "client_id": 34,    "user_id": 54,    "max_user_num": 46,    "upda_dte": "2158-06-23T00:05:05.568593749+09:00",    "upda_user_id": 84,    "crea_dte": "2174-03-13T21:13:29.7162108+09:00",    "crea_user_id": 33}



*/

// FAMILY_MST struct is a row record of the FAMILY_MST table in the anpidb database
type FAMILY_MST struct {
	//[ 0] FAMILY_ID                                      int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	FAMILYID int32 `gorm:"primary_key;AUTO_INCREMENT;column:FAMILY_ID;type:int;"`
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"`
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"`
	//[ 3] MAX_USER_NUM                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	MAXUSERNUM int32 `gorm:"column:MAX_USER_NUM;type:int;"`
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"`
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
	//[ 7] CREA_USER_ID                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID sql.NullInt64 `gorm:"column:CREA_USER_ID;type:int;"`
}

var FAMILY_MSTTableInfo = &TableInfo{
	Name: "FAMILY_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "FAMILY_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "FAMILYID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_id",
			ProtobufFieldName:  "family_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "USER_ID",
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
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "MAX_USER_NUM",
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
			GoFieldName:        "MAXUSERNUM",
			GoFieldType:        "int32",
			JSONFieldName:      "max_user_num",
			ProtobufFieldName:  "max_user_num",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "UPDA_DTE",
			Comment:            ``,
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
			Comment:            ``,
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
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CREAUSERID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "crea_user_id",
			ProtobufFieldName:  "crea_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FAMILY_MST) TableName() string {
	return "FAMILY_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FAMILY_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FAMILY_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FAMILY_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FAMILY_MST) TableInfo() *TableInfo {
	return FAMILY_MSTTableInfo
}
