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


CREATE TABLE `CONT_INFO_WARN` (
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `SUB_NO` int(11) NOT NULL COMMENT '枝番',
  `WARN_SUB_NO` int(11) NOT NULL DEFAULT '0' COMMENT '注警報枝番	 細分区域は0',
  `WARN_ANNO_KBN` char(1) NOT NULL COMMENT '注警報区分	 N:発表 C:発表中 D:解除',
  `WARN_CD` varchar(2) NOT NULL COMMENT '気象注警報コード',
  `CATE2_CD` varchar(10) NOT NULL DEFAULT '0' COMMENT 'カテゴリ2コード',
  `WARN_LVL` int(11) NOT NULL COMMENT '気象注警報レベル',
  `BEF_WARN_LVL` int(11) DEFAULT NULL COMMENT '前回気象注警報レベル',
  `WARN_LVL_ID` int(11) NOT NULL COMMENT '気象注警報レベルID	 WETH_WARN_MST.LEVEL_ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SID`,`SUB_NO`,`WARN_SUB_NO`),
  UNIQUE KEY `CONT_INFO_WARN_PKI` (`SID`,`SUB_NO`,`WARN_SUB_NO`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツ情報注警報'

JSON Sample
-------------------------------------
{    "sid": 75,    "sub_no": 27,    "warn_sub_no": 1,    "warn_anno_kbn": "dIUXyFsxWxogvxgZnFwnWsrce",    "warn_cd": "ULjcwpGmpviBHNcBmCcobZOlK",    "cate_2_cd": "VqgmeTJejlPxIUmhsnOJMvOTN",    "warn_lvl": 68,    "bef_warn_lvl": 45,    "warn_lvl_id": 38,    "crea_dte": "2294-07-25T07:11:01.064771568+09:00",    "crea_user_id": 83}



*/

// CONT_INFO_WARN_ struct is a row record of the CONT_INFO_WARN table in the anpidb database
type CONT_INFO_WARN_ struct {
	//[ 0] SID                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"primary_key;column:SID;type:int;"` // 情報番号
	//[ 1] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;"` // 枝番
	//[ 2] WARN_SUB_NO                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	WARNSUBNO int32 `gorm:"primary_key;column:WARN_SUB_NO;type:int;default:0;"` // 注警報枝番	 細分区域は0
	//[ 3] WARN_ANNO_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WARNANNOKBN string `gorm:"column:WARN_ANNO_KBN;type:char;size:1;"` // 注警報区分	 N:発表 C:発表中 D:解除
	//[ 4] WARN_CD                                        varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	WARNCD string `gorm:"column:WARN_CD;type:varchar;size:2;"` // 気象注警報コード
	//[ 5] CATE2_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	CATE2CD string `gorm:"column:CATE2_CD;type:varchar;size:10;default:0;"` // カテゴリ2コード
	//[ 6] WARN_LVL                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	WARNLVL int32 `gorm:"column:WARN_LVL;type:int;"` // 気象注警報レベル
	//[ 7] BEF_WARN_LVL                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	BEFWARNLVL sql.NullInt64 `gorm:"column:BEF_WARN_LVL;type:int;"` // 前回気象注警報レベル
	//[ 8] WARN_LVL_ID                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	WARNLVLID int32 `gorm:"column:WARN_LVL_ID;type:int;"` // 気象注警報レベルID	 WETH_WARN_MST.LEVEL_ID
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[10] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_INFO_WARNTableInfo = &TableInfo{
	Name: "CONT_INFO_WARN",
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
			Name:               "SUB_NO",
			Comment:            `枝番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "sub_no",
			ProtobufFieldName:  "sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index: 2,
			Name:  "WARN_SUB_NO",
			Comment: `注警報枝番	 細分区域は0`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "WARNSUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "warn_sub_no",
			ProtobufFieldName:  "warn_sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "WARN_ANNO_KBN",
			Comment: `注警報区分	 N:発表 C:発表中 D:解除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "WARNANNOKBN",
			GoFieldType:        "string",
			JSONFieldName:      "warn_anno_kbn",
			ProtobufFieldName:  "warn_anno_kbn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "WARN_CD",
			Comment:            `気象注警報コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "WARNCD",
			GoFieldType:        "string",
			JSONFieldName:      "warn_cd",
			ProtobufFieldName:  "warn_cd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "CATE2_CD",
			Comment:            `カテゴリ2コード`,
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
			Name:               "WARN_LVL",
			Comment:            `気象注警報レベル`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "WARNLVL",
			GoFieldType:        "int32",
			JSONFieldName:      "warn_lvl",
			ProtobufFieldName:  "warn_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "BEF_WARN_LVL",
			Comment:            `前回気象注警報レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "BEFWARNLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "bef_warn_lvl",
			ProtobufFieldName:  "bef_warn_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index: 8,
			Name:  "WARN_LVL_ID",
			Comment: `気象注警報レベルID	 WETH_WARN_MST.LEVEL_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "WARNLVLID",
			GoFieldType:        "int32",
			JSONFieldName:      "warn_lvl_id",
			ProtobufFieldName:  "warn_lvl_id",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_INFO_WARN_) TableName() string {
	return "CONT_INFO_WARN"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_INFO_WARN_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_INFO_WARN_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_INFO_WARN_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_INFO_WARN_) TableInfo() *TableInfo {
	return CONT_INFO_WARNTableInfo
}
