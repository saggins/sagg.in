package main

func getBlobByID(id int) Item {
	newid := string(id)
	return getRaws(newid)
}
