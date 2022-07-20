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


CREATE TABLE `WETH_XML_CITY_AREA_MST` (
  `XML_WETH_CITY_CD` char(7) NOT NULL COMMENT 'XML気象市区町村コード',
  `XML_WETH_CITY_NM` varchar(240) NOT NULL COMMENT 'XML気象市区町村名',
  `XML_WETH_SIGN_NM` varchar(60) DEFAULT NULL COMMENT 'XML気象関係表記名',
  `XML_VOLC_SIGN_NM` varchar(60) DEFAULT NULL COMMENT 'XML火山関係表記名	 2014/06追加',
  `XML_SUB_AREA_CD` char(6) DEFAULT NULL COMMENT 'XML一次細分区域コード',
  `XML_CITY_AREA_CD` char(6) DEFAULT NULL COMMENT 'XML市区町村等地域コード',
  `WARN_USE_FLG` char(1) NOT NULL COMMENT '注警報使用フラグ	 0:使用しない、1:使用する',
  `TORNA_USE_FLG` char(1) NOT NULL COMMENT '竜巻注意情報使用フラグ	 0:使用しない、1:使用する',
  `STORM_WARN_USE_FLG` char(1) NOT NULL COMMENT '高潮注警報使用フラグ	 0:使用しない、1:使用する',
  `WAVE_WARN_USE_FLG` char(1) NOT NULL COMMENT '波浪注警報使用フラグ	 0:使用しない、1:使用する',
  `LAND_USE_FLG` char(1) NOT NULL COMMENT '土砂災害情報使用フラグ	 0:使用しない、1:使用する',
  `RIVER_USE_FLG` char(1) NOT NULL COMMENT '河川洪水情報使用フラグ	 0:使用しない、1:使用する',
  `QUAKE_USE_FLG` char(1) NOT NULL COMMENT '地震情報使用フラグ	 0:使用しない、1:使用する',
  `VOLC_USE_FLG` char(1) NOT NULL COMMENT '火山情報使用フラグ	 0:使用しない、1:使用する',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) NOT NULL COMMENT '更新者コード',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) NOT NULL COMMENT '作成者コード',
  PRIMARY KEY (`XML_WETH_CITY_CD`),
  KEY `WETH_XML_CITY_AREA_MST_IDX_1` (`XML_SUB_AREA_CD`,`XML_CITY_AREA_CD`) USING BTREE,
  KEY `WETH_XML_CITY_AREA_MST_IDX_2` (`XML_CITY_AREA_CD`) USING BTREE,
  KEY `WETH_XML_CITY_AREA_MST_IDX_3` (`WARN_USE_FLG`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "xml_weth_city_cd": "NqSaxoxKFdxVemgvMEHDhMaYY",    "xml_weth_city_nm": "wlPPJvqohftjPTyDkqkBBowGq",    "xml_weth_sign_nm": "AiNKQuyWJTppdNgLiHRaeydJB",    "xml_volc_sign_nm": "ZYLIQpoTaNsXCUxBRbdQWuPDj",    "xml_sub_area_cd": "yGPGFbXkEGZVJREVDvWKZtooY",    "xml_city_area_cd": "GUNiSYyTEnTHyJGqlaxUrWvou",    "warn_use_flg": "bhixiDbyhFUROGQrAMVgJCGFo",    "torna_use_flg": "GNuUVqJCQdsKqfwjyKyuChlSD",    "storm_warn_use_flg": "CPoQrLgDjxgVbyaYGaSgjeiJh",    "wave_warn_use_flg": "eBJOuLsiOQRLIWebuoAGlgLMw",    "land_use_flg": "KPSnwcNGfZCTtGOkYOsMoubQN",    "river_use_flg": "LsndVytpEaHqeBGsRpuDWjMas",    "quake_use_flg": "LdOUtYbuwdlMVojdgVSodALeF",    "volc_use_flg": "sarUvreJOnkCuuYYRdUJPuJbD",    "dele_flg": "SaZavyafWSsGCovgdVCqcWxRi",    "upda_dte": "2307-03-03T00:13:33.41408312+09:00",    "upda_user_cd": "odDATSUgWbAvhSjuVOxVsZRZS",    "crea_dte": "2039-02-26T19:02:01.715466919+09:00",    "crea_user_cd": "qmuUaKviHgSsaTjaAKnEEVNla"}



