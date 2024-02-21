package transformer

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/sonic-mgmt-common/translib/db"
	log "github.com/golang/glog"
	"strconv"
	"strings"
)

const (
	RPC_STATUS_SUCCESSFUL string = "Successful"
	RPC_STATUS_FAILED     string = "Failed"
)

func init() {
	XlateFuncBind("YangToDb_otdr_scan_type_key_xfmr", YangToDb_otdr_scan_type_key_xfmr)
	XlateFuncBind("DbToYang_otdr_scan_type_key_xfmr", DbToYang_otdr_scan_type_key_xfmr)
	XlateFuncBind("YangToDb_otdr_scan_type_field_xfmr", YangToDb_otdr_scan_type_field_xfmr)
	XlateFuncBind("DbToYang_otdr_scan_type_field_xfmr", DbToYang_otdr_scan_type_field_xfmr)

	XlateFuncBind("YangToDb_otdr_result_key_xfmr", YangToDb_otdr_result_key_xfmr)
	XlateFuncBind("DbToYang_otdr_result_key_xfmr", DbToYang_otdr_result_key_xfmr)
	XlateFuncBind("DbToYang_otdr_completion_time_field_xfmr", DbToYang_otdr_completion_time_field_xfmr)

	XlateFuncBind("YangToDb_otdr_event_key_xfmr", YangToDb_otdr_event_key_xfmr)
	XlateFuncBind("DbToYang_otdr_event_key_xfmr", DbToYang_otdr_event_key_xfmr)
	XlateFuncBind("DbToYang_otdr_event_type_field_xfmr", DbToYang_otdr_event_type_field_xfmr)

	XlateFuncBind("rpc_otdr_scan_cb", rpc_otdr_scan_cb)
}

var YangToDb_otdr_scan_type_key_xfmr KeyXfmrYangToDb = func(inParams XfmrParams) (string, error) {
	if log.V(3) {
		log.Info("YangToDb_otdr_scan_type_key_xfmr: root: ", inParams.ygRoot,
			", uri: ", inParams.uri)
	}
	pathInfo := NewPathInfo(inParams.uri)
	osckey := pathInfo.Var("scan-type")

	return osckey, nil
}

var DbToYang_otdr_scan_type_key_xfmr KeyXfmrDbToYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	res_map := make(map[string]interface{}, 1)
	var err error

	if log.V(3) {
		log.Info("DbToYang_otdr_scan_type_key_xfmr: ", inParams.key, "uri:", inParams.uri)
	}

	res_map["scan-type"] = inParams.key

	return res_map, err
}

var YangToDb_otdr_scan_type_field_xfmr FieldXfmrYangToDb = func(inParams XfmrParams) (map[string]string, error) {
	res_map := make(map[string]string)
	var err error
	res_map["NULL"] = "NULL"
	return res_map, err
}

var DbToYang_otdr_scan_type_field_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	rmap["scan-type"] = inParams.key

	return rmap, err
}

var YangToDb_otdr_result_key_xfmr KeyXfmrYangToDb = func(inParams XfmrParams) (string, error) {
	if log.V(3) {
		log.Info("YangToDb_otdr_result_key_xfmr: root: ", inParams.ygRoot,
			", uri: ", inParams.uri)
	}

	pathInfo := NewPathInfo(inParams.uri)
	name := pathInfo.Var("name")
	completionTime := pathInfo.Var("completion-time")
	key := name + "|" + completionTime

	return key, nil
}

var DbToYang_otdr_result_key_xfmr KeyXfmrDbToYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 2 {
		rmap["completion-time"] = TableKeys[1]
	}

	return rmap, err
}

var DbToYang_otdr_completion_time_field_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 2 {
		rmap["completion-time"] = TableKeys[1]
	}

	return rmap, err
}

