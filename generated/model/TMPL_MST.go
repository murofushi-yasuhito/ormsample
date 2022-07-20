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


CREATE TABLE `TMPL_MST` (
  `TMPL_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'テンプレートID',
  `TMPL_KND` char(1) NOT NULL COMMENT 'テンプレート種別	 1:安否確認(自動) 2:安否確認(手動) 3:非常呼集 4:お知らせ',
  `TMPL_NM` varchar(60) NOT NULL COMMENT 'テンプレート名',
  `TMPL_TTL` varchar(100) DEFAULT NULL COMMENT 'メールタイトル',
  `TMPL_TTL_E` varchar(100) DEFAULT NULL COMMENT 'メールタイトル(英)',
  `TMPL_BODY` varchar(6000) DEFAULT NULL COMMENT 'メール本文',
  `TMPL_BODY_E` varchar(6000) DEFAULT NULL COMMENT 'メール本文(英)',
  `VALI_DAYS` int(11) DEFAULT NULL COMMENT '有効日数',
  `QES_TMPL_TYPE_ID` int(11) DEFAULT '0' COMMENT '設問テンプレート種別ID',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`TMPL_ID`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8 COMMENT='テンプレートマスタ'

JSON Sample
-------------------------------------
{    "tmpl_id": 36,    "tmpl_knd": "QZWAVENOfcnLPpjqSuKaHxmJX",    "tmpl_nm": "eMZJndiRcdRXDsmYWDSZfgWdH",    "tmpl_ttl": "xCXqDeoRnZOFSAqFdoRYZocji",    "tmpl_ttl_e": "abPvEcLtmlNhGmvaWjBOOdpAT",    "tmpl_body": "iEOGXnlCCbolOkecMFEScbOWe",    "tmpl_body_e": "QNefLVYbTCXkyJSyLdckwYXvV",    "vali_days": 59,    "qes_tmpl_type_id": 1,    "upda_dte": "2308-09-15T17:29:21.620859087+09:00",    "upda_user_id": 21,    "crea_dte": "2239-05-20T02:54:04.372709016+09:00",    "crea_user_id": 24}



*/

// TMPL_MST_ struct is a row record of the TMPL_MST table in the anpidb database
type TMPL_MST_ struct {
	//[ 0] TMPL_ID                                        int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	TMPLID int32 `gorm:"primary_key;AUTO_INCREMENT;column:TMPL_ID;type:int;"` // テンプレートID
	//[ 1] TMPL_KND                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	TMPLKND string `gorm:"column:TMPL_KND;type:char;size:1;"` // テンプレート種別	 1:安否確認(自動) 2:安否確認(手動) 3:非常呼集 4:お知らせ
	//[ 2] TMPL_NM                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	TMPLNM string `gorm:"column:TMPL_NM;type:varchar;size:60;"` // テンプレート名
	//[ 3] TMPL_TTL                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	TMPLTTL sql.NullString `gorm:"column:TMPL_TTL;type:varchar;size:100;"` // メールタイトル
	//[ 4] TMPL_TTL_E                                     varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	TMPLTTLE sql.NullString `gorm:"column:TMPL_TTL_E;type:varchar;size:100;"` // メールタイトル(英)
	//[ 5] TMPL_BODY                                      varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	TMPLBODY sql.NullString `gorm:"column:TMPL_BODY;type:varchar;size:6000;"` // メール本文
	//[ 6] TMPL_BODY_E                                    varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	TMPLBODYE sql.NullString `gorm:"column:TMPL_BODY_E;type:varchar;size:6000;"` // メール本文(英)
	//[ 7] VALI_DAYS                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	VALIDAYS sql.NullInt64 `gorm:"column:VALI_DAYS;type:int;"` // 有効日数
	//[ 8] QES_TMPL_TYPE_ID                               int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	QESTMPLTYPEID sql.NullInt64 `gorm:"column:QES_TMPL_TYPE_ID;type:int;default:0;"` // 設問テンプレート種別ID
	//[ 9] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[10] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[11] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[12] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var TMPL_MSTTableInfo = &TableInfo{
	Name: "TMPL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "TMPL_ID",
			Comment:            `テンプレートID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TMPLID",
			GoFieldType:        "int32",
			JSONFieldName:      "tmpl_id",
			ProtobufFieldName:  "tmpl_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index: 1,
			Name:  "TMPL_KND",
			Comment: `テンプレート種別	 1:安否確認(自動) 2:安否確認(手動) 3:非常呼集 4:お知らせ`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "TMPL_NM",
			Comment:            `テンプレート名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "TMPLNM",
			GoFieldType:        "string",
			JSONFieldName:      "tmpl_nm",
			ProtobufFieldName:  "tmpl_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "TMPL_TTL",
			Comment:            `メールタイトル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "TMPLTTL",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tmpl_ttl",
			ProtobufFieldName:  "tmpl_ttl",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "TMPL_TTL_E",
			Comment:            `メールタイトル(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "TMPLTTLE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tmpl_ttl_e",
			ProtobufFieldName:  "tmpl_ttl_e",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "TMPL_BODY",
			Comment:            `メール本文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(6000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       6000,
			GoFieldName:        "TMPLBODY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tmpl_body",
			ProtobufFieldName:  "tmpl_body",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "TMPL_BODY_E",
			Comment:            `メール本文(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(6000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       6000,
			GoFieldName:        "TMPLBODYE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "tmpl_body_e",
			ProtobufFieldName:  "tmpl_body_e",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "VALI_DAYS",
			Comment:            `有効日数`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "VALIDAYS",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "vali_days",
			ProtobufFieldName:  "vali_days",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "QES_TMPL_TYPE_ID",
			Comment:            `設問テンプレート種別ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "QESTMPLTYPEID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "qes_tmpl_type_id",
			ProtobufFieldName:  "qes_tmpl_type_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},
	},
}

// TableName sets the insert table name for this struct type
func (t *TMPL_MST_) TableName() string {
	return "TMPL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (t *TMPL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (t *TMPL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (t *TMPL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (t *TMPL_MST_) TableInfo() *TableInfo {
	return TMPL_MSTTableInfo
}
