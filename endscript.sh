#!/bin/bash

# Ostatni skrypt przed kompiowaniem plików

scriptspath="$HOME/bin/scripts"



chmod u+x $scriptspath/confbackupdate.sh

chmod u+x $scriptspath/latexfilesupdate.sh

$scriptspath/confbackupdate.sh
$scriptspath/latexfilesupdate.sh
