package ducklib

type User struct {
	Id        string   `json:"id"`
	Email     string   `json:"email"`
	Password  string   `json:"password"`
	Firstname string   `json:"firstname"`
	Lastname  string   `json:"lastname"`
	Revision  string   `json:"_rev"`
	Documents []string `json:"documents"`
}

type Response struct {
	Ok     bool    `json:"ok"`
	Reason *string `json:"reason,omitempty"`
	ID     *string `json:"id,omitempty"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (u *User) fromValueMap(mp map[string]interface{}) {

	u.Id = mp["_id"].(string)
	u.Revision = mp["_rev"].(string)

	u.Firstname = mp["firstname"].(string)
	u.Lastname = mp["lastname"].(string)
	u.Password = mp["password"].(string)
	u.Email = mp["email"].(string)

	if docs, prs := mp["documents"].([]interface{}); prs {
		u.Documents = make([]string, len(docs))
		for i, v := range docs {
			u.Documents[i] = v.(string)
		}
	}

}

type Document struct {
	ID         string      `json:"id"`
	Name       string      `json:"name"`
	Revision   string      `json:"_rev"`
	Statements []Statement `json:"statements"`
}

type Statement struct {
	UseScope     string `json:"useScope"`
	Qualifier    string `json:"qualifier"`
	DataCategory string `json:"dataCategory"`
	SourceScope  string `json:"sourceScope"`
	Action       string `json:"action"`
	ResultScope  string `json:"resultScope"`
	TrackingID   string `json:"trackingId"`
}

func (d *Document) fromValueMap(mp map[string]interface{}) {

	d.ID = mp["_id"].(string)
	d.Revision = mp["_rev"].(string)
	d.Name = mp["name"].(string)

	if stmts, prs := mp["statements"].([]interface{}); prs {
		d.Statements = make([]Statement, len(stmts))
		for i, stmt := range stmts {
			s := new(Statement)
			s.fromInterfaceMap(stmt.(map[string]interface{}))
			d.Statements[i] = *s
		}
	}

}

func (s *Statement) fromInterfaceMap(mp map[string]interface{}) {

	s.UseScope = getFieldValue(mp, "useScope")
	s.Qualifier = getFieldValue(mp, "qualifier")
	s.DataCategory = getFieldValue(mp, "dataCategory")
	s.SourceScope = getFieldValue(mp, "sourceScope")
	s.Action = getFieldValue(mp, "action")
	s.ResultScope = getFieldValue(mp, "resultScope")
	s.TrackingID = getFieldValue(mp, "trackingId")

}
func getFieldValue(mp map[string]interface{}, field string) string {

	if interf, ok := mp[field]; ok {
		if value, ok := interf.(string); ok {
			return value
		}
	}
	return ""
}
