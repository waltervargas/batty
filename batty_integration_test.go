//go:build integration

package batty_test

import (
	"testing"

	"github.com/waltervargas/batty"
)

func TestGetACPIOut(t *testing.T) {
	t.Parallel()
	output, err := batty.GetACPIOutput()
	if err != nil {
		t.Fatal(err)
	}
	status, err := batty.ParseACPIOutput(output)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Charge: %d%%", status.ChargePercent)
}
