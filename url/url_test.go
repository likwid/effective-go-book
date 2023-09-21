package url

import "testing"

func TestParse(t *testing.T) {
	const rawUrl = "https://foo.com/go"

	u, err := Parse(rawUrl)
	if err != nil {
		t.Fatalf("Parse(%q) err = %q, want nil", rawUrl, err)
	}

	got := u.Scheme  // arrange
	want := "https"  // act
	if got != want { // assert
		t.Errorf("Parse(%q).Scheme = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Host, "foo.com"; got != want {
		t.Errorf("Parse(%q).Host = %q; want %q", rawUrl, got, want)
	}

	if got, want := u.Path, "go"; got != want {
		t.Errorf("Parse(%q).Path = %q; want %q", rawUrl, got, want)
	}
}