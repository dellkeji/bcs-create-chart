## 简述

bcs create 插件目的是用于快速创建 bcs 场景下的默认chart内容。


## 安装
```
helm plugin install https://github.com/dellkeji/bcs-create-chart.git
```

## 卸载
```
helm plugin uninstall bcs-create 
```

## 使用

#### 创建chart
```
helm bcs-create example

// output the created chart
create chart success: ./example
```

#### 查看版本
```
helm bcs-create version

// output the version
Version  :  v0.0.1
GoVersion:  go1.17.6
```
