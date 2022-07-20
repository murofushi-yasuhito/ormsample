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


CREATE TABLE `ANKEN_TARGET_POST` (
  `ANKEN_TARGET_ID` int(11) NOT NULL COMMENT '送信先対象選択ID',
  `POST_SEQ` int(11) NOT NULL COMMENT '役職連番',
  `POST_COND_KBN` char(1) NOT NULL COMMENT '役職検索区分（0:指定なし 1:以上 2:以下）',
  `POST_ID` int(11) NOT NULL COMMENT '役職ID',
  `POST_NM` varchar(80) NOT NULL COMMENT '役職名',
  `POST_LEVEL` int(11) NOT NULL DEFAULT '0' COMMENT '職位順',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_TARGET_ID`,`POST_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件手動任意指定詳細'

JSON Sample
-------------------------------------
{    "anken_target_id": 7,    "post_seq": 49,    "post_cond_kbn": "jTviRiAKdJJuaufxUSeHcJfSZ",    "post_id": 69,    "post_nm": "VSyNOxIoZSKDKXgiPOblIxLUB",    "post_level": 47,    "upda_dte": "2125-10-09T13:22:12.619612006+09:00",    "upda_user_id": 74,    "crea_dte": "2112-06-27T08:26:51.754710413+09:00",    "crea_user_id": 85}



*/

// ANKEN_TARGET_POST_ struct is a row record of the ANKEN_TARGET_POST table in the anpidb database
type ANKEN_TARGET_POST_ struct {
	//[ 0] ANKEN_TARGET_ID                                int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENTARGETID int32 `gorm:"primary_key;column:ANKEN_TARGET_ID;type:int;"` // 送信先対象選択ID
	//[ 1] POST_SEQ                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	POSTSEQ int32 `gorm:"primary_key;column:POST_SEQ;type:int;"` // 役職連番
	//[ 2] POST_COND_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	POSTCONDKBN string `gorm:"column:POST_COND_KBN;type:char;size:1;"` // 役職検索区分（0:指定なし 1:以上 2:以下）
	//[ 3] POST_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	POSTID int32 `gorm:"column:POST_ID;type:int;"` // 役職ID
	//[ 4] POST_NM                                        varchar(80)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 80      default: []
	POSTNM string `gorm:"column:POST_NM;type:varchar;size:80;"` // 役職名
	//[ 5] POST_LEVEL                                     int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	POSTLEVEL int32 `gorm:"column:POST_LEVEL;type:int;default:0;"` // 職位順
	//[ 6] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 7] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[ 8] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 9] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_TARGET_POSTTableInfo = &TableInfo{
	Name: "ANKEN_TARGET_POST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_TARGET_ID",
			Comment:            `送信先対象選択ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENTARGETID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_target_id",
			ProtobufFieldName:  "anken_target_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "POST_SEQ",
			Comment:            `役職連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "POSTSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "post_seq",
			ProtobufFieldName:  "post_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "POST_COND_KBN",
			Comment:            `役職検索区分（0:指定なし 1:以上 2:以下）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "POSTCONDKBN",
			GoFieldType:        "string",
			JSONFieldName:      "post_cond_kbn",
			ProtobufFieldName:  "post_cond_kbn",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "POST_ID",
			Comment:            `役職ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "POSTID",
			GoFieldType:        "int32",
			JSONFieldName:      "post_id",
			ProtobufFieldName:  "post_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "POST_NM",
			Comment:            `役職名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(80)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       80,
			GoFieldName:        "POSTNM",
			GoFieldType:        "string",
			JSONFieldName:      "post_nm",
			ProtobufFieldName:  "post_nm",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "POST_LEVEL",
			Comment:            `職位順`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "POSTLEVEL",
			GoFieldType:        "int32",
			JSONFieldName:      "post_level",
			ProtobufFieldName:  "post_level",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_TARGET_POST_) TableName() string {
	return "ANKEN_TARGET_POST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_TARGET_POST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_TARGET_POST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_TARGET_POST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_TARGET_POST_) TableInfo() *TableInfo {
	return ANKEN_TARGET_POSTTableInfo
}
