#!/bin/bash

# Linkuje pliki do folderu ~/bin

binpath="$HOME/bin"
scriptspath="$HOME/bin/scripts"




cd $HOME/bin

rm -f confupdate confbackup latexfilesupdate Pokuta koans Python-koans \
   Swiety-Czas dooddania zwrociles



#####
cd $HOME/bin/scripts

# Skrypty w bashu
chmod u+x $scriptspath/*

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
ln -s $scriptspath/startscript.sh $binpath/startscript
ln -s $scriptspath/endscript.sh $binpath/endscript


ln -s $scriptspath/latexfilesupdate.sh $binpath/latexfilesupdate
ln -s $scriptspath/confupdate.sh $binpath/confupdate
ln -s $scriptspath/confbackupdate.sh $binpath/confbackupdate
ln -s $scriptspath/confreserbackup.sh $binpath/confreserbackup


ln -s $scriptspath/Pokuta.py $binpath/Pokuta
ln -s $scriptspath/koans.sh $binpath/koans
ln -s $scriptspath/Python-koans.sh $binpath/python-koans
ln -s $scriptspath/Swiety-Czas.py $binpath/Swiety-Czas


ln -s $scriptspath/jagiellonianupdate.sh $binpath/jagiellonianupdate
ln -s $scriptspath/geometria3Dwyklady.sh $binpath/geometria3Dwyklady
ln -s $scriptspath/geometria3Dlightwyklady.sh $binpath/geometria3Dlightwyklady
ln -s $scriptspath/geometria3Ddarkwyklady.sh $binpath/geometria3Ddarkwyklady
ln -s $scriptspath/geometria3Dkopiuj.sh $binpath/geometria3Dkopiuj
ln -s $scriptspath/symulacjefizykiwyklady.sh \
   $binpath/symulacjefizykiwyklady
ln -s $scriptspath/symulacjefizykilightwyklady.sh \
   $binpath/symulacjefizykilightwyklady
ln -s $scriptspath/symulacjefizykidarkwyklady.sh \
   $binpath/symulacjefizykidarkwyklady
ln -s $scriptspath/symulacjefizykikopiuj.sh $binpath/symulacjefizykikopiuj


ln -s $scriptspath/dropbox.py $binpath/dropbox

# Skrypty w Pythonie 3.x
ln -s $scriptspath/dooddania.py $binpath/dooddania
ln -s $scriptspath/rafaloplaty.py $binpath/rafaloplaty
ln -s $scriptspath/zwrociles.py $binpath/zwrociles

ln -s $scriptspath/prezentypieniadze.py $binpath/prezentypieniadze
