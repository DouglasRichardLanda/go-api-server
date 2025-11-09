package main

import (
	"fmt"
	"io"
	"net/http"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"os"
	"path/filepath"
	"strings"
)

func PDFHandlerFunctionMerge(res http.ResponseWriter, req *http.Request) {
	// parsing form data
	if err := req.ParseMultipartForm(50 << 20); err != nil {
		http.Error(res, "Cannot parse form", http.StatusBadRequest)
		return
	}

	// I put all of them here
	pdf_files := req.MultipartForm.File["pdfFiles"]

	// may be unnecessary but still
	type PDFFile struct {
		Name string
		Size int64
		File io.ReadCloser
	}

	var pdfFiles []PDFFile
	var tempFiles []string

	// detailed read out
	for _, single_pdf := range pdf_files {
		file, err := single_pdf.Open()
		if err != nil {
			http.Error(res, "Cannot open file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		tmpName := "/tmp/" + single_pdf.Filename
		out, err := os.Create(tmpName)
		if err != nil {
			http.Error(res, "Cannot create temp file", http.StatusInternalServerError)
			return
		}

		if _, err := io.Copy(out, file); err != nil {
			out.Close()
			http.Error(res, "Cannot save file", http.StatusInternalServerError)
			return
		}
		out.Close()

		pdfFiles = append(pdfFiles, PDFFile{
			Name: single_pdf.Filename,
			Size: single_pdf.Size,
			File: file,
		})

		tempFiles = append(tempFiles, tmpName)
	}

	// I want that the new pdf has names of all pdfs
	var names []string
	for _, f := range pdfFiles {
		base := strings.TrimSuffix(f.Name, filepath.Ext(f.Name)) // remove .pdf
		names = append(names, base)
	}
	mergedName := strings.Join(names, "_AND_")
	mergedFile := fmt.Sprintf("/tmp/%s.pdf", mergedName)

	// Merging PDFs using pdfcpu
	if err := api.MergeCreateFile(tempFiles, mergedFile, false, nil); err != nil {
		http.Error(res, "Cannot merge PDFs", http.StatusInternalServerError)
		return
	}

	// Send merged PDF to client
	merged, err := os.Open(mergedFile)
	if err != nil {
		http.Error(res, "Cannot open merged file", http.StatusInternalServerError)
		return
	}
	defer merged.Close()

	res.Header().Set("Content-Type", "application/pdf")
	res.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.pdf\"", mergedName))
	if _, err := io.Copy(res, merged); err != nil {
		http.Error(res, "Cannot send merged PDF", http.StatusInternalServerError)
		return
	}
}

