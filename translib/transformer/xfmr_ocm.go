package transformer

import (
	log "github.com/golang/glog"
	"strings"
)

func init() {
	XlateFuncBind("YangToDb_ocm_channel_key_xfmr", YangToDb_ocm_channel_key_xfmr)
	XlateFuncBind("DbToYang_ocm_channel_key_xfmr", DbToYang_ocm_channel_key_xfmr)
	XlateFuncBind("YangToDb_ocm_lower_frequency_xfmr", YangToDb_ocm_lower_frequency_xfmr)
	XlateFuncBind("DbToYang_ocm_lower_frequency_xfmr", DbToYang_ocm_lower_frequency_xfmr)
	XlateFuncBind("YangToDb_ocm_upper_frequency_xfmr", YangToDb_ocm_upper_frequency_xfmr)
	XlateFuncBind("DbToYang_ocm_upper_frequency_xfmr", DbToYang_ocm_upper_frequency_xfmr)
}

var YangToDb_ocm_channel_key_xfmr KeyXfmrYangToDb = func(inParams XfmrParams) (string, error) {
	if log.V(3) {
		log.Info("YangToDb_ocm_key_xfmr: root: ", inParams.ygRoot,
			", uri: ", inParams.uri)
	}
	pathInfo := NewPathInfo(inParams.uri)
	name := pathInfo.Var("name")
	lower := pathInfo.Var("lower-frequency")
	upper := pathInfo.Var("upper-frequency")
	key := name + "|" + lower + "|" + upper

	return key, nil
}

var DbToYang_ocm_channel_key_xfmr KeyXfmrDbToYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 3 {
		rmap["lower-frequency"] = TableKeys[1]
		rmap["upper-frequency"] = TableKeys[2]
	}
	log.Info("DbToYang_ocm_channel_key_xfmr : - ", rmap)

	return rmap, err
}

var YangToDb_ocm_lower_frequency_xfmr FieldXfmrYangToDb = func(inParams XfmrParams) (map[string]string, error) {
	var err error
	rmap := make(map[string]string)

	rmap["NULL"] = "NULL"

	return rmap, err
}

var DbToYang_ocm_lower_frequency_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 2 {
		rmap["lower-frequency"] = TableKeys[1]
	}

	return rmap, err
}

var YangToDb_ocm_upper_frequency_xfmr FieldXfmrYangToDb = func(inParams XfmrParams) (map[string]string, error) {
	var err error
	rmap := make(map[string]string)

	rmap["NULL"] = "NULL"

	return rmap, err
}

var DbToYang_ocm_upper_frequency_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	key := inParams.key
	TableKeys := strings.Split(key, "|")

	if len(TableKeys) >= 2 {
		rmap["upper-frequency"] = TableKeys[2]
	}

	return rmap, err
}
