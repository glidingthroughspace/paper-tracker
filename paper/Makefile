target = Studienarbeit.tex
latexmk_options = -synctex=1 -interaction=nonstopmode -pdflatex=lualatex -lualatex -file-line-error -pdf $(target)

Studienarbeit.pdf:
	latexmk $(latexmk_options)

watch:
	latexmk -pvc $(latexmk_options)

run: watch
	# Alias for watch target

clean:
	latexmk -c

.PHONY: Studienarbeit.pdf
