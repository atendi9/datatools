package datatools

import (
	"bytes"
	"encoding/json"
	"os"
	"slices"
	"testing"
)

func TestFindMostRelevantEmbeddings(t *testing.T) {
	type GeneratedEmbedding struct {
		Value []struct {
			Index      int       `json:"index"`
			Embeddings []float64 `json:"embs"`
		} `json:"emb"`
	}
	parseEmbeddingFile := func(fileName string) (ge GeneratedEmbedding) {
		b, _ := os.ReadFile(fileName)
		if err := json.NewDecoder(bytes.NewBuffer(b)).Decode(&ge); err != nil {
			t.Fatal(err)
		}
		return
	}
	embeddings := [][]float64{}
	embeddingsData := parseEmbeddingFile("file.embeddings")
	for _, data := range embeddingsData.Value {
		if len(data.Embeddings) > 0 {
			embeddings = append(embeddings, data.Embeddings)
		}
	}
	queryEmbeddingsData := parseEmbeddingFile("file.query")
	indexes, err := FindMostRelevant([][]float64{queryEmbeddingsData.Value[0].Embeddings}, embeddings)
	if err != nil {
		t.Fatal(err)
	}
	expected := []int{0, 7, 9, 5, 2, 3, 8, 1, 6, 4}
	if !slices.Equal(indexes, expected) {
		t.Fatalf("invalid code -> result: %v expected: %v", indexes, expected)
	}
}
