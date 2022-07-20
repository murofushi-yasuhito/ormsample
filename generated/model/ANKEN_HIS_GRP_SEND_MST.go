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


CREATE TABLE `ANKEN_HIS_GRP_SEND_MST` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `GRP_ID` int(11) NOT NULL COMMENT 'グループID',
  `GRP_SEND_SEQ` int(11) NOT NULL DEFAULT '1' COMMENT 'グループ送信先連番',
  `GRP_KBN` char(1) NOT NULL DEFAULT '1' COMMENT 'グループ区分	 1:任意指定 2:ユーザ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`GRP_ID`,`GRP_SEND_SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "anken_id": 63,    "grp_id": 62,    "grp_send_seq": 14,    "grp_kbn": "qVFKTAmRVPYeMUSAtJbipRPHo",    "upda_dte": "2176-08-17T13:45:20.393993553+09:00",    "upda_user_id": 83,    "crea_dte": "2089-01-17T07:27:37.310593319+09:00",    "crea_user_id": 91}



*/

// ANKEN_HIS_GRP_SEND_MST_ struct is a row record of the ANKEN_HIS_GRP_SEND_MST table in the anpidb database
type ANKEN_HIS_GRP_SEND_MST_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] GRP_ID                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	GRPID int32 `gorm:"primary_key;column:GRP_ID;type:int;"` // グループID
	//[ 2] GRP_SEND_SEQ                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [1]
	GRPSENDSEQ int32 `gorm:"primary_key;column:GRP_SEND_SEQ;type:int;default:1;"` // グループ送信先連番
	//[ 3] GRP_KBN                                        char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	GRPKBN string `gorm:"column:GRP_KBN;type:char;size:1;default:1;"` // グループ区分	 1:任意指定 2:ユーザ
	//[ 4] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 5] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 6] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 7] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_HIS_GRP_SEND_MSTTableInfo = &TableInfo{
	Name: "ANKEN_HIS_GRP_SEND_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANKEN_ID",
			Comment:            `案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_id",
			ProtobufFieldName:  "anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "GRP_ID",
			Comment:            `グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_id",
			ProtobufFieldName:  "grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "GRP_SEND_SEQ",
			Comment:            `グループ送信先連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPSENDSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_send_seq",
			ProtobufFieldName:  "grp_send_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "GRP_KBN",
			Comment: `グループ区分	 1:任意指定 2:ユーザ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "GRPKBN",
			GoFieldType:        "string",
			JSONFieldName:      "grp_kbn",
			ProtobufFieldName:  "grp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
func (a *ANKEN_HIS_GRP_SEND_MST_) TableName() string {
	return "ANKEN_HIS_GRP_SEND_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_HIS_GRP_SEND_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_HIS_GRP_SEND_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_HIS_GRP_SEND_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_HIS_GRP_SEND_MST_) TableInfo() *TableInfo {
	return ANKEN_HIS_GRP_SEND_MSTTableInfo
}
