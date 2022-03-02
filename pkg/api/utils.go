package api

func getPaginationFrom(meta Meta) int {
	return (meta.Page - 1) * meta.PerPage
}
