package main

import "fmt"

func mapFunc(){
    var nilMap map[string]int; // nil map
    totalWins := map[string]int{}; // zero length map
    teams := map[string][]string {
        "manu": []string {"bruno", "onana", "dalot"},
        "manc": []string {"debruyne", "haaland", "akanje"},
    };
    // nilMap["marcos"] = 10; // can't assing value to nil map
    totalWins["marcos"] = 10;
    totalWins["marcos"]++;
    fmt.Println(nilMap)
    fmt.Println(totalWins)
    fmt.Println(teams)
    fmt.Println(totalWins["marcos"])
    fmt.Println(totalWins["bruno"]) // as there is no key, default value for value type i.e. int is zero(0). so output: 0
    // v, ok := totalWins["marcos"]
    // fmt.Println(v, ok)
    delete(teams, "manc") // deletes if key is present and does nothing if there is no key
    fmt.Println(teams, len(teams))
    clear(teams)
    fmt.Println(teams, len(teams))
}
