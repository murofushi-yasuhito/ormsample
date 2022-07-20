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


CREATE TABLE `FAMILY_BOARD` (
  `FAMILY_BOARD_ID` int(11) NOT NULL AUTO_INCREMENT,
  `FAMILY_ID` int(11) NOT NULL,
  `FAMILY_USER_ID` int(11) NOT NULL,
  `STS_CD` char(2) DEFAULT NULL,
  `MESSAGE` varchar(3000) DEFAULT NULL,
  `REG_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `UPDA_USER_ID` int(11) NOT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `CREA_USER_ID` int(11) NOT NULL,
  PRIMARY KEY (`FAMILY_BOARD_ID`),
  KEY `FAMILY_BOARD_IDX1` (`FAMILY_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=229922 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "family_board_id": 36,    "family_id": 31,    "family_user_id": 67,    "sts_cd": "tiXqqVkWyYTHUevoRVyNkKyCA",    "message": "XwiHeoQeXaBkdIGTaDkXJwbeD",    "reg_dte": "2175-08-06T23:19:58.064200477+09:00",    "upda_dte": "2294-07-05T01:56:05.968972964+09:00",    "upda_user_id": 74,    "crea_dte": "2291-07-31T09:27:27.167704036+09:00",    "crea_user_id": 91}



*/

// FAMILY_BOARD_ struct is a row record of the FAMILY_BOARD table in the anpidb database
type FAMILY_BOARD_ struct {
	//[ 0] FAMILY_BOARD_ID                                int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	FAMILYBOARDID int32 `gorm:"primary_key;AUTO_INCREMENT;column:FAMILY_BOARD_ID;type:int;"`
	//[ 1] FAMILY_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FAMILYID int32 `gorm:"column:FAMILY_ID;type:int;"`
	//[ 2] FAMILY_USER_ID                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FAMILYUSERID int32 `gorm:"column:FAMILY_USER_ID;type:int;"`
	//[ 3] STS_CD                                         char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	STSCD sql.NullString `gorm:"column:STS_CD;type:char;size:2;"`
	//[ 4] MESSAGE                                        varchar(3000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 3000    default: []
	MESSAGE sql.NullString `gorm:"column:MESSAGE;type:varchar;size:3000;"`
	//[ 5] REG_DTE                                        timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	REGDTE time.Time `gorm:"column:REG_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"`
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"`
}

var FAMILY_BOARDTableInfo = &TableInfo{
	Name: "FAMILY_BOARD",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "FAMILY_BOARD_ID",
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
			GoFieldName:        "FAMILYBOARDID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_board_id",
			ProtobufFieldName:  "family_board_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "FAMILY_ID",
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
			GoFieldName:        "FAMILYID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_id",
			ProtobufFieldName:  "family_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "FAMILY_USER_ID",
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
			GoFieldName:        "FAMILYUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_user_id",
			ProtobufFieldName:  "family_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "STS_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "STSCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_cd",
			ProtobufFieldName:  "sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "MESSAGE",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(3000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       3000,
			GoFieldName:        "MESSAGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "message",
			ProtobufFieldName:  "message",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "REG_DTE",
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
			GoFieldName:        "REGDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "reg_dte",
			ProtobufFieldName:  "reg_dte",
			ProtobufType:       "uint64",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CREA_USER_ID",
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
			GoFieldName:        "CREAUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "crea_user_id",
			ProtobufFieldName:  "crea_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FAMILY_BOARD_) TableName() string {
	return "FAMILY_BOARD"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FAMILY_BOARD_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FAMILY_BOARD_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FAMILY_BOARD_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FAMILY_BOARD_) TableInfo() *TableInfo {
	return FAMILY_BOARDTableInfo
}
