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


CREATE TABLE `FAMILY_USER_MST` (
  `FAMILY_USER_ID` int(11) NOT NULL AUTO_INCREMENT,
  `FAMILY_USER_SEQ` int(11) NOT NULL,
  `FAMILY_ID` int(11) NOT NULL,
  `CLIENT_ID` int(11) NOT NULL,
  `USER_ID` int(11) NOT NULL,
  `FAMILY_USER_NM` varchar(60) NOT NULL,
  `LOGIN_ID` varchar(50) NOT NULL,
  `PASSWD` varbinary(140) DEFAULT NULL,
  `MAIL` varchar(266) NOT NULL,
  `CARR_KBN` char(2) DEFAULT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `UPDA_USER_ID` int(11) NOT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `CREA_USER_ID` int(11) NOT NULL,
  PRIMARY KEY (`FAMILY_USER_ID`),
  KEY `FAMILY_USER_MST_IDX_1` (`CLIENT_ID`),
  KEY `FAMILY_USER_MST_IDX2` (`FAMILY_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=352915 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "family_user_id": 55,    "family_user_seq": 35,    "family_id": 79,    "client_id": 15,    "user_id": 7,    "family_user_nm": "STURlssBfJFeaOZvOXhpeYyCM",    "login_id": "QYVyRZgZwQpxPtoMFFVOCUrXv",    "passwd": "KDtbFQ40CRsQDg4rKRFaEUEhNFIoJydBETULTVAdKQM8AkMpI0A8K0YtJx0ZUAJYWTlKAFtGJEkuLDcAL00EJRRZQxgh",    "mail": "RcmPFtthOfCoZYGcQAQAKUPwy",    "carr_kbn": "FZmucTWXJdOFrnSrOiOJduQvD",    "upda_dte": "2174-09-09T00:30:56.079993821+09:00",    "upda_user_id": 52,    "crea_dte": "2030-02-08T12:25:22.827023501+09:00",    "crea_user_id": 33}



*/

// FAMILY_USER_MST struct is a row record of the FAMILY_USER_MST table in the anpidb database
type FAMILY_USER_MST struct {
	//[ 0] FAMILY_USER_ID                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	FAMILYUSERID int32 `gorm:"primary_key;AUTO_INCREMENT;column:FAMILY_USER_ID;type:int;"`
	//[ 1] FAMILY_USER_SEQ                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FAMILYUSERSEQ int32 `gorm:"column:FAMILY_USER_SEQ;type:int;"`
	//[ 2] FAMILY_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	FAMILYID int32 `gorm:"column:FAMILY_ID;type:int;"`
	//[ 3] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"`
	//[ 4] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"`
	//[ 5] FAMILY_USER_NM                                 varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	FAMILYUSERNM string `gorm:"column:FAMILY_USER_NM;type:varchar;size:60;"`
	//[ 6] LOGIN_ID                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LOGINID string `gorm:"column:LOGIN_ID;type:varchar;size:50;"`
	//[ 7] PASSWD                                         varbinary            null: true   primary: false  isArray: false  auto: false  col: varbinary       len: -1      default: []
	PASSWD []byte `gorm:"column:PASSWD;type:varbinary;"`
	//[ 8] MAIL                                           varchar(266)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 266     default: []
	MAIL string `gorm:"column:MAIL;type:varchar;size:266;"`
	//[ 9] CARR_KBN                                       char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	CARRKBN sql.NullString `gorm:"column:CARR_KBN;type:char;size:2;"`
	//[10] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"`
	//[11] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"`
	//[12] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"`
	//[13] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"`
}

var FAMILY_USER_MSTTableInfo = &TableInfo{
	Name: "FAMILY_USER_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "FAMILY_USER_ID",
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
			GoFieldName:        "FAMILYUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "family_user_id",
			ProtobufFieldName:  "family_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "FAMILY_USER_SEQ",
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
			GoFieldName:        "FAMILYUSERSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "family_user_seq",
			ProtobufFieldName:  "family_user_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "FAMILY_USER_NM",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "FAMILYUSERNM",
			GoFieldType:        "string",
			JSONFieldName:      "family_user_nm",
			ProtobufFieldName:  "family_user_nm",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "LOGIN_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       50,
			GoFieldName:        "LOGINID",
			GoFieldType:        "string",
			JSONFieldName:      "login_id",
			ProtobufFieldName:  "login_id",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "PASSWD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varbinary",
			DatabaseTypePretty: "varbinary",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varbinary",
			ColumnLength:       -1,
			GoFieldName:        "PASSWD",
			GoFieldType:        "[]byte",
			JSONFieldName:      "passwd",
			ProtobufFieldName:  "passwd",
			ProtobufType:       "bytes",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "MAIL",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(266)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       266,
			GoFieldName:        "MAIL",
			GoFieldType:        "string",
			JSONFieldName:      "mail",
			ProtobufFieldName:  "mail",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "CARR_KBN",
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
			GoFieldName:        "CARRKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "carr_kbn",
			ProtobufFieldName:  "carr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (f *FAMILY_USER_MST) TableName() string {
	return "FAMILY_USER_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (f *FAMILY_USER_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (f *FAMILY_USER_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (f *FAMILY_USER_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (f *FAMILY_USER_MST) TableInfo() *TableInfo {
	return FAMILY_USER_MSTTableInfo
}
