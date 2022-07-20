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


CREATE TABLE `ANPI_AREA_SYNC` (
  `ANPI_SYNC_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '安否同期ID',
  `REF_DATA_ID` int(11) NOT NULL COMMENT '引用データID',
  `REF_DATA_KND` char(1) NOT NULL COMMENT '引用データ種別	 A:安否起動DLCT U:ユーザID H:拠点ID G:起動G',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `PROC_KND` char(1) NOT NULL COMMENT '処理種別	 I:新規追加 U:更新 D:削除',
  `PROC_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '処理フラグ	 0:未処理 1:処理済',
  `PROC_END_DTE` datetime DEFAULT NULL COMMENT '処理終了日時',
  `PROC_MSG` varchar(512) DEFAULT NULL COMMENT '処理メッセージ',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`ANPI_SYNC_ID`),
  KEY `ANPI_AREA_SYNC_IDX1` (`REF_DATA_ID`,`CLIENT_ID`) USING BTREE,
  KEY `ANPI_AREA_SYNC_IDX2` (`PROC_FLG`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=29541957 DEFAULT CHARSET=utf8 COMMENT='安否起動DLCT同期	 初期データはTRIGGERで登録'

JSON Sample
-------------------------------------
{    "anpi_sync_id": 94,    "ref_data_id": 51,    "ref_data_knd": "wKeYTlVmBGOOmWcGlgclyQgSu",    "client_id": 27,    "proc_knd": "SDOiLhfleOfiKPlvdtBOCXTCW",    "proc_flg": "PXxBaDEkDAKHnARAwWVIlkWBO",    "proc_end_dte": "2190-10-03T17:09:58.816125021+09:00",    "proc_msg": "GNbgdArhGAZdAGMgHpJCSIyOj",    "upda_dte": "2261-09-16T22:30:43.118634231+09:00",    "upda_user_id": 58,    "crea_dte": "2075-08-03T21:35:50.899044635+09:00",    "crea_user_id": 8}



*/

// ANPI_AREA_SYNC_ struct is a row record of the ANPI_AREA_SYNC table in the anpidb database
type ANPI_AREA_SYNC_ struct {
	//[ 0] ANPI_SYNC_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANPISYNCID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANPI_SYNC_ID;type:int;"` // 安否同期ID
	//[ 1] REF_DATA_ID                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	REFDATAID int32 `gorm:"column:REF_DATA_ID;type:int;"` // 引用データID
	//[ 2] REF_DATA_KND                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	REFDATAKND string `gorm:"column:REF_DATA_KND;type:char;size:1;"` // 引用データ種別	 A:安否起動DLCT U:ユーザID H:拠点ID G:起動G
	//[ 3] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 4] PROC_KND                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	PROCKND string `gorm:"column:PROC_KND;type:char;size:1;"` // 処理種別	 I:新規追加 U:更新 D:削除
	//[ 5] PROC_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	PROCFLG string `gorm:"column:PROC_FLG;type:char;size:1;default:0;"` // 処理フラグ	 0:未処理 1:処理済
	//[ 6] PROC_END_DTE                                   datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	PROCENDDTE time.Time `gorm:"column:PROC_END_DTE;type:datetime;"` // 処理終了日時
	//[ 7] PROC_MSG                                       varchar(512)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 512     default: []
	PROCMSG sql.NullString `gorm:"column:PROC_MSG;type:varchar;size:512;"` // 処理メッセージ
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var ANPI_AREA_SYNCTableInfo = &TableInfo{
	Name: "ANPI_AREA_SYNC",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANPI_SYNC_ID",
			Comment:            `安否同期ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPISYNCID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_sync_id",
			ProtobufFieldName:  "anpi_sync_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "REF_DATA_ID",
			Comment:            `引用データID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "REFDATAID",
			GoFieldType:        "int32",
			JSONFieldName:      "ref_data_id",
			ProtobufFieldName:  "ref_data_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index: 2,
			Name:  "REF_DATA_KND",
			Comment: `引用データ種別	 A:安否起動DLCT U:ユーザID H:拠点ID G:起動G`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "REFDATAKND",
			GoFieldType:        "string",
			JSONFieldName:      "ref_data_knd",
			ProtobufFieldName:  "ref_data_knd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index: 4,
			Name:  "PROC_KND",
			Comment: `処理種別	 I:新規追加 U:更新 D:削除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PROCKND",
			GoFieldType:        "string",
			JSONFieldName:      "proc_knd",
			ProtobufFieldName:  "proc_knd",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index: 5,
			Name:  "PROC_FLG",
			Comment: `処理フラグ	 0:未処理 1:処理済`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PROCFLG",
			GoFieldType:        "string",
			JSONFieldName:      "proc_flg",
			ProtobufFieldName:  "proc_flg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "PROC_END_DTE",
			Comment:            `処理終了日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "PROCENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "proc_end_dte",
			ProtobufFieldName:  "proc_end_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "PROC_MSG",
			Comment:            `処理メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(512)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       512,
			GoFieldName:        "PROCMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "proc_msg",
			ProtobufFieldName:  "proc_msg",
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
func (a *ANPI_AREA_SYNC_) TableName() string {
	return "ANPI_AREA_SYNC"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_AREA_SYNC_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_AREA_SYNC_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_AREA_SYNC_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_AREA_SYNC_) TableInfo() *TableInfo {
	return ANPI_AREA_SYNCTableInfo
}
