package batty_test

import (
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/waltervargas/batty"
)

func TestParseACPIOutput(t *testing.T) {
	t.Parallel()
	data, err := os.ReadFile("testdata/acpi.txt")
	if err != nil {
		t.Fatal(err)
	}
	want := batty.Status{
		ChargePercent: 82,
	}
	got, err := batty.ParseACPIOutput(string(data))
	if err != nil {
		t.Fatal(err)
	}
	if !cmp.Equal(want, got) {
		t.Error(cmp.Diff(want, got))
	}

}
