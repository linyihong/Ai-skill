package app

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/linyihong/Ai-skill/scripts/ai-skill-cli/portable/kge"
	"gopkg.in/yaml.v3"
)

func TestFeedbackReviewCanonicalSchemaParity(t *testing.T) {
	repo := discoveryRepoRoot(t)
	b, err := os.ReadFile(filepath.Join(repo, "enforcement/feedback-lessons.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	var contract struct {
		Review struct {
			Transfer  []string `yaml:"required_transfer_fields"`
			Decisions []string `yaml:"decision_values"`
			Snapshot  string   `yaml:"reference_snapshot"`
		} `yaml:"review_schema"`
	}
	if err := yaml.Unmarshal(b, &contract); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(contract.Review.Transfer, kge.FeedbackTransferFields()) {
		t.Fatalf("canonical transfer schema drift: %+v", contract.Review.Transfer)
	}
	if !reflect.DeepEqual(contract.Review.Decisions, kge.FeedbackIntelligenceDecisions()) {
		t.Fatalf("canonical intelligence decisions drift: %+v", contract.Review.Decisions)
	}
	if contract.Review.Snapshot != "git_index" {
		t.Fatal("reference snapshot must remain git_index")
	}
}
