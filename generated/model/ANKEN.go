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


CREATE TABLE `ANKEN` (
  `ANKEN_ID` int(11) NOT NULL AUTO_INCREMENT COMMENT '案件ID',
  `TOP_ANKEN_ID` int(11) NOT NULL DEFAULT '0' COMMENT '先頭案件ID	先頭は案件ID',
  `UP_ANKEN_ID` int(11) NOT NULL DEFAULT '0' COMMENT '上案件ID    先頭は0',
  `ANKEN_POS` int(11) NOT NULL DEFAULT '1' COMMENT '案件ポジション	先頭は1',
  `CLIENT_ID` int(11) NOT NULL COMMENT 'クライアントID',
  `EN_KBN` char(1) DEFAULT '0' COMMENT '英語化対応区分',
  `ANKEN_NO` int(11) NOT NULL COMMENT '案件NO',
  `ANKEN_CSE_DTE` datetime NOT NULL COMMENT '案件発生日時',
  `ANKEN_KND` char(1) NOT NULL COMMENT '案件種別	 1:安否確認 2:非常呼集 3:お知らせ',
  `SEND_TO_KBN` char(1) DEFAULT NULL COMMENT '1:自動起動 2:任意指定 3:グループ指定 4:引き続き案件 5:地域指定',
  `SID` int(11) DEFAULT '0' COMMENT '情報番号',
  `TOP_SID` int(11) DEFAULT '0' COMMENT 'TOP SID',
  `ANKEN_TTL` varchar(100) NOT NULL COMMENT '案件件名',
  `ANKEN_TTL_E` varchar(100) DEFAULT NULL COMMENT '案件件名(英)',
  `ANKEN_BODY` varchar(6000) DEFAULT NULL COMMENT '案件本文',
  `ANKEN_BODY_E` varchar(6000) DEFAULT NULL COMMENT '案件本文(英)',
  `INFO_LONG` text COMMENT 'コンテンツロング文',
  `SEND_USER_ID` int(11) NOT NULL COMMENT '送信者ID',
  `ANKEN_PROC_FLG` char(1) DEFAULT '0' COMMENT '案件処理フラグ 0:未処理 1:ユーザ登録済 2:メール配信済',
  `ANKEN_END_DTE` datetime DEFAULT NULL COMMENT '案件終了日時',
  `VALI_DTE` datetime DEFAULT NULL COMMENT '有効期日',
  `ANNO_DISP_KBN` char(1) DEFAULT '0' COMMENT 'お知らせ表示区分\r\n    お知らせのみ 1:表示 0:非表示',
  `UP_ANKEN_NO` int(11) DEFAULT '0' COMMENT '上案件NO',
  `CHIL_ANKEN_IDS` text COMMENT '子案件ID(カンマ区切)',
  `DEPT_START_KBN` char(1) DEFAULT '0' COMMENT '部署起動区分    0:全社権限 1:部署権限',
  `START_DEPT_DELI_KBN` char(2) DEFAULT NULL,
  `ANPI_DLCT_ID` int(11) unsigned zerofill NOT NULL DEFAULT '00000000000' COMMENT '安否DLCTID',
  `JOKYO_CLIENT_TMPL_ID` int(11) unsigned zerofill NOT NULL DEFAULT '00000000000' COMMENT '状況一覧に表示する設問',
  `START_RECV_KBN` char(1) DEFAULT '0' COMMENT '起動予約区分 1:利用する 0:利用しない 2:予約停止',
  `START_RECV_DTE` datetime DEFAULT NULL COMMENT '起動予約日時',
  `START_NOTIF_SEND_FLG` char(1) NOT NULL DEFAULT '1' COMMENT '起動通知送信フラグ 0：起動通知を送信しない 1：起動通知を送信する',
  `AUTO_RESEND_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '自動再送 0:再送なし 1:再送する 2:ユーザによる中断',
  `AUTO_RESEND_TIMES` int(11) NOT NULL DEFAULT '0' COMMENT '再送実施回数',
  `AUTO_RESEND_DTE` datetime DEFAULT NULL COMMENT '最後に自動再送した時刻',
  `INCL_RESIDENCE_KBN` char(1) NOT NULL DEFAULT '0' COMMENT '居住地を含める（0:含めない 1:含める）',
  `UPDA_DTE` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時',
  `UPDA_USER_ID` int(11) NOT NULL COMMENT '更新者ID',
  `CREA_DTE` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '作成日時',
  `CREA_USER_ID` int(11) NOT NULL COMMENT '作成者ID',
  PRIMARY KEY (`ANKEN_ID`),
  KEY `ANKEN_IDX1` (`CLIENT_ID`),
  KEY `ANKEN_IDX2` (`TOP_SID`,`ANPI_DLCT_ID`) USING BTREE,
  KEY `ANKEN_IDX3` (`SID`)
) ENGINE=InnoDB AUTO_INCREMENT=144660 DEFAULT CHARSET=utf8 COMMENT='案件'

JSON Sample
-------------------------------------
{    "anken_id": 70,    "top_anken_id": 25,    "up_anken_id": 7,    "anken_pos": 58,    "client_id": 90,    "en_kbn": "XmRYxnGPHxJIdInBwjpvVxjcJ",    "anken_no": 82,    "anken_cse_dte": "2187-04-18T10:40:46.609327062+09:00",    "anken_knd": "eaIaDsSQcKuRxOCNKtORkGeGy",    "send_to_kbn": "btSeGYZYNDxFVhCAtNRAapSjA",    "sid": 28,    "top_sid": 20,    "anken_ttl": "WxVLFfdKcKCyZAOkfnJSGMsef",    "anken_ttl_e": "hyJGOUuPUkYLdpVnBryRKPtox",    "anken_body": "rMtgBPephusWwZMtEWnNwgPGk",    "anken_body_e": "cuxtEpiUXJCPnKDmJOPZvEZQJ",    "info_long": "rFhWVShsirKGKhhEUISscHLfw",    "send_user_id": 15,    "anken_proc_flg": "xRHBHAZaCPBCtRVWMOIUcLVwW",    "anken_end_dte": "2028-03-16T08:05:09.494200541+09:00",    "vali_dte": "2193-02-16T23:31:54.147062544+09:00",    "anno_disp_kbn": "LvgWsWtfqoNZTqeiprSliNqJM",    "up_anken_no": 99,    "chil_anken_ids": "tSbAVbGgnWgoDEFfPBMcPEKId",    "dept_start_kbn": "PMdUQcEBxopisffJUYotyORhr",    "start_dept_deli_kbn": "hIFUPoayAiPahWjUbcbAnacOi",    "anpi_dlct_id": 22,    "jokyo_client_tmpl_id": 46,    "start_recv_kbn": "HNOUThVQobONwlhiZOBECbkti",    "start_recv_dte": "2292-05-09T00:50:12.847847718+09:00",    "start_notif_send_flg": "tSomXUJmNKZxnJoIBFkebBCyk",    "auto_resend_kbn": "dZFMGLvFZlxWbErGXlauePKDE",    "auto_resend_times": 60,    "auto_resend_dte": "2257-10-09T16:41:52.046641915+09:00",    "incl_residence_kbn": "XufxtNIEtbNKIpqwWhOePPGIy",    "upda_dte": "2109-11-14T03:47:53.293719922+09:00",    "upda_user_id": 46,    "crea_dte": "2231-07-20T12:50:46.34642102+09:00",    "crea_user_id": 6}


Comments
-------------------------------------
[26] column is set for unsigned
[27] column is set for unsigned



*/

// ANKEN_ struct is a row record of the ANKEN table in the anpidb database
type ANKEN_ struct {
	//[ 0] ANKEN_ID                                       int                  null: false  primary: true   isArray: false  auto: true   col: int             len: -1      default: []
	ANKENID int32 `gorm:"primary_key;AUTO_INCREMENT;column:ANKEN_ID;type:int;"` // 案件ID
	//[ 1] TOP_ANKEN_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TOPANKENID int32 `gorm:"column:TOP_ANKEN_ID;type:int;default:0;"` // 先頭案件ID	先頭は案件ID
	//[ 2] UP_ANKEN_ID                                    int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPANKENID int32 `gorm:"column:UP_ANKEN_ID;type:int;default:0;"` // 上案件ID    先頭は0
	//[ 3] ANKEN_POS                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [1]
	ANKENPOS int32 `gorm:"column:ANKEN_POS;type:int;default:1;"` // 案件ポジション	先頭は1
	//[ 4] CLIENT_ID                                      int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CLIENTID int32 `gorm:"column:CLIENT_ID;type:int;"` // クライアントID
	//[ 5] EN_KBN                                         char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	ENKBN sql.NullString `gorm:"column:EN_KBN;type:char;size:1;default:0;"` // 英語化対応区分
	//[ 6] ANKEN_NO                                       int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	ANKENNO int32 `gorm:"column:ANKEN_NO;type:int;"` // 案件NO
	//[ 7] ANKEN_CSE_DTE                                  datetime             null: false  primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ANKENCSEDTE time.Time `gorm:"column:ANKEN_CSE_DTE;type:datetime;"` // 案件発生日時
	//[ 8] ANKEN_KND                                      char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	ANKENKND string `gorm:"column:ANKEN_KND;type:char;size:1;"` // 案件種別	 1:安否確認 2:非常呼集 3:お知らせ
	//[ 9] SEND_TO_KBN                                    char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: []
	SENDTOKBN sql.NullString `gorm:"column:SEND_TO_KBN;type:char;size:1;"` // 1:自動起動 2:任意指定 3:グループ指定 4:引き続き案件 5:地域指定
	//[10] SID                                            int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	SID sql.NullInt64 `gorm:"column:SID;type:int;default:0;"` // 情報番号
	//[11] TOP_SID                                        int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	TOPSID sql.NullInt64 `gorm:"column:TOP_SID;type:int;default:0;"` // TOP SID
	//[12] ANKEN_TTL                                      varchar(100)         null: false  primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ANKENTTL string `gorm:"column:ANKEN_TTL;type:varchar;size:100;"` // 案件件名
	//[13] ANKEN_TTL_E                                    varchar(100)         null: true   primary: false  isArray: false  auto: false  col: varchar         len: 100     default: []
	ANKENTTLE sql.NullString `gorm:"column:ANKEN_TTL_E;type:varchar;size:100;"` // 案件件名(英)
	//[14] ANKEN_BODY                                     varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	ANKENBODY sql.NullString `gorm:"column:ANKEN_BODY;type:varchar;size:6000;"` // 案件本文
	//[15] ANKEN_BODY_E                                   varchar(6000)        null: true   primary: false  isArray: false  auto: false  col: varchar         len: 6000    default: []
	ANKENBODYE sql.NullString `gorm:"column:ANKEN_BODY_E;type:varchar;size:6000;"` // 案件本文(英)
	//[16] INFO_LONG                                      text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	INFOLONG sql.NullString `gorm:"column:INFO_LONG;type:text;size:65535;"` // コンテンツロング文
	//[17] SEND_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	SENDUSERID int32 `gorm:"column:SEND_USER_ID;type:int;"` // 送信者ID
	//[18] ANKEN_PROC_FLG                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	ANKENPROCFLG sql.NullString `gorm:"column:ANKEN_PROC_FLG;type:char;size:1;default:0;"` // 案件処理フラグ 0:未処理 1:ユーザ登録済 2:メール配信済
	//[19] ANKEN_END_DTE                                  datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	ANKENENDDTE time.Time `gorm:"column:ANKEN_END_DTE;type:datetime;"` // 案件終了日時
	//[20] VALI_DTE                                       datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	VALIDTE time.Time `gorm:"column:VALI_DTE;type:datetime;"` // 有効期日
	//[21] ANNO_DISP_KBN                                  char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	ANNODISPKBN sql.NullString `gorm:"column:ANNO_DISP_KBN;type:char;size:1;default:0;"` // お知らせ表示区分\r\n    お知らせのみ 1:表示 0:非表示
	//[22] UP_ANKEN_NO                                    int                  null: true   primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	UPANKENNO sql.NullInt64 `gorm:"column:UP_ANKEN_NO;type:int;default:0;"` // 上案件NO
	//[23] CHIL_ANKEN_IDS                                 text(65535)          null: true   primary: false  isArray: false  auto: false  col: text            len: 65535   default: []
	CHILANKENIDS sql.NullString `gorm:"column:CHIL_ANKEN_IDS;type:text;size:65535;"` // 子案件ID(カンマ区切)
	//[24] DEPT_START_KBN                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	DEPTSTARTKBN sql.NullString `gorm:"column:DEPT_START_KBN;type:char;size:1;default:0;"` // 部署起動区分    0:全社権限 1:部署権限
	//[25] START_DEPT_DELI_KBN                            char(2)              null: true   primary: false  isArray: false  auto: false  col: char            len: 2       default: []
	STARTDEPTDELIKBN sql.NullString `gorm:"column:START_DEPT_DELI_KBN;type:char;size:2;"`
	//[26] ANPI_DLCT_ID                                   uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [00000000000]
	ANPIDLCTID uint32 `gorm:"column:ANPI_DLCT_ID;type:uint;default:00000000000;"` // 安否DLCTID
	//[27] JOKYO_CLIENT_TMPL_ID                           uint                 null: false  primary: false  isArray: false  auto: false  col: uint            len: -1      default: [00000000000]
	JOKYOCLIENTTMPLID uint32 `gorm:"column:JOKYO_CLIENT_TMPL_ID;type:uint;default:00000000000;"` // 状況一覧に表示する設問
	//[28] START_RECV_KBN                                 char(1)              null: true   primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	STARTRECVKBN sql.NullString `gorm:"column:START_RECV_KBN;type:char;size:1;default:0;"` // 起動予約区分 1:利用する 0:利用しない 2:予約停止
	//[29] START_RECV_DTE                                 datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	STARTRECVDTE time.Time `gorm:"column:START_RECV_DTE;type:datetime;"` // 起動予約日時
	//[30] START_NOTIF_SEND_FLG                           char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [1]
	STARTNOTIFSENDFLG string `gorm:"column:START_NOTIF_SEND_FLG;type:char;size:1;default:1;"` // 起動通知送信フラグ 0：起動通知を送信しない 1：起動通知を送信する
	//[31] AUTO_RESEND_KBN                                char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	AUTORESENDKBN string `gorm:"column:AUTO_RESEND_KBN;type:char;size:1;default:0;"` // 自動再送 0:再送なし 1:再送する 2:ユーザによる中断
	//[32] AUTO_RESEND_TIMES                              int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: [0]
	AUTORESENDTIMES int32 `gorm:"column:AUTO_RESEND_TIMES;type:int;default:0;"` // 再送実施回数
	//[33] AUTO_RESEND_DTE                                datetime             null: true   primary: false  isArray: false  auto: false  col: datetime        len: -1      default: []
	AUTORESENDDTE time.Time `gorm:"column:AUTO_RESEND_DTE;type:datetime;"` // 最後に自動再送した時刻
	//[34] INCL_RESIDENCE_KBN                             char(1)              null: false  primary: false  isArray: false  auto: false  col: char            len: 1       default: [0]
	INCLRESIDENCEKBN string `gorm:"column:INCL_RESIDENCE_KBN;type:char;size:1;default:0;"` // 居住地を含める（0:含めない 1:含める）
	//[35] UPDA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [CURRENT_TIMESTAMP]
	UPDADTE time.Time `gorm:"column:UPDA_DTE;type:timestamp;default:CURRENT_TIMESTAMP;"` // 更新日時
	//[36] UPDA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	UPDAUSERID int32 `gorm:"column:UPDA_USER_ID;type:int;"` // 更新者ID
	//[37] CREA_DTE                                       timestamp            null: false  primary: false  isArray: false  auto: false  col: timestamp       len: -1      default: [0000-00-00 00:00:00]
	CREADTE time.Time `gorm:"column:CREA_DTE;type:timestamp;default:0000-00-00 00:00:00;"` // 作成日時
	//[38] CREA_USER_ID                                   int                  null: false  primary: false  isArray: false  auto: false  col: int             len: -1      default: []
	CREAUSERID int32 `gorm:"column:CREA_USER_ID;type:int;"` // 作成者ID

}

var ANKENTableInfo = &TableInfo{
	Name: "ANKEN",
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
			IsAutoIncrement:    true,
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
			Index: 1,
			Name:  "TOP_ANKEN_ID",
			Comment: `先頭案件ID	先頭は案件ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "top_anken_id",
			ProtobufFieldName:  "top_anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        2,
		},

		&ColumnInfo{
			Index:              2,
			Name:               "UP_ANKEN_ID",
			Comment:            `上案件ID    先頭は0`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPANKENID",
			GoFieldType:        "int32",
			JSONFieldName:      "up_anken_id",
			ProtobufFieldName:  "up_anken_id",
			ProtobufType:       "int32",
			ProtobufPos:        3,
		},

		&ColumnInfo{
			Index: 3,
			Name:  "ANKEN_POS",
			Comment: `案件ポジション	先頭は1`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENPOS",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_pos",
			ProtobufFieldName:  "anken_pos",
			ProtobufType:       "int32",
			ProtobufPos:        4,
		},

		&ColumnInfo{
			Index:              4,
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
			ProtobufPos:        5,
		},

		&ColumnInfo{
			Index:              5,
			Name:               "EN_KBN",
			Comment:            `英語化対応区分`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ENKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "en_kbn",
			ProtobufFieldName:  "en_kbn",
			ProtobufType:       "string",
			ProtobufPos:        6,
		},

		&ColumnInfo{
			Index:              6,
			Name:               "ANKEN_NO",
			Comment:            `案件NO`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "ANKENNO",
			GoFieldType:        "int32",
			JSONFieldName:      "anken_no",
			ProtobufFieldName:  "anken_no",
			ProtobufType:       "int32",
			ProtobufPos:        7,
		},

		&ColumnInfo{
			Index:              7,
			Name:               "ANKEN_CSE_DTE",
			Comment:            `案件発生日時`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ANKENCSEDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "anken_cse_dte",
			ProtobufFieldName:  "anken_cse_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        8,
		},

		&ColumnInfo{
			Index: 8,
			Name:  "ANKEN_KND",
			Comment: `案件種別	 1:安否確認 2:非常呼集 3:お知らせ`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ANKENKND",
			GoFieldType:        "string",
			JSONFieldName:      "anken_knd",
			ProtobufFieldName:  "anken_knd",
			ProtobufType:       "string",
			ProtobufPos:        9,
		},

		&ColumnInfo{
			Index:              9,
			Name:               "SEND_TO_KBN",
			Comment:            `1:自動起動 2:任意指定 3:グループ指定 4:引き続き案件 5:地域指定`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "SENDTOKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "send_to_kbn",
			ProtobufFieldName:  "send_to_kbn",
			ProtobufType:       "string",
			ProtobufPos:        10,
		},

		&ColumnInfo{
			Index:              10,
			Name:               "SID",
			Comment:            `情報番号`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "sid",
			ProtobufFieldName:  "sid",
			ProtobufType:       "int32",
			ProtobufPos:        11,
		},

		&ColumnInfo{
			Index:              11,
			Name:               "TOP_SID",
			Comment:            `TOP SID`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "TOPSID",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "top_sid",
			ProtobufFieldName:  "top_sid",
			ProtobufType:       "int32",
			ProtobufPos:        12,
		},

		&ColumnInfo{
			Index:              12,
			Name:               "ANKEN_TTL",
			Comment:            `案件件名`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ANKENTTL",
			GoFieldType:        "string",
			JSONFieldName:      "anken_ttl",
			ProtobufFieldName:  "anken_ttl",
			ProtobufType:       "string",
			ProtobufPos:        13,
		},

		&ColumnInfo{
			Index:              13,
			Name:               "ANKEN_TTL_E",
			Comment:            `案件件名(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(100)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       100,
			GoFieldName:        "ANKENTTLE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anken_ttl_e",
			ProtobufFieldName:  "anken_ttl_e",
			ProtobufType:       "string",
			ProtobufPos:        14,
		},

		&ColumnInfo{
			Index:              14,
			Name:               "ANKEN_BODY",
			Comment:            `案件本文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(6000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       6000,
			GoFieldName:        "ANKENBODY",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anken_body",
			ProtobufFieldName:  "anken_body",
			ProtobufType:       "string",
			ProtobufPos:        15,
		},

		&ColumnInfo{
			Index:              15,
			Name:               "ANKEN_BODY_E",
			Comment:            `案件本文(英)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "varchar",
			DatabaseTypePretty: "varchar(6000)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "varchar",
			ColumnLength:       6000,
			GoFieldName:        "ANKENBODYE",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anken_body_e",
			ProtobufFieldName:  "anken_body_e",
			ProtobufType:       "string",
			ProtobufPos:        16,
		},

		&ColumnInfo{
			Index:              16,
			Name:               "INFO_LONG",
			Comment:            `コンテンツロング文`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "INFOLONG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "info_long",
			ProtobufFieldName:  "info_long",
			ProtobufType:       "string",
			ProtobufPos:        17,
		},

		&ColumnInfo{
			Index:              17,
			Name:               "SEND_USER_ID",
			Comment:            `送信者ID`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "SENDUSERID",
			GoFieldType:        "int32",
			JSONFieldName:      "send_user_id",
			ProtobufFieldName:  "send_user_id",
			ProtobufType:       "int32",
			ProtobufPos:        18,
		},

		&ColumnInfo{
			Index:              18,
			Name:               "ANKEN_PROC_FLG",
			Comment:            `案件処理フラグ 0:未処理 1:ユーザ登録済 2:メール配信済`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ANKENPROCFLG",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anken_proc_flg",
			ProtobufFieldName:  "anken_proc_flg",
			ProtobufType:       "string",
			ProtobufPos:        19,
		},

		&ColumnInfo{
			Index:              19,
			Name:               "ANKEN_END_DTE",
			Comment:            `案件終了日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "ANKENENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "anken_end_dte",
			ProtobufFieldName:  "anken_end_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        20,
		},

		&ColumnInfo{
			Index:              20,
			Name:               "VALI_DTE",
			Comment:            `有効期日`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "VALIDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "vali_dte",
			ProtobufFieldName:  "vali_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        21,
		},

		&ColumnInfo{
			Index:              21,
			Name:               "ANNO_DISP_KBN",
			Comment:            `お知らせ表示区分\r\n    お知らせのみ 1:表示 0:非表示`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "ANNODISPKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "anno_disp_kbn",
			ProtobufFieldName:  "anno_disp_kbn",
			ProtobufType:       "string",
			ProtobufPos:        22,
		},

		&ColumnInfo{
			Index:              22,
			Name:               "UP_ANKEN_NO",
			Comment:            `上案件NO`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "UPANKENNO",
			GoFieldType:        "sql.NullInt64",
			JSONFieldName:      "up_anken_no",
			ProtobufFieldName:  "up_anken_no",
			ProtobufType:       "int32",
			ProtobufPos:        23,
		},

		&ColumnInfo{
			Index:              23,
			Name:               "CHIL_ANKEN_IDS",
			Comment:            `子案件ID(カンマ区切)`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "text",
			DatabaseTypePretty: "text(65535)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "text",
			ColumnLength:       65535,
			GoFieldName:        "CHILANKENIDS",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "chil_anken_ids",
			ProtobufFieldName:  "chil_anken_ids",
			ProtobufType:       "string",
			ProtobufPos:        24,
		},

		&ColumnInfo{
			Index:              24,
			Name:               "DEPT_START_KBN",
			Comment:            `部署起動区分    0:全社権限 1:部署権限`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "DEPTSTARTKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "dept_start_kbn",
			ProtobufFieldName:  "dept_start_kbn",
			ProtobufType:       "string",
			ProtobufPos:        25,
		},

		&ColumnInfo{
			Index:              25,
			Name:               "START_DEPT_DELI_KBN",
			Comment:            ``,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(2)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       2,
			GoFieldName:        "STARTDEPTDELIKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "start_dept_deli_kbn",
			ProtobufFieldName:  "start_dept_deli_kbn",
			ProtobufType:       "string",
			ProtobufPos:        26,
		},

		&ColumnInfo{
			Index:              26,
			Name:               "ANPI_DLCT_ID",
			Comment:            `安否DLCTID`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "ANPIDLCTID",
			GoFieldType:        "uint32",
			JSONFieldName:      "anpi_dlct_id",
			ProtobufFieldName:  "anpi_dlct_id",
			ProtobufType:       "uint32",
			ProtobufPos:        27,
		},

		&ColumnInfo{
			Index:              27,
			Name:               "JOKYO_CLIENT_TMPL_ID",
			Comment:            `状況一覧に表示する設問`,
			Notes:              `column is set for unsigned`,
			Nullable:           false,
			DatabaseTypeName:   "uint",
			DatabaseTypePretty: "uint",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "uint",
			ColumnLength:       -1,
			GoFieldName:        "JOKYOCLIENTTMPLID",
			GoFieldType:        "uint32",
			JSONFieldName:      "jokyo_client_tmpl_id",
			ProtobufFieldName:  "jokyo_client_tmpl_id",
			ProtobufType:       "uint32",
			ProtobufPos:        28,
		},

		&ColumnInfo{
			Index:              28,
			Name:               "START_RECV_KBN",
			Comment:            `起動予約区分 1:利用する 0:利用しない 2:予約停止`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STARTRECVKBN",
			GoFieldType:        "sql.NullString",
			JSONFieldName:      "start_recv_kbn",
			ProtobufFieldName:  "start_recv_kbn",
			ProtobufType:       "string",
			ProtobufPos:        29,
		},

		&ColumnInfo{
			Index:              29,
			Name:               "START_RECV_DTE",
			Comment:            `起動予約日時`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "STARTRECVDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "start_recv_dte",
			ProtobufFieldName:  "start_recv_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        30,
		},

		&ColumnInfo{
			Index:              30,
			Name:               "START_NOTIF_SEND_FLG",
			Comment:            `起動通知送信フラグ 0：起動通知を送信しない 1：起動通知を送信する`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "STARTNOTIFSENDFLG",
			GoFieldType:        "string",
			JSONFieldName:      "start_notif_send_flg",
			ProtobufFieldName:  "start_notif_send_flg",
			ProtobufType:       "string",
			ProtobufPos:        31,
		},

		&ColumnInfo{
			Index:              31,
			Name:               "AUTO_RESEND_KBN",
			Comment:            `自動再送 0:再送なし 1:再送する 2:ユーザによる中断`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "AUTORESENDKBN",
			GoFieldType:        "string",
			JSONFieldName:      "auto_resend_kbn",
			ProtobufFieldName:  "auto_resend_kbn",
			ProtobufType:       "string",
			ProtobufPos:        32,
		},

		&ColumnInfo{
			Index:              32,
			Name:               "AUTO_RESEND_TIMES",
			Comment:            `再送実施回数`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "int",
			DatabaseTypePretty: "int",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "int",
			ColumnLength:       -1,
			GoFieldName:        "AUTORESENDTIMES",
			GoFieldType:        "int32",
			JSONFieldName:      "auto_resend_times",
			ProtobufFieldName:  "auto_resend_times",
			ProtobufType:       "int32",
			ProtobufPos:        33,
		},

		&ColumnInfo{
			Index:              33,
			Name:               "AUTO_RESEND_DTE",
			Comment:            `最後に自動再送した時刻`,
			Notes:              ``,
			Nullable:           true,
			DatabaseTypeName:   "datetime",
			DatabaseTypePretty: "datetime",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "datetime",
			ColumnLength:       -1,
			GoFieldName:        "AUTORESENDDTE",
			GoFieldType:        "time.Time",
			JSONFieldName:      "auto_resend_dte",
			ProtobufFieldName:  "auto_resend_dte",
			ProtobufType:       "google.protobuf.Timestamp",
			ProtobufPos:        34,
		},

		&ColumnInfo{
			Index:              34,
			Name:               "INCL_RESIDENCE_KBN",
			Comment:            `居住地を含める（0:含めない 1:含める）`,
			Notes:              ``,
			Nullable:           false,
			DatabaseTypeName:   "char",
			DatabaseTypePretty: "char(1)",
			IsPrimaryKey:       false,
			IsAutoIncrement:    false,
			IsArray:            false,
			ColumnType:         "char",
			ColumnLength:       1,
			GoFieldName:        "INCLRESIDENCEKBN",
			GoFieldType:        "string",
			JSONFieldName:      "incl_residence_kbn",
			ProtobufFieldName:  "incl_residence_kbn",
			ProtobufType:       "string",
			ProtobufPos:        35,
		},

		&ColumnInfo{
			Index:              35,
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
			ProtobufPos:        36,
		},

		&ColumnInfo{
			Index:              36,
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
			ProtobufPos:        37,
		},

		&ColumnInfo{
			Index:              37,
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
			ProtobufPos:        38,
		},

		&ColumnInfo{
			Index:              38,
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
			ProtobufPos:        39,
		},
	},
}

// TableName sets the insert table name for this struct type
func (a *ANKEN_) TableName() string {
	return "ANKEN"
}

// BeforeSave invoked before saving, return an error if field is not populated.
func (a *ANKEN_) BeforeSave() error {
	return nil
}

// Prepare invoked before saving, can be used to populate fields etc.
func (a *ANKEN_) Prepare() {
}

// Validate invoked before performing action, return an error if field is not populated.
func (a *ANKEN_) Validate(action Action) error {
	return nil
}

// TableInfo return table meta data
func (a *ANKEN_) TableInfo() *TableInfo {
	return ANKENTableInfo
}
