export const timeStrToDateStr = (time: string) => {
    const date = new Date(time);
    const year = date.getFullYear()
    // 由于 getMonth 返回值会比正常月份小 1
    const month = date.getMonth() + 1;
    const day = date.getDate();

    return `${year}-${month}-${day}`
}

export const relativeDate = (time: string) => {
    const originalTime = new Date(time);
    const currentTime = new Date();

    const originalYear = originalTime.getFullYear()
    const originalMonth = originalTime.getMonth() + 1;
    const originalDay = originalTime.getDay();

    if (currentTime.getFullYear() !== originalYear || currentTime.getMonth() + 1 != originalMonth) {
        return `${originalYear}-${originalMonth}-${originalDay}`;
    }

    switch (currentTime.getDay() - originalDay) {
        case 0:
            return "今天";
        case 1:
            return "昨天";
        case 2:
            return "前天";
        default:
            return `${originalYear}-${originalMonth}-${originalDay}`;
    }
} 