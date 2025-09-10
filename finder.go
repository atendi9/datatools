package datatools

import (
	"errors"
	"sort"
)

// ErrInputMatricesCannotBeEmpty is returned when either the query matrix (q) or
// the embeddings matrix (e) is empty, or when their inner vectors have zero length.
var ErrInputMatricesCannotBeEmpty = errors.New("input matrices cannot be empty")

// FindMostRelevant returns a slice of indices ordered by relevance, or an error if the input
// matrices are invalid.
//   - q is the query embedding
//   - e is the matrix of embeddings to search in.
func FindMostRelevant(
	q, e [][]float64,
) ([]int, error) {
	if len(q) == 0 || len(e) == 0 || len(q[0]) == 0 || len(e[0]) == 0 {
		return nil, ErrInputMatricesCannotBeEmpty
	}

	qNorm := Normalize(q[0])
	eNorm := make([][]float64, len(e))

	for i := range e {
		eNorm[i] = Normalize(e[i])
	}

	similarities := make([]float64, 0)
	for i := range eNorm {
		dotProduct := DotProduct(qNorm, eNorm[i])
		similarities = append(similarities, dotProduct)
	}
	type similarityIndex struct {
		similarity float64
		index      int
	}
	similarityIndices := make([]similarityIndex, len(similarities))
	for i, sim := range similarities {
		similarityIndices[i] = similarityIndex{similarity: sim, index: i}
	}
	sort.Slice(similarityIndices, func(i, j int) bool {
		return similarityIndices[i].similarity > similarityIndices[j].similarity
	})

	topIndices := make([]int, len(similarities))
	for i := range similarities {
		topIndices[i] = similarityIndices[i].index
	}

	return topIndices, nil
}
