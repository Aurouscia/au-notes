1. 对于有回调设置的对象，应该先设置回调，再令其开始工作，否则首次开始会出问题，例如：
    ```js
        // 先设置回调
        manager.onStart = res => console.log('开始录音', res)
        // onRecognize可能实机测试时才有效果，ide中可能没有
        manager.onRecognize = res => console.log('实时结果', res)
        manager.onStop = res => console.log('最终文字', res.result)
        manager.onError = res => console.error('识别错误', res)

        // 再开始/结束录音
        manager.start({ duration: 60000, lang: 'zh_CN' })
    ```

2. 支付宝小程序中，已经在`pages.json`注册的vue文件，不能再作为组件使用（script不会执行）
3. 支付宝小程序不支持`v-show`
4. 支付宝小程序不支持模板组件调用用`style属性`
5. `::v-deep`样式穿透仅对`pages.json`中的页面起效，对于“组件”不起效
6. 大组件`v-for`嵌套小组件时，不要试图在小组件内更改状态，可能有响应性问题。
    - 正确做法：emit 一个类似于 id 的东西，让外层组件来改