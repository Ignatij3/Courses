func exists (name string) bool {
	f, err:=os.Open(name)
	if err==nil {
		f.Close()
		return true
	} else {
		return !os.IsNotExist(err)
	}
}
// Import "os"
