export const getRandom = (min: number, max: number) => {
    return Math.ceil(Math.random() * (max - min) + min);
}

// 防抖
export const debounce = (fn: Function, delay = 1000) => {
    let timer: number | null = null;
    return function () {
        if (timer) {
            clearTimeout(timer)
        }
        timer = window.setTimeout(() => {
            fn(arguments);
            timer = null
        }, delay);
    }
}

// 节流
export const thorttle = (fn: Function, delay = 1000) => {
    let timer: number | null = null;
    return function () {
        if (!timer) {
            timer = window.setTimeout(function () {
                console.log('arguments',arguments)
                fn(arguments);

                timer = null;
            }, delay);
        }
    };
};

