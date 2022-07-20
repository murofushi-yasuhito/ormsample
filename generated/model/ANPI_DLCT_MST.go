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


CREATE TABLE `ANPI_DLCT_MST` (
  `ANPI_DLCT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '安否起動DLCT_ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード	 画面選択カテゴリ',
  `LEVEL_VALUE` int(11) DEFAULT '99' COMMENT 'レベル',
  `LEVEL_KBN` char(1) DEFAULT '0' COMMENT 'レベル区分',
  `LEVEL_ID` int(11) DEFAULT NULL COMMENT 'レベルID',
  `DLCT_AREA_KBN` char(4) DEFAULT '0' COMMENT 'DLCTエリア区分',
  `TARG_AREA_EXCP_FLG` char(1) NOT NULL DEFAULT '1' COMMENT '安否起動の際に、指定したエリアを除外するか判断する 0:除外しない 1:除外する',
  `ANPI_DLCT_KBN` char(1) NOT NULL DEFAULT '1' COMMENT '1:ユーザDLCT 2:起動グループ',
  `ANPI_DLCT_CD` varchar(60) DEFAULT NULL COMMENT '安否DLCT_CD',
  `CLIENT_TMPL_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'クライアントテンプレートID',
  `DELE_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANPI_DLCT_ID`),
  KEY `ANPI_DLCT_MST_IDX1` (`CLIENT_ID`,`CATE1_CD`,`LEVEL_VALUE`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=961 DEFAULT CHARSET=utf8 COMMENT='安否起動DLCTマスタ'

JSON Sample
-------------------------------------
{    "anpi_dlct_id": 61,    "client_id": 3,    "cate_1_cd": "gXAmWLNsTwNWCHIRWVkGUFBQW",    "level_value": 31,    "level_kbn": "ayWTeKnrUDHFfZIkrtMNVlWjk",    "level_id": 73,    "dlct_area_kbn": "DCQLOrQJmHDJWhCLHTtpebeiu",    "targ_area_excp_flg": "rZTGOwvrVyFkKtXbCuPNQruxR",    "anpi_dlct_kbn": "ElbHcWEdhUadjSwLRLIAAsYVF",    "anpi_dlct_cd": "ohxKjvshpJkNmfaXPoYxTWnpt",    "client_tmpl_id": 82,    "dele_flg": "IjVflkbHqgPIUeJcNqBGkPFWX",    "upda_dte": "2088-11-04T13:15:00.29784012+09:00",    "upda_user_id": 73,    "crea_dte": "2146-11-01T11:41:47.774408395+09:00",    "crea_user_id": 73}



*/

// ANPI_DLCT_MST struct is a row record of the ANPI_DLCT_MST table in the anpidb database
type ANPI_DLCT_MST struct {
	//[ 0] ANPI_DLCT_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANPIDLCTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANPI_DLCT_ID;type:int;"` // 安否起動DLCT_ID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] CATE1_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード	 画面選択カテゴリ
	//[ 3] LEVEL_VALUE                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [99]
	LEVELVALUE sql.NullInt64 `gorm:"column:LEVEL_VALUE;type:int;default:99;"` // レベル
	//[ 4] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;default:0;"` // レベル区分
	//[ 5] LEVEL_ID                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LEVELID sql.NullInt64 `gorm:"column:LEVEL_ID;type:int;"` // レベルID
	//[ 6] DLCT_AREA_KBN                                  char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	DLCTAREAKBN sql.NullString `gorm:"column:DLCT_AREA_KBN;type:char;size:4;"` // DLCTエリア区分
	//[ 7] TARG_AREA_EXCP_FLG                             char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	TARGAREAEXCPFLG string `gorm:"column:TARG_AREA_EXCP_FLG;type:char;size:1;default:1;"` // 安否起動の際に、指定したエリアを除外するか判断する 0:除外しない 1:除外する
	//[ 8] ANPI_DLCT_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	ANPIDLCTKBN string `gorm:"column:ANPI_DLCT_KBN;type:char;size:1;default:1;"` // 1:ユーザDLCT 2:起動グループ
	//[ 9] ANPI_DLCT_CD                                   varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	ANPIDLCTCD sql.NullString `gorm:"column:ANPI_DLCT_CD;type:varchar;size:60;"` // 安否DLCT_CD
	//[10] CLIENT_TMPL_ID                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CLIENTTMPLID int32 `gorm:"column:CLIENT_TMPL_ID;type:int;default:0;"` // クライアントテンプレートID
	//[11] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANPI_DLCT_MSTTableInfo = &TableInfo{
	Name: "ANPI_DLCT_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANPI_DLCT_ID",
			Comment:            `安否起動DLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPIDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_dlct_id",
			ProtobufFieldName:  "anpi_dlct_id",
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
			Index: 2,
			Name:  "CATE1_CD",
			Comment: `カテゴリ1コード	 画面選択カテゴリ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "LEVEL_VALUE",
			Comment:            `レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELVALUE",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "level_value",
			ProtobufFieldName:  "level_value",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "LEVEL_KBN",
			Comment:            `レベル区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LEVELKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "level_kbn",
			ProtobufFieldName:  "level_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "LEVEL_ID",
			Comment:            `レベルID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "level_id",
			ProtobufFieldName:  "level_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DLCT_AREA_KBN",
			Comment:            `DLCTエリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "DLCTAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dlct_area_kbn",
			ProtobufFieldName:  "dlct_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "TARG_AREA_EXCP_FLG",
			Comment:            `安否起動の際に、指定したエリアを除外するか判断する 0:除外しない 1:除外する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "TARGAREAEXCPFLG",
			GoFieldType:        "string",
			JSONFieldName:      "targ_area_excp_flg",
			ProtobufFieldName:  "targ_area_excp_flg",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "ANPI_DLCT_KBN",
			Comment:            `1:ユーザDLCT 2:起動グループ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ANPIDLCTKBN",
			GoFieldType:        "string",
			JSONFieldName:      "anpi_dlct_kbn",
			ProtobufFieldName:  "anpi_dlct_kbn",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "ANPI_DLCT_CD",
			Comment:            `安否DLCT_CD`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "ANPIDLCTCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anpi_dlct_cd",
			ProtobufFieldName:  "anpi_dlct_cd",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "CLIENT_TMPL_ID",
			Comment:            `クライアントテンプレートID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTTMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_tmpl_id",
			ProtobufFieldName:  "client_tmpl_id",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANPI_DLCT_MST) TableName() string {
	return "ANPI_DLCT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_DLCT_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_DLCT_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_DLCT_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_DLCT_MST) TableInfo() *TableInfo {
	return ANPI_DLCT_MSTTableInfo
}
