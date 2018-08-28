package main

import (
	"bytes"
	"html/template"
	"path/filepath"
)

// consume list of templates and release their full path counterparts
func fileNames(tdir string, filenames ...string) []string {
	flist := []string{}
	for _, fname := range filenames {
		flist = append(flist, filepath.Join(tdir, fname))
	}
	return flist
}

// parse template with given data
func parseTmpl(tdir, tmpl string, data interface{}) string {
	buf := new(bytes.Buffer)
	filenames := fileNames(tdir, tmpl)
	t := template.Must(template.ParseFiles(filenames...))
	err := t.Execute(buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

// CMSTemplates structure
type CMSTemplates struct {
	main, dashboards, activities, sources, training, shifters, contacts, meetings string
}

// Main method for CMSTemplates structure
func (q CMSTemplates) Main(tdir string, tmplData map[string]interface{}) string {
	if q.main != "" {
		return q.main
	}
	q.main = parseTmpl(_tdir, "main.tmpl", tmplData)
	return q.main
}

// Dashboards method for CMSTemplates structure
func (q CMSTemplates) Dashboards(tdir string, tmplData map[string]interface{}) string {
	if q.dashboards != "" {
		return q.dashboards
	}
	q.dashboards = parseTmpl(_tdir, "dashboards.tmpl", tmplData)
	return q.dashboards
}

// Activities method for CMSTemplates structure
func (q CMSTemplates) Activities(tdir string, tmplData map[string]interface{}) string {
	if q.activities != "" {
		return q.activities
	}
	q.activities = parseTmpl(_tdir, "activities.tmpl", tmplData)
	return q.activities
}

// Sources method for CMSTemplates structure
func (q CMSTemplates) Sources(tdir string, tmplData map[string]interface{}) string {
	if q.sources != "" {
		return q.sources
	}
	q.sources = parseTmpl(_tdir, "sources.tmpl", tmplData)
	return q.sources
}

// Training method for CMSTemplates structure
func (q CMSTemplates) Training(tdir string, tmplData map[string]interface{}) string {
	if q.training != "" {
		return q.training
	}
	q.training = parseTmpl(_tdir, "training.tmpl", tmplData)
	return q.training
}

// Shifters method for CMSTemplates structure
func (q CMSTemplates) Shifters(tdir string, tmplData map[string]interface{}) string {
	if q.shifters != "" {
		return q.shifters
	}
	q.shifters = parseTmpl(_tdir, "shifters.tmpl", tmplData)
	return q.shifters
}

// Contacts method for CMSTemplates structure
func (q CMSTemplates) Contacts(tdir string, tmplData map[string]interface{}) string {
	if q.contacts != "" {
		return q.contacts
	}
	q.contacts = parseTmpl(_tdir, "contacts.tmpl", tmplData)
	return q.contacts
}

// Meetings method for CMSTemplates structure
func (q CMSTemplates) Meetings(tdir string, tmplData map[string]interface{}) string {
	if q.meetings != "" {
		return q.meetings
	}
	q.meetings = parseTmpl(_tdir, "meetings.tmpl", tmplData)
	return q.meetings
}
