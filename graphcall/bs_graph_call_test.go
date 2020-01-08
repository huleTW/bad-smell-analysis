package graphcall

import (
	. "github.com/onsi/gomega"
	"testing"
)

func TestBadSmellGraphCall_AnalysisGraphCallPath(t *testing.T) {
	g := NewGomegaWithT(t)

	app := NewBadSmellGraphCall()
	var nodes = map[string][]string{
		"A": {"C", "B", "D"},
		"B": {"D", "H"},
		"C": {"D", "M"},
		"D": {"E", "B"},
		"E": {"B", "D"}}
	results := app.AnalysisGraphCallPath(nodes)

	g.Expect(len(results)).To(Equal(9))
	g.Expect(contains(results, "A->B->D;A->D")).To(Equal(true))
	g.Expect(contains(results, "A->C->D;A->D")).To(Equal(true))
	g.Expect(contains(results, "A->D->B;A->B")).To(Equal(true))
	g.Expect(contains(results, "A->D->E->B;A->B")).To(Equal(true))
}
