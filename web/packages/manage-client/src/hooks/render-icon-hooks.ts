import { h } from "vue";
import { NIcon } from "naive-ui";

export default function useRenderIcon() {
    function renderIcon(icon: any, color?: string) {
        return () => h(NIcon, { color: color }, { default: () => h(icon) });
    }

    return {
        renderIcon
    }
}