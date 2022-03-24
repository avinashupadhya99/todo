package task

func ListTasks() [][]string {
	data := [][]string{
		[]string{"Task 1", "20-03-2022"},
		[]string{"Task 2", "21-03-2022"},
		[]string{"Task 3", "18-03-2022"},
		[]string{"Task 4", "19-03-2022"},
	}

	return data
}

func ListAllTasks() [][]string {
	data := [][]string{
		[]string{"Task 1", "20-03-2022", "No"},
		[]string{"Task 2", "21-03-2022", "No"},
		[]string{"Task 3", "18-03-2022", "No"},
		[]string{"Task 5", "19-03-2022", "No"},
		[]string{"Task 6", "12-03-2022", "Yes"},
		[]string{"Task 7", "10-03-2022", "Yes"},
		[]string{"Task 8", "15-03-2022", "Yes"},
	}

	return data
}
