#!/bin/bash

# Script for synchronizing bibliographies and other LaTeX files
# that are held in a directory
# "$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"
# across the file system.

# !!!!!!!!!!
# For this reason you should always change ONE files in the directory
# # "$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"


# Path to the directory with bibliographies and other LaTeX files
BIBLIOGRAPHYPATH="$HOME/Good-things/Bibliografia-i-pliki-LaTeXa"
# Path to the main directory with LaTeX files containing errors found
# in books and other works and commentaries to these works.
ERRORSANDCOMMENTSPATH="$HOME/Good-things/Błędy-i-uwagi"
# Path to directory with list of excercises
EXERCISESLISTSPATH="$HOME/Good-things/Listy-zadań"


##############################
# Paths to the various directories in which bibliographies and other LaTeX #
# files need be synchronized.

# Paths are divided in categories due to type of directories to which
# they lead. In one category you have directory with LaTeX files dealing
# with DEUS and philosophy in other with mathematics, etc.
# Every category starts with path of main directory, followed by paths
# to subdirectories in alphabetical order.
# With exception of category "DEUS and philosophy" all other categories
# are in the alphabetical order. Also path to subdirectories are listed
# in alphabetical order.



####################
# DEUS and philosophy paths
DEUSPHILOSOPHYPATH="$ERRORSANDCOMMENTSPATH/DEUS-i-filozofia-błędy-i-uwagi/"
####################



####################
# Computer science paths
COMPUTERSCIENCEPATH="$ERRORSANDCOMMENTSPATH/Pozostałe-dziedziny-błędy-i-uwagi/Informatyka-błędy-i-uwagi"
####################



####################
# Economics paths
ECONOMICSPATH="$ERRORSANDCOMMENTSPATH/Pozostałe-dziedziny-błędy-i-uwagi/Ekonomia-błędy-i-uwagi"
####################



#####################
# History paths
HISTORYPATH="$ERRORSANDCOMMENTSPATH/Historia-błędy-i-uwagi"
####################



####################
# Mathematics paths
MATHEMATICSPATH="$ERRORSANDCOMMENTSPATH/Matematyka-błędy-i-uwagi"


ALGEBRAPATH="$MATHEMATICSPATH/Algebra-błędy-i-uwagi"

DIFFERENTIALEQUATIONSPATH="$MATHEMATICSPATH/Równania-różniczkowe-błędy-i-uwagi"

DIFFERENTIALGEOMETRYPATH="$MATHEMATICSPATH/Geometria-różniczkowa-błędy-i-uwagi"

FUNCTIONALANALYSISPATH="$MATHEMATICSPATH/Analiza-funkcjonalna-błędy-i-uwagi"

LOGICANDSETTHEORYPATH="$MATHEMATICSPATH/Podstawy-matematyki-błędy-i-uwagi/Logika-i-teoria-mnogości-błędy-i-uwagi"

MATHEMATICALANALYSISPATH="$MATHEMATICSPATH/Analiza-matematyczna-błędy-i-uwagi"

NONCOMMUTATIVEGEOMETRYPATH="$MATHEMATICSPATH/Geometria-nieprzemienna-błędy-i-uwagi"

PROBABILTYTHEORYPATH="$MATHEMATICSPATH/Rachunek-prawdopodobieństwa-błędy-i-uwagi"


##########
# Path to mathematical exercises

MATHEXCERCISEPATH="$EXERCISESLISTSPATH/Matematyka-listy-zadań"
####################



####################
# Exercises to do paths
EXERCISESTODOLISTPATH="$HOME/Good-things/Various-writings/Zadania-zrób"
####################



####################
# Physics paths
PHYSICSPATH="$ERRORSANDCOMMENTSPATH/Fizyka-błędy-i-uwagi"


ANALYSISOFEXPERIMENTALDATA="$PHYSICSPATH/Analiza-danych-eksperymentalnych-błędy-i-uwagi"

CLERKMAXWELLELECTRODYNAMICSPATH="$PHYSICSPATH/Elektrodynamika-Clerka-Maxwella-błędy-i-uwagi"

