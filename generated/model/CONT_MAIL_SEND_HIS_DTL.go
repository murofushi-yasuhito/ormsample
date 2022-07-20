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


CREATE TABLE `CONT_MAIL_SEND_HIS_DTL` (
  `MAIL_SEND_ID` int(11) NOT NULL COMMENT 'メール送信ID',
  `SUB_NO` int(11) NOT NULL COMMENT '枝番',
  `ELMNT_LVL` int(11) DEFAULT NULL COMMENT '影響レベル',
  `AREA_KBN` char(4) DEFAULT NULL COMMENT 'エリア区分',
  `AREA_CD1` varchar(2) DEFAULT NULL COMMENT 'エリアコード1',
  `AREA_CD2` varchar(10) DEFAULT NULL COMMENT 'エリアコード2',
  `AREA_CD3` varchar(10) DEFAULT NULL COMMENT 'エリアコード3',
  `AREA_CD4` varchar(10) DEFAULT NULL COMMENT 'エリアコード4',
  `AREA_CD5` varchar(10) DEFAULT NULL COMMENT 'エリアコード5',
  `WETH_AREA_CD1` varchar(2) DEFAULT NULL COMMENT '気象予報区コード1',
  `WETH_AREA_CD2` varchar(7) DEFAULT NULL COMMENT '気象予報区コード2',
  `WETH_AREA_CD3` varchar(7) DEFAULT NULL COMMENT '気象予報区コード3',
  `MAIL_SEND_FLG` char(1) NOT NULL COMMENT 'メール送信フラグ',
  `CREA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`MAIL_SEND_ID`,`SUB_NO`),
  KEY `CONT_MAIl_SEND_HIS_DTL_IDX_1` (`AREA_CD1`,`AREA_CD2`,`AREA_CD3`,`WETH_AREA_CD1`,`WETH_AREA_CD2`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='コンテンツメール配信履歴詳細'

JSON Sample
-------------------------------------
{    "mail_send_id": 33,    "sub_no": 23,    "elmnt_lvl": 13,    "area_kbn": "GKSTgoMZFSZUtsiVsVxGHSfAv",    "area_cd_1": "vhluowVhuarCJHcHtXlRFYqDe",    "area_cd_2": "OuqVHtyMXjHjPfRBcsaeQynBy",    "area_cd_3": "JrKpQwAjeVbxaffkFRXyqbdQO",    "area_cd_4": "agdrYjgkDsQbLxvjbOxKKHArM",    "area_cd_5": "LYIAbtRAgtAHvmgKqjYiSEDQH",    "weth_area_cd_1": "ArcDQOLjjQiqlPphkmCxfffRC",    "weth_area_cd_2": "otbTqOviJScSBjWPNfYQRBhES",    "weth_area_cd_3": "OilVHyKlETyRDXGZnNrsoVGSD",    "mail_send_flg": "oKvqnYJKuZkIoudvwQWxSVvBb",    "crea_dte": "2285-05-14T01:51:45.585418673+09:00",    "crea_user_id": 9}



*/

// CONT_MAIL_SEND_HIS_DTL_ struct is a row record of the CONT_MAIL_SEND_HIS_DTL table in the anpidb database
type CONT_MAIL_SEND_HIS_DTL_ struct {
	//[ 0] MAIL_SEND_ID                                   int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	MAILSENDID int32 `gorm:"primary_key;column:MAIL_SEND_ID;type:int;"` // メール送信ID
	//[ 1] SUB_NO                                         int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SUBNO int32 `gorm:"primary_key;column:SUB_NO;type:int;"` // 枝番
	//[ 2] ELMNT_LVL                                      int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ELMNTLVL sql.NullInt64 `gorm:"column:ELMNT_LVL;type:int;"` // 影響レベル
	//[ 3] AREA_KBN                                       char(4)              null: true   primary: false  isArray: false  auto: false  col: char            len: 4       default: []
	AREAKBN sql.NullString `gorm:"column:AREA_KBN;type:char;size:4;"` // エリア区分
	//[ 4] AREA_CD1                                       varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	AREACD1 sql.NullString `gorm:"column:AREA_CD1;type:varchar;size:2;"` // エリアコード1
	//[ 5] AREA_CD2                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AREACD2 sql.NullString `gorm:"column:AREA_CD2;type:varchar;size:10;"` // エリアコード2
	//[ 6] AREA_CD3                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AREACD3 sql.NullString `gorm:"column:AREA_CD3;type:varchar;size:10;"` // エリアコード3
	//[ 7] AREA_CD4                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AREACD4 sql.NullString `gorm:"column:AREA_CD4;type:varchar;size:10;"` // エリアコード4
	//[ 8] AREA_CD5                                       varchar(10)          null: true   primary: false  isArray: false  auto: false  col: varchar         len: 10      default: []
	AREACD5 sql.NullString `gorm:"column:AREA_CD5;type:varchar;size:10;"` // エリアコード5
	//[ 9] WETH_AREA_CD1                                  varchar(2)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 2       default: []
	WETHAREACD1 sql.NullString `gorm:"column:WETH_AREA_CD1;type:varchar;size:2;"` // 気象予報区コード1
	//[10] WETH_AREA_CD2                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: []
	WETHAREACD2 sql.NullString `gorm:"column:WETH_AREA_CD2;type:varchar;size:7;"` // 気象予報区コード2
	//[11] WETH_AREA_CD3                                  varchar(7)           null: true   primary: false  isArray: false  auto: false  col: varchar         len: 7       default: []
	WETHAREACD3 sql.NullString `gorm:"column:WETH_AREA_CD3;type:varchar;size:7;"` // 気象予報区コード3
	//[12] MAIL_SEND_FLG                                  char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	MAILSENDFLG string `gorm:"column:MAIL_SEND_FLG;type:char;size:1;"` // メール送信フラグ
	//[13] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 作成日時
	//[14] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var CONT_MAIL_SEND_HIS_DTLTableInfo = &TableInfo{
	Name: "CONT_MAIL_SEND_HIS_DTL",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "MAIL_SEND_ID",
			Comment:            `メール送信ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "MAILSENDID",
			GoFieldType:        "int32",
			JSONFieldName:      "mail_send_id",
			ProtobufFieldName:  "mail_send_id",
			ProtobufType:       "int32",
			ProtobufPos:        1,
		},

		&ColumnInfo{
			Index:              1,
			Name:               "SUB_NO",
			Comment:            `枝番`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SUBNO",
			GoFieldType:        "int32",
			JSONFieldName:      "sub_no",
			ProtobufFieldName:  "sub_no",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "ELMNT_LVL",
			Comment:            `影響レベル`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ELMNTLVL",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "elmnt_lvl",
			ProtobufFieldName:  "elmnt_lvl",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "AREA_KBN",
			Comment:            `エリア区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(4)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       4,
			GoFieldName:        "AREAKBN",
			GoFieldType:        "sql.NullString",
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
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       2,
			GoFieldName:        "AREACD1",
			GoFieldType:        "sql.NullString",
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
			Name:               "MAIL_SEND_FLG",
			Comment:            `メール送信フラグ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "MAILSENDFLG",
			GoFieldType:        "string",
			JSONFieldName:      "mail_send_flg",
			ProtobufFieldName:  "mail_send_flg",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
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
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
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
			ProtobufPos:        15,
		},
	},
}

// TableName sets the insert table name for this struct type
func (c *CONT_MAIL_SEND_HIS_DTL_) TableName() string {
	return "CONT_MAIL_SEND_HIS_DTL"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_DTL_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (c *CONT_MAIL_SEND_HIS_DTL_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (c *CONT_MAIL_SEND_HIS_DTL_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (c *CONT_MAIL_SEND_HIS_DTL_) TableInfo() *TableInfo {
	return CONT_MAIL_SEND_HIS_DTLTableInfo
}
