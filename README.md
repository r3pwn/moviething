# "The Movie Thing"

## Use case?
Eh, whatever you want. We used it for signing up for movie waitlists 
at our school by having people swipe their student ID card. (Much 
quicker than having them say their name, student ID number, email,
etc.)

##Compiling for Windows
32-bit:
`GOOS=windows GOARCH=386 go build -o moviething.exe`
64-bit:
`GOOS=windows GOARCH=amd64 go build -o moviething64.exe`

Originally, it was going to use the "Terminal" package, but apparently
some of that stuff completely breaks Windows versions (I'm gonna go
ahead and blame Microsoft here.)

Oh, yeah, and this is licensed under APACHE2 or something.

```
Copyright 2018 Greg Willard (r3pwn)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```