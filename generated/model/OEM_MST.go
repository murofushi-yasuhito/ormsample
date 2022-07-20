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


CREATE TABLE `OEM_MST` (
  `OEM_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'OEMID',
  `OEM_NM` varchar(100) NOT NULL COMMENT 'OEM名',
  `DOMAIN_NM` varchar(256) DEFAULT NULL COMMENT 'ドメイン名',
  `HEAD_TXT` varchar(100) DEFAULT NULL COMMENT 'ヘッダテキスト',
  `HEAD_LOGO_FILE` varchar(100) DEFAULT NULL COMMENT 'ヘッダロゴファイル',
  `FOOT_TXT` varchar(100) DEFAULT NULL COMMENT 'フッタテキスト',
  `FOOT_LOGO_FILE` varchar(100) DEFAULT NULL COMMENT 'フッタロゴファイル',
  `CONTRACT_DISP_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '契約内容管理表示区分',
  `SUPPORT_DISP_KBN` char(1) NOT NULL DEFAULT '0' COMMENT 'サポートサイト表示区分（0：表示しない、１：表示する）',
  `BIKO` text COMMENT '備考',
  `UPDA_CMNT` text COMMENT '更新内容',
  `DELE_FLG` char(1) DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`OEM_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=72 DEFAULT CHARSET=utf8 COMMENT='OEMマスタ'

JSON Sample
-------------------------------------
{    "oem_id": 42,    "oem_nm": "VvmFpvdVyFNsnKmFZcmkEdrOR",    "domain_nm": "icEFiMsGpBIpyWIvwZZmgqLEC",    "head_txt": "YiVktcWRKNXkAoVsviOqBJlCe",    "head_logo_file": "JswapTQHNbGijNqxJlwLWtLKK",    "foot_txt": "uBArgxWiHSqCwyrocEFwuukwO",    "foot_logo_file": "eyJraAjIZtwHskTsVXHgXfUkl",    "contract_disp_kbn": "qLuWGlQRilLNKSTNnLGPEythq",    "support_disp_kbn": "DJQdDBMigyqANJCUgSrmflFPl",    "biko": "WEEXvAeQPHrryxFNMpwTjGOYg",    "upda_cmnt": "ndxwSdPgExxHOVPwtaVwuYMIF",    "dele_flg": "FAcIxXSJaWtmtvnobeWJEFOQm",    "upda_dte": "2110-11-20T17:45:23.931594148+09:00",    "upda_user_id": 28,    "crea_dte": "2263-04-03T09:18:15.052262484+09:00",    "crea_user_id": 34}



*/

// OEM_MST_ struct is a row record of the OEM_MST table in the anpidb database
type OEM_MST_ struct {
	//[ 0] OEM_ID                                         int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	OEMID int32 `gorm:"primary_key;AUTO_INCREMENT;column:OEM_ID;type:int;"` // OEMID
	//[ 1] OEM_NM                                         varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	OEMNM string `gorm:"column:OEM_NM;type:varchar;size:100;"` // OEM名
	//[ 2] DOMAIN_NM                                      varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	DOMAINNM sql.NullString `gorm:"column:DOMAIN_NM;type:varchar;size:256;"` // ドメイン名
	//[ 3] HEAD_TXT                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	HEADTXT sql.NullString `gorm:"column:HEAD_TXT;type:varchar;size:100;"` // ヘッダテキスト
	//[ 4] HEAD_LOGO_FILE                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	HEADLOGOFILE sql.NullString `gorm:"column:HEAD_LOGO_FILE;type:varchar;size:100;"` // ヘッダロゴファイル
	//[ 5] FOOT_TXT                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	FOOTTXT sql.NullString `gorm:"column:FOOT_TXT;type:varchar;size:100;"` // フッタテキスト
	//[ 6] FOOT_LOGO_FILE                                 varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	FOOTLOGOFILE sql.NullString `gorm:"column:FOOT_LOGO_FILE;type:varchar;size:100;"` // フッタロゴファイル
	//[ 7] CONTRACT_DISP_KBN                              char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	CONTRACTDISPKBN string `gorm:"column:CONTRACT_DISP_KBN;type:char;size:1;default:0;"` // 契約内容管理表示区分
	//[ 8] SUPPORT_DISP_KBN                               char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	SUPPORTDISPKBN string `gorm:"column:SUPPORT_DISP_KBN;type:char;size:1;default:0;"` // サポートサイト表示区分（0：表示しない、１：表示する）
	//[ 9] BIKO                                           text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	BIKO sql.NullString `gorm:"column:BIKO;type:text;size:65535;"` // 備考
	//[10] UPDA_CMNT                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	UPDACMNT sql.NullString `gorm:"column:UPDA_CMNT;type:text;size:65535;"` // 更新内容
	//[11] DELE_FLG                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG sql.NullString `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var OEM_MSTTableInfo = &TableInfo{
	Name: "OEM_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "OEM_ID",
			Comment:            `OEMID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "OEMID",
			GoFieldType:        "int32",
			JSONFieldName:      "oem_id",
			ProtobufFieldName:  "oem_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "OEM_NM",
			Comment:            `OEM名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "OEMNM",
			GoFieldType:        "string",
			JSONFieldName:      "oem_nm",
			ProtobufFieldName:  "oem_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DOMAIN_NM",
			Comment:            `ドメイン名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "DOMAINNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "domain_nm",
			ProtobufFieldName:  "domain_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "HEAD_TXT",
			Comment:            `ヘッダテキスト`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "HEADTXT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "head_txt",
			ProtobufFieldName:  "head_txt",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "HEAD_LOGO_FILE",
			Comment:            `ヘッダロゴファイル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "HEADLOGOFILE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "head_logo_file",
			ProtobufFieldName:  "head_logo_file",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "FOOT_TXT",
			Comment:            `フッタテキスト`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "FOOTTXT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "foot_txt",
			ProtobufFieldName:  "foot_txt",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "FOOT_LOGO_FILE",
			Comment:            `フッタロゴファイル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "FOOTLOGOFILE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "foot_logo_file",
			ProtobufFieldName:  "foot_logo_file",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CONTRACT_DISP_KBN",
			Comment:            `契約内容管理表示区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CONTRACTDISPKBN",
			GoFieldType:        "string",
			JSONFieldName:      "contract_disp_kbn",
			ProtobufFieldName:  "contract_disp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "SUPPORT_DISP_KBN",
			Comment:            `サポートサイト表示区分（0：表示しない、１：表示する）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "SUPPORTDISPKBN",
			GoFieldType:        "string",
			JSONFieldName:      "support_disp_kbn",
			ProtobufFieldName:  "support_disp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
func (o *OEM_MST_) TableName() string {
	return "OEM_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (o *OEM_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (o *OEM_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (o *OEM_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (o *OEM_MST_) TableInfo() *TableInfo {
	return OEM_MSTTableInfo
}
