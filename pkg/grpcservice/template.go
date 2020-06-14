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

# these configuration options have been removed from their standard location and
# placed here so that Pagoda Box could override them with the neccessary values
# to configure redundancy.

# IMPORTANT: these settings will always be overriden when the server boots. They
# are set dynamically and so should never change.

host    replication     {{.User}}        {{.IP}}/32            trust`

	PG_CONF = `#~lx-----------------------------------------------------------------------------
# PG CONFIG
#------------------------------------------------------------------------------

# these configuration options have been removed from their standard location and
# placed here so that Pagoda Box could override them with the neccessary values
# to configure redundancy.

# IMPORTANT: these settings will always be overriden when the server boots. They
# are set dynamically and so should never change.

listen_addresses = '0.0.0.0'      # what IP address(es) to listen on;
                                  # comma-separated list of addresses;
                                  # defaults to 'localhost'; use '*' for all
                                  # (change requires restart)
port = {{.Port}}                  # (change requires restart)
wal_level = replica				  # minimal、replica、logical 
                                  # (change requires restart)
archive_mode = on                 # allows archiving to be done
                                  # (change requires restart)
archive_command = 'exit 0'        # command to use to archive a logfile segment
                                  # placeholders: \%p = path of file to archive
                                  #               \%f = file name only
                                  # e.g. 'test ! -f /mnt/server/archivedir/\%f && cp \%p /mnt/server/archivedir/\%f'
max_wal_senders = 10              # max number of walsender processes
                                  # (change requires restart)
wal_keep_segments = 32          	# in logfile segments, 16MB each; 0 disables
hot_standby = on                  # "on" allows queries during recovery
                                  # (change requires restart)
fsync = off 
wal_sender_timeout = 60s
synchronous_commit = off 

# - Standby Server Config -
# These settings are ignored on a master server.
primary_conninfo = 'host={{.Host}} port={{.Port}} user={{.User}} password={{.Password}} application_name=slave'
`

	PG_EXTENSIONS_FOR_MASTER_CONF = `
# added only the the node running postgres as 'master'
synchronous_standby_names = slave # standby servers that provide sync rep
                                  # comma-separated list of application_name
                                  # from standby(s); '*' = all
`

	PG_STANDBY_SIGNAL = `#~lx-----------------------------------------------------------------------------
# PG STANDBY SIGNAL CONFIG
#------------------------------------------------------------------------------
# When standby_mode is enabled, the PostgreSQL server will work as a standby. It
# tries to connect to the primary according to the connection settings
# primary_conninfo has move to postgresql.conf
standby_mode = on
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
