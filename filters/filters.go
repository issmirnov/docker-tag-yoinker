package filters

// Filter is a filter function applied to a single record.
type Filter func(string) bool

// FilterBulk is a bulk filter function applied to an entire slice of records.
type FilterBulk func([]string) []string

// FilterSet is a function that applies a set of filters and returns the filtered records.
type FilterSet func([]string) []string

// ApplyFilters applies a set of filters to a record list.
// Each record will be checked against each filter.
// The filters are applied in the order they are passed in.
func ApplyFilters(records []string, filters ...Filter) []string {
	// Make sure there are actually filters to be applied.
	if len(filters) == 0 {
		return records
	}

	filteredRecords := make([]string, 0, len(records))

	// Range over the records and apply all the filters to each record.
	// If the record passes all the filters, add it to the final slice.
	for _, r := range records {
		keep := true

		for _, f := range filters {
			if !f(r) {
				keep = false
				break
			}
		}

		if keep {
			filteredRecords = append(filteredRecords, r)
		}
	}

	return filteredRecords
}
