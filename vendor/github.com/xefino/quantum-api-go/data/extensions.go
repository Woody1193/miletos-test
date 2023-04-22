package data

// ForDate returns true if the frequency is less than or equal to a day (Day, Week, Month, Quarter, Year)
func (enum Frequency) ForDate() bool {
	switch enum {
	case Frequency_Day, Frequency_Week, Frequency_Month, Frequency_Quarter, Frequency_Year:
		return true
	default:
		return false
	}
}

// ForTime returns true if the frequency is more than a day (Second, Minute, Hour)
func (enum Frequency) ForTime() bool {
	switch enum {
	case Frequency_Second, Frequency_Minute, Frequency_Hour:
		return true
	default:
		return false
	}
}
