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


CREATE TABLE `WETH_WARN_MST` (
  `WARN_CD` char(2) NOT NULL COMMENT '気象注警報コード',
  `WARN_NM` varchar(40) NOT NULL COMMENT '気象注警報名',
  `WARN_KBN` char(1) NOT NULL COMMENT '気象注警報区分	 0：特別警報、1：警報、2：注意報（又は全解除）',
  `LEVEL_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'レベルID',
  `LEVEL_VALUE` int(11) NOT NULL COMMENT 'レベル	 特別警報=1、警報=2、注意報=3',
  `CATE2_CD` varchar(10) NOT NULL DEFAULT '0' COMMENT '注意報カテゴリコード',
  `DISP_DSQ` int(11) NOT NULL COMMENT '表示順',
  `DELE_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`WARN_CD`),
  UNIQUE KEY `WETH_WARN_MST_PKI` (`WARN_CD`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='気象注意報マスタ'

JSON Sample
-------------------------------------
{    "warn_cd": "AnTCNlUITGwwsnfCGvlViKNKK",    "warn_nm": "iEtGRNeWKtIoQuMMsXUFCAxEU",    "warn_kbn": "QCyZhAVRUyuXLQvPbLfcdqbud",    "level_id": 27,    "level_value": 55,    "cate_2_cd": "UuYxtgpyxifulUrGopQsvBVbJ",    "disp_dsq": 93,    "dele_flg": "TWqdfsDrmQaxseFmXJmaPUkhG",    "upda_dte": "2074-11-24T11:37:36.920295871+09:00",    "upda_user_id": 44,    "crea_dte": "2108-08-01T17:02:35.765415328+09:00",    "crea_user_id": 21}



*/

// WETH_WARN_MST_ struct is a row record of the WETH_WARN_MST table in the anpidb database
type WETH_WARN_MST_ struct {
	//[ 0] WARN_CD                                        char(2)              null: false  primary: true   isArray: false  auto: false  col: char            len: 2       default: []
	WARNCD string `gorm:"primary_key;column:WARN_CD;type:char;size:2;"` // 気象注警報コード
	//[ 1] WARN_NM                                        varchar(40)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 40      default: []
	WARNNM string `gorm:"column:WARN_NM;type:varchar;size:40;"` // 気象注警報名
	//[ 2] WARN_KBN                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WARNKBN string `gorm:"column:WARN_KBN;type:char;size:1;"` // 気象注警報区分	 0：特別警報、1：警報、2：注意報（又は全解除）
	//[ 3] LEVEL_ID                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LEVELID int32 `gorm:"column:LEVEL_ID;type:int;default:0;"` // レベルID
	//[ 4] LEVEL_VALUE                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	LEVELVALUE int32 `gorm:"column:LEVEL_VALUE;type:int;"` // レベル	 特別警報=1、警報=2、注意報=3
	//[ 5] CATE2_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	CATE2CD string `gorm:"column:CATE2_CD;type:varchar;size:10;default:0;"` // 注意報カテゴリコード
	//[ 6] DISP_DSQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DISPDSQ int32 `gorm:"column:DISP_DSQ;type:int;"` // 表示順
	//[ 7] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var WETH_WARN_MSTTableInfo = &TableInfo{
	Name: "WETH_WARN_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "WARN_CD",
			Comment:            `気象注警報コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "WARNCD",
			GoFieldType:        "string",
			JSONFieldName:      "warn_cd",
			ProtobufFieldName:  "warn_cd",
			ProtobufType:       "string",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "WARN_NM",
			Comment:            `気象注警報名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(40)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       40,
			GoFieldName:        "WARNNM",
			GoFieldType:        "string",
			JSONFieldName:      "warn_nm",
			ProtobufFieldName:  "warn_nm",
			ProtobufType:       "string",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index: 2,
			Name:  "WARN_KBN",
			Comment: `気象注警報区分	 0：特別警報、1：警報、2：注意報（又は全解除）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WARNKBN",
			GoFieldType:        "string",
			JSONFieldName:      "warn_kbn",
			ProtobufFieldName:  "warn_kbn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "LEVEL_ID",
			Comment:            `レベルID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELID",
			GoFieldType:        "int32",
			JSONFieldName:      "level_id",
			ProtobufFieldName:  "level_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "LEVEL_VALUE",
			Comment: `レベル	 特別警報=1、警報=2、注意報=3`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELVALUE",
			GoFieldType:        "int32",
			JSONFieldName:      "level_value",
			ProtobufFieldName:  "level_value",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CATE2_CD",
			Comment:            `注意報カテゴリコード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE2CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_2_cd",
			ProtobufFieldName:  "cate_2_cd",
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},
	},
}

// TableName sets the insert table name for this struct type
func (w *WETH_WARN_MST_) TableName() string {
	return "WETH_WARN_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (w *WETH_WARN_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (w *WETH_WARN_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (w *WETH_WARN_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (w *WETH_WARN_MST_) TableInfo() *TableInfo {
	return WETH_WARN_MSTTableInfo
}
