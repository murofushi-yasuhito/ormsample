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


CREATE TABLE `ANPI_DLCT_UPDA_HIS` (
  `UPDA_HIS_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '更新履歴ID',
  `ANPI_DLCT_ID` int(11) NOT NULL COMMENT '安否起動DLCT_ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `UPDA_KBN` int(1) NOT NULL COMMENT '更新区分	 1:新規登録 2:変更 3:削除',
  `UPDA_TXT` text NOT NULL COMMENT '更新内容',
  `UPDA_USER_KBN` char(1) NOT NULL COMMENT '更新者区分	 U:ユーザ M:クライアントマネージャ A:運用仮者',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`UPDA_HIS_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=1932 DEFAULT CHARSET=utf8 COMMENT='安否起動DLCT更新履歴'

JSON Sample
-------------------------------------
{    "upda_his_id": 35,    "anpi_dlct_id": 90,    "client_id": 90,    "upda_kbn": 61,    "upda_txt": "XywvxgWJJZIXPWrkjCyxixnXI",    "upda_user_kbn": "hDZKAeZtOxCiBvsbNgExclNyS",    "upda_user_id": 48,    "upda_dte": "2260-08-10T03:15:15.04918799+09:00",    "crea_dte": "2039-10-30T02:17:24.851196253+09:00",    "crea_user_id": 64}



*/

// ANPI_DLCT_UPDA_HIS_ struct is a row record of the ANPI_DLCT_UPDA_HIS table in the anpidb database
type ANPI_DLCT_UPDA_HIS_ struct {
	//[ 0] UPDA_HIS_ID                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	UPDAHISID int32 `gorm:"primary_key;AUTO_INCREMENT;column:UPDA_HIS_ID;type:int;"` // 更新履歴ID
	//[ 1] ANPI_DLCT_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANPIDLCTID int32 `gorm:"column:ANPI_DLCT_ID;type:int;"` // 安否起動DLCT_ID
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] UPDA_KBN                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAKBN int32 `gorm:"column:UPDA_KBN;type:int;"` // 更新区分	 1:新規登録 2:変更 3:削除
	//[ 4] UPDA_TXT                                       text(65535)          null: false  primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	UPDATXT string `gorm:"column:UPDA_TXT;type:text;size:65535;"` // 更新内容
	//[ 5] UPDA_USER_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	UPDAUSERKBN string `gorm:"column:UPDA_USER_KBN;type:char;size:1;"` // 更新者区分	 U:ユーザ M:クライアントマネージャ A:運用仮者
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANPI_DLCT_UPDA_HISTableInfo = &TableInfo{
	Name: "ANPI_DLCT_UPDA_HIS",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "UPDA_HIS_ID",
			Comment:            `更新履歴ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDAHISID",
			GoFieldType:        "int32",
			JSONFieldName:      "upda_his_id",
			ProtobufFieldName:  "upda_his_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ANPI_DLCT_ID",
			Comment:            `安否起動DLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPIDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_dlct_id",
			ProtobufFieldName:  "anpi_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "UPDA_KBN",
			Comment: `更新区分	 1:新規登録 2:変更 3:削除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPDAKBN",
			GoFieldType:        "int32",
			JSONFieldName:      "upda_kbn",
			ProtobufFieldName:  "upda_kbn",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "UPDA_TXT",
			Comment:            `更新内容`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "UPDATXT",
			GoFieldType:        "string",
			JSONFieldName:      "upda_txt",
			ProtobufFieldName:  "upda_txt",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "UPDA_USER_KBN",
			Comment: `更新者区分	 U:ユーザ M:クライアントマネージャ A:運用仮者`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "UPDAUSERKBN",
			GoFieldType:        "string",
			JSONFieldName:      "upda_user_kbn",
			ProtobufFieldName:  "upda_user_kbn",
			ProtobufType:       "string",
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANPI_DLCT_UPDA_HIS_) TableName() string {
	return "ANPI_DLCT_UPDA_HIS"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_DLCT_UPDA_HIS_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_DLCT_UPDA_HIS_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_DLCT_UPDA_HIS_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_DLCT_UPDA_HIS_) TableInfo() *TableInfo {
	return ANPI_DLCT_UPDA_HISTableInfo
}
