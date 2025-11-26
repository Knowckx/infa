# In-fa
In-fa is just a collection of utils of golang

Repo was created after Typhoon In-fa tracking inland south of Shanghai in July 2021.



# 解决拉不到新版本的问题
go get -u github.com/Knowckx/infa@3d2ed45baa3c93c4b6c15c0b642ac56b0ce43952



- 直接引用本地的版本
go mod edit -replace github.com/Knowckx/infa=../infa
有问题 引用方的IDE拿不到刚才的改动 需要点进去刷新文件的内容
有时候行有时候不行 
需要重载界面