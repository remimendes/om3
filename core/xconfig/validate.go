package xconfig

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/opensvc/om3/core/driver"
	"github.com/opensvc/om3/core/naming"
	"github.com/opensvc/om3/core/rawconfig"
	"github.com/opensvc/om3/core/resourceid"
	"github.com/opensvc/om3/util/capabilities"
	"github.com/opensvc/om3/util/key"
	"github.com/opensvc/om3/util/render/tree"
	"github.com/opensvc/om3/util/stringslice"
)

type (
	ValidateAlerts []ValidateAlert
	ValidateAlert  struct {
		Path    naming.Path        `json:"path"`
		Level   ValidateAlertLevel `json:"level"`
		Kind    ValidateAlertKind  `json:"kind"`
		Key     key.T              `json:"key"`
		Driver  driver.ID          `json:"driver"`
		Comment string             `json:"comment"`
	}
	ValidateAlertKind  int
	ValidateAlertLevel int
)

const (
	validateAlertLevelWarn ValidateAlertLevel = iota
	validateAlertLevelError

	validateAlertKindScoping ValidateAlertKind = iota
	validateAlertKindUnknown
	validateAlertKindUnknownDriver
	validateAlertKindEval
	validateAlertKindCandidates
	validateAlertKindDeprecated
	validateAlertKindCapabilities
)

var (
	validateAlertLevelWarnStr  = "warning"
	validateAlertLevelErrorStr = "error"
	validateAlertLevelNames    = map[ValidateAlertLevel]string{
		validateAlertLevelWarn:  validateAlertLevelWarnStr,
		validateAlertLevelError: validateAlertLevelErrorStr,
	}
	validateAlertLevelFromNames = map[string]ValidateAlertLevel{
		validateAlertLevelWarnStr:  validateAlertLevelWarn,
		validateAlertLevelErrorStr: validateAlertLevelError,
	}
	validateAlertKindUnknownDriverStr = "driver does not exist"
	validateAlertKindScopingStr       = "keyword does not support scoping"
	validateAlertKindUnknownStr       = "keyword does not exist"
	validateAlertKindEvalStr          = "keyword does not evaluate"
	validateAlertKindCandidatesStr    = "keyword value is not in allowed candidates"
	validateAlertKindDeprecatedStr    = "keyword is deprecated"
	validateAlertKindCapabilitiesStr  = "driver is not in node capabilities"
	validateAlertKindNames            = map[ValidateAlertKind]string{
		validateAlertKindScoping:       validateAlertKindScopingStr,
		validateAlertKindUnknown:       validateAlertKindUnknownStr,
		validateAlertKindUnknownDriver: validateAlertKindUnknownDriverStr,
		validateAlertKindEval:          validateAlertKindEvalStr,
		validateAlertKindCandidates:    validateAlertKindCandidatesStr,
		validateAlertKindDeprecated:    validateAlertKindDeprecatedStr,
		validateAlertKindCapabilities:  validateAlertKindCapabilitiesStr,
	}
	validateAlertKindFromNames = map[string]ValidateAlertKind{
		validateAlertKindScopingStr:       validateAlertKindScoping,
		validateAlertKindUnknownStr:       validateAlertKindUnknown,
		validateAlertKindUnknownDriverStr: validateAlertKindUnknownDriver,
		validateAlertKindEvalStr:          validateAlertKindEval,
		validateAlertKindCandidatesStr:    validateAlertKindCandidates,
		validateAlertKindDeprecatedStr:    validateAlertKindDeprecated,
		validateAlertKindCapabilitiesStr:  validateAlertKindCapabilities,
	}
)

func (t T) NewValidateAlertScoping(k key.T, did driver.ID) ValidateAlert {
	return ValidateAlert{
		Path:   t.Path,
		Kind:   validateAlertKindScoping,
		Level:  validateAlertLevelError,
		Key:    k,
		Driver: did,
	}
}

func (t T) NewValidateAlertUnknownDriver(k key.T, did driver.ID) ValidateAlert {
	return ValidateAlert{
		Path:   t.Path,
		Kind:   validateAlertKindUnknownDriver,
		Level:  validateAlertLevelWarn,
		Key:    k,
		Driver: did,
	}
}

func (t T) NewValidateAlertUnknown(k key.T, did driver.ID) ValidateAlert {
	return ValidateAlert{
		Path:   t.Path,
		Kind:   validateAlertKindUnknown,
		Level:  validateAlertLevelWarn,
		Key:    k,
		Driver: did,
	}
}

func (t T) NewValidateAlertCandidates(k key.T, did driver.ID) ValidateAlert {
	return ValidateAlert{
		Path:   t.Path,
		Kind:   validateAlertKindCandidates,
		Level:  validateAlertLevelError,
		Key:    k,
		Driver: did,
	}
}

func (t T) NewValidateAlertEval(k key.T, did driver.ID, comment string) ValidateAlert {
	return ValidateAlert{
		Path:    t.Path,
		Kind:    validateAlertKindEval,
		Level:   validateAlertLevelError,
		Key:     k,
		Driver:  did,
		Comment: comment,
	}
}

func (t T) NewValidateAlertDeprecated(k key.T, did driver.ID, release, replacedBy string) ValidateAlert {
	comment := fmt.Sprintf("since %s", release)
	if replacedBy != "" {
		comment += fmt.Sprintf("replaced by %s", replacedBy)
	}
	return ValidateAlert{
		Path:    t.Path,
		Kind:    validateAlertKindDeprecated,
		Level:   validateAlertLevelWarn,
		Key:     k,
		Driver:  did,
		Comment: comment,
	}
}

