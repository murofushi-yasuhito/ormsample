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


CREATE TABLE `USER_DLCT_D_MST` (
  `USER_DLCT_ID` int(11) NOT NULL DEFAULT '0' COMMENT 'ユーザDLCT_ID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `AREA_KBN` char(4) NOT NULL COMMENT 'エリア区分',
  `AREA_CD1` varchar(2) NOT NULL DEFAULT '0' COMMENT 'エリアコード1',
  `AREA_CD2` varchar(10) DEFAULT '0' COMMENT 'エリアコード2',
  `AREA_CD3` varchar(10) DEFAULT '0' COMMENT 'エリアコード3',
  `AREA_CD4` varchar(10) DEFAULT '0' COMMENT 'エリアコード4',
  `AREA_CD5` varchar(10) DEFAULT '0' COMMENT 'エリアコード5',
  `WETH_AREA_CD1` varchar(2) DEFAULT '0' COMMENT '気象予報区コード1',
  `WETH_AREA_CD2` varchar(7) DEFAULT '0' COMMENT '気象予報区コード2',
  `WETH_AREA_CD3` varchar(7) DEFAULT '0' COMMENT '気象予報区コード3',
  `REL_ADDR_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '0：ユーザ作成 1：拠点 2：部署 3：居住地',
  `REL_ADDR_ID` int(11) NOT NULL DEFAULT '0',
  `REL_SEQ` int(11) NOT NULL DEFAULT '0' COMMENT '紐付けシーケンス番号',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`USER_DLCT_ID`,`SEQ`),
  KEY `USER_DLCT_D_MST_IDX_1` (`AREA_KBN`,`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`AREA_CD4`,`AREA_CD5`) USING BTREE,
  KEY `USER_DLCT_D_MST_IDX_2` (`REL_ADDR_KBN`,`REL_ADDR_ID`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='ユーザDLCT_Dマスタ'

JSON Sample
-------------------------------------
{    "user_dlct_id": 20,    "seq": 36,    "area_kbn": "NacKOoMhSkksVbsJKZPUyXGlj",    "area_cd_1": "QclHHmoBoFbNPiVPQveOlDfgF",    "area_cd_2": "vyVlghDLDEjgYtTdbQwAorHYR",    "area_cd_3": "mRMoIDEcJMJeuDUQlnoaCXvLh",    "area_cd_4": "EaxDFaAvslRSfdXUxSSQbCOwp",    "area_cd_5": "nNIrYkFHhvkXhLMAgwOoPWyfW",    "weth_area_cd_1": "EWWpXVgLOoVoshPXqBJvjEllk",    "weth_area_cd_2": "mEipcjQrKeSpGyIQXUaspfuHK",    "weth_area_cd_3": "mvacSmkiAtmwHBMqtcXdupqVd",    "rel_addr_kbn": "DELDTfOXsRMqsPIQprXedqjGv",    "rel_addr_id": 35,    "rel_seq": 0,    "crea_dte": "2254-06-19T21:40:09.900966304+09:00",    "crea_user_id": 95}



*/

// USER_DLCT_D_MST_ struct is a row record of the USER_DLCT_D_MST table in the anpidb database
type USER_DLCT_D_MST_ struct {
	//[ 0] USER_DLCT_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: [0]
	USERDLCTID int32 `gorm:"primary_key;column:USER_DLCT_ID;type:int;default:0;"` // ユーザDLCT_ID
	//[ 1] SEQ                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"primary_key;column:SEQ;type:int;"` // 連番
	//[ 2] AREA_KBN                                       char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:char;size:4;"` // エリア区分
	//[ 3] AREA_CD1                                       varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	AREACD1 string `gorm:"column:AREA_CD1;type:varchar;size:2;default:0;"` // エリアコード1
	//[ 4] AREA_CD2                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD2 sql.NullString `gorm:"column:AREA_CD2;type:varchar;size:10;default:0;"` // エリアコード2
	//[ 5] AREA_CD3                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD3 sql.NullString `gorm:"column:AREA_CD3;type:varchar;size:10;default:0;"` // エリアコード3
	//[ 6] AREA_CD4                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD4 sql.NullString `gorm:"column:AREA_CD4;type:varchar;size:10;default:0;"` // エリアコード4
	//[ 7] AREA_CD5                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD5 sql.NullString `gorm:"column:AREA_CD5;type:varchar;size:10;default:0;"` // エリアコード5
	//[ 8] WETH_AREA_CD1                                  varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	WETHAREACD1 sql.NullString `gorm:"column:WETH_AREA_CD1;type:varchar;size:2;default:0;"` // 気象予報区コード1
	//[ 9] WETH_AREA_CD2                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD2 sql.NullString `gorm:"column:WETH_AREA_CD2;type:varchar;size:7;default:0;"` // 気象予報区コード2
	//[10] WETH_AREA_CD3                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD3 sql.NullString `gorm:"column:WETH_AREA_CD3;type:varchar;size:7;default:0;"` // 気象予報区コード3
	//[11] REL_ADDR_KBN                                   char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	RELADDRKBN string `gorm:"column:REL_ADDR_KBN;type:char;size:1;default:0;"` // 0：ユーザ作成 1：拠点 2：部署 3：居住地
	//[12] REL_ADDR_ID                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RELADDRID int32 `gorm:"column:REL_ADDR_ID;type:int;default:0;"`
	//[13] REL_SEQ                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	RELSEQ int32 `gorm:"column:REL_SEQ;type:int;default:0;"` // 紐付けシーケンス番号
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var USER_DLCT_D_MSTTableInfo = &TableInfo{
	Name: "USER_DLCT_D_MST",
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
			Index:              2,
			Name:               "AREA_KBN",
			Comment:            `エリア区分`,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "AREA_CD1",
			Comment:            `エリアコード1`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "AREACD1",
			GoFieldType:        "string",
			JSONFieldName:      "area_cd_1",
			ProtobufFieldName:  "area_cd_1",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "AREA_CD2",
			Comment:            `エリアコード2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_2",
			ProtobufFieldName:  "area_cd_2",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "AREA_CD3",
			Comment:            `エリアコード3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_3",
			ProtobufFieldName:  "area_cd_3",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "AREA_CD4",
			Comment:            `エリアコード4`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD4",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_4",
			ProtobufFieldName:  "area_cd_4",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "AREA_CD5",
			Comment:            `エリアコード5`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(10)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       10,
			GoFieldName:        "AREACD5",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "area_cd_5",
			ProtobufFieldName:  "area_cd_5",
			ProtobufType:       "string",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
			Name:               "WETH_AREA_CD1",
			Comment:            `気象予報区コード1`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "WETHAREACD1",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_1",
			ProtobufFieldName:  "weth_area_cd_1",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "WETH_AREA_CD2",
			Comment:            `気象予報区コード2`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       7,
			GoFieldName:        "WETHAREACD2",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_2",
			ProtobufFieldName:  "weth_area_cd_2",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "WETH_AREA_CD3",
			Comment:            `気象予報区コード3`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(7)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       7,
			GoFieldName:        "WETHAREACD3",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "weth_area_cd_3",
			ProtobufFieldName:  "weth_area_cd_3",
			ProtobufType:       "string",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "REL_ADDR_KBN",
			Comment:            `0：ユーザ作成 1：拠点 2：部署 3：居住地`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "RELADDRKBN",
			GoFieldType:        "string",
			JSONFieldName:      "rel_addr_kbn",
			ProtobufFieldName:  "rel_addr_kbn",
			ProtobufType:       "string",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "REL_ADDR_ID",
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
			GoFieldName:        "RELADDRID",
			GoFieldType:        "int32",
			JSONFieldName:      "rel_addr_id",
			ProtobufFieldName:  "rel_addr_id",
			ProtobufType:       "int32",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "REL_SEQ",
			Comment:            `紐付けシーケンス番号`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "RELSEQ",
			GoFieldType:        "int32",
			JSONFieldName:      "rel_seq",
			ProtobufFieldName:  "rel_seq",
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
func (u *USER_DLCT_D_MST_) TableName() string {
	return "USER_DLCT_D_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (u *USER_DLCT_D_MST_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (u *USER_DLCT_D_MST_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (u *USER_DLCT_D_MST_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (u *USER_DLCT_D_MST_) TableInfo() *TableInfo {
	return USER_DLCT_D_MSTTableInfo
}
