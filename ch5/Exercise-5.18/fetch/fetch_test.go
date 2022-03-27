package fetch

import (
	"os"
	"testing"
)

func TestFetch(t *testing.T) {
	url := "https://golang.org"
	dir := "test_dir"

	if tmp, _ := os.Open(dir); tmp != nil {
		t.Fatalf("wrong test conditions, %s already exists", dir)
	}
	if err := os.Mkdir(dir, 0777); err != nil {
		t.Fatal("cant create dir for test")
	}
	if err := os.Chdir(dir); err != nil {
		t.Fatal("cant change working directory")
	}

	defer func() {
		if err := os.Chdir(".."); err != nil {
			t.Fatal("cant change working directory after test")
		}
		if err := os.RemoveAll(dir); err != nil {
			t.Fatalf("cant remove directory %s arter test", dir)
		}
	}()

	if local, n, err := Fetch(url); err == nil {
		if local == "" {
			t.Fatal("no file path in results")
		}
		if n == 0 {
			t.Fatal("downloaded file has 0 length")
		}
	} else {
		t.Fatal(err)
	}

}
