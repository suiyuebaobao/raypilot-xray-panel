<template>
  <router-view />
</template>

<script setup>
// App.vue — 根组件。
// 职责：作为路由出口，不持有业务逻辑。
</script>

<style>
:root {
  color-scheme: dark;
  --rp-bg: #070b14;
  --rp-bg-soft: #0b1220;
  --rp-panel: rgba(13, 22, 39, 0.84);
  --rp-panel-strong: rgba(16, 28, 51, 0.94);
  --rp-line: rgba(92, 241, 255, 0.2);
  --rp-line-strong: rgba(92, 241, 255, 0.42);
  --rp-cyan: #42f5ff;
  --rp-blue: #5c7cff;
  --rp-pink: #ff3df2;
  --rp-amber: #ffd166;
  --rp-text: #eaf7ff;
  --rp-muted: #8fa7bd;
  --rp-danger: #ff4d7d;
  --rp-success: #2fffb2;

  --el-color-primary: #40e7ff;
  --el-color-success: #2fffb2;
  --el-color-warning: #ffd166;
  --el-color-danger: #ff4d7d;
  --el-color-error: #ff4d7d;
  --el-color-info: #8fa7bd;
  --el-bg-color: #0b1220;
  --el-bg-color-page: #070b14;
  --el-bg-color-overlay: #101a2d;
  --el-text-color-primary: #eaf7ff;
  --el-text-color-regular: #c6d8e8;
  --el-text-color-secondary: #8fa7bd;
  --el-text-color-placeholder: #5d7186;
  --el-border-color: rgba(92, 241, 255, 0.18);
  --el-border-color-light: rgba(92, 241, 255, 0.14);
  --el-border-color-lighter: rgba(92, 241, 255, 0.1);
  --el-fill-color: rgba(18, 31, 55, 0.82);
  --el-fill-color-light: rgba(24, 39, 68, 0.72);
  --el-fill-color-lighter: rgba(28, 45, 77, 0.62);
  --el-fill-color-blank: rgba(9, 16, 30, 0.92);
  --el-mask-color: rgba(2, 6, 14, 0.78);
  --el-box-shadow: 0 16px 48px rgba(0, 0, 0, 0.38);
  --el-box-shadow-light: 0 12px 36px rgba(0, 0, 0, 0.28);
  --el-border-radius-base: 6px;
}

html, body {
  margin: 0;
  padding: 0;
  height: 100%;
  font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, 'Helvetica Neue', Arial, sans-serif;
  color: var(--rp-text);
  background:
    linear-gradient(rgba(66, 245, 255, 0.035) 1px, transparent 1px),
    linear-gradient(90deg, rgba(66, 245, 255, 0.035) 1px, transparent 1px),
    radial-gradient(circle at top left, rgba(92, 124, 255, 0.18), transparent 34rem),
    linear-gradient(135deg, #050814 0%, #081225 44%, #12091c 100%);
  background-size: 32px 32px, 32px 32px, auto, auto;
}
#app {
  height: 100%;
}

* {
  box-sizing: border-box;
}

a {
  color: var(--rp-cyan);
}

::selection {
  color: #061019;
  background: var(--rp-cyan);
}

.el-card,
.el-dialog,
.el-popper,
.el-message-box {
  background: var(--rp-panel-strong);
  border: 1px solid var(--rp-line);
  box-shadow: 0 18px 52px rgba(0, 0, 0, 0.42), inset 0 1px 0 rgba(255, 255, 255, 0.04);
  backdrop-filter: blur(14px);
}

.el-card {
  border-radius: 8px;
}

.el-button {
  border-radius: 6px;
  font-weight: 600;
}

