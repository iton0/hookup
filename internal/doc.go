package internal

import "errors"

var gitHooks = map[string]string{ // NOTE: 28 default hooks
	"applypatch-msg":        "#_applypatch_msg",
	"pre-applypatch":        "#_pre_applypatch",
	"post-applypatch":       "#_post_applypatch",
	"pre-commit":            "#_pre_commit",
	"pre-merge-commit":      "#_pre_merge_commit",
	"prepare-commit-msg":    "#_prepare_commit_msg",
	"commit-msg":            "#_commit_msg",
	"post-commit":           "#_post_commit",
	"pre-rebase":            "#_pre_rebase",
	"post-checkout":         "#_post_checkout",
	"post-merge":            "#_post_merge",
	"pre-push":              "#_pre_push",
	"pre-receive":           "#_pre_receive",
	"update":                "#_update",
	"proc-receive":          "#_proc_receive",
	"post-receive":          "#_post_receive",
	"post-update":           "#_post_update",
	"reference-transaction": "#_reference_transaction",
	"push-to-checkout":      "#_push_to_checkout",
	"pre-auto-gc":           "#_pre_auto_gc",
	"post-rewrite":          "#_post_rewrite",
	"sendemail-validate":    "#_sendemail_validate",
	"fsmonitor-watchman":    "#_fsmonitor_watchman",
	"p4-changelist":         "#_p4_changelist",
	"p4-prepare-changelist": "#_p4_prepare_changelist",
	"p4-post-changelist":    "#_p4_post_changelist",
	"p4-pre-submit":         "#_p4_pre_submit",
	"post-index-change":     "#_post_index_change",
}

func getGitHook(key string) (string, error) {
	value, exists := gitHooks[key]
	if !exists {
		return "", errors.New("hook not found: " + key)
	}
	return value, nil
}

func OpenBrowser() {
	const SITE = "https://git-scm.com/docs/githooks"
}
