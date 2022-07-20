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


CREATE TABLE `CATEGORY1_MST` (
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード',
  `CATE1_NM` varchar(20) NOT NULL COMMENT 'カテゴリ1名称',
  `DISP_DSQ` int(11) DEFAULT NULL COMMENT '表示順',
  `LEVEL_ID` int(11) DEFAULT NULL COMMENT 'レベルID',
  `LEVEL_KBN` char(1) DEFAULT NULL COMMENT 'レベル区分	 C:コンテンツレベル A:エリアレベル',
  `ANPI_FLG` char(1) DEFAULT '0' COMMENT '安否確認利用フラグ    1:利用する 0:利用しない',
  `SETT_AREA_KBN` char(4) DEFAULT '0' COMMENT '配信設定エリア区分',
  `ANPI_AREA_KBN` char(4) DEFAULT '0' COMMENT '安否設定エリア区分',
  `CONT_AREA_KBN` char(4) DEFAULT '0' COMMENT 'コンテンツエリア区分',
  `DELI_SET_DISP_ALL_KBN` char(1) DEFAULT NULL COMMENT '日本全国の表示区分',
  `DELI_SET_DISP_REGION_KBN` char(1) DEFAULT NULL COMMENT '地域の表示非表示',
  `DELI_SET_DISP_PREF_KBN` char(1) DEFAULT NULL COMMENT '都道府県の表示非表示',
  `DELI_SET_DISP_CITY_KBN` char(1) DEFAULT NULL COMMENT '市区町村の表示非表示',
  `DELI_SET_AREA_KBN` char(4) DEFAULT NULL COMMENT 'USER_DLCT_DのAREA_KBNに設定するAREA_KBN',
  `CATE_NOTE` varchar(100) DEFAULT NULL COMMENT 'カテゴリ説明',
  `VALID_KBN` char(1) DEFAULT '1' COMMENT '有効区分',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`CATE1_CD`),
  UNIQUE KEY `CATEGORY1_MST_PKI` (`CATE1_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='カテゴリ1マスタ'

JSON Sample
-------------------------------------
{    "cate_1_cd": "sXMiAQeSiCgkfHOpsCpHWREZD",    "cate_1_nm": "RTTMqHmDQJAgJFEWbvInJEymj",    "disp_dsq": 54,    "level_id": 37,    "level_kbn": "mxYkPKMgCFUeYmJGXVVePyXCl",    "anpi_flg": "HBakOXaXdWPRkeigUVwjAmrgU",    "sett_area_kbn": "JkaXdYkydYAgkykZMErcKWBnw",    "anpi_area_kbn": "NiOibQMQWhBIIgVyqwlMXQUwF",    "cont_area_kbn": "vxlFdaYKxukyFRIVFZsSDHQnQ",    "deli_set_disp_all_kbn": "eJgfJIKQrilWshDwFkcllJwyX",    "deli_set_disp_region_kbn": "YLZGIDWrxBoPnKQwsbKdPcQuU",    "deli_set_disp_pref_kbn": "AEPokUandqKxrhudgFaOLJGvN",    "deli_set_disp_city_kbn": "WZZMlIUiyoYANwKhVBsMggWTD",    "deli_set_area_kbn": "oVxcUhZOPsuLvJkDLJwcqluRR",    "cate_note": "ZuRSMqLjVNXKOXpnpNSgtqaBm",    "valid_kbn": "YtYXOjnEnWttrgogbNDfmMEEx",    "upda_dte": "2287-09-23T21:45:55.056797865+09:00",    "upda_user_id": 93,    "crea_dte": "2285-03-03T13:51:27.522601278+09:00",    "crea_user_id": 6}



*/

// CATEGORY1_MST_ struct is a row record of the CATEGORY1_MST table in the anpidb database
type CATEGORY1_MST_ struct {
	//[ 0] CATE1_CD                                       varchar(10)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"primary_key;column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード
	//[ 1] CATE1_NM                                       varchar(20)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 20      default: []
	CATE1NM string `gorm:"column:CATE1_NM;type:varchar;size:20;"` // カテゴリ1名称
	//[ 2] DISP_DSQ                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ sql.NullInt64 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 3] LEVEL_ID                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LEVELID sql.NullInt64 `gorm:"column:LEVEL_ID;type:int;"` // レベルID
	//[ 4] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;"` // レベル区分	 C:コンテンツレベル A:エリアレベル
	//[ 5] ANPI_FLG                                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	ANPIFLG sql.NullString `gorm:"column:ANPI_FLG;type:char;size:1;default:0;"` // 安否確認利用フラグ    1:利用する 0:利用しない
	//[ 6] SETT_AREA_KBN                                  char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: [0]
	SETTAREAKBN sql.NullString `gorm:"column:SETT_AREA_KBN;type:char;size:4;default:0;"` // 配信設定エリア区分
	//[ 7] ANPI_AREA_KBN                                  char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: [0]
	ANPIAREAKBN sql.NullString `gorm:"column:ANPI_AREA_KBN;type:char;size:4;default:0;"` // 安否設定エリア区分
	//[ 8] CONT_AREA_KBN                                  char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: [0]
	CONTAREAKBN sql.NullString `gorm:"column:CONT_AREA_KBN;type:char;size:4;default:0;"` // コンテンツエリア区分
	//[ 9] DELI_SET_DISP_ALL_KBN                          char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELISETDISPALLKBN sql.NullString `gorm:"column:DELI_SET_DISP_ALL_KBN;type:char;size:1;"` // 日本全国の表示区分
	//[10] DELI_SET_DISP_REGION_KBN                       char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELISETDISPREGIONKBN sql.NullString `gorm:"column:DELI_SET_DISP_REGION_KBN;type:char;size:1;"` // 地域の表示非表示
	//[11] DELI_SET_DISP_PREF_KBN                         char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELISETDISPPREFKBN sql.NullString `gorm:"column:DELI_SET_DISP_PREF_KBN;type:char;size:1;"` // 都道府県の表示非表示
	//[12] DELI_SET_DISP_CITY_KBN                         char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELISETDISPCITYKBN sql.NullString `gorm:"column:DELI_SET_DISP_CITY_KBN;type:char;size:1;"` // 市区町村の表示非表示
	//[13] DELI_SET_AREA_KBN                              char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	DELISETAREAKBN sql.NullString `gorm:"column:DELI_SET_AREA_KBN;type:char;size:4;"` // USER_DLCT_DのAREA_KBNに設定するAREA_KBN
	//[14] CATE_NOTE                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	CATENOTE sql.NullString `gorm:"column:CATE_NOTE;type:varchar;size:100;"` // カテゴリ説明
	//[15] VALID_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	VALIDKBN sql.NullString `gorm:"column:VALID_KBN;type:char;size:1;default:1;"` // 有効区分
	//[16] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[17] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[18] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[19] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var CATEGORY1_MSTTableInfo = &TableInfo{
	Name: "CATEGORY1_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CATE1_NM",
			Comment:            `カテゴリ1名称`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(20)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       20,
			GoFieldName:        "CATE1NM",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_nm",
			ProtobufFieldName:  "cate_1_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "LEVEL_KBN",
			Comment: `レベル区分	 C:コンテンツレベル A:エリアレベル`,
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
			Name:               "ANPI_FLG",
			Comment:            `安否確認利用フラグ    1:利用する 0:利用しない`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ANPIFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anpi_flg",
			ProtobufFieldName:  "anpi_flg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "SETT_AREA_KBN",
			Comment:            `配信設定エリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "SETTAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sett_area_kbn",
			ProtobufFieldName:  "sett_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "ANPI_AREA_KBN",
			Comment:            `安否設定エリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "ANPIAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anpi_area_kbn",
			ProtobufFieldName:  "anpi_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CONT_AREA_KBN",
			Comment:            `コンテンツエリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "CONTAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cont_area_kbn",
			ProtobufFieldName:  "cont_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "DELI_SET_DISP_ALL_KBN",
			Comment:            `日本全国の表示区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELISETDISPALLKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_set_disp_all_kbn",
			ProtobufFieldName:  "deli_set_disp_all_kbn",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "DELI_SET_DISP_REGION_KBN",
			Comment:            `地域の表示非表示`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELISETDISPREGIONKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_set_disp_region_kbn",
			ProtobufFieldName:  "deli_set_disp_region_kbn",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "DELI_SET_DISP_PREF_KBN",
			Comment:            `都道府県の表示非表示`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELISETDISPPREFKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_set_disp_pref_kbn",
			ProtobufFieldName:  "deli_set_disp_pref_kbn",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "DELI_SET_DISP_CITY_KBN",
			Comment:            `市区町村の表示非表示`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELISETDISPCITYKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_set_disp_city_kbn",
			ProtobufFieldName:  "deli_set_disp_city_kbn",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "DELI_SET_AREA_KBN",
			Comment:            `USER_DLCT_DのAREA_KBNに設定するAREA_KBN`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "DELISETAREAKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_set_area_kbn",
			ProtobufFieldName:  "deli_set_area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "CATE_NOTE",
			Comment:            `カテゴリ説明`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "CATENOTE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_note",
			ProtobufFieldName:  "cate_note",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "VALID_KBN",
			Comment:            `有効区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "VALIDKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "valid_kbn",
			ProtobufFieldName:  "valid_kbn",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
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
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CATEGORY1_MST_) TableName() string {
	return "CATEGORY1_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CATEGORY1_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CATEGORY1_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CATEGORY1_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CATEGORY1_MST_) TableInfo() *TableInfo {
	return CATEGORY1_MSTTableInfo
}