*/

// WETH_XML_CITY_AREA_MST_ struct is a row record of the WETH_XML_CITY_AREA_MST table in the anpidb database
type WETH_XML_CITY_AREA_MST_ struct {
	//[ 0] XML_WETH_CITY_CD                               char(7)              null: false  primary: true   isArray: false  auto: false  col: char            len: 7       default: []
	XMLWETHCITYCD string `gorm:"primary_key;column:XML_WETH_CITY_CD;type:char;size:7;"` // XML気象市区町村コード
	//[ 1] XML_WETH_CITY_NM                               varchar(240)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 240     default: []
	XMLWETHCITYNM string `gorm:"column:XML_WETH_CITY_NM;type:varchar;size:240;"` // XML気象市区町村名
	//[ 2] XML_WETH_SIGN_NM                               varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	XMLWETHSIGNNM sql.NullString `gorm:"column:XML_WETH_SIGN_NM;type:varchar;size:60;"` // XML気象関係表記名
	//[ 3] XML_VOLC_SIGN_NM                               varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	XMLVOLCSIGNNM sql.NullString `gorm:"column:XML_VOLC_SIGN_NM;type:varchar;size:60;"` // XML火山関係表記名	 2014/06追加
	//[ 4] XML_SUB_AREA_CD                                char(6)              null: true   primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	XMLSUBAREACD sql.NullString `gorm:"column:XML_SUB_AREA_CD;type:char;size:6;"` // XML一次細分区域コード
	//[ 5] XML_CITY_AREA_CD                               char(6)              null: true   primary: false  isArray: false  auto: false  col: char            len: 6       default: []
	XMLCITYAREACD sql.NullString `gorm:"column:XML_CITY_AREA_CD;type:char;size:6;"` // XML市区町村等地域コード
	//[ 6] WARN_USE_FLG                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WARNUSEFLG string `gorm:"column:WARN_USE_FLG;type:char;size:1;"` // 注警報使用フラグ	 0:使用しない、1:使用する
	//[ 7] TORNA_USE_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	TORNAUSEFLG string `gorm:"column:TORNA_USE_FLG;type:char;size:1;"` // 竜巻注意情報使用フラグ	 0:使用しない、1:使用する
	//[ 8] STORM_WARN_USE_FLG                             char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STORMWARNUSEFLG string `gorm:"column:STORM_WARN_USE_FLG;type:char;size:1;"` // 高潮注警報使用フラグ	 0:使用しない、1:使用する
	//[ 9] WAVE_WARN_USE_FLG                              char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WAVEWARNUSEFLG string `gorm:"column:WAVE_WARN_USE_FLG;type:char;size:1;"` // 波浪注警報使用フラグ	 0:使用しない、1:使用する
	//[10] LAND_USE_FLG                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	LANDUSEFLG string `gorm:"column:LAND_USE_FLG;type:char;size:1;"` // 土砂災害情報使用フラグ	 0:使用しない、1:使用する
	//[11] RIVER_USE_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	RIVERUSEFLG string `gorm:"column:RIVER_USE_FLG;type:char;size:1;"` // 河川洪水情報使用フラグ	 0:使用しない、1:使用する
	//[12] QUAKE_USE_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	QUAKEUSEFLG string `gorm:"column:QUAKE_USE_FLG;type:char;size:1;"` // 地震情報使用フラグ	 0:使用しない、1:使用する
	//[13] VOLC_USE_FLG                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	VOLCUSEFLG string `gorm:"column:VOLC_USE_FLG;type:char;size:1;"` // 火山情報使用フラグ	 0:使用しない、1:使用する
	//[14] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[15] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[16] UPDA_USER_CD                                   varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD string `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"` // 更新者コード
	//[17] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[18] CREA_USER_CD                                   varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD string `gorm:"column:CREA_USER_CD;type:varchar;size:10;"` // 作成者コード

}

