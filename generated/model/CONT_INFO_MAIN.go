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


CREATE TABLE `CONT_INFO_MAIN` (
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `TOP_SID` int(11) DEFAULT NULL COMMENT '親情報番号',
  `BEF_SID` int(11) DEFAULT NULL COMMENT '継続情報番号',
  `DELI_FLG` char(1) NOT NULL COMMENT '配信フラグ	 N:新規・続報 C:訂正 4:削除',
  `CATE_CD` varchar(12) NOT NULL COMMENT 'カテゴリコード',
  `CATE_CD1` varchar(12) DEFAULT NULL COMMENT 'カテゴリコード1',
  `CATE_CD2` varchar(12) DEFAULT NULL COMMENT 'カテゴリコード2',
  `CATE_CD3` varchar(12) DEFAULT NULL COMMENT 'カテゴリコード3',
  `CNFRM_DTE` datetime DEFAULT NULL COMMENT 'ステータス確認日時',
  `ELMNT_LVL` tinyint(4) DEFAULT NULL COMMENT 'RSQ影響レベル',
  `INFO_TITLE` varchar(256) DEFAULT NULL COMMENT '記事タイトル',
  `INFO_SHORT` text COMMENT 'ショート文',
  `INFO_LONG` text COMMENT 'ロング文',
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード',
  `CATE2_CD` varchar(10) DEFAULT NULL COMMENT 'カテゴリ2コード',
  `BEF_ELMNT_LVL` tinyint(4) DEFAULT '0' COMMENT '前回影響レベル',
  `AREA_LVL_UP_KBN` char(1) DEFAULT '0' COMMENT 'エリアレベルUP区分	 0:無し 1:エリアレベルアップ 2:エリア拡大 3:両方',
  `DELI_DAY_KND` char(1) DEFAULT NULL COMMENT '配信日種別',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SID`),
  KEY `CONT_INFO_MAIN_IDX_1` (`CNFRM_DTE`,`BEF_SID`),
  KEY `CONT_INFO_MAIN_IDX_2` (`CATE_CD`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツ情報'

JSON Sample
-------------------------------------
{    "sid": 56,    "top_sid": 16,    "bef_sid": 92,    "deli_flg": "MlFFCdjRMahQpHjNGNIuCDoKY",    "cate_cd": "woFchNgaOsXGEOykwRvIvufmi",    "cate_cd_1": "UmZIZTXjiZZEhJVLXfLnmKtpy",    "cate_cd_2": "lSrguDfMBVcAwawNLOAAcMksg",    "cate_cd_3": "MOaPFaMSffoceRjMfsgIqaRDY",    "cnfrm_dte": "2189-11-07T07:20:12.946781848+09:00",    "elmnt_lvl": 17,    "info_title": "JiTvEpGCItLvuGXKtvRMrEBuv",    "info_short": "KUDDRSAtRNSFXsvTWvBDgdPry",    "info_long": "HpEMVBZyNPynmvbTipxHlpuOq",    "cate_1_cd": "pbKErNhBElfcvmTKmudPuoHJe",    "cate_2_cd": "lExtBlutQptiwLMeMlDtmSAhC",    "bef_elmnt_lvl": 10,    "area_lvl_up_kbn": "hhZjXMRHUEeVhRRWyWTUyxZSc",    "deli_day_knd": "PctvSHBRiAZIscheSLnkMxvUU",    "crea_dte": "2165-03-16T07:55:36.723727753+09:00",    "crea_user_id": 16}



*/

// CONT_INFO_MAIN_ struct is a row record of the CONT_INFO_MAIN table in the anpidb database
type CONT_INFO_MAIN_ struct {
	//[ 0] SID                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"primary_key;column:SID;type:int;"` // 情報番号
	//[ 1] TOP_SID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	TOPSID sql.NullInt64 `gorm:"column:TOP_SID;type:int;"` // 親情報番号
	//[ 2] BEF_SID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BEFSID sql.NullInt64 `gorm:"column:BEF_SID;type:int;"` // 継続情報番号
	//[ 3] DELI_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELIFLG string `gorm:"column:DELI_FLG;type:char;size:1;"` // 配信フラグ	 N:新規・続報 C:訂正 4:削除
	//[ 4] CATE_CD                                        varchar(12)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD string `gorm:"column:CATE_CD;type:varchar;size:12;"` // カテゴリコード
	//[ 5] CATE_CD1                                       varchar(12)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD1 sql.NullString `gorm:"column:CATE_CD1;type:varchar;size:12;"` // カテゴリコード1
	//[ 6] CATE_CD2                                       varchar(12)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD2 sql.NullString `gorm:"column:CATE_CD2;type:varchar;size:12;"` // カテゴリコード2
	//[ 7] CATE_CD3                                       varchar(12)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 12      default: []
	CATECD3 sql.NullString `gorm:"column:CATE_CD3;type:varchar;size:12;"` // カテゴリコード3
	//[ 8] CNFRM_DTE                                      datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CNFRMDTE time.Time `gorm:"column:CNFRM_DTE;type:datetime;"` // ステータス確認日時
	//[ 9] ELMNT_LVL                                      tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: []
	ELMNTLVL sql.NullInt64 `gorm:"column:ELMNT_LVL;type:tinyint;"` // RSQ影響レベル
	//[10] INFO_TITLE                                     varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	INFOTITLE sql.NullString `gorm:"column:INFO_TITLE;type:varchar;size:256;"` // 記事タイトル
	//[11] INFO_SHORT                                     text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	INFOSHORT sql.NullString `gorm:"column:INFO_SHORT;type:text;size:65535;"` // ショート文
	//[12] INFO_LONG                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	INFOLONG sql.NullString `gorm:"column:INFO_LONG;type:text;size:65535;"` // ロング文
	//[13] CATE1_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード
	//[14] CATE2_CD                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE2CD sql.NullString `gorm:"column:CATE2_CD;type:varchar;size:10;"` // カテゴリ2コード
	//[15] BEF_ELMNT_LVL                                  tinyint              null: true   primary: false  isArray: false  auto: false  col: tinyint         len: -1      default: [0]
	BEFELMNTLVL sql.NullInt64 `gorm:"column:BEF_ELMNT_LVL;type:tinyint;default:0;"` // 前回影響レベル
	//[16] AREA_LVL_UP_KBN                                char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	AREALVLUPKBN sql.NullString `gorm:"column:AREA_LVL_UP_KBN;type:char;size:1;default:0;"` // エリアレベルUP区分	 0:無し 1:エリアレベルアップ 2:エリア拡大 3:両方
	//[17] DELI_DAY_KND                                   char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DELIDAYKND sql.NullString `gorm:"column:DELI_DAY_KND;type:char;size:1;"` // 配信日種別
	//[18] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[19] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_INFO_MAINTableInfo = &TableInfo{
	Name: "CONT_INFO_MAIN",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "SID",
			Comment:            `情報番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SID",
			GoFieldType:        "int32",
			JSONFieldName:      "sid",
			ProtobufFieldName:  "sid",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "TOP_SID",
			Comment:            `親情報番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPSID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "top_sid",
			ProtobufFieldName:  "top_sid",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "BEF_SID",
			Comment:            `継続情報番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BEFSID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bef_sid",
			ProtobufFieldName:  "bef_sid",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "DELI_FLG",
			Comment: `配信フラグ	 N:新規・続報 C:訂正 4:削除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELIFLG",
			GoFieldType:        "string",
			JSONFieldName:      "deli_flg",
			ProtobufFieldName:  "deli_flg",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "CATE_CD",
			Comment:            `カテゴリコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_cd",
			ProtobufFieldName:  "cate_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CATE_CD1",
			Comment:            `カテゴリコード1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD1",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_cd_1",
			ProtobufFieldName:  "cate_cd_1",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CATE_CD2",
			Comment:            `カテゴリコード2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_cd_2",
			ProtobufFieldName:  "cate_cd_2",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CATE_CD3",
			Comment:            `カテゴリコード3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(12)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       12,
			GoFieldName:        "CATECD3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_cd_3",
			ProtobufFieldName:  "cate_cd_3",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "CNFRM_DTE",
			Comment:            `ステータス確認日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CNFRMDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "cnfrm_dte",
			ProtobufFieldName:  "cnfrm_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "ELMNT_LVL",
			Comment:            `RSQ影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "ELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "elmnt_lvl",
			ProtobufFieldName:  "elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "INFO_TITLE",
			Comment:            `記事タイトル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(256)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       256,
			GoFieldName:        "INFOTITLE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "info_title",
			ProtobufFieldName:  "info_title",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "INFO_SHORT",
			Comment:            `ショート文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "INFOSHORT",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "info_short",
			ProtobufFieldName:  "info_short",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "INFO_LONG",
			Comment:            `ロング文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "INFOLONG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "info_long",
			ProtobufFieldName:  "info_long",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "CATE1_CD",
			Comment:            `カテゴリ1コード`,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "CATE2_CD",
			Comment:            `カテゴリ2コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "BEF_ELMNT_LVL",
			Comment:            `前回影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "tinyint",
			DatabaseTypePretty: "tinyint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "tinyint",
			ColumnLength:       -1,
			GoFieldName:        "BEFELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bef_elmnt_lvl",
			ProtobufFieldName:  "bef_elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index: 16,
			Name:  "AREA_LVL_UP_KBN",
			Comment: `エリアレベルUP区分	 0:無し 1:エリアレベルアップ 2:エリア拡大 3:両方`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AREALVLUPKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_lvl_up_kbn",
			ProtobufFieldName:  "area_lvl_up_kbn",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "DELI_DAY_KND",
			Comment:            `配信日種別`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELIDAYKND",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "deli_day_knd",
			ProtobufFieldName:  "deli_day_knd",
			ProtobufType:       "string",
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
func (c *CONT_INFO_MAIN_) TableName() string {
	return "CONT_INFO_MAIN"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_INFO_MAIN_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_INFO_MAIN_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_INFO_MAIN_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_INFO_MAIN_) TableInfo() *TableInfo {
	return CONT_INFO_MAINTableInfo
}
