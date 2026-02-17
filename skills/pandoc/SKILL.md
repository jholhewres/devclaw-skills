---
name: pandoc
description: "Document conversion between formats with Pandoc"
---
# Pandoc

Use the **bash** tool with pandoc for document format conversion.

## Common Conversions
```bash
pandoc input.md -o output.pdf
pandoc input.md -o output.docx
pandoc input.md -o output.html --standalone
pandoc input.docx -o output.md
pandoc input.html -o output.md
```

## With Styling
```bash
pandoc input.md -o output.pdf --template=template.tex
pandoc input.md -o output.html --css=style.css --standalone
pandoc input.md -o output.pdf -V geometry:margin=1in
```

## Batch
```bash
for f in *.md; do pandoc "$f" -o "${f%.md}.pdf"; done
```

## Metadata
```bash
pandoc input.md -o output.pdf --metadata title="Title" --metadata author="Author"
```

## Tips
- Use --standalone (-s) for complete HTML documents
- Use --toc for automatic table of contents
- PDF output requires LaTeX (texlive) or use --pdf-engine=wkhtmltopdf
- Use read_file to check input, write_file to save templates
