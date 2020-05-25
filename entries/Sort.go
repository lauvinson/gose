package entries

import (
	"bufio"
	"fmt"
	"gose/utils/algorithm"
	"os"
	"strconv"
)

func Run(al string, inPath string, outPath string) (err error) {
	file, err := os.Open(inPath)
	if err != nil {
		fmt.Println("failed to open the file:", inPath)
		return
	}
	defer file.Close()

	r := bufio.NewReader(file)
	values := make([]int, 0)
	for {
		bytes, prefix, err := r.ReadLine()
		if prefix || nil != err{
			break
		}
		str := string(bytes)
		value, err := strconv.Atoi(str)
		if nil != err {
			return err
		}
		values = append(values, value)
	}
	if al == "quick" {
		sorted, err := algorithm.Sort(values)
		if nil != err {
			return err
		}

		ofile, err := os.Create(outPath)
		if nil != err {
			fmt.Println("failed to create the file:", outPath)
			return err
		}

		defer ofile.Close()
		for _, v := range sorted {
			_, err = ofile.WriteString(strconv.Itoa(v) + "\n")
			if nil != err {
				return err
			}
		}
	}
	return
}
