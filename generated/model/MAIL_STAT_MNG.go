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


CREATE TABLE `MAIL_STAT_MNG` (
  `MAIL_STAT_MNG_ID` int(11) NOT NULL AUTO_INCREMENT,
  `RECORD_KBN` int(1) NOT NULL,
  `PROC_STATUS` int(1) DEFAULT NULL,
  `LOG_LAST_LINE` varchar(1024) DEFAULT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`MAIL_STAT_MNG_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "mail_stat_mng_id": 82,    "record_kbn": 13,    "proc_status": 91,    "log_last_line": "RItBCcRTTnrBYKZQmKOGhQEft",    "upda_dte": "2067-09-11T08:55:08.169657747+09:00",    "crea_dte": "2048-07-15T13:56:10.671313015+09:00"}



*/

// MAIL_STAT_MNG_ struct is a row record of the MAIL_STAT_MNG table in the anpidb database
type MAIL_STAT_MNG_ struct {
	//[ 0] MAIL_STAT_MNG_ID                               int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MAILSTATMNGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MAIL_STAT_MNG_ID;type:int;"`
	//[ 1] RECORD_KBN                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	RECORDKBN int32 `gorm:"column:RECORD_KBN;type:int;"`
	//[ 2] PROC_STATUS                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	PROCSTATUS sql.NullInt64 `gorm:"column:PROC_STATUS;type:int;"`
	//[ 3] LOG_LAST_LINE                                  varchar(1024)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 1024    default: []
	LOGLASTLINE sql.NullString `gorm:"column:LOG_LAST_LINE;type:varchar;size:1024;"`
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[ 5] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
}

var MAIL_STAT_MNGTableInfo = &TableInfo{
	Name: "MAIL_STAT_MNG",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MAIL_STAT_MNG_ID",
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
			GoFieldName:        "MAILSTATMNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_stat_mng_id",
			ProtobufFieldName:  "mail_stat_mng_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "RECORD_KBN",
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
			GoFieldName:        "RECORDKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "record_kbn",
			ProtobufFieldName:  "record_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "PROC_STATUS",
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
			GoFieldName:        "PROCSTATUS",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "proc_status",
			ProtobufFieldName:  "proc_status",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "LOG_LAST_LINE",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(1024)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       1024,
			GoFieldName:        "LOGLASTLINE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "log_last_line",
			ProtobufFieldName:  "log_last_line",
			ProtobufType:       "string",
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
			ProtobufPos:        6,
		},
	},
}

// TableName sets the insert table name for this struct type
func (m *MAIL_STAT_MNG_) TableName() string {
	return "MAIL_STAT_MNG"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MAIL_STAT_MNG_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MAIL_STAT_MNG_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MAIL_STAT_MNG_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MAIL_STAT_MNG_) TableInfo() *TableInfo {
	return MAIL_STAT_MNGTableInfo
}
