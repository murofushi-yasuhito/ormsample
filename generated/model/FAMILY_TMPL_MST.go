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


CREATE TABLE `FAMILY_TMPL_MST` (
  `FAMILY_TMPL_ID` int(11) NOT NULL AUTO_INCREMENT,
  `FAMILY_ID` int(11) DEFAULT NULL,
  `TMPL_KND` char(1) NOT NULL,
  `STS_CD` char(1) NOT NULL,
  `STS_MSG` varchar(100) NOT NULL,
  `STS_MSG_E` varchar(100) NOT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `UPDA_USER_ID` int(11) NOT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `CREA_USER_ID` int(11) NOT NULL,
  PRIMARY KEY (`FAMILY_TMPL_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "family_tmpl_id": 50,    "family_id": 76,    "tmpl_knd": "bPGmGnVhmipwyKUoQhImaEmdX",    "sts_cd": "JMUYnGYSlPTbFwUlkcOJOQNeF",    "sts_msg": "bjkcIELsubdvOCGRCMBELfoWE",    "sts_msg_e": "vDXDXIluoxGjYqOuofZLVNjQf",    "upda_dte": "2038-08-02T16:12:29.885807106+09:00",    "upda_user_id": 98,    "crea_dte": "2045-02-08T05:53:58.938789338+09:00",    "crea_user_id": 19}



*/

// FAMILY_TMPL_MST_ struct is a row record of the FAMILY_TMPL_MST table in the anpidb database
type FAMILY_TMPL_MST_ struct {
	//[ 0] FAMILY_TMPL_ID                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	FAMILYTMPLID int32 `gorm:"primary_key;AUTO_INCREMENT;column:FAMILY_TMPL_ID;type:int;"`
	//[ 1] FAMILY_ID                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FAMILYID sql.NullInt64 `gorm:"column:FAMILY_ID;type:int;"`
	//[ 2] TMPL_KND                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	TMPLKND string `gorm:"column:TMPL_KND;type:char;size:1;"`
	//[ 3] STS_CD                                         char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STSCD string `gorm:"column:STS_CD;type:char;size:1;"`
	//[ 4] STS_MSG                                        varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	STSMSG string `gorm:"column:STS_MSG;type:varchar;size:100;"`
	//[ 5] STS_MSG_E                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	STSMSGE string `gorm:"column:STS_MSG_E;type:varchar;size:100;"`
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"`
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"`
}

var FAMILY_TMPL_MSTTableInfo = &TableInfo{
	Name: "FAMILY_TMPL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "FAMILY_TMPL_ID",
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
			GoFieldName:        "FAMILYTMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_tmpl_id",
			ProtobufFieldName:  "family_tmpl_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "FAMILY_ID",
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
			GoFieldName:        "FAMILYID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "family_id",
			ProtobufFieldName:  "family_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "TMPL_KND",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "TMPLKND",
			GoFieldType:        "string",
			JSONFieldName:      "tmpl_knd",
			ProtobufFieldName:  "tmpl_knd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "STS_CD",
			Comment:            ``,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "STS_MSG",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "STSMSG",
			GoFieldType:        "string",
			JSONFieldName:      "sts_msg",
			ProtobufFieldName:  "sts_msg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "STS_MSG_E",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "STSMSGE",
			GoFieldType:        "string",
			JSONFieldName:      "sts_msg_e",
			ProtobufFieldName:  "sts_msg_e",
			ProtobufType:       "string",
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
func (f *FAMILY_TMPL_MST_) TableName() string {
	return "FAMILY_TMPL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FAMILY_TMPL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FAMILY_TMPL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FAMILY_TMPL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FAMILY_TMPL_MST_) TableInfo() *TableInfo {
	return FAMILY_TMPL_MSTTableInfo
}
