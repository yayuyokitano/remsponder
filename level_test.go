package remsponder

import (
	"testing"

	"github.com/yayuyokitano/kitaipu"
)

func TestLevel(t *testing.T) {
	r := Remsponder{}
	interaction := kitaipu.Command{
		GuildID: "12345",
		Member: kitaipu.Member{
			User: kitaipu.User{
				ID: "12345",
			},
		},
		Data: kitaipu.CommandData{
			Options: kitaipu.Options{
				{
					Name: "display",
					Options: kitaipu.Options{
						{
							Name:  "user",
							Value: "12345",
						},
					},
				},
			},
			Resolved: kitaipu.Resolved{
				Users: map[string]kitaipu.User{
					"12345": {
						ID:       "12345",
						Username: "test",
					},
				},
			},
		},
	}

	resp, err := r.Level(interaction)
	if err != nil {
		t.Error(err)
	}

	if resp.Content != "test has 0 xp, level 0, 0% until level 1 (100xp remaining)" {
		t.Errorf("Expected %s, got %s", "test has 0 xp, level 0, 0% until level 1 (100xp remaining)", resp.Content)
	}

	interaction = kitaipu.Command{
		GuildID: "868012888827256882",
		Member: kitaipu.Member{
			User: kitaipu.User{
				ID: "12345",
			},
		},
		Data: kitaipu.CommandData{
			Options: kitaipu.Options{
				{
					Name: "display",
					Options: kitaipu.Options{
						{
							Name:  "user",
							Value: "196249128286552064",
						},
					},
				},
			},
			Resolved: kitaipu.Resolved{
				Users: map[string]kitaipu.User{
					"196249128286552064": {
						ID:       "196249128286552064",
						Username: "Themex",
					},
				},
			},
		},
	}

	resp, err = r.Level(interaction)
	if err != nil {
		t.Error(err)
	}

	if resp.Content != "Themex has 24 xp, level 0, 24% until level 1 (76xp remaining)" {
		t.Errorf("Expected %s, got %s", "Themex has 24 xp, level 0, 24% until level 1 (76xp remaining)", resp.Content)
	}

	interaction = kitaipu.Command{
		GuildID: "868012888827256882",
		Member: kitaipu.Member{
			User: kitaipu.User{
				ID:       "196249128286552064",
				Username: "Themex",
			},
		},
		Data: kitaipu.CommandData{
			Options: kitaipu.Options{
				{
					Name: "display",
				},
			},
		},
	}

	resp, err = r.Level(interaction)
	if err != nil {
		t.Error(err)
	}

	if resp.Content != "Themex has 24 xp, level 0, 24% until level 1 (76xp remaining)" {
		t.Errorf("Expected %s, got %s", "Themex has 24 xp, level 0, 24% until level 1 (76xp remaining)", resp.Content)
	}

}

func TestGetLevel(t *testing.T) {
	if (getLevel(0)) != 0 {
		t.Error("Expected 0, got ", getLevel(0))
	}
	if getLevel(100) != 1 {
		t.Error("Expected 1, got ", getLevel(100))
	}
	if getLevel(255) != 2 {
		t.Error("Expected 2, got ", getLevel(255))
	}
	if getLevel(254) != 1 {
		t.Error("Expected 1, got ", getLevel(254))
	}

	if getRequiredXP(0) != 0 {
		t.Error("Expected 0, got ", getRequiredXP(0))
	}
	if getRequiredXP(1) != 100 {
		t.Error("Expected 100, got ", getRequiredXP(1))
	}
	if getRequiredXP(2) != 255 {
		t.Error("Expected 255, got ", getRequiredXP(2))
	}
}
