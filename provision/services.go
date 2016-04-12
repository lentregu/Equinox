package main

import (
	"fmt"
	"regexp"

	"github.com/lentregu/Equinox/oxford"
)

type printOption int

const (
	pretty printOption = iota
	normal
)

type wordOption int

const (
	oneWord = iota
	multipleWords
)

//const oneWordRegExp ="^\S.*\S$"
const oneWordRegExp = `^[^\t\n\f\r ]*$`

const multipleWordsRegExp = `^.*$`

var faceService = oxford.NewFace("567c560aa85245418459b82634bc7a98")


func addFace() (string, error) {

	faceListID, err := readString("FaceList ID", oneWordRegExp)
	imageFileName, err := readString("Face", oneWordRegExp)
	if err != nil {
		return "", err
	}
	//return faceService.AddFaceURL(faceListID, imageFileName)
	return faceService.AddFace(faceListID, imageFileName)
}

func getFacesInAList() (string, error) {

	faceListID, _ := readString("FaceList ID", oneWordRegExp)
	return faceService.GetFacesInAList(faceListID)
}

func getFaceList() (string, error) {

	return faceService.GetFaceList()
}

func createFaceList() (string, error) {

	faceListID, err := readString("FaceList Name", oneWordRegExp)

	if err != nil {
		return "", err
	}

	return faceService.CreateFaceList(faceListID)
}

func readString(name string, wordRegExp string) (string, error) {
	var value string
	fmt.Print(name + ": ")

	validExpression := regexp.MustCompile(wordRegExp)

	line, _, err := stdin.ReadLine()
	if err != nil {
		err = fmt.Errorf("Error reading value for %s: %s", name, err.Error())
	} else {

		value = fmt.Sprintf("%s", line)

		if !validExpression.MatchString(value) && wordRegExp == oneWordRegExp {
			err = fmt.Errorf("ERROR Not spaces are allowed for %s field\n", name)
		}
	}

	return value, err
}
