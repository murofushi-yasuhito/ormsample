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


CREATE TABLE `USER_DEPT_MST` (
  `USER_DEPT_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT 'ユーザ部署ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `USER_DEPT_SEQ` int(11) NOT NULL COMMENT 'ユーザ部署連番',
  `DEPT_ID` int(11) NOT NULL COMMENT '所属している末端の部署ID',
  `DEPT_ID1` int(11) NOT NULL COMMENT '部署ID1',
  `DEPT_ID2` int(11) DEFAULT '0' COMMENT '部署ID2',
  `DEPT_ID3` int(11) DEFAULT '0' COMMENT '部署ID3',
  `DEPT_ID4` int(11) DEFAULT '0' COMMENT '部署ID4',
  `DEPT_ID5` int(11) DEFAULT '0' COMMENT '部署ID5',
  `DEPT_ID6` int(11) DEFAULT '0' COMMENT '部署ID6',
  `DEPT_ID7` int(11) DEFAULT '0' COMMENT '部署ID7',
  `DEPT_ID8` int(11) DEFAULT '0' COMMENT '部署ID8',
  `DEPT_ID9` int(11) DEFAULT '0' COMMENT '部署ID9',
  `DEPT_ID10` int(11) DEFAULT '0' COMMENT '部署ID10',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県コード',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村コード',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `AUTO_START_FLG` char(1) NOT NULL DEFAULT '1',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_DEPT_ID`),
  KEY `USER_DEPT_MST_IDX_1` (`USER_ID`) USING BTREE,
  KEY `USER_DEPT_MST_IDX_2` (`CLIENT_ID`,`DEPT_ID`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=10273947 DEFAULT CHARSET=utf8

JSON Sample
-------------------------------------
{    "user_dept_id": 89,    "user_id": 96,    "user_dept_seq": 68,    "dept_id": 85,    "dept_id_1": 73,    "dept_id_2": 41,    "dept_id_3": 50,    "dept_id_4": 36,    "dept_id_5": 12,    "dept_id_6": 64,    "dept_id_7": 15,    "dept_id_8": 9,    "dept_id_9": 7,    "dept_id_10": 46,    "pref_cd": "JgcxWNSxTXupxFWHnVyVIwhMG",    "city_cd": "IJqYhrAGxdabyuJGIelDimOHs",    "client_id": 73,    "auto_start_flg": "NUWAcLRApVyjsEucIVPTtayvs",    "upda_dte": "2275-08-03T02:39:11.048038145+09:00",    "upda_user_id": 55,    "crea_dte": "2270-05-25T07:50:34.005668428+09:00",    "crea_user_id": 15}



*/

// USER_DEPT_MST_ struct is a row record of the USER_DEPT_MST table in the anpidb database
type USER_DEPT_MST_ struct {
	//[ 0] USER_DEPT_ID                                   int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	USERDEPTID int32 `gorm:"primary_key;AUTO_INCREMENT;column:USER_DEPT_ID;type:int;"` // ユーザ部署ID
	//[ 1] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 2] USER_DEPT_SEQ                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERDEPTSEQ int32 `gorm:"column:USER_DEPT_SEQ;type:int;"` // ユーザ部署連番
	//[ 3] DEPT_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID int32 `gorm:"column:DEPT_ID;type:int;"` // 所属している末端の部署ID
	//[ 4] DEPT_ID1                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	DEPTID1 int32 `gorm:"column:DEPT_ID1;type:int;"` // 部署ID1
	//[ 5] DEPT_ID2                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID2 sql.NullInt64 `gorm:"column:DEPT_ID2;type:int;default:0;"` // 部署ID2
	//[ 6] DEPT_ID3                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID3 sql.NullInt64 `gorm:"column:DEPT_ID3;type:int;default:0;"` // 部署ID3
	//[ 7] DEPT_ID4                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID4 sql.NullInt64 `gorm:"column:DEPT_ID4;type:int;default:0;"` // 部署ID4
	//[ 8] DEPT_ID5                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID5 sql.NullInt64 `gorm:"column:DEPT_ID5;type:int;default:0;"` // 部署ID5
	//[ 9] DEPT_ID6                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID6 sql.NullInt64 `gorm:"column:DEPT_ID6;type:int;default:0;"` // 部署ID6
	//[10] DEPT_ID7                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID7 sql.NullInt64 `gorm:"column:DEPT_ID7;type:int;default:0;"` // 部署ID7
	//[11] DEPT_ID8                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID8 sql.NullInt64 `gorm:"column:DEPT_ID8;type:int;default:0;"` // 部署ID8
	//[12] DEPT_ID9                                       int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID9 sql.NullInt64 `gorm:"column:DEPT_ID9;type:int;default:0;"` // 部署ID9
	//[13] DEPT_ID10                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	DEPTID10 sql.NullInt64 `gorm:"column:DEPT_ID10;type:int;default:0;"` // 部署ID10
	//[14] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[15] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[16] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[17] AUTO_START_FLG                                 char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	AUTOSTARTFLG string `gorm:"column:AUTO_START_FLG;type:char;size:1;default:1;"`
	//[18] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[19] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[20] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[21] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_DEPT_MSTTableInfo = &TableInfo{
	Name: "USER_DEPT_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "USER_DEPT_ID",
			Comment:            `ユーザ部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERDEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_dept_id",
			ProtobufFieldName:  "user_dept_id",
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
			Index:              2,
			Name:               "USER_DEPT_SEQ",
			Comment:            `ユーザ部署連番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERDEPTSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "user_dept_seq",
			ProtobufFieldName:  "user_dept_seq",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "DEPT_ID",
			Comment:            `所属している末端の部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id",
			ProtobufFieldName:  "dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "DEPT_ID1",
			Comment:            `部署ID1`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID1",
			GoFieldType:        "int32",
			JSONFieldName:      "dept_id_1",
			ProtobufFieldName:  "dept_id_1",
			ProtobufType:       "int32",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "DEPT_ID2",
			Comment:            `部署ID2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID2",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_2",
			ProtobufFieldName:  "dept_id_2",
			ProtobufType:       "int32",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "DEPT_ID3",
			Comment:            `部署ID3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID3",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_3",
			ProtobufFieldName:  "dept_id_3",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "DEPT_ID4",
			Comment:            `部署ID4`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID4",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_4",
			ProtobufFieldName:  "dept_id_4",
			ProtobufType:       "int32",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "DEPT_ID5",
			Comment:            `部署ID5`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID5",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_5",
			ProtobufFieldName:  "dept_id_5",
			ProtobufType:       "int32",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "DEPT_ID6",
			Comment:            `部署ID6`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID6",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_6",
			ProtobufFieldName:  "dept_id_6",
			ProtobufType:       "int32",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "DEPT_ID7",
			Comment:            `部署ID7`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID7",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_7",
			ProtobufFieldName:  "dept_id_7",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "DEPT_ID8",
			Comment:            `部署ID8`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID8",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_8",
			ProtobufFieldName:  "dept_id_8",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "DEPT_ID9",
			Comment:            `部署ID9`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID9",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_9",
			ProtobufFieldName:  "dept_id_9",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "DEPT_ID10",
			Comment:            `部署ID10`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "DEPTID10",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "dept_id_10",
			ProtobufFieldName:  "dept_id_10",
			ProtobufType:       "int32",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "PREF_CD",
			Comment:            `都道府県コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "PREFCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "pref_cd",
			ProtobufFieldName:  "pref_cd",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "CITY_CD",
			Comment:            `市区町村コード`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(5)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       5,
			GoFieldName:        "CITYCD",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "city_cd",
			ProtobufFieldName:  "city_cd",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
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
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "AUTO_START_FLG",
			Comment:            ``,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AUTOSTARTFLG",
			GoFieldType:        "string",
			JSONFieldName:      "auto_start_flg",
			ProtobufFieldName:  "auto_start_flg",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
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
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
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
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
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
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
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
			ProtobufPos:        22,
		},
	},
}

// TableName sets the insert table name for this struct type
func (u *USER_DEPT_MST_) TableName() string {
	return "USER_DEPT_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_DEPT_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_DEPT_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_DEPT_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_DEPT_MST_) TableInfo() *TableInfo {
	return USER_DEPT_MSTTableInfo
}
