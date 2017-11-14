# 3v master page 接口文档
## 3v master page项目简介
3v master page是3v电子优惠券一个服务，包含了会员优惠券查询，详细信息查看，优惠券使用等功能。
点击微信公众下方菜单栏会员里面的我的优惠券，即可进入到3v master page服务。

## 1. 进入3v优惠券主页
#### 简介
* 3v优惠券主页地址。

#### 方法
GET
#### 地址
http://sr.metrowechat.com/pilot3v/index

#### 跨域
配置了对所有域名的GET方法的跨域

#### URL参数

|序号|参数|类型|必须|样例或备注|
|--- |---|---|---|---|
|1|state| string|可选|微信oauth认证后返回的参数|
|2|code| string|可选|换取access_token的票据，定时过期|
|3|openid| string|可选|会员openid，微信公众号入口|
|4|unionid| string|可选|会员unionid，其他微信服务入口|
|5|mappid| string|可选|会员mappid，阿里App入口|

## 2. 获取会员优惠券列表
#### 简介
* 此接口返回会员拥有的所有优惠券信息

#### 方法
GET

#### 地址
http://sr.metrowechat.com/pilot3v/getlist

#### 跨域
配置了对所有域名的GET方法的跨域

#### URL参数
|序号|参数|类型|必须|样例或备注|
|--- |---|---|---|---|
|1|openid|string|必选|oJ9_hjsJxK8PPO03h8_2eremk0ic|

## 3. 获取会员单张优惠券信息
#### 简介
调用此接口，关闭与Esoon的对话，之后的对话将由自动回复处理

#### 方法
POST

#### 地址
http://sr.metrowechat.com/pilot3v/getcoupon

#### 跨域
配置了对所有域名的GET方法的跨域

#### URL参数
|序号|参数|类型|必须|样例或备注|
|--- |---|---|---|---|
|1|epc| string|可选|3v优惠券编号|
|2|gcn| string|可选|3v活动编号|
|3|openid| string|可选|会员openid|
|4|gln| string|可选|3v活动商店编号|


#### 返回
|HTTP code|msg|解释|
|---|---|---|
|200|OK|正确关闭客服对话|
|406|Not Supported|http 方法不符合要求|
|400|params error| 缺少openid 参数|