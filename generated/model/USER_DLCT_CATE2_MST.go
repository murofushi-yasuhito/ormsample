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


CREATE TABLE `USER_DLCT_CATE2_MST` (
  `USER_DLCT_ID` int(11) NOT NULL COMMENT 'ユーザDLCT_ID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `CATE2_CD` varchar(10) NOT NULL COMMENT 'カテゴリ2コード	 画面選択サブカテゴリ',
  `LEVEL_VALUE` int(11) DEFAULT '99' COMMENT 'レベル',
  `LEVEL_KBN` char(1) DEFAULT '0' COMMENT 'レベル区分',
  `LEVEL_ID` int(11) DEFAULT '0' COMMENT 'レベルID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_DLCT_ID`,`SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザDLCTカテゴリ2マスタ'

JSON Sample
-------------------------------------
{    "user_dlct_id": 17,    "seq": 75,    "cate_2_cd": "ttIYYMGcWLfywytuhTNjxcvdG",    "level_value": 8,    "level_kbn": "JenlqsEBrIacIhITpUtFCpyLv",    "level_id": 24,    "crea_dte": "2063-11-20T18:29:53.587562+09:00",    "crea_user_id": 85}



*/

// USER_DLCT_CATE2_MST_ struct is a row record of the USER_DLCT_CATE2_MST table in the anpidb database
type USER_DLCT_CATE2_MST_ struct {
	//[ 0] USER_DLCT_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERDLCTID int32 `gorm:"primary_key;column:USER_DLCT_ID;type:int;"` // ユーザDLCT_ID
	//[ 1] SEQ                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"primary_key;column:SEQ;type:int;"` // 連番
	//[ 2] CATE2_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE2CD string `gorm:"column:CATE2_CD;type:varchar;size:10;"` // カテゴリ2コード	 画面選択サブカテゴリ
	//[ 3] LEVEL_VALUE                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [99]
	LEVELVALUE sql.NullInt64 `gorm:"column:LEVEL_VALUE;type:int;default:99;"` // レベル
	//[ 4] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;default:0;"` // レベル区分
	//[ 5] LEVEL_ID                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	LEVELID sql.NullInt64 `gorm:"column:LEVEL_ID;type:int;default:0;"` // レベルID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_DLCT_CATE2_MSTTableInfo = &TableInfo{
	Name: "USER_DLCT_CATE2_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_DLCT_ID",
			Comment:            `ユーザDLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_dlct_id",
			ProtobufFieldName:  "user_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "SEQ",
			Comment:            `連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "seq",
			ProtobufFieldName:  "seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index: 2,
			Name:  "CATE2_CD",
			Comment: `カテゴリ2コード	 画面選択サブカテゴリ`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "LEVEL_VALUE",
			Comment:            `レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "LEVELVALUE",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "level_value",
			ProtobufFieldName:  "level_value",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "LEVEL_KBN",
			Comment:            `レベル区分`,
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
func (u *USER_DLCT_CATE2_MST_) TableName() string {
	return "USER_DLCT_CATE2_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_DLCT_CATE2_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_DLCT_CATE2_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_DLCT_CATE2_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_DLCT_CATE2_MST_) TableInfo() *TableInfo {
	return USER_DLCT_CATE2_MSTTableInfo
}
