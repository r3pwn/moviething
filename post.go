package main

import (
   "strings"
   "net/url"
   "net/http"
)

var UrlShowOne string = "https://docs.google.com/forms/d/e/1FAIpQLScrLOwjzg8CL6EON0l2cSBzh6lLcgAJx3SHCLByHXw-mr65HA/formResponse"
var FnameShowOne string = "entry.838312440"
var LnameShowOne string = "entry.1393804301"
var NumberShowOne string = "entry.1876171918"
var EmailShowOne string = "entry.86023736"

var UrlShowTwo string = "https://docs.google.com/forms/d/e/1FAIpQLSdmBqNXLUyJ95p9_6bTl8XaksyB54PLX59-775iAj_APn6e2w/formResponse"
var FnameShowTwo string = "entry.2013781449"
var LnameShowTwo string = "entry.1852780729"
var NumberShowTwo string = "entry.2097210479"
var EmailShowTwo string = "entry.1400159114"

func Post(firstName, lastName, number, email string) {
   hc := http.Client{}

   form := url.Values{}
   form.Add("entry.2005620554", firstName)
   form.Add("entry.1045781291", lastName)
   form.Add("entry.1065046570", number)
   form.Add("entry.1166974658", email)
   req, _ := http.NewRequest("POST", URL, strings.NewReader(form.Encode()))
   req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
   hc.Do(req)
}
