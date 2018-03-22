package afas

type UpdateConnectors []UpdateConnector

type UpdateConnector struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}

type GetConnectors []GetConnector

type GetConnector struct {
	ID          string `json:"id"`
	Description string `json:"description"`
}
