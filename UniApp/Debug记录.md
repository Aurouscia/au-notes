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