INTRODUCTIONSTOPHYSICSPATH="$PHYSICSPATH/Wprowadzenie-do-fizyki-błędy-i-uwagi"

NEWTONIANMECHANICSPATH="$PHYSICSPATH/Mechanika-Newtona-błędy-i-uwagi"

PHYSICSANDOTHERDISCIPLINESPATH="$PHYSICSPATH/Fizyka-i-inne-dziedziny-błędy-i-uwagi"

PHYSICSBOOKSWITHEXCERCISESPATH="$PHYSICSPATH/Fizyka-zbiory-zadań-błędy-i-uwagi"

PHYSICSOFSPACETIMEPATH="$PHYSICSPATH/Fizyka-czasoprzestrzeni-błędy-i-uwagi"

QFTPATH="$PHYSICSPATH/QFT-błędy-i-uwagi"

QUANTUMMECHANICSPATH="$PHYSICSPATH/Mechanika-kwantowa-błędy-i-uwagi"

TERMOSTATICSETCPATH="$PHYSICSPATH/Termostatyka-termodynamika-i-fizyka-statystyczna-błędy-i-uwagi"
####################




####################
# Various temporary paths

NOTESONEPATH="$HOME/Good-things/Various-writings/Notatki-do-oddziaływań-elektrosłabych"
####################





##############################
# Synchronization of bibliographies and various LaTeX files across
# the file system.
# Outside the cathegory "DEUS and philosophy", all other cathegories
# are in alphabetical order. Also subcategories inside cathegory



####################
# Synchronization of file "DEUSPhilBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/DEUSPhilBooks.bib"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH
####################



####################
# Synchronization of file "HistoryBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/HistoryBooks.bib"


# History directory
rsync $latexfile $HISTORYPATH
####################



####################
# Synchronization of file "MathComScienceBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/MathComScienceBooks.bib"


# Computer science directory
rsync $LATEXFILEPATH $COMPUTERSCIENCEPATH

# Excercises lists
rsync $LATEXFILEPATH $EXERCISELISTSPATH


##########
# Mathematical directories

rsync $LATEXFILEPATH $ALGEBRAPATH

# Math excercises path
rsync $LATEXFILEPATH $MATHEXCERCISEPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH
####################



####################
# Synchronization of file "VariousFieldsBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/VariousFieldsBooks.bib"


# Economics directory
rsync $LATEXFILEPATH $ECONOMICSPATH

# History directory
rsync $LATEXFILEPATH $HISTORYPATH
####################



####################
# Synchronization of file "PhilNaturBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/PhilNaturBooks.bib"


##########
# Mathematical directories

# Functional analysis directory
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH


##########
# Physics directories

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Clerk Maxwell electrodynamics
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines directory
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXCERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH
####################



####################
# Synchronization of file "VariousFieldsBooks.bib"
LATEXFILEPATH="$BIBLIOGRAPHYPATH/VariousFieldsBooks.bib"


# History directory
rsync $LATEXFILEPATH $HISTORYPATH
####################


















# ####################
# # Kopiowanie pliku "LibHistoria.bib"
# latexfile="$biblatexpath/LibHistoria.bib"


# rsync $latexfile $historiapath/


# rsync $latexfile $errorspath/Do-wydawnictwa/





####################
# Kopiowanie pliku "LibPhilNatur.bib"
# latexfile="$biblatexpath/LibPhilNatur.bib"


# rsync $latexfile $physicspath/Fizyka-czasoprzetrzeni-błędy-i-uwagi/
# rsync $latexfile $physicspath/Mechanika-błędy-i-uwagi/
# rsync $latexfile $physicspath/Mechanika-kwantowa-błędy-i-uwagi/
# rsync $latexfile $physicspath/Teoria-pola-błędy-i-uwagi/
# rsync $latexfile $physicspath/Wprowadzenie-do-fizyki-błędy-i-uwagi/


# rsync $latexfile $presentimprovepath
# rsync $latexfile $presentsomepath


# rsync $latexfile $exercisespath


# rsync $latexfile $errorspath/Do-wydawnictwa/





# # #####
# # # Article
# # rsync $latexfile $articlepath






















