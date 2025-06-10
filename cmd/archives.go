package cmd

import (
	"archive/tar"
	"archive/zip"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

func Unarchive(dir, archive string) {
	mtype, err := mimetype.DetectFile(archive)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(mtype.String(), mtype.Extension())

	switch mtype.String() {
	case "application/gzip":
		extractTarGz(dir, archive)
	case "application/zip":
		log.Fatal("It is a .zip file")
		extractZip(dir, archive)
	default:
		fmt.Printf("Type: %s\n", mtype.String())
		log.Fatal("Unrecognized archive type.")
	}
}

func extractTarGz(dir, archive string) {
	gzipStream, err := os.Open(archive)
	if err != nil {
		log.Fatalf("Unable to open archive: %v", err)
	}
	defer gzipStream.Close()

	uncompressedStream, err := gzip.NewReader(gzipStream)
	if err != nil {
		log.Fatalf("Failed to create GZIP reader: %v", err)
	}
	defer uncompressedStream.Close()

	tarReader := tar.NewReader(uncompressedStream)

	for {
		header, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Next() failed: %v", err)
		}

		relPath := strings.TrimPrefix(header.Name, "./")
		headerPath := filepath.Join(dir, relPath)

		switch header.Typeflag {
		case tar.TypeDir:
			if err := os.MkdirAll(headerPath, 0755); err != nil {
				log.Fatalf("Failed to create directory %s: %v", headerPath, err)
			}
		case tar.TypeReg:
			if err := os.MkdirAll(filepath.Dir(headerPath), 0755); err != nil {
				log.Fatalf("Failed to create parent directory for file %s: %v", headerPath, err)
			}

			outFile, err := os.Create(headerPath)
			if err != nil {
				log.Fatalf("Failed to create file %s: %v", headerPath, err)
			}
			if _, err := io.Copy(outFile, tarReader); err != nil {
				outFile.Close()
				log.Fatalf("Failed to write file %s: %v", headerPath, err)
			}
			outFile.Close()
		default:
			log.Printf("Skipping unknown type: %c in %s", header.Typeflag, header.Name)
		}
	}
}

func extractZip(dir, archive string) {
	r, err := zip.OpenReader(archive)
	if err != nil {
		log.Fatal("Could not open zip file: ", err)
	}
	defer r.Close()

	for _, f := range r.File {
		fmt.Printf("Contents of %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}
	}
}
