package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "regexp"
    "strings"
    "time"
    "runtime"
)

func clear() {
   if runtime.GOOS == "windows"{
      c := exec.Command("cls");
      c.Stdout = os.Stdout
      if err := c.Run(); err != nil {
         fmt.Println("Error Clearing Screen");
      }
   } else if runtime.GOOS == "linux" {
      c := exec.Command("clear");
      c.Stdout = os.Stdout
      if err := c.Run(); err != nil {
         fmt.Println("Error Clearing Screen");
      }
   }
}

func printHeader() {
    fmt.Println("  _   _     _")
    fmt.Println(" | | | |   (_)")
    fmt.Println(" | |_| |__  _ _ __   __ _")
    fmt.Println(" | __| '_ \\| | '_ \\ / _` |")
    fmt.Println(" | |_| | | | | | | | (_| |")
    fmt.Println("  \\__|_| |_|_|_| |_|\\__, |")
    fmt.Println("                     __/ |")
    fmt.Println("                    |___/")
    fmt.Println("")
}

func getStudentData(cardData string) {
    // regex pattern to find last/first name
    nameReg, _ := regexp.Compile("\\^([a-zA-Z/]+)")
    name := nameReg.FindStringSubmatch(cardData)[1]
    
    // regex pattern to find student id num
    numReg, _ := regexp.Compile("(?:0{6,8})([0-9]+)")
    stuNum := numReg.FindStringSubmatch(cardData)[1]
    
    // split the name on the "/"
    nameSlice := strings.Split(name, "/")
    
    // assemble the student email
    email := string(nameSlice[1][0]) + nameSlice[0] + stuNum + "@floridapoly.edu"
    email = strings.ToLower(email)
    
    Post(nameSlice[1], nameSlice[0], stuNum, email)
    
    fmt.Println("Swipe for " + nameSlice[1] + " submitted successfully...")
    time.Sleep(time.Second)
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    for {
        printHeader()
        fmt.Println("Please swipe a card...")
        cardData, _ := reader.ReadString('\n')
        getStudentData(string(cardData))
        clear()
    }
}
