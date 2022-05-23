#!/bin/bash


jagiellonianpath="$HOME/Good-things/Work/jagiellonian"
wykladygrafika3Dpath="$HOME/Good-things/Work/Grafika3D-projekt/Wykłady-Grafika3D-PL"
fizykwykladypath="$HOME/Good-things/Work/SymulacjeFizyki-projekt/Wyklady-SymulacjeFizyki-PL"
wykladywszystkiepath="$wykladygrafika3Dpath/Wykłady-wszystkie"
presentationpath="$HOME/Good-things/Prezentacje/Prezentacje-do-poprawy"



cd $jagiellonianpath
latex beamerthemejagiellonian.ins


####################
# Copying files to directory with exercise to do
rsync $jagiellonianpath/*.sty $wykladygrafika3Dpath
rsync $jagiellonianpath/*.sty $wykladygrafika3Dpath/Wyklady-kompilacja/
rsync $jagiellonianpath/*.sty $wykladygrafika3Dpath/Slajdy-w-obróbce/

rsync $jagiellonianpath/*.sty $fizykwykladypath
rsync $jagiellonianpath/*.sty $fizykwykladypath/Wyklady-kompilacja/
rsync $jagiellonianpath/*.sty $fizykwykladypath/Slajdy-w-obróbce/

# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-01/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-02/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-03/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-04/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-05/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-06/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-07/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-08/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-09/
# rsync $jagiellonianpath/*sty $wykladywszystkiepath/Wykład-10/
# rsync $jagiellonianpath/*sty $presentationpath

rsync $jagiellonianpath/*sty $presentationpath/jagiellonian/


rsync $jagiellonianpath/*sty $presentationpath/Julia-2010s-proposition/



rm $jagiellonianpath/beamer*jagiellonian.sty
rm $jagiellonianpath/beamercolorthemejagiellonian-highcontrast.sty
rm $jagiellonianpath/pgfplotsthemetol.sty
