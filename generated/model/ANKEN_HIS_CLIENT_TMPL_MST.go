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


CREATE TABLE `ANKEN_HIS_CLIENT_TMPL_MST` (
  `ANKEN_ID` int(11) NOT NULL,
  `CLIENT_TMPL_ID` int(11) NOT NULL COMMENT 'クライアントテンプレートID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `TMPL_KND` char(1) NOT NULL,
  `TMPL_KND_SEQ` int(11) NOT NULL,
  `TMPL_NM` varchar(60) NOT NULL COMMENT 'テンプレート名',
  `TMPL_TTL` varchar(100) DEFAULT NULL COMMENT 'メールタイトル',
  `TMPL_TTL_E` varchar(100) DEFAULT NULL COMMENT 'メールタイトル(英)',
  `TMPL_BODY` varchar(6000) DEFAULT NULL COMMENT 'メール本文',
  `TMPL_BODY_E` varchar(6000) DEFAULT NULL COMMENT 'メール本文(英)',
  `VALI_DAYS` int(11) DEFAULT NULL COMMENT '有効日数',
  `QES_TMPL_TYPE_NM` varchar(256) DEFAULT NULL COMMENT '状況一覧表示名',
  `QES_TMPL_TYPE_NM_E` varchar(256) DEFAULT NULL COMMENT '状況一覧表示名',
  `JOKYO_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '状況一覧表示区分 0:表示しない 1:表示する',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  KEY `CLIENT_TMPL_MST_IDX_1` (`CLIENT_ID`),
  KEY `CLIENT_TMPL_MST_IDX_2` (`ANKEN_ID`) USING BTREE,
  KEY `ANKEN_HIS_CLIENT_TMPL_MST_IDX_1` (`ANKEN_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='クライアントテンプレートマスタ'

JSON Sample
-------------------------------------
{    "anken_id": 43,    "client_tmpl_id": 88,    "client_id": 29,    "tmpl_knd": "nQXnJqbIHJPwRAvZBdCEMKKfV",    "tmpl_knd_seq": 80,    "tmpl_nm": "VYyekauUkEdTDbLwrSOaNphFq",    "tmpl_ttl": "EauEaPVDZBMwDFsMfrEfMBcdo",    "tmpl_ttl_e": "WMDUeCuhOJfSLJZsNLJuLXGHf",    "tmpl_body": "vOXKQQbCsrebevuNsivglwBvI",    "tmpl_body_e": "aTIkSgasAfYdjAmDOqqXUIPua",    "vali_days": 53,    "qes_tmpl_type_nm": "LnHETPqIiVuZnIqnoHYoCyNUJ",    "qes_tmpl_type_nm_e": "QNENrQfWbJaFeLmNfmQLDdXHX",    "jokyo_kbn": "JscGUjoFBeDysmfDeSnburuLJ",    "upda_dte": "2252-09-23T23:33:56.74403329+09:00",    "upda_user_id": 38,    "crea_dte": "2221-10-30T09:08:54.83656936+09:00",    "crea_user_id": 92}


Comments
-------------------------------------
[ 0] Warning table: ANKEN_HIS_CLIENT_TMPL_MST does not have a primary key defined, setting col position 1 ANKEN_ID as primary key




*/

// ANKEN_HIS_CLIENT_TMPL_MST_ struct is a row record of the ANKEN_HIS_CLIENT_TMPL_MST table in the anpidb database
type ANKEN_HIS_CLIENT_TMPL_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"`
	//[ 1] CLIENT_TMPL_ID                                 int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTTMPLID int32 `gorm:"column:CLIENT_TMPL_ID;type:int;"` // クライアントテンプレートID
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] TMPL_KND                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	TMPLKND string `gorm:"column:TMPL_KND;type:char;size:1;"`
	//[ 4] TMPL_KND_SEQ                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TMPLKNDSEQ int32 `gorm:"column:TMPL_KND_SEQ;type:int;"`
	//[ 5] TMPL_NM                                        varchar(60)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	TMPLNM string `gorm:"column:TMPL_NM;type:varchar;size:60;"` // テンプレート名
	//[ 6] TMPL_TTL                                       varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	TMPLTTL sql.NullString `gorm:"column:TMPL_TTL;type:varchar;size:100;"` // メールタイトル
	//[ 7] TMPL_TTL_E                                     varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	TMPLTTLE sql.NullString `gorm:"column:TMPL_TTL_E;type:varchar;size:100;"` // メールタイトル(英)
	//[ 8] TMPL_BODY                                      varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	TMPLBODY sql.NullString `gorm:"column:TMPL_BODY;type:varchar;size:6000;"` // メール本文
	//[ 9] TMPL_BODY_E                                    varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	TMPLBODYE sql.NullString `gorm:"column:TMPL_BODY_E;type:varchar;size:6000;"` // メール本文(英)
	//[10] VALI_DAYS                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	VALIDAYS sql.NullInt64 `gorm:"column:VALI_DAYS;type:int;"` // 有効日数
	//[11] QES_TMPL_TYPE_NM                               varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	QESTMPLTYPENM sql.NullString `gorm:"column:QES_TMPL_TYPE_NM;type:varchar;size:256;"` // 状況一覧表示名
	//[12] QES_TMPL_TYPE_NM_E                             varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	QESTMPLTYPENME sql.NullString `gorm:"column:QES_TMPL_TYPE_NM_E;type:varchar;size:256;"` // 状況一覧表示名
	//[13] JOKYO_KBN                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	JOKYOKBN string `gorm:"column:JOKYO_KBN;type:char;size:1;default:0;"` // 状況一覧表示区分 0:表示しない 1:表示する
	//[14] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[15] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[16] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[17] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_CLIENT_TMPL_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_CLIENT_TMPL_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:   0,
			Name:    "ANKEN_ID",
			Comment: ``,
			Notes: `Warning table: ANKEN_HIS_CLIENT_TMPL_MST does not have a primary key defined, setting col position 1 ANKEN_ID as primary key
`,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "TMPL_KND_SEQ",
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
			GoFieldName:        "TMPLKNDSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "tmpl_knd_seq",
			ProtobufFieldName:  "tmpl_knd_seq",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "QES_TMPL_TYPE_NM",
			Comment:            `状況一覧表示名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "QESTMPLTYPENM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_tmpl_type_nm",
			ProtobufFieldName:  "qes_tmpl_type_nm",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "QES_TMPL_TYPE_NM_E",
			Comment:            `状況一覧表示名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "QESTMPLTYPENME",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "qes_tmpl_type_nm_e",
			ProtobufFieldName:  "qes_tmpl_type_nm_e",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "JOKYO_KBN",
			Comment:            `状況一覧表示区分 0:表示しない 1:表示する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "JOKYOKBN",
			GoFieldType:        "string",
			JSONFieldName:      "jokyo_kbn",
			ProtobufFieldName:  "jokyo_kbn",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_HIS_CLIENT_TMPL_MST_) TableName() string {
	return "ANKEN_HIS_CLIENT_TMPL_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_CLIENT_TMPL_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_CLIENT_TMPL_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_CLIENT_TMPL_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_CLIENT_TMPL_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_CLIENT_TMPL_MSTTableInfo
}
