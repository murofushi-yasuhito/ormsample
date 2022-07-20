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


CREATE TABLE `CONT_MAIL_SEND_HIS_WARN` (
  `MAIL_SEND_ID` int(11) NOT NULL COMMENT 'メール送信ID',
  `SUB_NO` int(11) NOT NULL COMMENT '枝番',
  `WARN_SUB_NO` int(11) NOT NULL COMMENT '注警報枝番',
  `WARN_ANNO_KBN` char(1) NOT NULL COMMENT '注警報区分	 N:発表 C:発表中 D:解除',
  `WARN_CD` varchar(2) NOT NULL COMMENT '注警報コード',
  `WARN_CATE_CD` char(2) DEFAULT NULL COMMENT '注意報カテゴリコード',
  `WARN_LVL` int(11) NOT NULL COMMENT '注意報レベル',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MAIL_SEND_ID`,`SUB_NO`,`WARN_SUB_NO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツメール配信履歴注警報'

JSON Sample
-------------------------------------
{    "mail_send_id": 50,    "sub_no": 73,    "warn_sub_no": 47,    "warn_anno_kbn": "fdHNDierwrfsXLZmQGkHnnUdY",    "warn_cd": "koApbWrJeJyTMUXkwIOfuMZIe",    "warn_cate_cd": "SQZUdhCUlPhgopeGoAMhyCQcN",    "warn_lvl": 44,    "crea_dte": "2040-09-20T20:34:07.383477215+09:00",    "crea_user_id": 10}



*/

// CONT_MAIL_SEND_HIS_WARN_ struct is a row record of the CONT_MAIL_SEND_HIS_WARN table in the anpidb database
type CONT_MAIL_SEND_HIS_WARN_ struct {
	//[ 0] MAIL_SEND_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	MAILSENDID int32 `gorm:"primary_key;column:MAIL_SEND_ID;type:int;"` // メール送信ID
	//[ 1] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;"` // 枝番
	//[ 2] WARN_SUB_NO                                    int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	WARNSUBNO int32 `gorm:"primary_key;column:WARN_SUB_NO;type:int;"` // 注警報枝番
	//[ 3] WARN_ANNO_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	WARNANNOKBN string `gorm:"column:WARN_ANNO_KBN;type:char;size:1;"` // 注警報区分	 N:発表 C:発表中 D:解除
	//[ 4] WARN_CD                                        varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	WARNCD string `gorm:"column:WARN_CD;type:varchar;size:2;"` // 注警報コード
	//[ 5] WARN_CATE_CD                                   char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	WARNCATECD sql.NullString `gorm:"column:WARN_CATE_CD;type:char;size:2;"` // 注意報カテゴリコード
	//[ 6] WARN_LVL                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	WARNLVL int32 `gorm:"column:WARN_LVL;type:int;"` // 注意報レベル
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_MAIL_SEND_HIS_WARNTableInfo = &TableInfo{
	Name: "CONT_MAIL_SEND_HIS_WARN",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MAIL_SEND_ID",
			Comment:            `メール送信ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MAILSENDID",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_send_id",
			ProtobufFieldName:  "mail_send_id",
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
			Index:              2,
			Name:               "WARN_SUB_NO",
			Comment:            `注警報枝番`,
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
			Comment:            `注警報コード`,
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
			Name:               "WARN_CATE_CD",
			Comment:            `注意報カテゴリコード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "WARNCATECD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "warn_cate_cd",
			ProtobufFieldName:  "warn_cate_cd",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "WARN_LVL",
			Comment:            `注意報レベル`,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_MAIL_SEND_HIS_WARN_) TableName() string {
	return "CONT_MAIL_SEND_HIS_WARN"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_WARN_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_MAIL_SEND_HIS_WARN_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_WARN_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_MAIL_SEND_HIS_WARN_) TableInfo() *TableInfo {
	return CONT_MAIL_SEND_HIS_WARNTableInfo
}
