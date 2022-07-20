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


CREATE TABLE `START_GRP_MST` (
  `START_GRP_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '起動グループID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `START_GRP_NM` varchar(60) DEFAULT NULL COMMENT '起動グループ名',
  `ANPI_DLCT_ID` int(11) NOT NULL,
  `AREA_KBN` varchar(4) NOT NULL,
  `GRP_ID` int(11) NOT NULL,
  `BATCH_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '一括処理中フラグ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`START_GRP_ID`),
  KEY `START_GRP_MST_IDX_1` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2560 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "start_grp_id": 3,    "client_id": 75,    "start_grp_nm": "LUvplqXrdgwwQaQEyBjGHauhK",    "anpi_dlct_id": 53,    "area_kbn": "lRirpVxGkrdQrNbZLtiteTUYA",    "grp_id": 36,    "batch_flg": "pQEhnODNZQPZpwVOLscswShXX",    "upda_dte": "2154-03-20T10:01:16.129937757+09:00",    "upda_user_id": 81,    "crea_dte": "2087-10-13T12:08:24.604267017+09:00",    "crea_user_id": 33}



*/

// START_GRP_MST_ struct is a row record of the START_GRP_MST table in the anpidb database
type START_GRP_MST_ struct {
	//[ 0] START_GRP_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	STARTGRPID int32 `gorm:"primary_key;AUTO_INCREMENT;column:START_GRP_ID;type:int;"` // 起動グループID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] START_GRP_NM                                   varchar(60)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 60      default: []
	STARTGRPNM sql.NullString `gorm:"column:START_GRP_NM;type:varchar;size:60;"` // 起動グループ名
	//[ 3] ANPI_DLCT_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANPIDLCTID int32 `gorm:"column:ANPI_DLCT_ID;type:int;"`
	//[ 4] AREA_KBN                                       varchar(4)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:varchar;size:4;"`
	//[ 5] GRP_ID                                         int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	GRPID int32 `gorm:"column:GRP_ID;type:int;"`
	//[ 6] BATCH_FLG                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	BATCHFLG string `gorm:"column:BATCH_FLG;type:char;size:1;default:0;"` // 一括処理中フラグ
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[10] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var START_GRP_MSTTableInfo = &TableInfo{
	Name: "START_GRP_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "START_GRP_ID",
			Comment:            `起動グループID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "STARTGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_id",
			ProtobufFieldName:  "start_grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "START_GRP_NM",
			Comment:            `起動グループ名`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(60)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       60,
			GoFieldName:        "STARTGRPNM",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "start_grp_nm",
			ProtobufFieldName:  "start_grp_nm",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "ANPI_DLCT_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPIDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_dlct_id",
			ProtobufFieldName:  "anpi_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "AREA_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       4,
			GoFieldName:        "AREAKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_kbn",
			ProtobufFieldName:  "area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "GRP_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "GRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "grp_id",
			ProtobufFieldName:  "grp_id",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
func (s *START_GRP_MST_) TableName() string {
	return "START_GRP_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (s *START_GRP_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (s *START_GRP_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (s *START_GRP_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (s *START_GRP_MST_) TableInfo() *TableInfo {
	return START_GRP_MSTTableInfo
}
