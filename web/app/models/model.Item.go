package models

type Item struct {
	Blobs      []string `json:"blobs"`
	Title      string   `json:"title"`
	ID         string   `json:"id"`
	BlobsTitle []string `json:"blobstitle"`
}