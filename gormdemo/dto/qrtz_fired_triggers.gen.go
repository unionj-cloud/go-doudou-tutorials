// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.
// Code generated by gorm.io/gen. YOU CAN EDIT.

package dto

//go:generate go-doudou name --file $GOFILE

// QrtzFiredTrigger mapped from table <qrtz_fired_triggers>
type QrtzFiredTrigger struct {
	SCHEDNAME        string `json:"sCHEDNAME"`
	ENTRYID          string `json:"eNTRYID"`
	TRIGGERNAME      string `json:"tRIGGERNAME"`
	TRIGGERGROUP     string `json:"tRIGGERGROUP"`
	INSTANCENAME     string `json:"iNSTANCENAME"`
	FIREDTIME        int64  `json:"fIREDTIME"`
	SCHEDTIME        int64  `json:"sCHEDTIME"`
	PRIORITY         int32  `json:"pRIORITY"`
	STATE            string `json:"sTATE"`
	JOBNAME          string `json:"jOBNAME"`
	JOBGROUP         string `json:"jOBGROUP"`
	ISNONCONCURRENT  string `json:"iSNONCONCURRENT"`
	REQUESTSRECOVERY string `json:"rEQUESTSRECOVERY"`
}
