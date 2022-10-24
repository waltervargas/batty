package batty

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
)

type Status struct {
	ChargePercent int
}

/*
 Compiling regular expressions is slow and uses a lot of memory, so we avoid
 wasted work by compiling them once in a package-level var statement, and using
 the compiled value in a function. - John Arundel -- The Power of Go: Tools
*/
var acpiOutputRegex = regexp.MustCompile("([0-9]+)%")

func ParseACPIOutput(output string) (Status, error) {
	matches := acpiOutputRegex.FindStringSubmatch(output)
	if len(matches) < 2 {
		return Status{}, fmt.Errorf("failed to parse acpi -b output: %q", output)
	}
	charge, err := strconv.Atoi(matches[1])
	if err != nil {
		return Status{}, fmt.Errorf("failed to parse charge percentage: %q", output)
	}
	return Status{ChargePercent: charge}, nil
}

func GetACPIOutput() (string, error) {
	data, err := exec.Command("/usr/bin/acpi", "-b").CombinedOutput()
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func Run() {
	// TODO: Add support for macos
	data, err := GetACPIOutput()
	if err != nil {
		panic(err)
	}

	status, err := ParseACPIOutput(data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Battery Charge Precent: %d%%\n", status.ChargePercent)
}
