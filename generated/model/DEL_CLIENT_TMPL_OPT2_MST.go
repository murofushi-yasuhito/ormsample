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


CREATE TABLE `DEL_CLIENT_TMPL_OPT2_MST` (
  `CLIENT_TMPL_OPT2_MNG_ID` int(11) NOT NULL COMMENT 'クライアントテンプレートID',
  `SEQ` int(11) NOT NULL COMMENT '連番',
  `STS_CD` char(1) NOT NULL COMMENT '状態コード',
  `STS_MSG` varchar(100) DEFAULT NULL COMMENT '状態メッセージ',
  `STS_MSG_E` varchar(100) DEFAULT NULL COMMENT '状態メッセージ(英)',
  `RES_MSG` varchar(100) DEFAULT NULL COMMENT '選択後メッセージ',
  `RES_MSG_E` varchar(100) DEFAULT NULL COMMENT '選択後メッセージ(英)',
  `USE_FLG` char(1) DEFAULT NULL,
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`CLIENT_TMPL_OPT2_MNG_ID`,`SEQ`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='クライアントテンプレート選択肢マスタ'

JSON Sample
-------------------------------------
{    "client_tmpl_opt_2_mng_id": 99,    "seq": 65,    "sts_cd": "KKfNjdIGkJnBLKijIqhTLkKyX",    "sts_msg": "cUqqKmXieqcUTJKsqZFTSGlhU",    "sts_msg_e": "uJnqwuabkiiKxVTRUYtbZVHQf",    "res_msg": "QNWemauRRUdBgygFTEbOtjiTE",    "res_msg_e": "IFXnrXkKEtthAfrJGQWauXJUj",    "use_flg": "hujpwEOebolArHFgVdObWMqrc",    "upda_dte": "2165-12-28T05:47:28.198288409+09:00",    "upda_user_id": 54,    "crea_dte": "2045-05-07T18:36:43.626809008+09:00",    "crea_user_id": 9}



*/

// DEL_CLIENT_TMPL_OPT2_MST struct is a row record of the DEL_CLIENT_TMPL_OPT2_MST table in the anpidb database
type DEL_CLIENT_TMPL_OPT2_MST struct {
	//[ 0] CLIENT_TMPL_OPT2_MNG_ID                        int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTTMPLOPT2MNGID int32 `gorm:"primary_key;column:CLIENT_TMPL_OPT2_MNG_ID;type:int;"` // クライアントテンプレートID
	//[ 1] SEQ                                            int                  null: false  primary: true   isArray: false  auto: false  col: int             len: -1      default: []
	SEQ int32 `gorm:"primary_key;column:SEQ;type:int;"` // 連番
	//[ 2] STS_CD                                         char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	STSCD string `gorm:"column:STS_CD;type:char;size:1;"` // 状態コード
	//[ 3] STS_MSG                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	STSMSG sql.NullString `gorm:"column:STS_MSG;type:varchar;size:100;"` // 状態メッセージ
	//[ 4] STS_MSG_E                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	STSMSGE sql.NullString `gorm:"column:STS_MSG_E;type:varchar;size:100;"` // 状態メッセージ(英)
	//[ 5] RES_MSG                                        varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RESMSG sql.NullString `gorm:"column:RES_MSG;type:varchar;size:100;"` // 選択後メッセージ
	//[ 6] RES_MSG_E                                      varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	RESMSGE sql.NullString `gorm:"column:RES_MSG_E;type:varchar;size:100;"` // 選択後メッセージ(英)
	//[ 7] USE_FLG                                        char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	USEFLG sql.NullString `gorm:"column:USE_FLG;type:char;size:1;"`
	//[ 8] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[ 9] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[10] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[11] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var DEL_CLIENT_TMPL_OPT2_MSTTableInfo = &TableInfo{
	Name: "DEL_CLIENT_TMPL_OPT2_MST",
	Columns: []*ColumnInfo{

		&ColumnInfo{
			Index:              0,
			Name:               "CLIENT_TMPL_OPT2_MNG_ID",
			Comment:            `クライアントテンプレートID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       true,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "CLIENTTMPLOPT2MNGID",
			GoFieldType:        "int32",
			JSONFieldName:      "client_tmpl_opt_2_mng_id",
			ProtobufFieldName:  "client_tmpl_opt_2_mng_id",
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
			Name:               "STS_CD",
			Comment:            `状態コード`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STSCD",
			GoFieldType:        "string",
			JSONFieldName:      "sts_cd",
			ProtobufFieldName:  "sts_cd",
			ProtobufType:       "string",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index:              3,
			Name:               "STS_MSG",
			Comment:            `状態メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "STSMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_msg",
			ProtobufFieldName:  "sts_msg",
			ProtobufType:       "string",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
			Name:               "STS_MSG_E",
			Comment:            `状態メッセージ(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "STSMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "sts_msg_e",
			ProtobufFieldName:  "sts_msg_e",
			ProtobufType:       "string",
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "RES_MSG",
			Comment:            `選択後メッセージ`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RESMSG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg",
			ProtobufFieldName:  "res_msg",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "RES_MSG_E",
			Comment:            `選択後メッセージ(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "RESMSGE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "res_msg_e",
			ProtobufFieldName:  "res_msg_e",
			ProtobufType:       "string",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "USE_FLG",
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
			GoFieldName:        "USEFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "use_flg",
			ProtobufFieldName:  "use_flg",
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
func (d *DEL_CLIENT_TMPL_OPT2_MST) TableName() string {
	return "DEL_CLIENT_TMPL_OPT2_MST"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (d *DEL_CLIENT_TMPL_OPT2_MST) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (d *DEL_CLIENT_TMPL_OPT2_MST) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (d *DEL_CLIENT_TMPL_OPT2_MST) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (d *DEL_CLIENT_TMPL_OPT2_MST) TableInfo() *TableInfo {
	return DEL_CLIENT_TMPL_OPT2_MSTTableInfo
}
