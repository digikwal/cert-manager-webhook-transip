package transip

import "testing"

func TestNormalizeHostedDomain(t *testing.T) {
	tests := []struct {
		name    string
		zone    string
		want    string
		wantErr bool
	}{
		{name: "strips trailing dot", zone: "lslx.nl.", want: "lslx.nl"},
		{name: "keeps valid domain", zone: "example.com", want: "example.com"},
		{name: "empty zone", zone: "", wantErr: true},
		{name: "tld only", zone: "nl.", wantErr: true},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := normalizeHostedDomain(tc.zone)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil (value=%q)", got)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got != tc.want {
				t.Fatalf("unexpected normalized zone: got %q, want %q", got, tc.want)
			}
		})
	}
}
