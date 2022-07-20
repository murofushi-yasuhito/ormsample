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


CREATE TABLE `CATE_MST` (
  `CATE_VER_NO` int(11) NOT NULL COMMENT 'カテゴリVerNo',
  `CATE_CD` varchar(12) NOT NULL COMMENT 'カテゴリコード',
  `CATE_NM` varchar(80) NOT NULL COMMENT 'カテゴリ名',
  `CATE_LVL` int(11) NOT NULL COMMENT 'カテゴリレベル',
  `CATE_STAT_FLG` char(1) NOT NULL COMMENT 'カテゴリステータスフラグ',
  `CHILD_CATE_EXIST_FLG` char(1) NOT NULL COMMENT '子カテゴリ有無',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DELI_AREA_KBN_CD` char(4) DEFAULT NULL COMMENT '配信エリア区分コード',
  `DEFA_VALI_MIN` int(11) NOT NULL COMMENT '標準有効時間',
  `DELE_FLG` char(1) NOT NULL COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_CD` varchar(10) DEFAULT NULL,
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_CD` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`CATE_VER_NO`,`CATE_CD`),
  KEY `CATE_MST_IDX_1` (`CATE_VER_NO`,`CATE_LVL`,`CATE_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "cate_ver_no": 93,    "cate_cd": "lpRkrMpVRBySnybbHyflHAabW",    "cate_nm": "cCwsZcUtaqdmplsXBvbnVweYa",    "cate_lvl": 28,    "cate_stat_flg": "cvGMlkJoKyuBxMUNFujwYuZlY",    "child_cate_exist_flg": "WiAirRFDmhXcmJAgpmDKNfMbj",    "disp_dsq": 38,    "deli_area_kbn_cd": "TCasPVlhbWrTXPTMGOKVjBQnV",    "defa_vali_min": 28,    "dele_flg": "xDHpTAyYRqWDDHorlwTIgRUxo",    "upda_dte": "2087-01-16T23:00:06.908586394+09:00",    "upda_user_cd": "kwBmrpqTciUDvhqpfQIutNvAP",    "crea_dte": "2061-03-19T21:26:38.009676134+09:00",    "crea_user_cd": "hqFMZppSYelJfKeuLAdOTyirX"}



*/

// CATE_MST_ struct is a row record of the CATE_MST table in the anpidb database
type CATE_MST_ struct {
	//[ 0] CATE_VER_NO                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CATEVERNO int32 `gorm:"primary_key;column:CATE_VER_NO;type:int;"` // カテゴリVerNo
	//[ 1] CATE_CD                                        varchar(12)          null: false  primary: true   isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD string `gorm:"primary_key;column:CATE_CD;type:varchar;size:12;"` // カテゴリコード
	//[ 2] CATE_NM                                        varchar(80)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 80      default: []
	CATENM string `gorm:"column:CATE_NM;type:varchar;size:80;"` // カテゴリ名
	//[ 3] CATE_LVL                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CATELVL int32 `gorm:"column:CATE_LVL;type:int;"` // カテゴリレベル
	//[ 4] CATE_STAT_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CATESTATFLG string `gorm:"column:CATE_STAT_FLG;type:char;size:1;"` // カテゴリステータスフラグ
	//[ 5] CHILD_CATE_EXIST_FLG                           char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CHILDCATEEXISTFLG string `gorm:"column:CHILD_CATE_EXIST_FLG;type:char;size:1;"` // 子カテゴリ有無
	//[ 6] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 7] DELI_AREA_KBN_CD                               char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	DELIAREAKBNCD sql.NullString `gorm:"column:DELI_AREA_KBN_CD;type:char;size:4;"` // 配信エリア区分コード
	//[ 8] DEFA_VALI_MIN                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEFAVALIMIN int32 `gorm:"column:DEFA_VALI_MIN;type:int;"` // 標準有効時間
	//[ 9] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;"` // 削除フラグ
	//[10] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[11] UPDA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	UPDAUSERCD sql.NullString `gorm:"column:UPDA_USER_CD;type:varchar;size:10;"`
	//[12] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[13] CREA_USER_CD                                   varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CREAUSERCD sql.NullString `gorm:"column:CREA_USER_CD;type:varchar;size:10;"`
}

var CATE_MSTTableInfo = &TableInfo{
	Name: "CATE_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CATE_VER_NO",
			Comment:            `カテゴリVerNo`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CATEVERNO",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_ver_no",
			ProtobufFieldName:  "cate_ver_no",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "CATE_CD",
			Comment:            `カテゴリコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_cd",
			ProtobufFieldName:  "cate_cd",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CATE_NM",
			Comment:            `カテゴリ名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(80)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       80,
			GoFieldName:        "CATENM",
			GoFieldType:        "string",
			JSONFieldName:      "cate_nm",
			ProtobufFieldName:  "cate_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "CATE_LVL",
			Comment:            `カテゴリレベル`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CATELVL",
			GoFieldType:        "int32",
			JSONFieldName:      "cate_lvl",
			ProtobufFieldName:  "cate_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CATE_STAT_FLG",
			Comment:            `カテゴリステータスフラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CATESTATFLG",
			GoFieldType:        "string",
			JSONFieldName:      "cate_stat_flg",
			ProtobufFieldName:  "cate_stat_flg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CHILD_CATE_EXIST_FLG",
			Comment:            `子カテゴリ有無`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CHILDCATEEXISTFLG",
			GoFieldType:        "string",
			JSONFieldName:      "child_cate_exist_flg",
			ProtobufFieldName:  "child_cate_exist_flg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DISP_DSQ",
			Comment:            `表示順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DISPDSQ",
			GoFieldType:        "int32",
			JSONFieldName:      "disp_dsq",
			ProtobufFieldName:  "disp_dsq",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "DELI_AREA_KBN_CD",
			Comment:            `配信エリア区分コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "DELIAREAKBNCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_area_kbn_cd",
			ProtobufFieldName:  "deli_area_kbn_cd",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "DEFA_VALI_MIN",
			Comment:            `標準有効時間`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEFAVALIMIN",
			GoFieldType:        "int32",
			JSONFieldName:      "defa_vali_min",
			ProtobufFieldName:  "defa_vali_min",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "UPDA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "UPDAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "upda_user_cd",
			ProtobufFieldName:  "upda_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "CREA_USER_CD",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CREAUSERCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "crea_user_cd",
			ProtobufFieldName:  "crea_user_cd",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CATE_MST_) TableName() string {
	return "CATE_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CATE_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CATE_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CATE_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CATE_MST_) TableInfo() *TableInfo {
	return CATE_MSTTableInfo
}
