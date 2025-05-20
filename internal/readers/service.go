package readers

import "fmt"

func GetReadersService() ([]Reader, error) {
	return GetAllReaders()
}

func GetReaderService(readerId int) (Reader, error) {
	return GetReaderById(readerId)
}

func UpdateReaderService(readerId int, fullName string) error {
	return ChangeReaderById(readerId, fullName)
}

func AddReaderService(reader *Reader) error {

	err := AddReader(reader)
	if err != nil {
		fmt.Println(err)
	}
	return err
}
