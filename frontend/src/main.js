import { createApp } from 'vue'
import { createPinia } from 'pinia'
import { ElButton } from 'element-plus/es/components/button/index.mjs'
import { ElCard } from 'element-plus/es/components/card/index.mjs'
import { ElForm, ElFormItem } from 'element-plus/es/components/form/index.mjs'
import { ElIcon } from 'element-plus/es/components/icon/index.mjs'
import { ElInput } from 'element-plus/es/components/input/index.mjs'
import { provideGlobalConfig } from 'element-plus/es/components/config-provider/index.mjs'
import {
  CircleCheck,
  CircleClose,
  Connection,
  Document,
  Goods,
  HomeFilled,
  Key,
  Loading,
  Lock,
  Message,
  Monitor,
  Refresh,
  Ticket,
  User,
  UserFilled,
} from '@element-plus/icons-vue'
import zhCn from 'element-plus/dist/locale/zh-cn.mjs'
import 'element-plus/theme-chalk/base.css'
import 'element-plus/theme-chalk/el-button.css'
import 'element-plus/theme-chalk/el-card.css'
import 'element-plus/theme-chalk/el-form.css'
import 'element-plus/theme-chalk/el-form-item.css'
import 'element-plus/theme-chalk/el-icon.css'
import 'element-plus/theme-chalk/el-input.css'
import 'element-plus/theme-chalk/el-message.css'

import App from './App.vue'
import router from './router'

const app = createApp(App)

provideGlobalConfig({ locale: zhCn }, app, true)

const coreComponents = [
  ElButton,
  ElCard,
  ElForm,
  ElFormItem,
  ElIcon,
  ElInput,
]

for (const component of coreComponents) {
  app.use(component, { locale: zhCn })
}

const icons = {
  CircleCheck,
  CircleClose,
  Connection,
  Document,
  Goods,
  HomeFilled,
  Key,
  Loading,
  Lock,
  Message,
  Monitor,
  Refresh,
  Ticket,
  User,
  UserFilled,
}

for (const [key, component] of Object.entries(icons)) {
  app.component(key, component)
}

app.use(createPinia())

let extendedComponentsInstalled = false

async function installExtendedComponents() {
  if (extendedComponentsInstalled) return

  const [
    { ElAlert },
    { ElCheckbox },
    { ElCol },
    { ElAside, ElContainer, ElHeader, ElMain },
    { ElDatePicker },
    { ElDescriptions, ElDescriptionsItem },
    { ElDialog },
    { ElDivider },
    { ElEmpty },
    { ElInputNumber },
    { ElLoading },
    { ElMenu, ElMenuItem },
    { ElOption, ElSelect },
    { ElPagination },
    { ElProgress },
    { ElRadio, ElRadioButton, ElRadioGroup },
    { ElRow },
    { ElSwitch },
    { ElTabPane, ElTabs },
    { ElTable, ElTableColumn },
    { ElTag },
  ] = await Promise.all([
    import('element-plus/es/components/alert/index.mjs'),
    import('element-plus/es/components/checkbox/index.mjs'),
    import('element-plus/es/components/col/index.mjs'),
    import('element-plus/es/components/container/index.mjs'),
    import('element-plus/es/components/date-picker/index.mjs'),
    import('element-plus/es/components/descriptions/index.mjs'),
    import('element-plus/es/components/dialog/index.mjs'),
    import('element-plus/es/components/divider/index.mjs'),
    import('element-plus/es/components/empty/index.mjs'),
    import('element-plus/es/components/input-number/index.mjs'),
    import('element-plus/es/components/loading/index.mjs'),
    import('element-plus/es/components/menu/index.mjs'),
    import('element-plus/es/components/select/index.mjs'),
    import('element-plus/es/components/pagination/index.mjs'),
    import('element-plus/es/components/progress/index.mjs'),
    import('element-plus/es/components/radio/index.mjs'),
    import('element-plus/es/components/row/index.mjs'),
    import('element-plus/es/components/switch/index.mjs'),
    import('element-plus/es/components/tabs/index.mjs'),
    import('element-plus/es/components/table/index.mjs'),
    import('element-plus/es/components/tag/index.mjs'),
    import('element-plus/theme-chalk/el-alert.css'),
    import('element-plus/theme-chalk/el-checkbox.css'),
    import('element-plus/theme-chalk/el-aside.css'),
    import('element-plus/theme-chalk/el-col.css'),
    import('element-plus/theme-chalk/el-container.css'),
    import('element-plus/theme-chalk/el-date-picker.css'),
    import('element-plus/theme-chalk/el-descriptions.css'),
    import('element-plus/theme-chalk/el-descriptions-item.css'),
    import('element-plus/theme-chalk/el-dialog.css'),
    import('element-plus/theme-chalk/el-divider.css'),
    import('element-plus/theme-chalk/el-empty.css'),
    import('element-plus/theme-chalk/el-header.css'),
    import('element-plus/theme-chalk/el-input-number.css'),
    import('element-plus/theme-chalk/el-loading.css'),
    import('element-plus/theme-chalk/el-main.css'),
    import('element-plus/theme-chalk/el-menu.css'),
    import('element-plus/theme-chalk/el-menu-item.css'),
    import('element-plus/theme-chalk/el-message-box.css'),
    import('element-plus/theme-chalk/el-option.css'),
    import('element-plus/theme-chalk/el-overlay.css'),
    import('element-plus/theme-chalk/el-pagination.css'),
    import('element-plus/theme-chalk/el-progress.css'),
    import('element-plus/theme-chalk/el-radio-button.css'),
    import('element-plus/theme-chalk/el-radio-group.css'),
    import('element-plus/theme-chalk/el-row.css'),
    import('element-plus/theme-chalk/el-select.css'),
    import('element-plus/theme-chalk/el-switch.css'),
    import('element-plus/theme-chalk/el-tab-pane.css'),
    import('element-plus/theme-chalk/el-table.css'),
    import('element-plus/theme-chalk/el-table-column.css'),
    import('element-plus/theme-chalk/el-tabs.css'),
    import('element-plus/theme-chalk/el-tag.css'),
  ])

  const extendedComponents = [
    ElAlert,
    ElCheckbox,
    ElAside,
    ElCol,
    ElContainer,
    ElDatePicker,
    ElDescriptions,
    ElDescriptionsItem,
    ElDialog,
    ElDivider,
    ElEmpty,
    ElHeader,
    ElInputNumber,
    ElMain,
    ElMenu,
    ElMenuItem,
    ElOption,
    ElPagination,
    ElProgress,
    ElRadio,
    ElRadioButton,
    ElRadioGroup,
    ElRow,
    ElSelect,
    ElSwitch,
    ElTabPane,
    ElTable,
    ElTableColumn,
    ElTabs,
    ElTag,
  ]

  for (const component of extendedComponents) {
    app.use(component, { locale: zhCn })
  }
  app.use(ElLoading)
  extendedComponentsInstalled = true
}

router.beforeEach(async (to) => {
  if (!to.meta.guest) {
    await installExtendedComponents()
  }
})

app.use(router)

app.mount('#app')
