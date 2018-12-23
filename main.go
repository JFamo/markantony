package main

import s "strings"
import "fmt"
var p = fmt.Println

func main() {

    var x string = "GFN{lbhPnaEbgngrnaqGungfCergglPbby}"
    x = s.Repeat(s.Replace(x,"P","x",1),2)
    x = s.Join(s.Split(x,"g"),"ipd")
    x = s.Replace(x, "{", "3", 12)
    x = s.Replace(x, "}", "7", 12)
    y := string(s.Count(x,"i") + 69)
    var z []string
    z = append(z,x)
    z = append(z,y)
    x = s.Join(z,"")
    x = s.Replace(x,"i",y,93)
    p(x)

    var out string = x
    out = s.Replace(out,"O","i",93)
    out = s.Replace(out,"3","{",100)
    out = s.Replace(out,"7","}",100)
    out = s.Replace(out,"ipd","g",100)
    out = s.Replace(out,"x","P",100)
    p(out)
}
    
