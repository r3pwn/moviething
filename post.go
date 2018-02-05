package main

import (
   "strings"
   "net/url"
   "net/http"
)

var URL string = "https://docs.google.com/forms/d/e/1FAIpQLSdNpMBFYVMIOrcj8SixavtCdvuv7P3oGEPGSAoJYvH6J6CMgw/formResponse"

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
