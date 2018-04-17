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
    "github.com/qwerty0981/FormAssistant"
)

func clear() {
   if runtime.GOOS == "windows"{
      c := exec.Command("cls");
      c.Stdout = os.Stdout
      if err := c.Run(); err != nil {
         fmt.Println("Error Clearing Screen");
      }
   } else if runtime.GOOS == "linux" || runtime.GOOS == "darwin" {
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

func createEmail(fname, lname, num string) string {
    return strings.ToLower(string(fname[0]) + lname + num + "@floridapoly.edu")
}

func getStudentData(cardData string) (map[string]string) {
    values := make(map[string]string) 
   // regex pattern to find last/first name
    nameReg, _ := regexp.Compile("\\^([a-zA-Z/]+)")
    name := nameReg.FindStringSubmatch(cardData)[1]

    // regex pattern to find student id num
    numReg, _ := regexp.Compile("(?:0{6,8})([0-9]+)")
    stuNum := numReg.FindStringSubmatch(cardData)[1]

    // split the name on the "/"
    nameSlice := strings.Split(name, "/")

    values["firstName"] = nameSlice[1]
    values["lastName"] = nameSlice[0]
    values["email"] = createEmail(nameSlice[1], nameSlice[0], stuNum)
    values["number"] = stuNum

    return values
}

func main() {
    clear()
    earlyShowing := form.GoogleForm("https://docs.google.com/forms/d/e/1FAIpQLScrLOwjzg8CL6EON0l2cSBzh6lLcgAJx3SHCLByHXw-mr65HA/formResponse")
    lateShowing := form.GoogleForm("https://docs.google.com/forms/d/e/1FAIpQLSdmBqNXLUyJ95p9_6bTl8XaksyB54PLX59-775iAj_APn6e2w/formResponse")

    earlyShowing.AddEndpoint("firstName", "entry.838312440")
    earlyShowing.AddEndpoint("lastName", "entry.1393804301")
    earlyShowing.AddEndpoint("email", "entry.86023736")
    earlyShowing.AddEndpoint("number", "entry.1876171918")

    lateShowing.AddEndpoint("firstName", "entry.2013781449")
    lateShowing.AddEndpoint("lastName", "entry.1852780729")
    lateShowing.AddEndpoint("email", "entry.1400159114")
    lateShowing.AddEndpoint("number", "entry.2097210479")

    reader := bufio.NewReader(os.Stdin)
    for {
        printHeader()
        fmt.Println("Please swipe a card...")
        cardData, _ := reader.ReadString('\n')
        values := getStudentData(string(cardData))
        if len(os.Args) > 1 && os.Args[1] == "early" {
            fmt.Println("early showing picked")
            earlyShowing.Post(values)
        } else if len(os.Args) > 1 && os.Args[1] == "late" {
            fmt.Println("late showing picked")
            lateShowing.Post(values)
        } else {
            fmt.Println("Enter 1 for early showing, 2 for late showing...")
            answer, _ := reader.ReadString('\n')
            answer = strings.TrimRight(string(answer), "\n")
            if answer == "1" {
                fmt.Println("early showing picked")
                earlyShowing.Post(values)
            } else if answer == "2" {
                fmt.Println("late showing picked")
                lateShowing.Post(values)
            }
        }
        fmt.Println("Swipe for " + values["firstName"] + " submitted successfully...")
        time.Sleep(time.Second)
        clear()
    }
}
