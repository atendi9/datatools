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
func FindMostRelevant(q, e [][]float64) ([]int, error) {
	if len(q) == 0 || len(e) == 0 || len(q[0]) == 0 || len(e[0]) == 0 {
		return nil, ErrInputMatricesCannotBeEmpty
	}

	qNorm := Normalize(q[0])
	similarityIndices := make([]struct {
		similarity float64
		index      int
	}, len(e))

	for i, emb := range e {
		eNorm := Normalize(emb)
		similarityIndices[i] = struct {
			similarity float64
			index      int
		}{
			similarity: DotProduct(qNorm, eNorm),
			index:      i,
		}
	}

	sort.Slice(similarityIndices, func(i, j int) bool {
		return similarityIndices[i].similarity > similarityIndices[j].similarity
	})

	topIndices := make([]int, len(e))
	for i, si := range similarityIndices {
		topIndices[i] = si.index
	}

	return topIndices, nil
}
