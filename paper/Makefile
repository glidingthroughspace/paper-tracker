sourcefiles := $(wildcard **/*.tex)
sourcefiles += $(wildcard **/*.bib)
sourcefiles += $(wildcard *.tex)
sourcefiles += $(wildcard *.bib)

Studienarbeit.pdf: $(sourcefiles)
	echo $(sourcefiles)
	latexmk -synctex=1 -interaction=nonstopmode -pdflatex=lualatex -lualatex -file-line-error -pdf Studienarbeit.tex