func (t T) NewValidateAlertCapabilities(k key.T, did driver.ID) ValidateAlert {
	return ValidateAlert{
		Path:   t.Path,
		Kind:   validateAlertKindCapabilities,
		Level:  validateAlertLevelWarn,
		Key:    k,
		Driver: did,
	}
}

func (t ValidateAlertKind) String() string {
	if s, ok := validateAlertKindNames[t]; ok {
		return s
	} else {
		return ""
	}
}

func (t ValidateAlertLevel) String() string {
	if s, ok := validateAlertLevelNames[t]; ok {
		return s
	} else {
		return ""
	}
}

func (t ValidateAlertLevel) MarshalJSON() ([]byte, error) {
	if s, ok := validateAlertLevelNames[t]; ok {
		return json.Marshal(s)
	} else {
		return nil, fmt.Errorf("unknown validate alert level: %d", t)
	}
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *ValidateAlertLevel) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t, _ = validateAlertLevelFromNames[j]
	return nil
}

func (t ValidateAlertKind) MarshalJSON() ([]byte, error) {
	if s, ok := validateAlertKindNames[t]; ok {
		return json.Marshal(s)
	} else {
		return nil, fmt.Errorf("unknown validate alert kind: %d", t)
	}
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (t *ValidateAlertKind) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*t, _ = validateAlertKindFromNames[j]
	return nil
}

func (t ValidateAlerts) HasError() bool {
	return t.has(validateAlertLevelError)
}

func (t ValidateAlerts) HasWarn() bool {
	return t.has(validateAlertLevelWarn)
}

func (t ValidateAlerts) has(lvl ValidateAlertLevel) bool {
	for _, alert := range t {
		if alert.Level == lvl {
			return true
		}
	}
	return false
}

func (t ValidateAlerts) Render() string {
	tr := t.Tree()
	return tr.Render()
}

func (t ValidateAlerts) Tree() *tree.Tree {
	tr := tree.New()
	if len(t) == 0 {
		return tr
	}
	node := tr.AddNode()
	node.AddColumn().AddText("alert level").SetColor(rawconfig.Color.Secondary)
	node.AddColumn().AddText("key").SetColor(rawconfig.Color.Secondary)
	node.AddColumn().AddText("driver").SetColor(rawconfig.Color.Secondary)
	node.AddColumn().AddText("kind").SetColor(rawconfig.Color.Secondary)
	node.AddColumn().AddText("comment").SetColor(rawconfig.Color.Secondary)
	for _, alert := range t {
		n := tr.AddNode()
		color := rawconfig.Color.Warning
		if alert.Level == validateAlertLevelError {
			color = rawconfig.Color.Error
		}
		driver := alert.Driver.String()
		if driver == "" {
			driver = "-"
		}
		comment := alert.Comment
		if comment == "" {
			comment = "-"
		}
		n.AddColumn().AddText(alert.Level.String()).SetColor(color)
		n.AddColumn().AddText(alert.Key.String())
		n.AddColumn().AddText(driver)
		n.AddColumn().AddText(alert.Kind.String())
		n.AddColumn().AddText(comment)
	}
	return tr
}

func (t T) Validate() (ValidateAlerts, error) {
	alerts := make(ValidateAlerts, 0)
	for _, s := range t.file.Sections() {
		var did driver.ID
		section := s.Name()
		sectionType := t.GetString(key.New(section, "type"))
		if rid, err := resourceid.Parse(section); err == nil {
			did = driver.NewID(rid.DriverGroup(), sectionType)
			if did.Name != "" {
				if sectionType == "" {
					sectionType = did.Name
				}
				if !driver.Exists(did) {
					alerts = append(alerts, t.NewValidateAlertUnknownDriver(key.T{Section: section}, did))
					continue
				}
				if !capabilities.Has(did.Cap()) {
					alerts = append(alerts, t.NewValidateAlertCapabilities(key.T{Section: section}, did))
					continue
				}
			}
		}
		for option := range s.KeysHash() {
			k := key.Parse(section + "." + option)
			if k.BaseOption() == "type" {
				continue
			}
			kw, err := getKeyword(k, sectionType, t.Referrer)
			if err != nil {
				alerts = append(alerts, t.NewValidateAlertUnknown(k, did))
				continue
			}
			if strings.Contains(k.Option, "@") && !kw.Scopable {
				alerts = append(alerts, t.NewValidateAlertScoping(k, did))
			}
			v, err := t.evalStringAs(k, kw, "")
			if err != nil {
				alerts = append(alerts, t.NewValidateAlertEval(k, did, fmt.Sprint(err)))
				continue
			}
			if kw.Deprecated != "" {
				alerts = append(alerts, t.NewValidateAlertDeprecated(k, did, kw.Deprecated, kw.ReplacedBy))
			}
			if (len(kw.Candidates) > 0) && !stringslice.Has(v, kw.Candidates) {
				alerts = append(alerts, t.NewValidateAlertCandidates(k, did))
			}
		}
	}
	if alerts.HasError() {
		return alerts, fmt.Errorf("")
	}
	return alerts, nil
}

func ValidateFile(p string, ref Referrer) error {
	cfg, err := NewObject(p)
	if err != nil {
		return err
	}
	cfg.Referrer = ref
	if _, err := cfg.Validate(); err != nil {
		return err
	}
	return nil
}

func (t ValidateAlerts) String() string {
	l := make([]string, len(t))
	for i, alert := range t {
		l[i] = alert.String()
	}
	return strings.Join(l, "\n")
}

func (t ValidateAlert) String() string {
	buff := fmt.Sprintf("[%s] path %s key %s: %s", t.Level, t.Path, t.Key, t.Kind)
	if t.Comment != "" {
		buff += ", " + t.Comment
	}
	return buff
}
