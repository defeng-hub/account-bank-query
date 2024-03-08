package utils

import "log"

func WordReplace(search, replace string, filepath, outpath string) (err error) {
	r, err := ReadDocxFile(filepath)
	if err != nil {
		log.Fatalf("error reading docx file: %s", err)
		return err
	}
	defer r.Close()

	doc := r.Editable()

	// Replace text
	doc.Replace(search, replace, -1)

	// Save the modified document
	err = doc.WriteToFile(outpath)
	if err != nil {
		log.Fatalf("error writing docx file: %s", err)
		return err
	}
	return nil
}
