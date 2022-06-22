type Vehicle struct {
	brand  string
	wheels int
}

func (v *Vehicle) GetBrand() string {
	return v.brand

}

func (v *Vehicle) GetSpeed() int {
	return v.wheels * 5
}

type Car struct {
	Vehicle
}

type Motorcycle struct {
	Base Vehicle
}





value, err := SomeNotSafeOperation()

if err != nil {
	log.Fatal(err)
}

panic("Oh, no")





package api

type API struct {
	storage *storage.storage
	server  *http.Server
}
