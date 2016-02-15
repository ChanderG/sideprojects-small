# Resume from markdown

Inspired from [https://github.com/mszep/pandoc_resumer].
See also: [https://mszep.github.io/pandoc_resume/]

### Running
Setup dev environment.

#### Dev environment
You need `pando`c and `latex`(fully). 
Use include `Dockerfile` to build your own image with required software.

#### Creating pdf via context
Use the following commands:

```
pandoc --standalone --template style.tex --from markdown --to context -V papersize=A4 -o resume.tex resume.md
```

```
context resume.tex
```

to use the `style.tex` to style the resume.
