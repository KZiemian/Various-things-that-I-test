#!/bin/bash

# Plik do uaktualniania bibligorafii w obecnym seminarium

bibliographypath="$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"

presentimprovepath="$HOME/Good-things/Prezentacje/Prezentacje-do-poprawy"
presentsomepath="$HOME/Good-things/Prezentacje/Prezentacje-jako-tako"

currentpresentationpath="$presentimprovepath/Promieniowanie-Hawkinga-czarnej-dziury-powstałej-w-wyniku-sferycznie-symetrycznego-kolapsu/"


# Plik ArticPhilNatur.bib
articphilnaturpath="$bibliographypath/ArticPhilNatur.bib"

rsync $articphilnaturpath $currentpresentationpath
