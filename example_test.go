package lastfm_test

import (
        "fmt"
        "github.com/PuloV/ExcelReader"
)

// Create new api client object
// with your api key as argument.
func ExampleExcelReaderNew() {
        reader , error := ExcelReader.new("test.xls")
        if error != nil {
                fmt.Println(error)
        }
        fmt.Printf("File open without problem !")

        // File open without problem !
}
