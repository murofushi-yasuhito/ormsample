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


CREATE TABLE `USER_DLCT_MST` (
  `USER_DLCT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザDLCT_ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `CATE1_CD` varchar(10) NOT NULL COMMENT 'カテゴリ1コード	 画面選択カテゴリ',
  `LEVEL_VALUE` int(11) DEFAULT '99' COMMENT 'レベル',
  `LEVEL_KBN` char(1) DEFAULT '0' COMMENT 'レベル区分',
  `DELI_TIME_HM` int(11) DEFAULT NULL COMMENT '配信時刻時分	 天気予報配信時分hhmm',
  `CLIENT_DLCT_ID` int(11) DEFAULT NULL COMMENT 'クライアントDLCT_ID',
  `CLIENT_DLCT_UPD_DTE` datetime DEFAULT NULL COMMENT 'クライアントDLCT更新日時',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `DEFAULT_DLCT_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'デフォルトDLCTから作成の場合DEFALT_DLCT_ID',
  `CONT_PROTECT` char(1) NOT NULL DEFAULT '0' COMMENT '0:デフォルトDLCTレベル連動可 1:デフォルトDLCTレベル連動不可',
  `VALID_KBN` char(1) NOT NULL DEFAULT '1' COMMENT '有効区分',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_DLCT_ID`),
  KEY `USER_DLCT_MST_IDX_1` (`CATE1_CD`) USING BTREE,
  KEY `USER_DLCT_MST_IDX_2` (`USER_ID`) USING BTREE,
  KEY `USER_DLCT_MST_IDX_3` (`CLIENT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=19638153 DEFAULT CHARSET=utf8 COMMENT='ユーザDLCTマスタ'

JSON Sample
-------------------------------------
{    "user_dlct_id": 11,    "user_id": 91,    "cate_1_cd": "KRheSbfNcgNxrtpIoedcMnCqu",    "level_value": 59,    "level_kbn": "VZowhZVlqOsSvdYXiZjWYggHg",    "deli_time_hm": 69,    "client_dlct_id": 73,    "client_dlct_upd_dte": "2169-07-28T08:11:17.299882589+09:00",    "client_id": 51,    "default_dlct_id": 69,    "cont_protect": "CHdAJYdmSSXLmdhqXesPnjAni",    "valid_kbn": "fVeHSnbOhWJnmwMuHysIljRiT",    "upda_dte": "2156-01-21T15:17:51.750113404+09:00",    "upda_user_id": 47,    "crea_dte": "2238-06-22T05:34:30.03691325+09:00",    "crea_user_id": 71}



*/

// USER_DLCT_MST_ struct is a row record of the USER_DLCT_MST table in the anpidb database
type USER_DLCT_MST_ struct {
	//[ 0] USER_DLCT_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	USERDLCTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:USER_DLCT_ID;type:int;"` // ユーザDLCT_ID
	//[ 1] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 2] CATE1_CD                                       varchar(10)          null: false  primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	CATE1CD string `gorm:"column:CATE1_CD;type:varchar;size:10;"` // カテゴリ1コード	 画面選択カテゴリ
	//[ 3] LEVEL_VALUE                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [99]
	LEVELVALUE sql.NullInt64 `gorm:"column:LEVEL_VALUE;type:int;default:99;"` // レベル
	//[ 4] LEVEL_KBN                                      char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	LEVELKBN sql.NullString `gorm:"column:LEVEL_KBN;type:char;size:1;default:0;"` // レベル区分
	//[ 5] DELI_TIME_HM                                   int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DELITIMEHM sql.NullInt64 `gorm:"column:DELI_TIME_HM;type:int;"` // 配信時刻時分	 天気予報配信時分hhmm
	//[ 6] CLIENT_DLCT_ID                                 int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTDLCTID sql.NullInt64 `gorm:"column:CLIENT_DLCT_ID;type:int;"` // クライアントDLCT_ID
	//[ 7] CLIENT_DLCT_UPD_DTE                            datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	CLIENTDLCTUPDDTE time.Time `gorm:"column:CLIENT_DLCT_UPD_DTE;type:datetime;"` // クライアントDLCT更新日時
	//[ 8] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 9] DEFAULT_DLCT_ID                                int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEFAULTDLCTID int32 `gorm:"column:DEFAULT_DLCT_ID;type:int;default:0;"` // デフォルトDLCTから作成の場合DEFALT_DLCT_ID
	//[10] CONT_PROTECT                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	CONTPROTECT string `gorm:"column:CONT_PROTECT;type:char;size:1;default:0;"` // 0:デフォルトDLCTレベル連動可 1:デフォルトDLCTレベル連動不可
	//[11] VALID_KBN                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	VALIDKBN string `gorm:"column:VALID_KBN;type:char;size:1;default:1;"` // 有効区分
	//[12] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[13] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_DLCT_MSTTableInfo = &TableInfo{
	Name: "USER_DLCT_MST",
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
			IsAutoIncrement:    true,
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
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_id",
			ProtobufFieldName:  "user_id",
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
			Index: 5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "CLIENT_DLCT_ID",
			Comment:            `クライアントDLCT_ID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTDLCTID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "client_dlct_id",
			ProtobufFieldName:  "client_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "CLIENT_DLCT_UPD_DTE",
			Comment:            `クライアントDLCT更新日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTDLCTUPDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "client_dlct_upd_dte",
			ProtobufFieldName:  "client_dlct_upd_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "DEFAULT_DLCT_ID",
			Comment:            `デフォルトDLCTから作成の場合DEFALT_DLCT_ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEFAULTDLCTID",
			GoFieldType:        "int32",
			JSONFieldName:      "default_dlct_id",
			ProtobufFieldName:  "default_dlct_id",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "CONT_PROTECT",
			Comment:            `0:デフォルトDLCTレベル連動可 1:デフォルトDLCTレベル連動不可`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "CONTPROTECT",
			GoFieldType:        "string",
			JSONFieldName:      "cont_protect",
			ProtobufFieldName:  "cont_protect",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "VALID_KBN",
			Comment:            `有効区分`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "VALIDKBN",
			GoFieldType:        "string",
			JSONFieldName:      "valid_kbn",
			ProtobufFieldName:  "valid_kbn",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
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
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
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
			ProtobufPos:        16,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *USER_DLCT_MST_) TableName() string {
	return "USER_DLCT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_DLCT_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_DLCT_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_DLCT_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_DLCT_MST_) TableInfo() *TableInfo {
	return USER_DLCT_MSTTableInfo
}
