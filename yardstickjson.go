package pystick

import (
	"encoding/json"
	"fmt"
)

//Classes Users struct which contains
type Classes struct {
	Boats []Boat `json:"classes"`
}

//Boat ...
type Boat struct {
	Name string `json:"name"`
	Hcap int    `json:"hcap"`
}

//Yardstick takes an class name arg and returns a class handicap
func Yardstick(boat string) int {
	fmt.Println(BoatClasses)
	// Open our jsonFile
	//jsonFile, err := os.Open("hcap.json")
	// if we os.Open returns an error then handle it
	//if err != nil {
	//	fmt.Println(err)

	//fmt.Println(jsonFile)

	//fmt.Println("Successfully Opened hcap.json")
	// defer the closing of our jsonFile so that we c	an parse it later on
	//	defer jsonFile.Close()

	//	byteValue, _ := ioutil.ReadAll(jsonFile)
	//fmt.Println(byteValue)

	var classes Classes

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'boat' which we defined above
	//json.Unmarshal(byteValue, &classes)
	json.Unmarshal([]byte(BoatClasses), &classes)
	//fmt.Println(classes.Boats)
	Pys := 0
	for i := 0; i < len(classes.Boats); i++ {
		//fmt.Println("Class type: " + classes.Boats[i].Name)
		//fmt.Println("Handicap: " + strconv.Itoa(classes.Boats[i].Hcap))

		if classes.Boats[i].Name == boat {
			Pys = classes.Boats[i].Hcap
			return Pys
		}

	}
	return Pys
}
