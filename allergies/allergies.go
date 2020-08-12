package allergies

const (
	Eggs uint = 1 << iota
	Peanuts
	Shellfish
	Strawberries
	Tomatoes
	Chocolate
	Pollen
	Cats
)

func Allergies(i uint) []string {
	var result []string

	if hasAllergy(i, Eggs) {
		result = append(result, "eggs")
	}

	if hasAllergy(i, Peanuts) {
		result = append(result, "peanuts")
	}

	if hasAllergy(i, Shellfish) {
		result = append(result, "shellfish")
	}

	if hasAllergy(i, Strawberries) {
		result = append(result, "strawberries")
	}

	if hasAllergy(i, Tomatoes) {
		result = append(result, "tomatoes")
	}

	if hasAllergy(i, Chocolate) {
		result = append(result, "chocolate")
	}

	if hasAllergy(i, Pollen) {
		result = append(result, "pollen")
	}

	if hasAllergy(i, Cats) {
		result = append(result, "cats")
	}

	return result
}

func hasAllergy(i uint, allergy uint) bool { return i&allergy != 0 }

func AllergicTo(score uint, substance string) bool {
	switch substance {
	case "cats":
		return hasAllergy(score, Cats)
	case "pollen":
		return hasAllergy(score, Pollen)
	case "chocolate":
		return hasAllergy(score, Chocolate)
	case "tomatoes":
		return hasAllergy(score, Tomatoes)
	case "strawberries":
		return hasAllergy(score, Strawberries)
	case "shellfish":
		return hasAllergy(score, Shellfish)
	case "peanuts":
		return hasAllergy(score, Peanuts)
	case "eggs":
		return hasAllergy(score, Eggs)
	default:
		return false
	}
}