####################
# Synchronization of file "latexgeneralcommands.sty"
# It contains various general purpose macros.
LATEXFILEPATH="$BIBLIOGRAPHYPATH/latexgeneralcommands.sty"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH

# Computer science directory
rsync $LATEXFILEPATH $COMPUTERSCIENCEPATH

# Economics directory
rsync $LATEXFILEPATH $ECONOMICSPATH

# Excercises lists
rsync $LATEXFILEPATH $EXERCISELISTSPATH

# History directory
rsync $LATEXFILEPATH $HISTORYPATH

# Math excercises lists
rsync $LATEXFILEPATH $MATHEXCERCISEPATH


##########
# Mathematical directories

# Algebra path
rsync $LATEXFILEPATH $ALGEBRAPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

# Differential geometry directory
rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

# Functional analysis directory
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

# Logic and set theory directory
rsync $LATEXFILEPATH $LOGICANDSETTHEORYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

# Noncommutative geometry directory
rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

# Probability theory directory
rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH


##########
# Physics directories

# Analysis of experimental data directory
rsync $LATEXFILEPATH $ANALYSISOFEXPERIMENTALDATA

# Clerk Maxwell electrodynamics
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXCERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH


##########
# Various notes directories

rsync $LATEXFILEPATH $NOTESONEPATH
####################



####################
# Synchronization of file "mathcommands.sty"
# It contains macros for mathematics formulas.
LATEXFILEPATH="$BIBLIOGRAPHYPATH/mathcommands.sty"


# DEUS and philosophy directory
rsync $LATEXFILEPATH $DEUSPHILOSOPHYPATH


##########
# Mathematical directories

# Algebra path
rsync $LATEXFILEPATH $ALGEBRAPATH

# Differential equations directory
rsync $LATEXFILEPATH $DIFFERENTIALEQUATIONSPATH

# Differential geometry directory
rsync $LATEXFILEPATH $DIFFERENTIALGEOMETRYPATH

# Functional analysis path
rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH

# Logic and set theory directory
rsync $LATEXFILEPATH $LOGICANDSETTHEORYPATH

# Mathematical analysis directory
rsync $LATEXFILEPATH $MATHEMATICALANALYSISPATH

# Noncommutative geometry directory
rsync $LATEXFILEPATH $NONCOMMUTATIVEGEOMETRYPATH

# Probability theory directory
rsync $LATEXFILEPATH $PROBABILTYTHEORYPATH


##########
# Physics directories

# Analysis of experimental data directory
rsync $LATEXFILEPATH $ANALYSISOFEXPERIMENTALDATA

# Clerk Maxwell electrodynamics
rsync $LATEXFILEPATH $CLERKMAXWELLELECTRODYNAMICSPATH

# Introductions to physics directory
rsync $LATEXFILEPATH $INTRODUCTIONSTOPHYSICSPATH

# Newtonian mechanics directory
rsync $LATEXFILEPATH $NEWTONIANMECHANICSPATH

# Physics and other disciplines
rsync $LATEXFILEPATH $PHYSICSANDOTHERDISCIPLINESPATH

# Books with physics excercise directory
rsync $LATEXFILEPATH $PHYSICSBOOKSWITHEXCERCISESPATH

# Physics of spacetime directory
rsync $LATEXFILEPATH $PHYSICSOFSPACETIMEPATH

# Quantum mechanics directory
rsync $LATEXFILEPATH $QUANTUMMECHANICSPATH

# QFT directory
rsync $LATEXFILEPATH $QFTPATH

# Thermostatic, thermodynamics and statistical physics
rsync $LATEXFILEPATH $TERMOSTATICSETCPATH


##########
# Various notes directories

rsync $LATEXFILEPATH $NOTESONEPATH
####################




####################
# Synchronization of file "functionalanalysiscommands.sty"
# It contains various macros for symbols that often appeare in the context
# of functional analysis
LATEXFILEPATH="$BIBLIOGRAPHYPATH/functionalanalysiscommands.sty"


rsync $LATEXFILEPATH $FUNCTIONALANALYSISPATH









# ##########
