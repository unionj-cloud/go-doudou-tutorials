package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Strudent struct {
	Name string
}

func main() {
	//os.Setenv("GOSTATS", "http://localhost:6060")
	//defer os.Unsetenv("GOSTATS")
	//statsClient := client.NewGoStatsClient()
	//_, data, err := statsClient.GetShelves_ShelfBooks_Book(context.Background(), nil, 20, "钢铁是怎样炼成的")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(data)

	var stu *Strudent
	err := json.NewDecoder(strings.NewReader(`{"name":"wubin"}`)).Decode(&stu)
	if err != nil {
		panic(err)
	}
	fmt.Println(stu)
}
