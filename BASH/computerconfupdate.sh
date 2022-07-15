#!/bin/bash

# Script that set ups other scripts and copy configuration files from
# backups directory

SCRIPTSPATH="$HOME/bin/scripts"


chmod u+x $SCRIPTSPATH/makesymbolicslinks.sh
$SCRIPTSPATH/makesymbolicslinks.sh

chmod u+x $SCRIPTSPATH/confcopyfrombackups.sh

$SCRIPTSPATH/confcopyfrombackups.sh
