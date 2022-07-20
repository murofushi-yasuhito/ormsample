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


CREATE TABLE `SUMM_GRP_MST` (
  `SUMM_GRP_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '集計グループID',
  `SUMM_SEQ` int(11) NOT NULL COMMENT '集計グループ連番',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `SUMM_GRP_NM` char(50) NOT NULL COMMENT '集計グループ名',
  `BATCH_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '一括処理中フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`SUMM_GRP_ID`),
  KEY `SUMM_GRP_MST_IDX_1` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1511 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "summ_grp_id": 55,    "summ_seq": 63,    "client_id": 11,    "summ_grp_nm": "beEWKvdqKMJpZQerFaWVmILTW",    "batch_flg": "oJCGhxhmopUxQsvVFbxFJAdEX",    "upda_dte": "2177-05-13T17:27:32.351864768+09:00",    "upda_user_id": 11,    "crea_dte": "2272-03-25T07:16:54.352681752+09:00",    "crea_user_id": 43}



*/

// SUMM_GRP_MST_ struct is a row record of the SUMM_GRP_MST table in the anpidb database
type SUMM_GRP_MST_ struct {
	//[ 0] SUMM_GRP_ID                                    int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	SUMMGRPID int32 `gorm:"primary_key;AUTO_INCREMENT;column:SUMM_GRP_ID;type:int;"` // 集計グループID
	//[ 1] SUMM_SEQ                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SUMMSEQ int32 `gorm:"column:SUMM_SEQ;type:int;"` // 集計グループ連番
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] SUMM_GRP_NM                                    char(50)             null: false  primary: false  isArray: false  auto: false  col: char            len: 50      default: []
	SUMMGRPNM string `gorm:"column:SUMM_GRP_NM;type:char;size:50;"` // 集計グループ名
	//[ 4] BATCH_FLG                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	BATCHFLG string `gorm:"column:BATCH_FLG;type:char;size:1;default:0;"` // 一括処理中フラグ
	//[ 5] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 6] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 7] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[ 8] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var SUMM_GRP_MSTTableInfo = &TableInfo{
	Name: "SUMM_GRP_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "SUMM_GRP_ID",
			Comment:            `集計グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUMMGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "summ_grp_id",
			ProtobufFieldName:  "summ_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "SUMM_SEQ",
			Comment:            `集計グループ連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUMMSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "summ_seq",
			ProtobufFieldName:  "summ_seq",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "CLIENT_ID",
			Comment:            `クライアントID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_id",
			ProtobufFieldName:  "client_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "SUMM_GRP_NM",
			Comment:            `集計グループ名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(50)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       50,
			GoFieldName:        "SUMMGRPNM",
			GoFieldType:        "string",
			JSONFieldName:      "summ_grp_nm",
			ProtobufFieldName:  "summ_grp_nm",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "BATCH_FLG",
			Comment:            `一括処理中フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "BATCHFLG",
			GoFieldType:        "string",
			JSONFieldName:      "batch_flg",
			ProtobufFieldName:  "batch_flg",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
func (s *SUMM_GRP_MST_) TableName() string {
	return "SUMM_GRP_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *SUMM_GRP_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *SUMM_GRP_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *SUMM_GRP_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *SUMM_GRP_MST_) TableInfo() *TableInfo {
	return SUMM_GRP_MSTTableInfo
}
