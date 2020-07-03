package grpcservice

import (
	"bytes"
	"text/template"
)

const (
	PG_HBA_CONF = `#~lx-----------------------------------------------------------------------------
# PG HBA CONFIG
#------------------------------------------------------------------------------

#test for HA remote access
host 	all  		all			{{.IP}}/32            trust			

# IMPORTANT: these settings will always be overriden when the server boots. They
# are set dynamically and so should never change.

host    replication     {{.User}}        {{.IP}}/32            trust`

	PG_CONF = `#~lx-----------------------------------------------------------------------------
# PG CONFIG
#------------------------------------------------------------------------------
# IMPORTANT: these settings will always be overriden when the server boots. They
# are set dynamically and so should never change.

listen_addresses = '*'
port = {{.Port}}
wal_level = replica
max_wal_senders = 10
wal_keep_segments = 32
hot_standby = on
fsync = off
wal_sender_timeout = 60s
synchronous_commit = off

`

	PG_EXTENSIONS_FOR_MASTER_CONF = `
# added only the the node running postgres as 'master'
synchronous_standby_names = 'slave'
`
	PG_EXTENSIONS_FOR_SLAVE_CONF = `
# added the the node running postgres as 'slave'
primary_conninfo = 'application_name=slave host={{.Host}} port={{.Port}} user={{.User}} password={{.Password}} sslmode=prefer sslcompression=0 gssencmode=prefer krbsrvname=postgres target_session_attrs=any'
`
)

func compileTemplate(templateStr string, values interface{}) (string, error) {
	buf := new(bytes.Buffer)
	t := template.Must(template.New("compiled_template").Parse(templateStr))
	if err := t.Execute(buf, values); err != nil {
		return "", err
	}

	return buf.String(), nil
}
