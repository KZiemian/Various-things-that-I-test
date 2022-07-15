#!/bin/bash

# Linkuje pliki do folderu ~/bin

BINPATH="$HOME/bin"
SCRIPTSPATH="$HOME/bin/scripts"
scriptspath="$HOME/bin/scripts"




cd $HOME/bin

rm -f confupdate confbackup latexfilesupdate Pokuta koans Python-koans \
   Swiety-Czas dooddania zwrociles



#####
cd $HOME/bin/scripts

# Skrypty w bashu
chmod u+x $SCRIPTSPATH/*

# chmod u+x Podlaczeni.sh
# chmod u+x makebackupfiles.sh
# chmod u+x MojSystem.sh
# chmod u+x Convert.py
# chmod u+x Pokuta.py
# chmod u+x koans.sh
# chmod u+x Python-koans.sh
# chmod u+x Swiety-Czas.py
# # chmod u+x dropbox.py

# # Skrypty w Pythonie 3.x
# chmod u+x dooddania.py zwrociles.py





#####
# Skrypty w bashu
ln -s $SCRIPTSPATH/startscript.sh $BINPATH/startscript
ln -s $SCRIPTSPATH/endscript.sh $BINPATH/endscript


ln -s $SCRIPTSPATH/latexfilesupdate.sh $BINPATH/latexfilesupdate
ln -s $SCRIPTSPATH/confupdate.sh $BINPATH/confupdate
ln -s $SCRIPTSPATH/confbackupdate.sh $BINPATH/confbackupdate
ln -s $SCRIPTSPATH/confreserbackup.sh $BINPATH/confreserbackup


ln -s $SCRIPTSPATH/Pokuta.py $BINPATH/Pokuta
ln -s $SCRIPTSPATH/koans.sh $BINPATH/koans
ln -s $scriptspath/Python-koans.sh $BINPATH/python-koans
ln -s $scriptspath/Swiety-Czas.py $BINPATH/Swiety-Czas


ln -s $scriptspath/jagiellonianthemeupdate.sh $BINPATH/jagiellonianthemeupdate
ln -s $scriptspath/geometria3Dwyklady.sh $BINPATH/geometria3Dwyklady
ln -s $scriptspath/geometria3Dlightwyklady.sh $BINPATH/geometria3Dlightwyklady
ln -s $scriptspath/geometria3Ddarkwyklady.sh $BINPATH/geometria3Ddarkwyklady
ln -s $scriptspath/geometria3Dkopiuj.sh $BINPATH/geometria3Dkopiuj
ln -s $scriptspath/symulacjefizykiwyklady.sh \
   $BINPATH/symulacjefizykiwyklady
ln -s $scriptspath/symulacjefizykilightwyklady.sh \
   $BINPATH/symulacjefizykilightwyklady
ln -s $scriptspath/symulacjefizykidarkwyklady.sh \
   $BINPATH/symulacjefizykidarkwyklady
ln -s $scriptspath/symulacjefizykikopiuj.sh $BINPATH/symulacjefizykikopiuj


ln -s $scriptspath/dropbox.py $BINPATH/dropbox

# Skrypty w Pythonie 3.x
ln -s $scriptspath/dooddania.py $BINPATH/dooddania
ln -s $scriptspath/rafaloplaty.py $BINPATH/rafaloplaty
ln -s $scriptspath/zwrociles.py $BINPATH/zwrociles

ln -s $scriptspath/prezentypieniadze.py $BINPATH/prezentypieniadze
