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


CREATE TABLE `ANPI_USER_AREA_MST_CHECK` (
  `ANPI_USER_AREA_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '安否起動ユーザエリアID',
  `ANPI_DLCT_ID` int(11) NOT NULL COMMENT '安否起動DLCT_ID',
  `USER_ID` int(11) NOT NULL COMMENT 'ユーザID',
  `AREA_KBN` char(4) NOT NULL COMMENT 'エリア区分',
  `AREA_CD1` varchar(2) NOT NULL DEFAULT '0' COMMENT 'エリアコード1',
  `AREA_CD2` varchar(10) DEFAULT '0' COMMENT 'エリアコード2',
  `AREA_CD3` varchar(10) DEFAULT '0' COMMENT 'エリアコード3',
  `AREA_CD4` varchar(10) DEFAULT '0' COMMENT 'エリアコード4',
  `AREA_CD5` varchar(10) DEFAULT '0' COMMENT 'エリアコード5',
  `WETH_AREA_CD1` varchar(2) DEFAULT '0' COMMENT '気象予報区コード1',
  `WETH_AREA_CD2` varchar(7) DEFAULT '0' COMMENT '気象予報区コード2',
  `WETH_AREA_CD3` varchar(7) DEFAULT '0' COMMENT '気象予報区コード3',
  `DELE_FLG` char(1) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `START_GRP_ID` int(11) NOT NULL DEFAULT '0',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANPI_USER_AREA_ID`),
  KEY `ANPI_USER_AREA_MST_IDX1` (`ANPI_DLCT_ID`,`USER_ID`,`AREA_KBN`,`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`WETH_AREA_CD1`,`WETH_AREA_CD2`) USING BTREE,
  KEY `ANPI_USER_AREA_MST_IDX_2` (`ANPI_DLCT_ID`,`AREA_CD1`,`AREA_CD2`,`AREA_CD3`) USING BTREE,
  KEY `ANPI_USER_AREA_MST_IDX_3` (`AREA_KBN`,`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`WETH_AREA_CD3`),
  KEY `ANPI_USER_AREA_MST_IDX_4` (`USER_ID`,`WETH_AREA_CD2`),
  KEY `ANPI_USER_AREA_MST_IDX_5_TEST` (`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`WETH_AREA_CD1`,`WETH_AREA_CD2`,`WETH_AREA_CD3`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=31910 DEFAULT CHARSET=utf8 COMMENT='安否起動ユーザエリアマスタ'

JSON Sample
-------------------------------------
{    "anpi_user_area_id": 23,    "anpi_dlct_id": 85,    "user_id": 65,    "area_kbn": "UrSkvVBsiRPxKZZbWdBRnrdXZ",    "area_cd_1": "vgLMvSMKPGmUSBDIDbfDoaEaP",    "area_cd_2": "qtsxfmbbWtBKSMXZvAhqySYMe",    "area_cd_3": "HwBQDipSvvVFccCdOmqqqjmse",    "area_cd_4": "RvLoNshwjXDuhATrYxwiEZFrj",    "area_cd_5": "iBELLjNckidaaKyfVcRxrWmQR",    "weth_area_cd_1": "XqXmyjKCiEeonjqqbePPfpNsy",    "weth_area_cd_2": "rLcSKWmDXFCIxLOguBArOJsjk",    "weth_area_cd_3": "hOGJvjDgmAfnrjfuaVsBLHyRv",    "dele_flg": "LopkdpurjKyosxIiyKRkDuWyd",    "start_grp_id": 92,    "crea_dte": "2133-03-16T18:01:04.147532598+09:00",    "crea_user_id": 51}



*/

// ANPI_USER_AREA_MST_CHECK struct is a row record of the ANPI_USER_AREA_MST_CHECK table in the anpidb database
type ANPI_USER_AREA_MST_CHECK struct {
	//[ 0] ANPI_USER_AREA_ID                              int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANPIUSERAREAID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANPI_USER_AREA_ID;type:int;"` // 安否起動ユーザエリアID
	//[ 1] ANPI_DLCT_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANPIDLCTID int32 `gorm:"column:ANPI_DLCT_ID;type:int;"` // 安否起動DLCT_ID
	//[ 2] USER_ID                                        int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	USERID int32 `gorm:"column:USER_ID;type:int;"` // ユーザID
	//[ 3] AREA_KBN                                       char(4)              null: false  primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN string `gorm:"column:AREA_KBN;type:char;size:4;"` // エリア区分
	//[ 4] AREA_CD1                                       varchar(2)           null: false  primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	AREACD1 string `gorm:"column:AREA_CD1;type:varchar;size:2;default:0;"` // エリアコード1
	//[ 5] AREA_CD2                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD2 sql.NullString `gorm:"column:AREA_CD2;type:varchar;size:10;default:0;"` // エリアコード2
	//[ 6] AREA_CD3                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD3 sql.NullString `gorm:"column:AREA_CD3;type:varchar;size:10;default:0;"` // エリアコード3
	//[ 7] AREA_CD4                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD4 sql.NullString `gorm:"column:AREA_CD4;type:varchar;size:10;default:0;"` // エリアコード4
	//[ 8] AREA_CD5                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: [0]
	AREACD5 sql.NullString `gorm:"column:AREA_CD5;type:varchar;size:10;default:0;"` // エリアコード5
	//[ 9] WETH_AREA_CD1                                  varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: [0]
	WETHAREACD1 sql.NullString `gorm:"column:WETH_AREA_CD1;type:varchar;size:2;default:0;"` // 気象予報区コード1
	//[10] WETH_AREA_CD2                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD2 sql.NullString `gorm:"column:WETH_AREA_CD2;type:varchar;size:7;default:0;"` // 気象予報区コード2
	//[11] WETH_AREA_CD3                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: [0]
	WETHAREACD3 sql.NullString `gorm:"column:WETH_AREA_CD3;type:varchar;size:7;default:0;"` // 気象予報区コード3
	//[12] DELE_FLG                                       char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DELEFLG string `gorm:"column:DELE_FLG;type:char;size:1;default:0;"` // 削除フラグ
	//[13] START_GRP_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	STARTGRPID int32 `gorm:"column:START_GRP_ID;type:int;default:0;"`
	//[14] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[15] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANPI_USER_AREA_MST_CHECKTableInfo = &TableInfo{
	Name: "ANPI_USER_AREA_MST_CHECK",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "ANPI_USER_AREA_ID",
			Comment:            `安否起動ユーザエリアID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    true,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANPIUSERAREAID",
			GoFieldType:        "int32",
			JSONFieldName:      "anpi_user_area_id",
			ProtobufFieldName:  "anpi_user_area_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "ANPI_DLCT_ID",
			Comment:            `安否起動DLCT_ID`,
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
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
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
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
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
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
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
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
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
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
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
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index:              8,
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
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
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
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
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
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
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
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "DELE_FLG",
			Comment:            `削除フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DELEFLG",
			GoFieldType:        "string",
			JSONFieldName:      "dele_flg",
			ProtobufFieldName:  "dele_flg",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "START_GRP_ID",
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
			GoFieldName:        "STARTGRPID",
			GoFieldType:        "int32",
			JSONFieldName:      "start_grp_id",
			ProtobufFieldName:  "start_grp_id",
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
func (a *ANPI_USER_AREA_MST_CHECK) TableName() string {
	return "ANPI_USER_AREA_MST_CHECK"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANPI_USER_AREA_MST_CHECK) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANPI_USER_AREA_MST_CHECK) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANPI_USER_AREA_MST_CHECK) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANPI_USER_AREA_MST_CHECK) TableInfo() *TableInfo {
	return ANPI_USER_AREA_MST_CHECKTableInfo
}
