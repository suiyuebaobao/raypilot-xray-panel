package model

import "testing"

func TestDefaultSalesLandingConfig_UsesShortTitle(t *testing.T) {
	cfg := DefaultSalesLandingConfig()
	if cfg.Title != "高速 VPN 节点" {
		t.Fatalf("expected short sales title, got %q", cfg.Title)
	}
}

func TestNormalizeSalesLandingConfig_FiltersEmptyItems(t *testing.T) {
	cfg := NormalizeSalesLandingConfig(SalesLandingConfig{
		Brand:      "  LeiYun  ",
		Title:      "  稳定节点  ",
		NavLinks:   []SalesLandingLink{{Label: "  套餐  ", To: "  #plans  "}, {}},
		Badges:     []string{" 高速 ", " "},
		PrimaryCTA: SalesLandingLink{Label: "购买", To: "/register"},
		Plans: []SalesLandingPlan{
			{Name: "推荐", Features: []string{" 可用 ", " "}},
			{},
		},
		FinalCTA: SalesLandingCTA{
			Title:       "开通",
			Text:        "选择套餐",
			ButtonLabel: "注册",
			ButtonTo:    "/register",
			FooterLinks: []SalesLandingLink{{Label: "登录", To: "/login"}, {}},
		},
		FooterText: " 服务 ",
	})

	if cfg.Brand != "LeiYun" || cfg.Title != "稳定节点" || cfg.FooterText != "服务" {
		t.Fatalf("basic text was not trimmed: %#v", cfg)
	}
	if len(cfg.NavLinks) != 1 || cfg.NavLinks[0].Label != "套餐" || cfg.NavLinks[0].To != "#plans" {
		t.Fatalf("nav links were not normalized: %#v", cfg.NavLinks)
	}
	if len(cfg.Badges) != 1 || cfg.Badges[0] != "高速" {
		t.Fatalf("badges were not normalized: %#v", cfg.Badges)
	}
	if len(cfg.Plans) != 1 || len(cfg.Plans[0].Features) != 1 || cfg.Plans[0].Features[0] != "可用" {
		t.Fatalf("plans were not normalized: %#v", cfg.Plans)
	}
	if len(cfg.FinalCTA.FooterLinks) != 1 || cfg.FinalCTA.FooterLinks[0].Label != "登录" {
		t.Fatalf("footer links were not normalized: %#v", cfg.FinalCTA.FooterLinks)
	}
}

func TestNormalizeSalesLandingConfig_EmptyPlanFeaturesRemainEditable(t *testing.T) {
	cfg := NormalizeSalesLandingConfig(SalesLandingConfig{
		Title: "节点",
		Plans: []SalesLandingPlan{
			{Name: "无权益卡片"},
		},
	})
	if cfg.Plans[0].Features == nil {
		t.Fatalf("expected empty features slice, got nil")
	}
}

func TestNormalizeSalesLandingConfig_SanitizesUnsafeLinks(t *testing.T) {
	cfg := NormalizeSalesLandingConfig(SalesLandingConfig{
		NavLinks:     []SalesLandingLink{{Label: "危险", To: "javascript:alert(1)"}, {Label: "官网", To: "https://example.com/path"}},
		PrimaryCTA:   SalesLandingLink{Label: "购买", To: "javascript:alert(1)"},
		SecondaryCTA: SalesLandingLink{Label: "登录", To: "/login"},
		FinalCTA: SalesLandingCTA{
			Title:       "开通",
			Text:        "选择套餐",
			ButtonLabel: "注册",
			ButtonTo:    "data:text/html,boom",
			FooterLinks: []SalesLandingLink{{Label: "锚点", To: "#plans"}, {Label: "坏链接", To: "vbscript:msgbox(1)"}},
		},
	})

	if cfg.NavLinks[0].To != "#" {
		t.Fatalf("expected unsafe nav link to become #, got %q", cfg.NavLinks[0].To)
	}
	if cfg.NavLinks[1].To != "https://example.com/path" {
		t.Fatalf("expected https nav link to pass, got %q", cfg.NavLinks[1].To)
	}
	if cfg.PrimaryCTA.To != DefaultSalesLandingConfig().PrimaryCTA.To {
		t.Fatalf("expected unsafe primary CTA to fallback, got %q", cfg.PrimaryCTA.To)
	}
	if cfg.FinalCTA.ButtonTo != DefaultSalesLandingConfig().FinalCTA.ButtonTo {
		t.Fatalf("expected unsafe final CTA to fallback, got %q", cfg.FinalCTA.ButtonTo)
	}
	if cfg.FinalCTA.FooterLinks[0].To != "#plans" || cfg.FinalCTA.FooterLinks[1].To != "#" {
		t.Fatalf("footer links were not sanitized: %#v", cfg.FinalCTA.FooterLinks)
	}
}

func TestParseSalesLandingConfig_InvalidJSONReturnsDefault(t *testing.T) {
	cfg := ParseSalesLandingConfig("{bad json")
	if cfg.Title != DefaultSalesLandingConfig().Title {
		t.Fatalf("expected default config on invalid json, got %#v", cfg)
	}
}
