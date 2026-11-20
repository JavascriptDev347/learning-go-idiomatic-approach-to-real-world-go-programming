package main

import "time"

type RFC822ZTime struct {
	time.Time
}

func (rt RFC822ZTime) MarshalJSON() ([]byte, error) {
	out := rt.Time.Format(time.RFC822Z)
	return []byte(`"` + out + `"`), nil
}

func (rt *RFC822ZTime) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}
	t, err := time.Parse(`"`+time.RFC822Z+`"`, string(b))
	if err != nil {
		return err
	}
	*rt = RFC822ZTime{t}
	return nil
}
func main() {

}
