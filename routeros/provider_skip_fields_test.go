package routeros

import (
	"strings"
	"testing"
)

// Fields that are intentionally both modelled and skipped on read: write-only
// action arguments, or values the resource manages separately.
var skipFieldCollisionAllowlist = map[string]map[string]struct{}{
	"/certificate":                              {"import": {}, "sign": {}, "sign_via_scep": {}},
	"/container":                                {"running": {}},
	"/interface/ethernet":                       {"factory_name": {}},
	"/interface/ethernet/switch":                {"switch_id": {}},
	"/interface/ethernet/switch/port":           {"name": {}},
	"/interface/ethernet/switch/port-isolation": {"name": {}},
	"/interface/w60g":                           {"tx_sector": {}},
	"/ip/hotspot":                               {"keepalive_timeout": {}},
	"/ip/hotspot/service-port":                  {"name": {}},
	"/system/script":                            {"launch_trigger": {}},
	"/tool/sniffer":                             {"enabled": {}},
}

// A field listed in MetaSkipFields is never read back from RouterOS, so a
// modelled field of the same name drifts on every plan.
func TestSkipFieldsDoNotShadowSchema(t *testing.T) {
	for name, r := range Provider().ResourcesMap {
		sf, ok := r.Schema[MetaSkipFields]
		if !ok {
			continue
		}
		path := r.Schema[MetaResourcePath].Default.(string)
		for _, f := range strings.Split(sf.Default.(string), ",") {
			f = strings.Trim(strings.TrimSpace(f), `"`)
			if _, modelled := r.Schema[f]; !modelled {
				continue
			}
			if _, allowed := skipFieldCollisionAllowlist[path][f]; allowed {
				continue
			}
			t.Errorf("%s (%s): %q is both a schema field and a skip field", name, path, f)
		}
	}
}
