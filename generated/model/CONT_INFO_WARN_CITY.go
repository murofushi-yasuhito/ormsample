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


CREATE TABLE `CONT_INFO_WARN_CITY` (
  `SID` int(11) NOT NULL COMMENT '情報番号',
  `SUB_NO` int(11) NOT NULL COMMENT '枝番',
  `ELMNT_LVL` int(11) DEFAULT NULL COMMENT 'RSQ影響レベル',
  `INFO_TITLE` varchar(256) DEFAULT NULL COMMENT '記事タイトル',
  `INFO_SHORT` text COMMENT 'ショート文',
  `INFO_LONG` text COMMENT 'ロング文',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SID`,`SUB_NO`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツ情報注警報市区町村'

JSON Sample
-------------------------------------
{    "sid": 83,    "sub_no": 14,    "elmnt_lvl": 64,    "info_title": "WRAdcXNDsJYFLJFUOxlTxauxt",    "info_short": "WXnqcDDoCZKClnLekWPVmUlZN",    "info_long": "YpkMaQBUNKgbtNhesojcpdGQt",    "crea_dte": "2118-12-29T21:12:05.832838741+09:00",    "crea_user_id": 73}



*/

// CONT_INFO_WARN_CITY_ struct is a row record of the CONT_INFO_WARN_CITY table in the anpidb database
type CONT_INFO_WARN_CITY_ struct {
	//[ 0] SID                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SID int32 `gorm:"primary_key;column:SID;type:int;"` // 情報番号
	//[ 1] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;"` // 枝番
	//[ 2] ELMNT_LVL                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ELMNTLVL sql.NullInt64 `gorm:"column:ELMNT_LVL;type:int;"` // RSQ影響レベル
	//[ 3] INFO_TITLE                                     varchar(256)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 256     default: []
	INFOTITLE sql.NullString `gorm:"column:INFO_TITLE;type:varchar;size:256;"` // 記事タイトル
	//[ 4] INFO_SHORT                                     text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	INFOSHORT sql.NullString `gorm:"column:INFO_SHORT;type:text;size:65535;"` // ショート文
	//[ 5] INFO_LONG                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	INFOLONG sql.NullString `gorm:"column:INFO_LONG;type:text;size:65535;"` // ロング文
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_INFO_WARN_CITYTableInfo = &TableInfo{
	Name: "CONT_INFO_WARN_CITY",
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
			Index:              2,
			Name:               "ELMNT_LVL",
			Comment:            `RSQ影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "elmnt_lvl",
			ProtobufFieldName:  "elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_INFO_WARN_CITY_) TableName() string {
	return "CONT_INFO_WARN_CITY"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_INFO_WARN_CITY_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_INFO_WARN_CITY_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_INFO_WARN_CITY_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_INFO_WARN_CITY_) TableInfo() *TableInfo {
	return CONT_INFO_WARN_CITYTableInfo
}