var WETH_XML_CITY_AREA_MSTTableInfo = &TableInfo{
	Name: "WETH_XML_CITY_AREA_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "XML_WETH_CITY_CD",
			Comment:            `XML気象市区町村コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(7)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       7,
			GoFieldName:        "XMLWETHCITYCD",
			GoFieldType:        "string",
			JSONFieldName:      "xml_weth_city_cd",
			ProtobufFieldName:  "xml_weth_city_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "XML_WETH_CITY_NM",
			Comment:            `XML気象市区町村名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(240)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       240,
			GoFieldName:        "XMLWETHCITYNM",
			GoFieldType:        "string",
			JSONFieldName:      "xml_weth_city_nm",
			ProtobufFieldName:  "xml_weth_city_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "XML_WETH_SIGN_NM",
			Comment:            `XML気象関係表記名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "XMLWETHSIGNNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "xml_weth_sign_nm",
			ProtobufFieldName:  "xml_weth_sign_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "XML_VOLC_SIGN_NM",
			Comment: `XML火山関係表記名	 2014/06追加`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "XMLVOLCSIGNNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "xml_volc_sign_nm",
			ProtobufFieldName:  "xml_volc_sign_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "XML_SUB_AREA_CD",
			Comment:            `XML一次細分区域コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "XMLSUBAREACD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "xml_sub_area_cd",
			ProtobufFieldName:  "xml_sub_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "XML_CITY_AREA_CD",
			Comment:            `XML市区町村等地域コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(6)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       6,
			GoFieldName:        "XMLCITYAREACD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "xml_city_area_cd",
			ProtobufFieldName:  "xml_city_area_cd",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index: 6,
			Name:  "WARN_USE_FLG",
			Comment: `注警報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WARNUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "warn_use_flg",
			ProtobufFieldName:  "warn_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index: 7,
			Name:  "TORNA_USE_FLG",
			Comment: `竜巻注意情報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "TORNAUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "torna_use_flg",
			ProtobufFieldName:  "torna_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index: 8,
			Name:  "STORM_WARN_USE_FLG",
			Comment: `高潮注警報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STORMWARNUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "storm_warn_use_flg",
			ProtobufFieldName:  "storm_warn_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
			Name:  "WAVE_WARN_USE_FLG",
			Comment: `波浪注警報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WAVEWARNUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "wave_warn_use_flg",
			ProtobufFieldName:  "wave_warn_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index: 10,
			Name:  "LAND_USE_FLG",
			Comment: `土砂災害情報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "LANDUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "land_use_flg",
			ProtobufFieldName:  "land_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index: 11,
			Name:  "RIVER_USE_FLG",
			Comment: `河川洪水情報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "RIVERUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "river_use_flg",
			ProtobufFieldName:  "river_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index: 12,
			Name:  "QUAKE_USE_FLG",
			Comment: `地震情報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "QUAKEUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "quake_use_flg",
			ProtobufFieldName:  "quake_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index: 13,
			Name:  "VOLC_USE_FLG",
			Comment: `火山情報使用フラグ	 0:使用しない、1:使用する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "VOLCUSEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "volc_use_flg",
			ProtobufFieldName:  "volc_use_flg",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "UPDA_USER_CD",
			Comment:            `更新者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "CREA_USER_CD",
			Comment:            `作成者コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "string",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_XML_CITY_AREA_MST_) TableName() string {
	return "WETH_XML_CITY_AREA_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_XML_CITY_AREA_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_XML_CITY_AREA_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_XML_CITY_AREA_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_XML_CITY_AREA_MST_) TableInfo() *TableInfo {
	return WETH_XML_CITY_AREA_MSTTableInfo
}
