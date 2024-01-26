package pagination

var defaultLimit = 10

func SetPagination(page, limit int) (int, int, int) {
	if page <= 0 {
		page = 1
	}

	if limit <= 0 {
		limit = defaultLimit
	}

	offset := (page - 1) * limit

	return offset, limit, page
}

func SetLastPage(page, limit, total int) (lastPage int) {

	if total > 0 {
		lastPage = total / limit

		if total%limit != 0 {
			lastPage = lastPage + 1
		}
	} else {
		lastPage = 0
	}

	return lastPage
}
