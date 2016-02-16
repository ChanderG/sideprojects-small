#! /bin/bash

# assuming that resume.md is the input

pandoc --standalone --template style.tex --from markdown --to context -V papersize=A4 -o resume.tex resume.md; context resume.tex
