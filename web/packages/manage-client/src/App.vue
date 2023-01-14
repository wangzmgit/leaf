<template>
  <n-config-provider :theme-overrides="themeOverrides" :locale="zhCN" :date-locale="dateZhCN">
      <n-notification-provider>
          <n-message-provider>
              <router-view></router-view>
          </n-message-provider>
      </n-notification-provider>
  </n-config-provider>
</template>

<script setup lang="ts">
import { getTheme } from "./theme";
import { zhCN, dateZhCN } from "naive-ui";
import type { GlobalThemeOverrides } from 'naive-ui';
import { NNotificationProvider, NConfigProvider, NMessageProvider } from 'naive-ui';

const theme = getTheme();

const themeOverrides: GlobalThemeOverrides = {
  common: {
      primaryColor: theme.primaryColor
  },
  Button: {
      // default
      textColor: theme.primaryColor,
      textColorFocus: theme.primaryHoverColor,
      borderFocus: theme.primaryHoverColor,
      textColorHover: theme.primaryHoverColor,
      borderHover: theme.primaryHoverColor,
      textColorPressed: theme.primaryHoverColor,
      borderPressed: theme.primaryHoverColor,


      // primary
      borderHoverPrimary: theme.primaryHoverColor,
      colorHoverPrimary: theme.primaryHoverColor,
      colorFocusPrimary: theme.primaryHoverColor,
      borderFocusPrimary: theme.primaryHoverColor,
      colorPressedPrimary: theme.primaryHoverColor,
      borderPressedPrimary: theme.primaryHoverColor,

      //text-default
      textColorText: '#999',
      textColorTextHover: theme.primaryColor,
      textColorTextPressed: theme.primaryColor
  },
  Select: {
      peers: {
          InternalSelection: {
              textColor: theme.primaryColor,
              borderFocus: theme.primaryColor,
              borderHover: theme.primaryColor
          },
      },
  },
  Input: {
      borderFocus: theme.primaryColor,
      borderHover: theme.primaryColor,
  },
  Menu: {
      itemTextColorHoverHorizontal: theme.primaryHoverColor,
  }
}
</script>

<style>
body {
  margin: 0;
  padding: 0;
  font-family: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto,
      "Helvetica Neue", Arial, "Noto Sans", sans-serif, "Apple Color Emoji",
      "Segoe UI Emoji", "Segoe UI Symbol", "Noto Color Emoji";
}
</style>
