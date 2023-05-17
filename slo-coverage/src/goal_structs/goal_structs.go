package goal_structs

// Commenting out Assets - don't need that struct for this application.
type Goals struct {
	Data struct {
		Goals []struct {
			ArchivedAt any    `json:"archivedAt"`
			ID         string `json:"id"`
			Name       string `json:"name"`
			Episodes   []struct {
				ArchivedAt any    `json:"archivedAt"`
				ID         string `json:"id"`
				Name       string `json:"name"`
				Jobs       []struct {
					ArchivedAt any    `json:"archivedAt"`
					ID         string `json:"id"`
					Name       string `json:"name"`
				} `json:"jobs"`
			} `json:"episodes"`
		} `json:"goals"`
		// Assets []struct {
		// 	ArchivedAt any    `json:"archivedAt"`
		// 	ID         string `json:"id"`
		// 	Name       string `json:"name"`
		// 	Owner      string `json:"owner"`
		// 	Components []struct {
		// 		ArchivedAt  any    `json:"archivedAt"`
		// 		ID          string `json:"id"`
		// 		Name        string `json:"name"`
		// 		Description any    `json:"description"`
		// 	} `json:"components"`
		// } `json:"assets"`
	} `json:"data"`
}