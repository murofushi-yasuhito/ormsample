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


CREATE TABLE `ADMIN_MNG_MST` (
  `ADMIN_MNG_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '運用マネージャID',
  `LOGIN_ID` varchar(50) NOT NULL COMMENT 'ログインID',
  `PASSWD` varbinary(140) DEFAULT NULL,
  `ADMIN_MNG_NM` varchar(60) NOT NULL COMMENT '氏名',
  `ADMIN_MNG_AUTH` char(1) NOT NULL DEFAULT 'M' COMMENT '運用マネージャ権限 仕様確認',
  `BIKO` text COMMENT '備考',
  `DELE_FLG` char(1) DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_CMNT` text COMMENT '更新内容',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`ADMIN_MNG_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=605 DEFAULT CHARSET=utf8 COMMENT='運用マネージャマスタ'

JSON Sample
-------------------------------------
{    "admin_mng_id": 99,    "login_id": "GVCkLvaXxqmMaWgjMxlKsxiqA",    "passwd": "FWFUNVobHxIzQRk2HAApHkI8ODseMT5CVjc1HTMKFwFBHVdTKB8/ECReKAdHRkoeJRoqKy4+IAY=",    "admin_mng_nm": "jVkOWXscJFNsXIfFIwBAbZdZb",    "admin_mng_auth": "LAqAKIBqMlGETBYCPbrAaDXCc",    "biko": "SeREwRIQbGTGiVQLJHjVXXnZs",    "dele_flg": "LvyPDEolZcInlcIVTIOwfkmBw",    "upda_cmnt": "JsdwhtMQfMhOoIsnNxErSvQwd",    "upda_dte": "2023-10-17T22:34:14.772042217+09:00",    "upda_user_id": 13,    "crea_dte": "2222-07-17T19:38:39.99151784+09:00",    "crea_user_id": 49}



*/

// ADMIN_MNG_MST struct is a row record of the ADMIN_MNG_MST table in the anpidb database
type ADMIN_MNG_MST struct {
	//[ 0] ADMIN_MNG_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ADMINMNGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ADMIN_MNG_ID;type:int;"` // 運用マネージャID
	//[ 1] LOGIN_ID                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LOGINID string `gorm:"column:LOGIN_ID;type:varchar;size:50;"` // ログインID
	//[ 2] PASSWD                                         varbinary            null: true   primary: false  isArray: false  auto: false  col: varbinary       len: -1      default: []
	PASSWD []byte `gorm:"column:PASSWD;type:varbinary;"`
	//[ 3] ADMIN_MNG_NM                                   varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	ADMINMNGNM string `gorm:"column:ADMIN_MNG_NM;type:varchar;size:60;"` // 氏名
	//[ 4] ADMIN_MNG_AUTH                                 char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [M]
	ADMINMNGAUTH string `gorm:"column:ADMIN_MNG_AUTH;type:char;size:1;default:M;"` // 運用マネージャ権限 仕様確認
	//[ 5] BIKO                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	BIKO sql.NullString `gorm:"column:BIKO;type:text;size:65535;"` // 備考
	//[ 6] DELE_FLG                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG sql.NullString `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[ 7] UPDA_CMNT                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	UPDACMNT sql.NullString `gorm:"column:UPDA_CMNT;type:text;size:65535;"` // 更新内容
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var ADMIN_MNG_MSTTableInfo = &TableInfo{
	Name: "ADMIN_MNG_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ADMIN_MNG_ID",
			Comment:            `運用マネージャID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ADMINMNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "admin_mng_id",
			ProtobufFieldName:  "admin_mng_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "LOGIN_ID",
			Comment:            `ログインID`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ADMIN_MNG_NM",
			Comment:            `氏名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "ADMINMNGNM",
			GoFieldType:        "string",
			JSONFieldName:      "admin_mng_nm",
			ProtobufFieldName:  "admin_mng_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "ADMIN_MNG_AUTH",
			Comment:            `運用マネージャ権限 仕様確認`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ADMINMNGAUTH",
			GoFieldType:        "string",
			JSONFieldName:      "admin_mng_auth",
			ProtobufFieldName:  "admin_mng_auth",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "BIKO",
			Comment:            `備考`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "BIKO",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "biko",
			ProtobufFieldName:  "biko",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "UPDA_CMNT",
			Comment:            `更新内容`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "UPDACMNT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "upda_cmnt",
			ProtobufFieldName:  "upda_cmnt",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ADMIN_MNG_MST) TableName() string {
	return "ADMIN_MNG_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ADMIN_MNG_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ADMIN_MNG_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ADMIN_MNG_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ADMIN_MNG_MST) TableInfo() *TableInfo {
	return ADMIN_MNG_MSTTableInfo
}
