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


CREATE TABLE `CLIENT_MNG_MST` (
  `CLIENT_MNG_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'クライアントマネージャID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `LOGIN_ID` varchar(50) NOT NULL COMMENT 'ログインID',
  `PASSWD` varbinary(140) DEFAULT NULL,
  `CLIENT_MNG_NM` varchar(60) NOT NULL COMMENT '氏名',
  `CLIENT_MNG_AUTH` char(2) NOT NULL DEFAULT '4' COMMENT 'クライアントマネージャ権限	 仮）権限検討',
  `DELE_FLG` char(1) DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`CLIENT_MNG_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=537 DEFAULT CHARSET=utf8 COMMENT='クライアントマネージャマスタ'

JSON Sample
-------------------------------------
{    "client_mng_id": 99,    "client_id": 11,    "login_id": "fQFxiPvHFFpdDAMuciRvVFMsY",    "passwd": "FysmFw==",    "client_mng_nm": "aoSynouCnwCOJrlxcKtFJYibU",    "client_mng_auth": "rAyMQNuBoVjVKHtnWSocYSdjc",    "dele_flg": "svtkIJdFclMlnkHxEobhPWMDt",    "upda_dte": "2272-07-01T05:32:03.945040073+09:00",    "upda_user_id": 92,    "crea_dte": "2231-10-05T06:55:42.631534519+09:00",    "crea_user_id": 39}



*/

// CLIENT_MNG_MST_ struct is a row record of the CLIENT_MNG_MST table in the anpidb database
type CLIENT_MNG_MST_ struct {
	//[ 0] CLIENT_MNG_ID                                  int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	CLIENTMNGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:CLIENT_MNG_ID;type:int;"` // クライアントマネージャID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] LOGIN_ID                                       varchar(50)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 50      default: []
	LOGINID string `gorm:"column:LOGIN_ID;type:varchar;size:50;"` // ログインID
	//[ 3] PASSWD                                         varbinary            null: true   primary: false  isArray: false  auto: false  col: varbinary       len: -1      default: []
	PASSWD []byte `gorm:"column:PASSWD;type:varbinary;"`
	//[ 4] CLIENT_MNG_NM                                  varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	CLIENTMNGNM string `gorm:"column:CLIENT_MNG_NM;type:varchar;size:60;"` // 氏名
	//[ 5] CLIENT_MNG_AUTH                                char(2)              null: false  primary: false  isArray: false  auto: false  col: char            len: 2       default: [4]
	CLIENTMNGAUTH string `gorm:"column:CLIENT_MNG_AUTH;type:char;size:2;default:4;"` // クライアントマネージャ権限	 仮）権限検討
	//[ 6] DELE_FLG                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG sql.NullString `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[10] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CLIENT_MNG_MSTTableInfo = &TableInfo{
	Name: "CLIENT_MNG_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CLIENT_MNG_ID",
			Comment:            `クライアントマネージャID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTMNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_mng_id",
			ProtobufFieldName:  "client_mng_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CLIENT_ID",
			Comment:            `クライアントID`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CLIENT_MNG_NM",
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
			GoFieldName:        "CLIENTMNGNM",
			GoFieldType:        "string",
			JSONFieldName:      "client_mng_nm",
			ProtobufFieldName:  "client_mng_nm",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "CLIENT_MNG_AUTH",
			Comment: `クライアントマネージャ権限	 仮）権限検討`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "CLIENTMNGAUTH",
			GoFieldType:        "string",
			JSONFieldName:      "client_mng_auth",
			ProtobufFieldName:  "client_mng_auth",
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CLIENT_MNG_MST_) TableName() string {
	return "CLIENT_MNG_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CLIENT_MNG_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CLIENT_MNG_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CLIENT_MNG_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CLIENT_MNG_MST_) TableInfo() *TableInfo {
	return CLIENT_MNG_MSTTableInfo
}
