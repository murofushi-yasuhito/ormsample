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


CREATE TABLE `ANKEN_SEND_USER_DEPT` (
  `ANKEN_ID` int(11) NOT NULL COMMENT '案件ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `USER_DEPT_ID` int(11) NOT NULL COMMENT 'ユーザ部署ID',
  `USER_DEPT_SEQ` int(11) NOT NULL COMMENT 'ユーザ部署連番',
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
  `DEPT_NM1` varchar(200) DEFAULT NULL COMMENT '部署名1',
  `DEPT_NM2` varchar(200) DEFAULT NULL COMMENT '部署名2',
  `DEPT_NM3` varchar(200) DEFAULT NULL COMMENT '部署名3',
  `DEPT_NM4` varchar(200) DEFAULT NULL COMMENT '部署名4',
  `DEPT_NM5` varchar(200) DEFAULT NULL COMMENT '部署名5',
  `DEPT_NM6` varchar(200) DEFAULT NULL COMMENT '部署名6',
  `DEPT_NM7` varchar(200) DEFAULT NULL COMMENT '部署名7',
  `DEPT_NM8` varchar(200) DEFAULT NULL COMMENT '部署名8',
  `DEPT_NM9` varchar(200) DEFAULT NULL COMMENT '部署名9',
  `DEPT_NM10` varchar(200) DEFAULT NULL COMMENT '部署名10',
  `PREF_CD` char(2) DEFAULT NULL COMMENT '都道府県コード',
  `CITY_CD` char(5) DEFAULT NULL COMMENT '市区町村コード',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`,`USER_ID`,`USER_DEPT_ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='案件送信ユーザ部署'

JSON Sample
-------------------------------------
{    "anken_id": 57,    "user_id": 20,    "user_dept_id": 54,    "user_dept_seq": 10,    "dept_id_1": 20,    "dept_id_2": 10,    "dept_id_3": 66,    "dept_id_4": 35,    "dept_id_5": 15,    "dept_id_6": 35,    "dept_id_7": 38,    "dept_id_8": 74,    "dept_id_9": 73,    "dept_id_10": 17,    "dept_nm_1": "wcIhkexVqndtXqXvCCoJabmSu",    "dept_nm_2": "pUAgcHSrJmoUhFcEtFWHSUeUI",    "dept_nm_3": "dUcmMbtKLNdtlgTBrrfNdfMVc",    "dept_nm_4": "eFpaMufZDGFZkbuNJBsopvxbM",    "dept_nm_5": "IZkJdRVlewIVhqjBdUCeETHJB",    "dept_nm_6": "hriphyTJrhLlBLsxihJOhaGQu",    "dept_nm_7": "OhkxcpridUjcRYJiTWevkmSLy",    "dept_nm_8": "UxgQKbpdhXTiaMAjRLrgQBiNG",    "dept_nm_9": "tgHFtgDgDLkfwjaQXZvRHnbOd",    "dept_nm_10": "ikNyUhZHahWLEQelKcqCbQpUS",    "pref_cd": "vQBOqpMjhUbxocExVinVolLCT",    "city_cd": "DukaquwcIEifyuBGbFbvGmsFy",    "client_id": 3,    "crea_dte": "2172-01-08T15:22:37.603418817+09:00",    "crea_user_id": 59}



*/

// ANKEN_SEND_USER_DEPT_ struct is a row record of the ANKEN_SEND_USER_DEPT table in the anpidb database
type ANKEN_SEND_USER_DEPT_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] USER_ID                                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"primary_key;column:USER_ID;type:int;"` // ユーザID
	//[ 2] USER_DEPT_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	USERDEPTID int32 `gorm:"primary_key;column:USER_DEPT_ID;type:int;"` // ユーザ部署ID
	//[ 3] USER_DEPT_SEQ                                  int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERDEPTSEQ int32 `gorm:"column:USER_DEPT_SEQ;type:int;"` // ユーザ部署連番
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
	//[14] DEPT_NM1                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM1 sql.NullString `gorm:"column:DEPT_NM1;type:varchar;size:200;"` // 部署名1
	//[15] DEPT_NM2                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM2 sql.NullString `gorm:"column:DEPT_NM2;type:varchar;size:200;"` // 部署名2
	//[16] DEPT_NM3                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM3 sql.NullString `gorm:"column:DEPT_NM3;type:varchar;size:200;"` // 部署名3
	//[17] DEPT_NM4                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM4 sql.NullString `gorm:"column:DEPT_NM4;type:varchar;size:200;"` // 部署名4
	//[18] DEPT_NM5                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM5 sql.NullString `gorm:"column:DEPT_NM5;type:varchar;size:200;"` // 部署名5
	//[19] DEPT_NM6                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM6 sql.NullString `gorm:"column:DEPT_NM6;type:varchar;size:200;"` // 部署名6
	//[20] DEPT_NM7                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM7 sql.NullString `gorm:"column:DEPT_NM7;type:varchar;size:200;"` // 部署名7
	//[21] DEPT_NM8                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM8 sql.NullString `gorm:"column:DEPT_NM8;type:varchar;size:200;"` // 部署名8
	//[22] DEPT_NM9                                       varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM9 sql.NullString `gorm:"column:DEPT_NM9;type:varchar;size:200;"` // 部署名9
	//[23] DEPT_NM10                                      varchar(200)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 200     default: []
	DEPTNM10 sql.NullString `gorm:"column:DEPT_NM10;type:varchar;size:200;"` // 部署名10
	//[24] PREF_CD                                        char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	PREFCD sql.NullString `gorm:"column:PREF_CD;type:char;size:2;"` // 都道府県コード
	//[25] CITY_CD                                        char(5)              null: true   primary: false  isArray: false  auto: false  col: char            len: 5       default: []
	CITYCD sql.NullString `gorm:"column:CITY_CD;type:char;size:5;"` // 市区町村コード
	//[26] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[27] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[28] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKEN_SEND_USER_DEPTTableInfo = &TableInfo{
	Name: "ANKEN_SEND_USER_DEPT",
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
			Name:               "USER_ID",
			Comment:            `ユーザID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
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
			Name:               "USER_DEPT_ID",
			Comment:            `ユーザ部署ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "USERDEPTID",
			GoFieldType:        "int32",
			JSONFieldName:      "user_dept_id",
			ProtobufFieldName:  "user_dept_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			Name:               "DEPT_NM1",
			Comment:            `部署名1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM1",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_1",
			ProtobufFieldName:  "dept_nm_1",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "DEPT_NM2",
			Comment:            `部署名2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_2",
			ProtobufFieldName:  "dept_nm_2",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "DEPT_NM3",
			Comment:            `部署名3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_3",
			ProtobufFieldName:  "dept_nm_3",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "DEPT_NM4",
			Comment:            `部署名4`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM4",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_4",
			ProtobufFieldName:  "dept_nm_4",
			ProtobufType:       "string",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "DEPT_NM5",
			Comment:            `部署名5`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM5",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_5",
			ProtobufFieldName:  "dept_nm_5",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "DEPT_NM6",
			Comment:            `部署名6`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM6",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_6",
			ProtobufFieldName:  "dept_nm_6",
			ProtobufType:       "string",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "DEPT_NM7",
			Comment:            `部署名7`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM7",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_7",
			ProtobufFieldName:  "dept_nm_7",
			ProtobufType:       "string",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "DEPT_NM8",
			Comment:            `部署名8`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM8",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_8",
			ProtobufFieldName:  "dept_nm_8",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "DEPT_NM9",
			Comment:            `部署名9`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM9",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_9",
			ProtobufFieldName:  "dept_nm_9",
			ProtobufType:       "string",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "DEPT_NM10",
			Comment:            `部署名10`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(200)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       200,
			GoFieldName:        "DEPTNM10",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_nm_10",
			ProtobufFieldName:  "dept_nm_10",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
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
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
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
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
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
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
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
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
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
			ProtobufPos:        29,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_SEND_USER_DEPT_) TableName() string {
	return "ANKEN_SEND_USER_DEPT"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_SEND_USER_DEPT_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_SEND_USER_DEPT_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_SEND_USER_DEPT_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_SEND_USER_DEPT_) TableInfo() *TableInfo {
	return ANKEN_SEND_USER_DEPTTableInfo
}
