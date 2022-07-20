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


CREATE TABLE `MST_UPD_MNG` (
  `MST_UPD_MNG_ID` int(11) NOT NULL AUTO_INCREMENT,
  `EXCLUSION_KBN` char(1) NOT NULL COMMENT '排他タイプ　1:システム単位　2:クライアント単位',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `LOGIC_CD` char(11) NOT NULL COMMENT '処理ID',
  `DATA_FLG` int(1) NOT NULL COMMENT 'データフラグ　1:新規 2:編集 3:削除',
  `PROC_STS` char(1) NOT NULL COMMENT '実行状態　1:実行待ち？ 2:処理中 3:中断 4:終了',
  `CONFIRM_FLG` char(1) DEFAULT NULL COMMENT '画面確認　1:未確認 2:確認済み',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MST_UPD_MNG_ID`),
  KEY `MST_UPD_MNG_IDX1` (`EXCLUSION_KBN`,`CLIENT_ID`,`PROC_STS`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=933 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "mst_upd_mng_id": 12,    "exclusion_kbn": "OiuXJWUXwuKtQNSkMpdWmRGOn",    "client_id": 52,    "logic_cd": "QVRCtPtDFrjNWdRUJtBhscbcI",    "data_flg": 46,    "proc_sts": "DQmonPQOlnbjfbagiYxHMewhq",    "confirm_flg": "SxWJcssPelidXpycuhstkmYoE",    "upda_dte": "2203-06-13T17:34:23.841178963+09:00",    "upda_user_id": 96,    "crea_dte": "2297-04-26T10:13:48.772805662+09:00",    "crea_user_id": 29}



*/

// MST_UPD_MNG_ struct is a row record of the MST_UPD_MNG table in the anpidb database
type MST_UPD_MNG_ struct {
	//[ 0] MST_UPD_MNG_ID                                 int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	MSTUPDMNGID int32 `gorm:"primary_key;AUTO_INCREMENT;column:MST_UPD_MNG_ID;type:int;"`
	//[ 1] EXCLUSION_KBN                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	EXCLUSIONKBN string `gorm:"column:EXCLUSION_KBN;type:char;size:1;"` // 排他タイプ　1:システム単位　2:クライアント単位
	//[ 2] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 3] LOGIC_CD                                       char(11)             null: false  primary: false  isArray: false  auto: false  col: char            len: 11      default: []
	LOGICCD string `gorm:"column:LOGIC_CD;type:char;size:11;"` // 処理ID
	//[ 4] DATA_FLG                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DATAFLG int32 `gorm:"column:DATA_FLG;type:int;"` // データフラグ　1:新規 2:編集 3:削除
	//[ 5] PROC_STS                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	PROCSTS string `gorm:"column:PROC_STS;type:char;size:1;"` // 実行状態　1:実行待ち？ 2:処理中 3:中断 4:終了
	//[ 6] CONFIRM_FLG                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	CONFIRMFLG sql.NullString `gorm:"column:CONFIRM_FLG;type:char;size:1;"` // 画面確認　1:未確認 2:確認済み
	//[ 7] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 8] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[ 9] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[10] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var MST_UPD_MNGTableInfo = &TableInfo{
	Name: "MST_UPD_MNG",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MST_UPD_MNG_ID",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MSTUPDMNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "mst_upd_mng_id",
			ProtobufFieldName:  "mst_upd_mng_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "EXCLUSION_KBN",
			Comment:            `排他タイプ　1:システム単位　2:クライアント単位`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "EXCLUSIONKBN",
			GoFieldType:        "string",
			JSONFieldName:      "exclusion_kbn",
			ProtobufFieldName:  "exclusion_kbn",
			ProtobufType:       "string",
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
			Name:               "LOGIC_CD",
			Comment:            `処理ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(11)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       11,
			GoFieldName:        "LOGICCD",
			GoFieldType:        "string",
			JSONFieldName:      "logic_cd",
			ProtobufFieldName:  "logic_cd",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "DATA_FLG",
			Comment:            `データフラグ　1:新規 2:編集 3:削除`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DATAFLG",
			GoFieldType:        "int32",
			JSONFieldName:      "data_flg",
			ProtobufFieldName:  "data_flg",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "PROC_STS",
			Comment:            `実行状態　1:実行待ち？ 2:処理中 3:中断 4:終了`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "PROCSTS",
			GoFieldType:        "string",
			JSONFieldName:      "proc_sts",
			ProtobufFieldName:  "proc_sts",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CONFIRM_FLG",
			Comment:            `画面確認　1:未確認 2:確認済み`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CONFIRMFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "confirm_flg",
			ProtobufFieldName:  "confirm_flg",
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
func (m *MST_UPD_MNG_) TableName() string {
	return "MST_UPD_MNG"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (m *MST_UPD_MNG_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (m *MST_UPD_MNG_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (m *MST_UPD_MNG_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (m *MST_UPD_MNG_) TableInfo() *TableInfo {
	return MST_UPD_MNGTableInfo
}