.el-button--primary {
  color: #061019;
  border-color: rgba(66, 245, 255, 0.86);
  background: linear-gradient(135deg, #42f5ff 0%, #6f7dff 100%);
  box-shadow: 0 0 18px rgba(66, 245, 255, 0.24);
}

.el-button--primary:hover,
.el-button--primary:focus {
  border-color: #a8fbff;
  background: linear-gradient(135deg, #88fbff 0%, #9a6cff 100%);
  box-shadow: 0 0 28px rgba(66, 245, 255, 0.34);
}

.el-button.is-text,
.el-button--primary.is-text,
.el-button--danger.is-text,
.el-button--warning.is-text {
  color: var(--rp-cyan);
}

.el-input__wrapper,
.el-textarea__inner,
.el-select__wrapper,
.el-input-number .el-input__wrapper {
  background: rgba(5, 12, 24, 0.78);
  border-radius: 6px;
  box-shadow: 0 0 0 1px rgba(92, 241, 255, 0.16) inset;
}

.el-input__wrapper:hover,
.el-textarea__inner:hover,
.el-select__wrapper:hover {
  box-shadow: 0 0 0 1px rgba(92, 241, 255, 0.34) inset;
}

.el-input__wrapper.is-focus,
.el-textarea__inner:focus,
.el-select__wrapper.is-focused {
  box-shadow: 0 0 0 1px rgba(66, 245, 255, 0.86) inset, 0 0 18px rgba(66, 245, 255, 0.18);
}

.el-table {
  --el-table-bg-color: rgba(9, 16, 30, 0.72);
  --el-table-tr-bg-color: rgba(9, 16, 30, 0.54);
  --el-table-header-bg-color: rgba(16, 30, 55, 0.94);
  --el-table-row-hover-bg-color: rgba(66, 245, 255, 0.08);
  --el-table-border-color: rgba(92, 241, 255, 0.12);
  color: var(--rp-text);
  border-radius: 8px;
  overflow: hidden;
  background: rgba(9, 16, 30, 0.72);
}

.el-table th.el-table__cell {
  color: #bdf8ff;
  font-weight: 700;
  background: rgba(16, 30, 55, 0.94);
}

.el-table .el-table__cell {
  border-color: rgba(92, 241, 255, 0.11);
}

.el-tag {
  border-radius: 999px;
  border-color: rgba(66, 245, 255, 0.32);
  background: rgba(66, 245, 255, 0.08);
}

.el-tag--success {
  border-color: rgba(47, 255, 178, 0.36);
  color: var(--rp-success);
  background: rgba(47, 255, 178, 0.08);
}

.el-tag--warning {
  border-color: rgba(255, 209, 102, 0.42);
  color: var(--rp-amber);
  background: rgba(255, 209, 102, 0.08);
}

.el-tag--danger {
  border-color: rgba(255, 77, 125, 0.42);
  color: var(--rp-danger);
  background: rgba(255, 77, 125, 0.08);
}

.el-progress-bar__outer {
  background: rgba(92, 241, 255, 0.1);
}

.el-progress-bar__inner {
  background: linear-gradient(90deg, var(--rp-cyan), var(--rp-pink));
}

.el-pagination button,
.el-pager li {
  background: rgba(8, 15, 28, 0.8);
  color: var(--rp-muted);
}

.el-pager li.is-active {
  color: #061019;
  background: var(--rp-cyan);
}

.el-tabs__item {
  color: var(--rp-muted);
}

.el-tabs__item.is-active {
  color: var(--rp-cyan);
}

.el-tabs__active-bar {
  background: linear-gradient(90deg, var(--rp-cyan), var(--rp-pink));
}

.el-overlay {
  backdrop-filter: blur(8px);
}

[class^="admin-"],
[class$="-page"],
.user-profile {
  color: var(--rp-text);
}

[class^="admin-"] > h2,
[class$="-page"] > h2,
.user-profile > h2,
.user-profile > h3,
.header h2,
.header-row h2 {
  margin-top: 0;
  color: var(--rp-text);
  font-weight: 800;
  letter-spacing: 0;
}

.header,
.header-row,
.toolbar,
.batch-toolbar,
.nodes-toolbar,
.backend-toolbar,
.token-list-header {
  border-color: rgba(92, 241, 255, 0.16) !important;
  color: var(--rp-muted) !important;
}

.toolbar,
.batch-toolbar,
.nodes-toolbar,
.backend-toolbar,
.deploy-steps,
.usage-metric,
.sub-link-item {
  background: rgba(9, 16, 30, 0.62) !important;
  border: 1px solid rgba(92, 241, 255, 0.16) !important;
  border-radius: 8px !important;
  box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.035);
}

.el-card__header {
  border-bottom: 1px solid rgba(92, 241, 255, 0.14);
  color: #bdf8ff;
  font-weight: 700;
}

.el-descriptions {
  --el-descriptions-table-border: rgba(92, 241, 255, 0.14);
}

.el-descriptions__label.el-descriptions__cell.is-bordered-label {
  color: #bdf8ff;
  background: rgba(16, 30, 55, 0.82);
}

.el-descriptions__content.el-descriptions__cell.is-bordered-content {
  color: var(--rp-text);
  background: rgba(9, 16, 30, 0.58);
}

.el-form-item__label,
.el-checkbox,
.el-radio,
.el-radio-button__inner {
  color: var(--rp-muted);
}

.el-radio-button__inner {
  background: rgba(8, 15, 28, 0.8);
  border-color: rgba(92, 241, 255, 0.18);
}

.el-radio-button__original-radio:checked + .el-radio-button__inner {
  color: #061019;
  background: var(--rp-cyan);
  border-color: var(--rp-cyan);
  box-shadow: -1px 0 0 0 var(--rp-cyan);
}

.el-alert {
  background: rgba(9, 16, 30, 0.72);
  border: 1px solid rgba(92, 241, 255, 0.16);
}

.el-alert--info {
  color: var(--rp-muted);
}

.el-empty__description p,
.el-divider__text {
  color: var(--rp-muted);
  background: var(--rp-bg-soft);
}

.el-select-dropdown,
.el-picker-panel,
.el-dropdown-menu {
  background: var(--rp-panel-strong);
  border: 1px solid rgba(92, 241, 255, 0.18);
  box-shadow: 0 18px 46px rgba(0, 0, 0, 0.4);
}

.el-select-dropdown__item,
.el-dropdown-menu__item,
.el-date-table td,
.el-month-table td,
.el-year-table td {
  color: var(--rp-muted);
}

.el-select-dropdown__item.is-hovering,
.el-select-dropdown__item:hover,
.el-dropdown-menu__item:hover {
  color: var(--rp-cyan);
  background: rgba(66, 245, 255, 0.08);
}

.el-select-dropdown__item.is-selected,
.el-date-table td.current:not(.disabled) .el-date-table-cell__text {
  color: #061019;
  background: var(--rp-cyan);
}

.el-dialog__title,
.el-message-box__title {
  color: var(--rp-text);
}

.el-dialog__body,
.el-message-box__content {
  color: #bfd3e4;
}

.el-divider {
  border-color: rgba(92, 241, 255, 0.14);
}

.pagination {
  color: var(--rp-muted);
}

.muted,
.empty-text,
.step-msg,
.step-message,
.backend-count,
.target-host-cell small,
.token-row span,
.traffic-cell small,
.traffic-line span,
.user-plan-cell small,
.link-label,
.link-url,
.sub-quick-link .link-label {
  color: var(--rp-muted) !important;
}

.traffic-line strong,
.target-host-cell span,
.usage-metric strong {
  color: var(--rp-cyan) !important;
}

.runtime-log-panel {
  background:
    linear-gradient(rgba(66, 245, 255, 0.035) 1px, transparent 1px),
    rgba(4, 9, 18, 0.92) !important;
  background-size: 100% 24px !important;
  border: 1px solid rgba(92, 241, 255, 0.18);
  box-shadow: inset 0 0 28px rgba(66, 245, 255, 0.05);
}
</style>
