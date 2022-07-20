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


CREATE TABLE `DEFAULT_DLCT_MST` (
  `DEFAULT_DLCT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザDLCT_ID',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード	 画面選択カテゴリ',
  `DEPT_ADDR_KBN` char(1) DEFAULT NULL,
  `HUB_ADDR_KBN` char(1) DEFAULT NULL,
  `HOME_ADDR_KBN` char(1) DEFAULT NULL,
  `AREA_KBN` char(4) NOT NULL,
  `LEVEL_VALUE` int(11) DEFAULT '99' COMMENT 'レベル',
  `LEVEL_KBN` char(1) DEFAULT '0' COMMENT 'レベル区分',
  `DELI_TIME_HM` int(11) DEFAULT NULL COMMENT '配信時刻時分	 天気予報配信時分hhmm',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL DEFAULT '0' COMMENT '作成者ID',
  PRIMARY KEY (`DEFAULT_DLCT_ID`),
  KEY `USER_DLCT_MST_IDX_1` (`CATE1_CD`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=996 DEFAULT CHARSET=utf8 COMMENT='ユーザDLCTマスタ'

JSON Sample
-------------------------------------
{    "default_dlct_id": 25,    "client_id": 38,    "cate_1_cd": "wHIvPLwXrJIDOxdLRZDYspqYB",    "dept_addr_kbn": "ptqugVWQGLYQxyqvckeKFiDcY",    "hub_addr_kbn": "nsjMIGNNHIEnNDjTguujCeiCO",    "home_addr_kbn": "JMbIawmHcvkJjPkHDBkpvqjZN",    "area_kbn": "JiGObYhwhdAWSEkDiSiPiuTkG",    "level_value": 77,    "level_kbn": "TnEpxlBdMUOWVmwLwqHjIsfXK",    "deli_time_hm": 72,    "upda_dte": "2039-04-02T12:34:29.129889601+09:00",    "upda_user_id": 6,    "crea_dte": "2178-08-14T22:29:29.177514789+09:00",    "crea_user_id": 76}



*/

// DEFAULT_DLCT_MST_ struct is a row record of the DEFAULT_DLCT_MST table in the anpidb database
type DEFAULT_DLCT_MST_ struct {
	//[ 0] DEFAULT_DLCT_ID                                int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	DEFAULTDLCTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:DEFAULT_DLCT_ID;type:int;"` // ユーザDLCT_ID
	//[ 1] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 2] CATE1_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード	 画面選択カテゴリ
	//[ 3] DEPT_ADDR_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	DEPTADDRKBN sql.NullString `gorm:"column:DEPT_ADDR_KBN;type:char;size:1;"`
	//[ 4] HUB_ADDR_KBN                                   char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	HUBADDRKBN sql.NullString `gorm:"column:HUB_ADDR_KBN;type:char;size:1;"`
	//[ 5] HOME_ADDR_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	HOMEADDRKBN sql.NullString `gorm:"column:HOME_ADDR_KBN;type:char;size:1;"`
	//[ 6] AREA_KBN                                       char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:char;size:4;"`
	//[ 7] LEVEL_VALUE                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [99]
	LEVELVALUE sql.NullInt64 `gorm:"column:LEVEL_VALUE;type:int;default:99;"` // レベル
	//[ 8] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;default:0;"` // レベル区分
	//[ 9] DELI_TIME_HM                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DELITIMEHM sql.NullInt64 `gorm:"column:DELI_TIME_HM;type:int;"` // 配信時刻時分	 天気予報配信時分hhmm
	//[10] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[11] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;default:0;"` // 更新者ID
	//[12] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[13] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;default:0;"` // 作成者ID

}

var DEFAULT_DLCT_MSTTableInfo = &TableInfo{
	Name: "DEFAULT_DLCT_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "DEFAULT_DLCT_ID",
			Comment:            `ユーザDLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEFAULTDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "default_dlct_id",
			ProtobufFieldName:  "default_dlct_id",
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
			Index: 2,
			Name:  "CATE1_CD",
			Comment: `カテゴリ1コード	 画面選択カテゴリ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "CATE1CD",
			GoFieldType:        "string",
			JSONFieldName:      "cate_1_cd",
			ProtobufFieldName:  "cate_1_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DEPT_ADDR_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEPTADDRKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_addr_kbn",
			ProtobufFieldName:  "dept_addr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "HUB_ADDR_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "HUBADDRKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "hub_addr_kbn",
			ProtobufFieldName:  "hub_addr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "HOME_ADDR_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "HOMEADDRKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "home_addr_kbn",
			ProtobufFieldName:  "home_addr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "AREA_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "AREAKBN",
			GoFieldType:        "string",
			JSONFieldName:      "area_kbn",
			ProtobufFieldName:  "area_kbn",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index: 9,
			Name:  "DELI_TIME_HM",
			Comment: `配信時刻時分	 天気予報配信時分hhmm`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DELITIMEHM",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "deli_time_hm",
			ProtobufFieldName:  "deli_time_hm",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},
	},
}

// TableName sets the insert table name for this struct type
func (d *DEFAULT_DLCT_MST_) TableName() string {
	return "DEFAULT_DLCT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEFAULT_DLCT_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEFAULT_DLCT_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEFAULT_DLCT_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DEFAULT_DLCT_MST_) TableInfo() *TableInfo {
	return DEFAULT_DLCT_MSTTableInfo
}
