#!/bin/bash


lecturespath="$HOME/Good-things/Work/Geometria3D-projekt/Wykłady-Geometria3D-PL"
latexcommand="pdflatex -file-line-error -interaction=nonstopmode"



cd $lecturespath/Wykłady-kompilacja/




$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_01.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_02.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_03.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_04.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_05.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_06.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_07.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_08.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_09.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_10.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_11.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_12.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_13.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_14.tex}"
$latexcommand "\def\PresentationStyle{light}\input{Wyklad_Geometria3D_15.tex}"



mv Wyklad_Geometria3D_01.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_02.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_03.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_04.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_05.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_06.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_07.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_08.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_09.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_10.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_11.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_12.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_13.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_14.pdf ../Wykłady-PDF/light/
mv Wyklad_Geometria3D_15.pdf ../Wykłady-PDF/light/
