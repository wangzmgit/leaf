import { defineStore } from 'pinia';

export const useLoginStore = defineStore({
    id: "showLogin",
    state: () => {
        return {
            showLogin: false,
        }
    },
    actions: {
        setLoginState(val: boolean) {
            this.showLogin = val;
        }
    },
})
