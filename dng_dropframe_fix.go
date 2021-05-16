package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	currentFrame := 0

	for _, f := range files {

		val := f.Name()

		// Check if DNG
		filetype_check := val[len(val)-4:]
		if filetype_check != ".dng" {
			continue
		}

		//
		newVal := val[len(val)-10:]
		prefix := val[:len(val)-10]
		// fmt.Println(prefix)
		frameNumberasString := strings.TrimSuffix(newVal, ".dng")
		fmt.Println(frameNumberasString)
		frameNumberasInt, err := strconv.Atoi(frameNumberasString)
		if err == nil {
			if frameNumberasInt == currentFrame {
				fmt.Println(strconv.Itoa(currentFrame) + " is present.")
			} else {
				fmt.Println(strconv.Itoa(currentFrame) + " is not present. Creating!")
				//Read all the contents of the  original file
				bytesRead, err := ioutil.ReadFile(val)
				if err != nil {
					log.Fatal(err)
				}
				currentFramePadded := fmt.Sprintf("%06d", currentFrame)
				newframe := prefix + currentFramePadded + ".dng"
				fmt.Println(prefix + currentFramePadded + ".dng created!")
				//Copy all the contents to the desitination file
				err = ioutil.WriteFile(newframe, bytesRead, 0755)
				if err != nil {
					log.Fatal(err)
				}
				currentFrame++
			}
			currentFrame++
		} else {

		}

	}
}
