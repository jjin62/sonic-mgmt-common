package transformer

import (
	log "github.com/golang/glog"
)

func init() {
	XlateFuncBind("YangToDb_osc_key_xfmr", YangToDb_osc_key_xfmr)
	XlateFuncBind("DbToYang_osc_key_xfmr", DbToYang_osc_key_xfmr)
	XlateFuncBind("YangToDb_osc_interface_xfmr", YangToDb_osc_interface_xfmr)
	XlateFuncBind("DbToYang_osc_interface_xfmr", DbToYang_osc_interface_xfmr)
}

var YangToDb_osc_key_xfmr KeyXfmrYangToDb = func(inParams XfmrParams) (string, error) {
	if log.V(3) {
		log.Info("YangToDb_osc_key_xfmr: root: ", inParams.ygRoot,
			", uri: ", inParams.uri)
	}
	pathInfo := NewPathInfo(inParams.uri)
	osckey := pathInfo.Var("interface")

	return osckey, nil
}

var DbToYang_osc_key_xfmr KeyXfmrDbToYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	res_map := make(map[string]interface{}, 1)
	var err error

	if log.V(3) {
		log.Info("DbToYang_osc_key_xfmr: ", inParams.key)
	}

	res_map["interface"] = inParams.key

	return res_map, err
}

var YangToDb_osc_interface_xfmr FieldXfmrYangToDb = func(inParams XfmrParams) (map[string]string, error) {
	res_map := make(map[string]string)
	var err error
	return res_map, err
}

var DbToYang_osc_interface_xfmr FieldXfmrDbtoYang = func(inParams XfmrParams) (map[string]interface{}, error) {
	var err error
	rmap := make(map[string]interface{})
	rmap["interface"] = inParams.key

	return rmap, err
}
