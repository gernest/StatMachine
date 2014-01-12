package statmachine

type ByDate []Result

func(r ByDate) Len() int{
	return len(r)
}

func(r ByDate) Swap(i, j int){
	r[i], r[j] = r[j], r[i]
}

func(r ByDate) Less(i, j int) bool{
	return r[i].date.Before(r[j].date)
}