var YangToDb_otdr_event_key_xfmr KeyXfmrYangToDb = func(inParams XfmrParams) (string, error) {
	if log.V(3) {
		log.Info("YangToDb_otdr_event_key_xfmr: root: ", inParams.ygRoot,
			", uri: ", inParams.uri)
	}
	pathInfo := NewPathInfo(inParams.uri)
	name := pathInfo.Var("name")
	completionTime := pathInfo.Var("completion-time")
	eventId := pathInfo.Var("event-id")
	key := name + "|" + completionTime + "|" + eventId

	return key, nil
}

var DbToYang_otdr_event_key_xfmr KeyXfmrDbToYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 3 {
		//rmap["completion-time"] = TableKeys[1]
		eventId, result := strconv.ParseUint(TableKeys[2], 10, 8)
		//log.Info("DbToYang_otdr_event_key_xfmr : eventId: ", eventId, "result: ", result, "TableKeys[2]: ", TableKeys[2])
		if result == nil {
			rmap["event-id"] = uint8(eventId)
		} else {
			log.Info("Error DbToYang_otdr_event_key_xfmr : key: ", inParams.key, "result: ", result)
		}
	}

	return rmap, err
}

var DbToYang_otdr_event_type_field_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 3 {
		//rmap["completion-time"] = TableKeys[1]

		eventId, result := strconv.ParseUint(TableKeys[2], 10, 8)
		//log.Info("DbToYang_otdr_event_key_xfmr : eventId: ", eventId, "result: ", result)
		if result == nil {
			rmap["event-id"] = uint8(eventId)
		} else {
			log.Info("Error DbToYang_otdr_event_type_field_xfmr : key: ", inParams.key, "result: ", result)
		}
	}

	return rmap, err
}

var rpc_otdr_scan_cb RpcCallpoint = func(body []byte, dbs [db.MaxDB]*db.DB) ([]byte, error) {
	var err error
	var operand struct {
		Input struct {
			Name         string `json:"name"`
			Operation    string `json:"operation"`
			OtdrScanType string `json:"otdr-scan-type"`
		} `json:"openconfig-gnoi-otdr:input"`
	}

	var result struct {
		Output struct {
			Status        string `json:"status"`
			StatusMessage string `json:"status-message"`
		} `json:"openconfig-gnoi-otdr:output"`
	}

	//parse input paras
	err = json.Unmarshal(body, &operand)
	if err != nil {
		log.Errorf("umarshal input data failed")

		result.Output.Status = RPC_STATUS_FAILED

		msg := fmt.Sprintf("[FAILED] to umarshal input data")
		result.Output.StatusMessage = msg
		response, err := json.Marshal(&result)

		return response, err
	}

	//
	log.Infof("rpc_otdr_scan_cb... name: %v, operation: %v, otdr-scan-type: %v", operand.Input.Name, operand.Input.Operation, operand.Input.OtdrScanType)

	//TODO: check parameters

	//TODO: check if current otdr is already scanning ...

	stateDB := dbs[db.StateDB]
	channel := "OTDR_NOTIFICATION"
	field := "scan"
	value := "true"
	message := fmt.Sprintf("[\"set\",\"%s\",\"%s\",\"%s\"]", operand.Input.Name, field, value)
	// ["set", "LineT", "scan", "true"]
	// ["set", "LineT/LineR/OtdrT", "scan", "true/false", "scan-type", "short/medium/long/auto/customized"]
	log.Infof("publish %s %s", channel, message)
	err = stateDB.Publish(channel, message)
	if err != nil {
		log.Errorf("publish OTDR_NOTIFICATION failed")
		result.Output.Status = RPC_STATUS_FAILED

		msg := fmt.Sprintf("[FAILED] publish OTDR_NOTIFICATION")
		result.Output.StatusMessage = msg
	} else {
		result.Output.Status = RPC_STATUS_SUCCESSFUL

		msg := fmt.Sprintf("OTDR is scanning ...")
		result.Output.StatusMessage = msg
	}

	response, err2 := json.Marshal(&result)

	return response, err2
}
