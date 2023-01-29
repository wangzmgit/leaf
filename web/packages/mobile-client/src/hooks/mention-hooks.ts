export default function useMention() {

    const handleMention = (content: string) => {
        let res = [];
        const state = {
            ReadContent: 'ReadContent',
            ReadUser: 'ReadUser'
        }
        let currentRead = '';//当前读取的内容
        let currentState = state.ReadContent;//当前状态

        content += ' ';//如果@用户结尾没有空格将无法识别
        for (let i = 0; i < content.length; i++) {
            if (currentState === state.ReadContent && content[i] === '@') {
                currentState = state.ReadUser;
                res.push({
                    class: '',
                    value: currentRead,
                    key: null
                });
                currentRead = '';
                continue;
            }

            if (currentState === state.ReadUser && content[i] === ' ') {
                currentState = state.ReadContent;
                res.push({
                    class: 'mention-user',
                    value: `@${currentRead}`,
                    key: currentRead
                });
                currentRead = '';
                continue;
            }

            currentRead += content[i];
        }
        res.push({
            class: '',
            value: currentRead,
            key: null
        })
        return res
    }

    return {
        handleMention
    }
}

