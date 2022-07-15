#!/bin/bash


# Eksportuje konfiguracje Linuxa do odpowiedniego katalogu

emacspath="$HOME/.emacs.d"
emacslisppath="$HOME/.emacs.d/elisp"
orgpath="$HOME/Good-things/Keep-at-hand/org"
scriptspath="$HOME/bin/scripts"

confbackupspath="$HOME/Good-things/configuration"
emacsbackupspath="$confbackupspath/Emacs-configuration"





rsync $HOME/.bashrc $HOME/.bash_aliases \
      $scriptspath/linkuj.sh \
      $scriptspath/latexfilesupdate.sh \
      $scriptspath/confupdate.sh $scriptspath/confbackupdate.sh \
      $scriptspath/confreserbackup.sh \
      $scriptspath/startscript.sh $scriptspath/endscript.sh \
      $scriptspath/koans.sh \
      $scriptspath/jagiellonianupdate.sh \
      $scriptspath/geometria3Dwyklady.sh $scriptspath/geometria3Dkopiuj.sh \
      $scriptspath/geometria3Dlightwyklady.sh \
      $scriptspath/geometria3Ddarkwyklady.sh \
      $scriptspath/symulacjefizykiwyklady.sh \
      $scriptspath/symulacjefizykikopiuj.sh \
      $scriptspath/symulacjefizykilightwyklady.sh \
      $scriptspath/symulacjefizykidarkwyklady.sh \
      $scriptspath/dooddania.py $scriptspath/kugazpieniadze.py \
      $scriptspath/zwrociles.py $scriptspath/prezentypieniadze.py \
      $scriptspath/rafaloplaty.py \
      $confbackupspath



touch $emacspath/init_1.el $emacspath/init_2.el



rsync $emacspath/init.el $emacspath/config.org \
      $emacspath/configadditional.org $emacspath/configovpowered.org \
      $emacspath/customize.el $emacslisppath/additional-functions.el \
      $emacspath/init_1.el $emacspath/init_2.el \
      $emacsbackupspath

rsync $emacspath/bookmarks $confbackupspath/Emacs-bookmarks

rsync $orgpath/ksiazki.org $orgpath/prezentypieniadze.org \
      $orgpath/todo.org $orgpath/z-roznych-zeszytow.org \
      $confbackupspath
