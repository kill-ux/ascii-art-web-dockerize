package ascii

import "os"

func VerifyExt(fileName string) bool {
	if len(fileName) > 4 && fileName[len(fileName)-4:] == ".txt" {
		return true
	}
	return false
}

func WriteF(name string, cnt []byte) error {
	file, err := os.Create(name)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.Write(cnt)
	return err
}
