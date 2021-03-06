#!/bin/bash


JAGIELLONIANTHEMEPATH="$HOME/Good-things/jagiellonian-theme/source"
JAGIELLONIANPRESENTATIONPATH="$HOME/Good-things/Presentation_in_Jagiellonian/"
JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH="$JAGIELLONIANPRESENTATIONPATH/preambule/JagiellonianCustomizationCommands.tex"
JAGIELLONIANCUSTOMIZATIONGENERALPATH="$JAGIELLONIANPRESENTATIONPATH/preambule/JagiellonianCustomizationGeneral.tex"
JAGIELLONIANLANGUAGESETTINGSPATH="$JAGIELLONIANPRESENTATIONPATH/preambule/LanguageSettings/JagiellonianPolishLanguageSettings.tex"
JAGIELLONIANTEXTPOSPATH="$JAGIELLONIANPRESENTATIONPATH/preambule/TextposConfiguration/TextposConfiguration.tex"


wykladygrafika3Dpath="$HOME/Good-things/Work/Grafika3D-projekt/Wykłady-Grafika3D-PL"
fizykwykladypath="$HOME/Good-things/Work/SymulacjeFizyki-projekt/Wyklady-SymulacjeFizyki-PL"
wykladywszystkiepath="$wykladygrafika3Dpath/Wykłady-wszystkie"
PRESENTATIONSNEEDSIMPROVEMENTSPATH="$HOME/Good-things/Prezentacje/Prezentacje-do-poprawy"



cd $JAGIELLONIANTHEMEPATH
latex beamerthemejagiellonian.ins


####################
# Synchronization of Jagiellonian theme files in POWR project
# rsync $JAGIELLONIANTHEMEPATH/*.sty $wykladygrafika3Dpath
# rsync $JAGIELLONIANTHEMEPATH/*.sty $wykladygrafika3Dpath/Wyklady-kompilacja/
# rsync $JAGIELLONIANTHEMEPATH/*.sty $wykladygrafika3Dpath/Slajdy-w-obróbce/

# rsync $JAGIELLONIANTHEMEPATH/*.sty $fizykwykladypath
# rsync $JAGIELLONIANTHEMEPATH/*.sty $fizykwykladypath/Wyklady-kompilacja/
# rsync $JAGIELLONIANTHEMEPATH/*.sty $fizykwykladypath/Slajdy-w-obróbce/

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


####################
# Synchronization of Jagiellonian theme files presetations
rsync $JAGIELLONIANTHEMEPATH/*sty \
      $PRESENTATIONSNEEDSIMPROVEMENTSPATH/jagiellonian-theme-files/


# Path to "Alistera McGratha krytyka memetyki" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Alistera-McGratha-krytyka-memetyki"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Computational condensed matter physics with Julia" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Computational-condensed-matter-physics-Julia"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Czemu rozważanie cząstek w zakrzywionej czasoprzestrzeni
# może być błędem?" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Czemu-rozważanie-cząstek-w-zakrzywionej-czasoprzestrzeni-może-być-błędem"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Filozofia współczesna. Wprowadzenie" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Filozofia-współczesna-Wprowadzenie"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Filozofia i matematyczne wyzwania nieralatywitycznej mechaniki
# kwantowej" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Filozoficzne-i-matematyczne-wyzwania-nierelatywistycznej-mechaniki-kwantowej"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Grafen, czyli dynamika relatywistyczna przy trochę mniejszych
# prędkościach" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Grafen-czyli-dynamika-relatywistyczna-przy"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Julia: 2010s proposition for scientific (and other) programming"
# presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Herdegens-algebraic-approach-to-Casimir-effect"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Julia: 2010s proposition for scientific (and other) programming"
# presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Julia-2010s-proposition"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Matematyczna strona mechaniki kwantowej" presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Matematyczna-strona-mechaniki-kwantowej"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Wstęp do LaTeXa. Kilka uwag o używaniu LaTeXa"
# presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Kilka-uwag-o-używaniu-LaTeXa"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/



# Path to "Podejście Herdegena do efektu Casimira-I"
# presentation
PRESENTATIONPATH="$PRESENTATIONSNEEDSIMPROVEMENTSPATH/Podejście-Herdegena-do-efektu-Casimira-I"

rsync $JAGIELLONIANTHEMEPATH/*sty $PRESENTATIONPATH
rsync $JAGIELLONIANCUSTOMIZATIONCOMMANDSPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANCUSTOMIZATIONGENERALPATH $PRESENTATIONPATH/preambule/
rsync $JAGIELLONIANLANGUAGESETTINGSPATH \
      $PRESENTATIONPATH/preambule/LanguageSettings/
rsync $JAGIELLONIANTEXTPOSPATH \
      $PRESENTATIONPATH/preambule/TextposConfiguration/





####################
# Removing files that aren't needed anymore

rm $JAGIELLONIANTHEMEPATH/beamer*jagiellonian.sty
rm $JAGIELLONIANTHEMEPATH/beamercolorthemejagiellonian-highcontrast.sty
rm $JAGIELLONIANTHEMEPATH/pgfplotsthemetol.sty
