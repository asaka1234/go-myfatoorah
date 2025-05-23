# go-myfatoorah

Doc
==============
1. https://docs.myfatoorah.com/docs/overview
2. demo account: https://demo.myfatoorah.com/



鉴权
==============
1. pre-order时: 把accessKey放到了header里一起传过去. 用这个来唯一鉴权(没有签名)
2. webhook回调时.psp会做一个sign签名, 并将结果放在MyFatoorah-Signature这样一个head里, merchant需要验证一下


回调地址
==============
在psp端指定的callback地址, 所以无法在pre-order中动态指定


Payment Process
===============
1. Generate a invoice link  (付款链接, 类似收银台)
2. User pay for this via invoice link (用户在收银台web页面中付款)
3. Myfatoorah call the notify-api of merchant (回调通知)


Summary
==============
1. 生成一个收银台支付地址(invoice link)
   1. merchant会发一个post请求给psp, 告知它:支付的金额,货币这些.  （head带上Token，好让psp知道是哪个商户的请求）
   2. psp会生成一个transaction, 并返回这个transId 和 对应收银台url 给 merchant, 这个url就可以展示给user来支付了
2. 用户支付后,会通知在后台里设置的webhook地址. 注意:这个回调请求里会有一个header,名字为: MyFatoorah-Signature, 这里边是一个返回值的签名，需要用webhook里配置的secret来把返回的数据签名, 这个值必须跟MyFatoorah-Signature的value保持一致才算验证成功.


Comment
===============
only support deposit, not support withdraw.