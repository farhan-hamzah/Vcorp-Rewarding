package main
import "fmt"

const (
	bigParcelPaper  int = 60 * 60
	smallParcelPaper int = 40 * 60
	paperSize        int = 80 * 100
	paperPrice       int = 750
	bigParcelCost    int = 75000
	smallParcelCost  int = 45000
)

func main() {
	var employee string
	var highLevel, entryLevel int
	var totalCostVal, totalPaperNeeded, totalSheets int
	var employeeList string

	fmt.Scan(&employee)

	for employee != "#" {
		employeeList += employee + "\n"
		cekEmployees(employee, &highLevel, &entryLevel)
		fmt.Scan(&employee)
	}

	totalEmployeeVal := totalEmployees(highLevel, entryLevel)
	totalCostVal = totalCost(highLevel, entryLevel)

	totalPaperNeeded = (highLevel * bigParcelPaper) + (entryLevel * smallParcelPaper)
	totalSheets = totalPaperNeeded / paperSize
	if totalPaperNeeded%paperSize != 0 {
		totalSheets++
	}
	totalCostVal += totalSheets * paperPrice

	fmt.Print(employeeList)
	fmt.Printf("\nVcorp membutuhkan anggaran sebesar: Rp%d, untuk memberikan parcel kepada %d pegawai.\n", totalCostVal, totalEmployeeVal)
	fmt.Printf("Yang terdiri dari %d bigParcel dan %d smallParcel.\n", highLevel, entryLevel)
}

func cekEmployees(pegawai string, highLevel, entryLevel *int) {
	if pegawai == "manager" || pegawai == "supervisor" {
		*highLevel += 1
	} else {
		*entryLevel += 1
	}
}

func totalEmployees(highLevel, entryLevel int) int {
	return highLevel + entryLevel
}

func totalCost(highLevel, entryLevel int) int {
	return (highLevel * bigParcelCost) + (entryLevel * smallParcelCost)
}
