#!/bin/bash

# Tworzy rezerwowe backupy

basicchangepath="$HOME/Good-things/Basic-changing"
confpath="$HOME/Good-things/configuration"
emacspath="$HOME/.emacs.d"
emacslisppath="$HOME/.emacs.d/elisp"
orgpath="$HOME/Good-things/Keep-at-hand/org"
scriptspath="$HOME/bin/scripts"


confreservepath="$confpath/Reserve-config"
emacsreservepath="$confpath/Reserve-config/Emacs-reserve-config"





# Reserve backups
rsync $HOME/.bashrc $HOME/.bash_aliases \
      $scriptspath/linkuj.sh \
      $scriptspath/latexfilesupdate.sh \
      $scriptspath/confupdate.sh $scriptspath/confbackupdate.sh \
      $scriptspath/confreserbackup.sh \
      $scriptspath/startscript.sh $scriptspath/endscript.sh \
      $scriptspath/koans.sh \
      $scriptspath/jagiellonianupdate.sh \
      $scriptspath/grafika3Dwyklady.sh $scriptspath/grafika3Dkopiuj.sh \
      $scriptspath/grafika3Dlightwyklady.sh \
      $scriptspath/grafika3Ddarkwyklady.sh \
      $scriptspath/symulacjefizykiwyklady.sh \
      $scriptspath/symulacjefizykilightwyklady.sh \
      $scriptspath/symulacjefizykilightwyklady.sh \
      $scriptspath/symulacjefizykikopiuj.sh \
      $scriptspath/dooddania.py $scriptspath/kugazpieniadze.py \
      $scriptspath/zwrociles.py $scriptspath/prezentypieniadze.py \
      $scriptspath/rafaloplaty.py \
      $confreservepath





# Emacs
rsync $emacspath/init.el $emacspath/config.org \
      $emacspath/configadditional.org $emacspath/configovpowered.org \
      $emacspath/customize.el $emacslisppath/additional-functions.el \
      $emacspath/init_1.el $emacspath/init_2.el \
      $emacsreservepath





# Emacs bookmarks
rsync $emacspath/bookmarks $confreservepath/Emacs-reserve-bookmarks





# Pliki ORG
rsync $orgpath/ksiazki.org $orgpath/prezentypieniadze.org \
      $orgpath/todo.org $orgpath/z-roznych-zeszytow.org \
      $confreservepath
