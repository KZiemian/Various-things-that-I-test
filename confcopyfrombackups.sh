#!/bin/bash

# Copy files from directory with backups and copy them to the places in file
# where they can fullfill they role.

CONFBACKUPSPATH="$HOME/Good-things/configuration"
emacsbackupspath="$confbackupspath/Emacs-configuration"


scriptspath="$HOME/bin/scripts"
emacspath="$HOME/.emacs.d"
emacslisppath="$HOME/.emacs.d/elisp"
orgpath="$HOME/Good-things/Keep-at-hand/org"





rsync $CONFBACKUPSPATH/.bashrc $CONFBACKUPSPATH/.bash_aliases $HOME

rsync $CONFBACKUPSPATH/latexfilesupdate.sh $CONFBACKUPSPATH/koans.sh \
      $CONFBACKUPSPATH/confbackupdate.sh \
      $CONFBACKUPSPATH/confreserbackup.sh \
      $CONFBACKUPSPATH/linkuj.sh \
      $CONFBACKUPSPATH/jagiellonianupdate.sh \
      $CONFBACKUPSPATH/grafika3Dwyklady.sh \
      $CONFBACKUPSPATH/grafika3Dkopiuj.sh \
      $CONFBACKUPSPATH/grafika3Dlightwyklady.sh \
      $CONFBACKUPSPATH/grafika3Ddarkwyklady.sh \
      $CONFBACKUPSPATH/symulacjefizykiwyklady.sh \
      $CONFBACKUPSPATH/symulacjefizykilightwyklady.sh \
      $CONFBACKUPSPATH/symulacjefizykidarkwyklady.sh \
      $CONFBACKUPSPATH/symulacjefizykikopiuj.sh \
      $CONFBACKUPSPATH/dooddania.py $CONFBACKUPSPATH/kugazpieniadze.py \
      $CONFBACKUPSPATH/zwrociles.py $CONFBACKUPSPATH/prezentypieniadze.py \
      $CONFBACKUPSPATH/rafaloplaty.py \
      $scriptspath

rsync $emacsbackupspath/init.el $emacsbackupspath/config.org \
      $emacsbackupspath/configadditional.org \
      $emacsbackupspath/configovpowered.org $emacsbackupspath/customize.el \
      $emacsbackupspath/init_1.el $emacsbackupspath/init_2.el \
      $emacspath

rsync $confbackupspath/Emacs-bookmarks/bookmarks $emacspath

rsync $confbackupspath/ksiazki.org $confbackupspath/prezentypieniadze.org \
      $confbackupspath/todo.org $confbackupspath/z-roznych-zeszytow.org \
      $orgpath

rsync $emacsbackupspath/additional-functions.el $emacslisppath


rsync $confbackupspath/confupdate.sh $confbackupspath/startscript.sh \
      $confbackupspath/endscript.sh \
      $scriptspath
